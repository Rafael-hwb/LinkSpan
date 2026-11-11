package service

import (
	"context"
	"encoding/json"
	"linkspan/admin/internal/dto/req"
	"linkspan/admin/internal/remote"
	"linkspan/internal/common/convention/result"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type ShortLinkService struct {
	groups *GroupService
	core   *remote.ShortLinkClient
}

func NewShortLinkService(groups *GroupService, core *remote.ShortLinkClient) *ShortLinkService {
	return &ShortLinkService{groups: groups, core: core}
}

func (s *ShortLinkService) Create(ctx context.Context, request req.ShortLinkCreateReq) (result.Result[json.RawMessage], error) {
	request.GID = strings.TrimSpace(request.GID)

	if err := s.groups.RequireOwnedGroup(ctx, request.GID); err != nil {
		return result.Result[json.RawMessage]{}, err
	}

	return s.core.Call(ctx, http.MethodPost, "/create", request)
}

func (s *ShortLinkService) Page(ctx context.Context, request req.ShortLinkPageReq) (result.Result[json.RawMessage], error) {
	request.GID = strings.TrimSpace(request.GID)
	if err := s.groups.RequireOwnedGroup(ctx, request.GID); err != nil {
		return result.Result[json.RawMessage]{}, err
	}

	query := url.Values{}
	query.Set("gid", request.GID)
	query.Set("current", strconv.FormatInt(int64(request.Current), 10))
	query.Set("size", strconv.FormatInt(int64(request.Size), 10))

	return s.core.Call(ctx, http.MethodGet, "/page?"+query.Encode(), nil)
}

func (s *ShortLinkService) Update(ctx context.Context, request req.ShortLinkUpdateReq) (result.Result[json.RawMessage], error) {
	request.OriginGID = strings.TrimSpace(request.OriginGID)
	request.GID = strings.TrimSpace(request.GID)
	if err := s.groups.RequireOwnedGroup(ctx, request.OriginGID); err != nil {
		return result.Result[json.RawMessage]{}, err
	}
	if err := s.groups.RequireOwnedGroup(ctx, request.GID); err != nil {
		return result.Result[json.RawMessage]{}, err
	}
	return s.core.Call(ctx, http.MethodPost, "/update", request)
}
