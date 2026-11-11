package service

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"

	"linkspan/internal/common/convention/errorcode"
	"linkspan/internal/common/convention/exception"
	"linkspan/project/internal/dao/entity"
	"linkspan/project/internal/dto/req"
	"linkspan/project/internal/dto/resp"
	"linkspan/project/internal/repository"

	"github.com/go-sql-driver/mysql"
)

type ShortLinkService struct {
	links  *repository.ShortLinkRepository
	routes *repository.ShortLinkGotoRepository
}

func NewShortLinkService(links *repository.ShortLinkRepository, routes *repository.ShortLinkGotoRepository) *ShortLinkService {
	return &ShortLinkService{links: links, routes: routes}
}

func (s *ShortLinkService) Create(ctx context.Context, request req.ShortLinkCreateReq) (resp.ShortLinkCreateResp, error) {
	var response resp.ShortLinkCreateResp
	link, err := buildShortLink(request)
	if err != nil {
		return response, err
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	for attempt := 0; attempt < 10; attempt++ {
		var random [4]byte
		if _, err := rand.Read(random[:]); err != nil {
			return response, fmt.Errorf("生成短链接后缀失败: %w", err)
		}
		link.ShortURI = hex.EncodeToString(random[:])
		link.FullShortURL = link.Domain + "/" + link.ShortURI

		route := entity.ShortLinkGotoDO{GID: link.GID, FullShortURL: link.FullShortURL}
		err := s.links.WithTransaction(ctx, func(tx *sql.Tx) error {
			if err := s.links.Insert(ctx, tx, link); err != nil {
				return err
			}
			return s.routes.Insert(ctx, tx, route)
		})
		if err == nil {
			return resp.ShortLinkCreateResp{
				GID: link.GID, OriginURL: link.OriginURL, FullShortURL: "http://" + link.FullShortURL,
			}, nil
		}
		var mysqlErr *mysql.MySQLError
		if !errors.As(err, &mysqlErr) || mysqlErr.Number != 1062 {
			return response, err
		}
	}
	return response, errors.New("短链接标识重试次数已耗尽")
}

func buildShortLink(request req.ShortLinkCreateReq) (entity.ShortLinkDO, error) {
	var link entity.ShortLinkDO
	domain := strings.ToLower(strings.TrimSpace(request.Domain))
	if domain != "localhost:8002" {
		return link, &exception.ServiceError{Code: errorcode.ParameterInvalid.Code, Message: "domain 须为 localhost:8002"}
	}
	origin := strings.TrimSpace(request.OriginURL)
	target, err := url.Parse(origin)
	if err != nil || target.Hostname() == "" || (target.Scheme != "http" && target.Scheme != "https") {
		return link, &exception.ServiceError{Code: errorcode.ParameterInvalid.Code, Message: "originUrl 须为完整的 HTTP 或 HTTPS 地址"}
	}
	gid := strings.TrimSpace(request.GID)
	if gid == "" {
		return link, &exception.ServiceError{Code: errorcode.ParameterInvalid.Code, Message: "gid 不能为空"}
	}
	if request.CreatedType == nil || (*request.CreatedType != 0 && *request.CreatedType != 1) || request.ValidDateType == nil {
		return link, &exception.ServiceError{Code: errorcode.ParameterInvalid.Code, Message: "请检查创建类型和有效期类型"}
	}

	var validDate *time.Time
	switch *request.ValidDateType {
	case 0:
		if request.ValidDate != nil {
			return link, &exception.ServiceError{Code: errorcode.ParameterInvalid.Code, Message: "永久有效时 validDate 须为 null"}
		}
	case 1:
		if request.ValidDate == nil {
			return link, &exception.ServiceError{Code: errorcode.ParameterInvalid.Code, Message: "自定义有效期需要 validDate"}
		}
		const layout = "2006-01-02 15:04:05"
		parsed, err := time.ParseInLocation(layout, *request.ValidDate, time.FixedZone("UTC+8", 8*60*60))
		if err != nil || parsed.Format(layout) != *request.ValidDate || !parsed.After(time.Now()) {
			return link, &exception.ServiceError{Code: errorcode.ParameterInvalid.Code, Message: "validDate 须为未来的北京时间，格式为 yyyy-MM-dd HH:mm:ss"}
		}
		validDate = &parsed
	default:
		return link, &exception.ServiceError{Code: errorcode.ParameterInvalid.Code, Message: "validDateType 须为 0 或 1"}
	}
	return entity.ShortLinkDO{
		Domain: domain, OriginURL: origin, GID: gid, CreatedType: *request.CreatedType,
		ValidDateType: *request.ValidDateType, ValidDate: validDate, Describe: request.Describe,
	}, nil
}

func (s *ShortLinkService) Page(ctx context.Context, request req.ShortLinkPageReq) (resp.ShortLinkPageResp, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	size := int64(request.Size)
	offset := (int64(request.Current) - 1) * size
	links, total, err := s.links.PageByGID(ctx, request.GID, size, offset)

	if err != nil {
		return resp.ShortLinkPageResp{}, err
	}

	result := resp.ShortLinkPageResp{
		Records: make([]resp.ShortLinkPageItem, 0, len(links)),
		Total:   total, Current: request.Current, Size: request.Size,
	}

	for _, link := range links {
		result.Records = append(result.Records, resp.ShortLinkPageItem{
			GID:          link.GID,
			FullShortURL: "http://" + link.FullShortURL,
			OriginURL:    link.OriginURL,
			Describe:     link.Describe,
		})
	}

	return result, nil
}

func (s *ShortLinkService) Update(ctx context.Context, request req.ShortLinkUpdateReq) error {
	gid := strings.TrimSpace(request.GID)
	if gid != strings.TrimSpace(request.OriginGID) {
		return &exception.ServiceError{Code: errorcode.ParameterInvalid.Code, Message: "本节仅支持同一分组内修改"}
	}

	const prefix = "localhost:8002/"
	fullShortURL := strings.TrimPrefix(strings.TrimSpace(request.FullShortURL), "http://")
	suffix, ok := strings.CutPrefix(fullShortURL, prefix)
	if !ok || len(suffix) != 8 {
		return &exception.ServiceError{Code: errorcode.ParameterInvalid.Code, Message: "短地址格式错误"}
	}
	if _, err := hex.DecodeString(suffix); err != nil || suffix != strings.ToLower(suffix) {
		return &exception.ServiceError{Code: errorcode.ParameterInvalid.Code, Message: "短地址后缀格式错误"}
	}

	createdType := 0
	link, err := buildShortLink(req.ShortLinkCreateReq{
		Domain: "localhost:8002", OriginURL: request.OriginURL, GID: gid,
		CreatedType: &createdType, ValidDateType: request.ValidDateType,
		ValidDate: request.ValidDate, Describe: request.Describe,
	})
	if err != nil {
		return err
	}
	link.FullShortURL = fullShortURL

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return s.links.Update(ctx, link)
}
