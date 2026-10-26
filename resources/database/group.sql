USE linkspan;

-- 分组分表
CREATE TABLE t_group_0 (
    id          BIGINT NOT NULL AUTO_INCREMENT COMMENT 'ID',
    gid         VARCHAR(32)  DEFAULT NULL COMMENT '分组标识',
    name        VARCHAR(64)  DEFAULT NULL COMMENT '分组名称',
    username    VARCHAR(256) DEFAULT NULL COMMENT '创建分组用户名',
    sort_order  INT          DEFAULT NULL COMMENT '分组排序',
    create_time DATETIME     DEFAULT NULL COMMENT '创建时间',
    update_time DATETIME     DEFAULT NULL COMMENT '修改时间',
    del_flag    TINYINT      DEFAULT NULL COMMENT '删除标识 0：未删除 1：已删除',

    PRIMARY KEY (id),
    KEY idx_username (username) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;