-- 沿用原项目结构，移除整数显示宽度。已存在的表保留，结构需单独核对。
USE linkspan;

CREATE TABLE IF NOT EXISTS `t_link_0`
(
    `id`              bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `domain`          varchar(128)                                   DEFAULT NULL COMMENT '域名',
    `short_uri`       varchar(8) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '短链接',
    `full_short_url`  varchar(128)                                   DEFAULT NULL COMMENT '完整短链接',
    `origin_url`      varchar(1024)                                  DEFAULT NULL COMMENT '原始链接',
    `click_num`       int DEFAULT '0' COMMENT '点击量',
    `gid`             varchar(32)                                    DEFAULT 'default' COMMENT '分组标识',
    `favicon`         varchar(256)                                   DEFAULT NULL COMMENT '网站图标',
    `enable_status`   tinyint DEFAULT NULL COMMENT '启用标识 0：启用 1：未启用',
    `created_type`    tinyint DEFAULT NULL COMMENT '创建类型 0：接口创建 1：控制台创建',
    `valid_date_type` tinyint DEFAULT NULL COMMENT '有效期类型 0：永久有效 1：自定义',
    `valid_date`      datetime                                       DEFAULT NULL COMMENT '有效期',
    `describe`        varchar(1024)                                  DEFAULT NULL COMMENT '描述',
    `total_pv`        int DEFAULT NULL COMMENT '历史PV',
    `total_uv`        int DEFAULT NULL COMMENT '历史UV',
    `total_uip`       int DEFAULT NULL COMMENT '历史UIP',
    `create_time`     datetime                                       DEFAULT NULL COMMENT '创建时间',
    `update_time`     datetime                                       DEFAULT NULL COMMENT '修改时间',
    `del_time`        bigint DEFAULT '0' COMMENT '删除时间戳',
    `del_flag`        tinyint DEFAULT NULL COMMENT '删除标识 0：未删除 1：已删除',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_unique_full-short-url` (`full_short_url`,`del_time`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_link_1`
(
    `id`              bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `domain`          varchar(128)                                   DEFAULT NULL COMMENT '域名',
    `short_uri`       varchar(8) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '短链接',
    `full_short_url`  varchar(128)                                   DEFAULT NULL COMMENT '完整短链接',
    `origin_url`      varchar(1024)                                  DEFAULT NULL COMMENT '原始链接',
    `click_num`       int DEFAULT '0' COMMENT '点击量',
    `gid`             varchar(32)                                    DEFAULT 'default' COMMENT '分组标识',
    `favicon`         varchar(256)                                   DEFAULT NULL COMMENT '网站图标',
    `enable_status`   tinyint DEFAULT NULL COMMENT '启用标识 0：启用 1：未启用',
    `created_type`    tinyint DEFAULT NULL COMMENT '创建类型 0：接口创建 1：控制台创建',
    `valid_date_type` tinyint DEFAULT NULL COMMENT '有效期类型 0：永久有效 1：自定义',
    `valid_date`      datetime                                       DEFAULT NULL COMMENT '有效期',
    `describe`        varchar(1024)                                  DEFAULT NULL COMMENT '描述',
    `total_pv`        int DEFAULT NULL COMMENT '历史PV',
    `total_uv`        int DEFAULT NULL COMMENT '历史UV',
    `total_uip`       int DEFAULT NULL COMMENT '历史UIP',
    `create_time`     datetime                                       DEFAULT NULL COMMENT '创建时间',
    `update_time`     datetime                                       DEFAULT NULL COMMENT '修改时间',
    `del_time`        bigint DEFAULT '0' COMMENT '删除时间戳',
    `del_flag`        tinyint DEFAULT NULL COMMENT '删除标识 0：未删除 1：已删除',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_unique_full-short-url` (`full_short_url`,`del_time`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_link_10`
(
    `id`              bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `domain`          varchar(128)                                   DEFAULT NULL COMMENT '域名',
    `short_uri`       varchar(8) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '短链接',
    `full_short_url`  varchar(128)                                   DEFAULT NULL COMMENT '完整短链接',
    `origin_url`      varchar(1024)                                  DEFAULT NULL COMMENT '原始链接',
    `click_num`       int DEFAULT '0' COMMENT '点击量',
    `gid`             varchar(32)                                    DEFAULT 'default' COMMENT '分组标识',
    `favicon`         varchar(256)                                   DEFAULT NULL COMMENT '网站图标',
    `enable_status`   tinyint DEFAULT NULL COMMENT '启用标识 0：启用 1：未启用',
    `created_type`    tinyint DEFAULT NULL COMMENT '创建类型 0：接口创建 1：控制台创建',
    `valid_date_type` tinyint DEFAULT NULL COMMENT '有效期类型 0：永久有效 1：自定义',
    `valid_date`      datetime                                       DEFAULT NULL COMMENT '有效期',
    `describe`        varchar(1024)                                  DEFAULT NULL COMMENT '描述',
    `total_pv`        int DEFAULT NULL COMMENT '历史PV',
    `total_uv`        int DEFAULT NULL COMMENT '历史UV',
    `total_uip`       int DEFAULT NULL COMMENT '历史UIP',
    `create_time`     datetime                                       DEFAULT NULL COMMENT '创建时间',
    `update_time`     datetime                                       DEFAULT NULL COMMENT '修改时间',
    `del_time`        bigint DEFAULT '0' COMMENT '删除时间戳',
    `del_flag`        tinyint DEFAULT NULL COMMENT '删除标识 0：未删除 1：已删除',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_unique_full-short-url` (`full_short_url`,`del_time`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_link_11`
(
    `id`              bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `domain`          varchar(128)                                   DEFAULT NULL COMMENT '域名',
    `short_uri`       varchar(8) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '短链接',
    `full_short_url`  varchar(128)                                   DEFAULT NULL COMMENT '完整短链接',
    `origin_url`      varchar(1024)                                  DEFAULT NULL COMMENT '原始链接',
    `click_num`       int DEFAULT '0' COMMENT '点击量',
    `gid`             varchar(32)                                    DEFAULT 'default' COMMENT '分组标识',
    `favicon`         varchar(256)                                   DEFAULT NULL COMMENT '网站图标',
    `enable_status`   tinyint DEFAULT NULL COMMENT '启用标识 0：启用 1：未启用',
    `created_type`    tinyint DEFAULT NULL COMMENT '创建类型 0：接口创建 1：控制台创建',
    `valid_date_type` tinyint DEFAULT NULL COMMENT '有效期类型 0：永久有效 1：自定义',
    `valid_date`      datetime                                       DEFAULT NULL COMMENT '有效期',
    `describe`        varchar(1024)                                  DEFAULT NULL COMMENT '描述',
    `total_pv`        int DEFAULT NULL COMMENT '历史PV',
    `total_uv`        int DEFAULT NULL COMMENT '历史UV',
    `total_uip`       int DEFAULT NULL COMMENT '历史UIP',
    `create_time`     datetime                                       DEFAULT NULL COMMENT '创建时间',
    `update_time`     datetime                                       DEFAULT NULL COMMENT '修改时间',
    `del_time`        bigint DEFAULT '0' COMMENT '删除时间戳',
    `del_flag`        tinyint DEFAULT NULL COMMENT '删除标识 0：未删除 1：已删除',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_unique_full-short-url` (`full_short_url`,`del_time`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_link_12`
(
    `id`              bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `domain`          varchar(128)                                   DEFAULT NULL COMMENT '域名',
    `short_uri`       varchar(8) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '短链接',
    `full_short_url`  varchar(128)                                   DEFAULT NULL COMMENT '完整短链接',
    `origin_url`      varchar(1024)                                  DEFAULT NULL COMMENT '原始链接',
    `click_num`       int DEFAULT '0' COMMENT '点击量',
    `gid`             varchar(32)                                    DEFAULT 'default' COMMENT '分组标识',
    `favicon`         varchar(256)                                   DEFAULT NULL COMMENT '网站图标',
    `enable_status`   tinyint DEFAULT NULL COMMENT '启用标识 0：启用 1：未启用',
    `created_type`    tinyint DEFAULT NULL COMMENT '创建类型 0：接口创建 1：控制台创建',
    `valid_date_type` tinyint DEFAULT NULL COMMENT '有效期类型 0：永久有效 1：自定义',
    `valid_date`      datetime                                       DEFAULT NULL COMMENT '有效期',
    `describe`        varchar(1024)                                  DEFAULT NULL COMMENT '描述',
    `total_pv`        int DEFAULT NULL COMMENT '历史PV',
    `total_uv`        int DEFAULT NULL COMMENT '历史UV',
    `total_uip`       int DEFAULT NULL COMMENT '历史UIP',
    `create_time`     datetime                                       DEFAULT NULL COMMENT '创建时间',
    `update_time`     datetime                                       DEFAULT NULL COMMENT '修改时间',
    `del_time`        bigint DEFAULT '0' COMMENT '删除时间戳',
    `del_flag`        tinyint DEFAULT NULL COMMENT '删除标识 0：未删除 1：已删除',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_unique_full-short-url` (`full_short_url`,`del_time`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_link_13`
(
    `id`              bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `domain`          varchar(128)                                   DEFAULT NULL COMMENT '域名',
    `short_uri`       varchar(8) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '短链接',
    `full_short_url`  varchar(128)                                   DEFAULT NULL COMMENT '完整短链接',
    `origin_url`      varchar(1024)                                  DEFAULT NULL COMMENT '原始链接',
    `click_num`       int DEFAULT '0' COMMENT '点击量',
    `gid`             varchar(32)                                    DEFAULT 'default' COMMENT '分组标识',
    `favicon`         varchar(256)                                   DEFAULT NULL COMMENT '网站图标',
    `enable_status`   tinyint DEFAULT NULL COMMENT '启用标识 0：启用 1：未启用',
    `created_type`    tinyint DEFAULT NULL COMMENT '创建类型 0：接口创建 1：控制台创建',
    `valid_date_type` tinyint DEFAULT NULL COMMENT '有效期类型 0：永久有效 1：自定义',
    `valid_date`      datetime                                       DEFAULT NULL COMMENT '有效期',
    `describe`        varchar(1024)                                  DEFAULT NULL COMMENT '描述',
    `total_pv`        int DEFAULT NULL COMMENT '历史PV',
    `total_uv`        int DEFAULT NULL COMMENT '历史UV',
    `total_uip`       int DEFAULT NULL COMMENT '历史UIP',
    `create_time`     datetime                                       DEFAULT NULL COMMENT '创建时间',
    `update_time`     datetime                                       DEFAULT NULL COMMENT '修改时间',
    `del_time`        bigint DEFAULT '0' COMMENT '删除时间戳',
    `del_flag`        tinyint DEFAULT NULL COMMENT '删除标识 0：未删除 1：已删除',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_unique_full-short-url` (`full_short_url`,`del_time`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_link_14`
(
    `id`              bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `domain`          varchar(128)                                   DEFAULT NULL COMMENT '域名',
    `short_uri`       varchar(8) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '短链接',
    `full_short_url`  varchar(128)                                   DEFAULT NULL COMMENT '完整短链接',
    `origin_url`      varchar(1024)                                  DEFAULT NULL COMMENT '原始链接',
    `click_num`       int DEFAULT '0' COMMENT '点击量',
    `gid`             varchar(32)                                    DEFAULT 'default' COMMENT '分组标识',
    `favicon`         varchar(256)                                   DEFAULT NULL COMMENT '网站图标',
    `enable_status`   tinyint DEFAULT NULL COMMENT '启用标识 0：启用 1：未启用',
    `created_type`    tinyint DEFAULT NULL COMMENT '创建类型 0：接口创建 1：控制台创建',
    `valid_date_type` tinyint DEFAULT NULL COMMENT '有效期类型 0：永久有效 1：自定义',
    `valid_date`      datetime                                       DEFAULT NULL COMMENT '有效期',
    `describe`        varchar(1024)                                  DEFAULT NULL COMMENT '描述',
    `total_pv`        int DEFAULT NULL COMMENT '历史PV',
    `total_uv`        int DEFAULT NULL COMMENT '历史UV',
    `total_uip`       int DEFAULT NULL COMMENT '历史UIP',
    `create_time`     datetime                                       DEFAULT NULL COMMENT '创建时间',
    `update_time`     datetime                                       DEFAULT NULL COMMENT '修改时间',
    `del_time`        bigint DEFAULT '0' COMMENT '删除时间戳',
    `del_flag`        tinyint DEFAULT NULL COMMENT '删除标识 0：未删除 1：已删除',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_unique_full-short-url` (`full_short_url`,`del_time`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_link_15`
(
    `id`              bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `domain`          varchar(128)                                   DEFAULT NULL COMMENT '域名',
    `short_uri`       varchar(8) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '短链接',
    `full_short_url`  varchar(128)                                   DEFAULT NULL COMMENT '完整短链接',
    `origin_url`      varchar(1024)                                  DEFAULT NULL COMMENT '原始链接',
    `click_num`       int DEFAULT '0' COMMENT '点击量',
    `gid`             varchar(32)                                    DEFAULT 'default' COMMENT '分组标识',
    `favicon`         varchar(256)                                   DEFAULT NULL COMMENT '网站图标',
    `enable_status`   tinyint DEFAULT NULL COMMENT '启用标识 0：启用 1：未启用',
    `created_type`    tinyint DEFAULT NULL COMMENT '创建类型 0：接口创建 1：控制台创建',
    `valid_date_type` tinyint DEFAULT NULL COMMENT '有效期类型 0：永久有效 1：自定义',
    `valid_date`      datetime                                       DEFAULT NULL COMMENT '有效期',
    `describe`        varchar(1024)                                  DEFAULT NULL COMMENT '描述',
    `total_pv`        int DEFAULT NULL COMMENT '历史PV',
    `total_uv`        int DEFAULT NULL COMMENT '历史UV',
    `total_uip`       int DEFAULT NULL COMMENT '历史UIP',
    `create_time`     datetime                                       DEFAULT NULL COMMENT '创建时间',
    `update_time`     datetime                                       DEFAULT NULL COMMENT '修改时间',
    `del_time`        bigint DEFAULT '0' COMMENT '删除时间戳',
    `del_flag`        tinyint DEFAULT NULL COMMENT '删除标识 0：未删除 1：已删除',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_unique_full-short-url` (`full_short_url`,`del_time`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_link_2`
(
    `id`              bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `domain`          varchar(128)                                   DEFAULT NULL COMMENT '域名',
    `short_uri`       varchar(8) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '短链接',
    `full_short_url`  varchar(128)                                   DEFAULT NULL COMMENT '完整短链接',
    `origin_url`      varchar(1024)                                  DEFAULT NULL COMMENT '原始链接',
    `click_num`       int DEFAULT '0' COMMENT '点击量',
    `gid`             varchar(32)                                    DEFAULT 'default' COMMENT '分组标识',
    `favicon`         varchar(256)                                   DEFAULT NULL COMMENT '网站图标',
    `enable_status`   tinyint DEFAULT NULL COMMENT '启用标识 0：启用 1：未启用',
    `created_type`    tinyint DEFAULT NULL COMMENT '创建类型 0：接口创建 1：控制台创建',
    `valid_date_type` tinyint DEFAULT NULL COMMENT '有效期类型 0：永久有效 1：自定义',
    `valid_date`      datetime                                       DEFAULT NULL COMMENT '有效期',
    `describe`        varchar(1024)                                  DEFAULT NULL COMMENT '描述',
    `total_pv`        int DEFAULT NULL COMMENT '历史PV',
    `total_uv`        int DEFAULT NULL COMMENT '历史UV',
    `total_uip`       int DEFAULT NULL COMMENT '历史UIP',
    `create_time`     datetime                                       DEFAULT NULL COMMENT '创建时间',
    `update_time`     datetime                                       DEFAULT NULL COMMENT '修改时间',
    `del_time`        bigint DEFAULT '0' COMMENT '删除时间戳',
    `del_flag`        tinyint DEFAULT NULL COMMENT '删除标识 0：未删除 1：已删除',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_unique_full-short-url` (`full_short_url`,`del_time`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_link_3`
(
    `id`              bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `domain`          varchar(128)                                   DEFAULT NULL COMMENT '域名',
    `short_uri`       varchar(8) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '短链接',
    `full_short_url`  varchar(128)                                   DEFAULT NULL COMMENT '完整短链接',
    `origin_url`      varchar(1024)                                  DEFAULT NULL COMMENT '原始链接',
    `click_num`       int DEFAULT '0' COMMENT '点击量',
    `gid`             varchar(32)                                    DEFAULT 'default' COMMENT '分组标识',
    `favicon`         varchar(256)                                   DEFAULT NULL COMMENT '网站图标',
    `enable_status`   tinyint DEFAULT NULL COMMENT '启用标识 0：启用 1：未启用',
    `created_type`    tinyint DEFAULT NULL COMMENT '创建类型 0：接口创建 1：控制台创建',
    `valid_date_type` tinyint DEFAULT NULL COMMENT '有效期类型 0：永久有效 1：自定义',
    `valid_date`      datetime                                       DEFAULT NULL COMMENT '有效期',
    `describe`        varchar(1024)                                  DEFAULT NULL COMMENT '描述',
    `total_pv`        int DEFAULT NULL COMMENT '历史PV',
    `total_uv`        int DEFAULT NULL COMMENT '历史UV',
    `total_uip`       int DEFAULT NULL COMMENT '历史UIP',
    `create_time`     datetime                                       DEFAULT NULL COMMENT '创建时间',
    `update_time`     datetime                                       DEFAULT NULL COMMENT '修改时间',
    `del_time`        bigint DEFAULT '0' COMMENT '删除时间戳',
    `del_flag`        tinyint DEFAULT NULL COMMENT '删除标识 0：未删除 1：已删除',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_unique_full-short-url` (`full_short_url`,`del_time`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_link_4`
(
    `id`              bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `domain`          varchar(128)                                   DEFAULT NULL COMMENT '域名',
    `short_uri`       varchar(8) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '短链接',
    `full_short_url`  varchar(128)                                   DEFAULT NULL COMMENT '完整短链接',
    `origin_url`      varchar(1024)                                  DEFAULT NULL COMMENT '原始链接',
    `click_num`       int DEFAULT '0' COMMENT '点击量',
    `gid`             varchar(32)                                    DEFAULT 'default' COMMENT '分组标识',
    `favicon`         varchar(256)                                   DEFAULT NULL COMMENT '网站图标',
    `enable_status`   tinyint DEFAULT NULL COMMENT '启用标识 0：启用 1：未启用',
    `created_type`    tinyint DEFAULT NULL COMMENT '创建类型 0：接口创建 1：控制台创建',
    `valid_date_type` tinyint DEFAULT NULL COMMENT '有效期类型 0：永久有效 1：自定义',
    `valid_date`      datetime                                       DEFAULT NULL COMMENT '有效期',
    `describe`        varchar(1024)                                  DEFAULT NULL COMMENT '描述',
    `total_pv`        int DEFAULT NULL COMMENT '历史PV',
    `total_uv`        int DEFAULT NULL COMMENT '历史UV',
    `total_uip`       int DEFAULT NULL COMMENT '历史UIP',
    `create_time`     datetime                                       DEFAULT NULL COMMENT '创建时间',
    `update_time`     datetime                                       DEFAULT NULL COMMENT '修改时间',
    `del_time`        bigint DEFAULT '0' COMMENT '删除时间戳',
    `del_flag`        tinyint DEFAULT NULL COMMENT '删除标识 0：未删除 1：已删除',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_unique_full-short-url` (`full_short_url`,`del_time`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_link_5`
(
    `id`              bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `domain`          varchar(128)                                   DEFAULT NULL COMMENT '域名',
    `short_uri`       varchar(8) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '短链接',
    `full_short_url`  varchar(128)                                   DEFAULT NULL COMMENT '完整短链接',
    `origin_url`      varchar(1024)                                  DEFAULT NULL COMMENT '原始链接',
    `click_num`       int DEFAULT '0' COMMENT '点击量',
    `gid`             varchar(32)                                    DEFAULT 'default' COMMENT '分组标识',
    `favicon`         varchar(256)                                   DEFAULT NULL COMMENT '网站图标',
    `enable_status`   tinyint DEFAULT NULL COMMENT '启用标识 0：启用 1：未启用',
    `created_type`    tinyint DEFAULT NULL COMMENT '创建类型 0：接口创建 1：控制台创建',
    `valid_date_type` tinyint DEFAULT NULL COMMENT '有效期类型 0：永久有效 1：自定义',
    `valid_date`      datetime                                       DEFAULT NULL COMMENT '有效期',
    `describe`        varchar(1024)                                  DEFAULT NULL COMMENT '描述',
    `total_pv`        int DEFAULT NULL COMMENT '历史PV',
    `total_uv`        int DEFAULT NULL COMMENT '历史UV',
    `total_uip`       int DEFAULT NULL COMMENT '历史UIP',
    `create_time`     datetime                                       DEFAULT NULL COMMENT '创建时间',
    `update_time`     datetime                                       DEFAULT NULL COMMENT '修改时间',
    `del_time`        bigint DEFAULT '0' COMMENT '删除时间戳',
    `del_flag`        tinyint DEFAULT NULL COMMENT '删除标识 0：未删除 1：已删除',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_unique_full-short-url` (`full_short_url`,`del_time`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_link_6`
(
    `id`              bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `domain`          varchar(128)                                   DEFAULT NULL COMMENT '域名',
    `short_uri`       varchar(8) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '短链接',
    `full_short_url`  varchar(128)                                   DEFAULT NULL COMMENT '完整短链接',
    `origin_url`      varchar(1024)                                  DEFAULT NULL COMMENT '原始链接',
    `click_num`       int DEFAULT '0' COMMENT '点击量',
    `gid`             varchar(32)                                    DEFAULT 'default' COMMENT '分组标识',
    `favicon`         varchar(256)                                   DEFAULT NULL COMMENT '网站图标',
    `enable_status`   tinyint DEFAULT NULL COMMENT '启用标识 0：启用 1：未启用',
    `created_type`    tinyint DEFAULT NULL COMMENT '创建类型 0：接口创建 1：控制台创建',
    `valid_date_type` tinyint DEFAULT NULL COMMENT '有效期类型 0：永久有效 1：自定义',
    `valid_date`      datetime                                       DEFAULT NULL COMMENT '有效期',
    `describe`        varchar(1024)                                  DEFAULT NULL COMMENT '描述',
    `total_pv`        int DEFAULT NULL COMMENT '历史PV',
    `total_uv`        int DEFAULT NULL COMMENT '历史UV',
    `total_uip`       int DEFAULT NULL COMMENT '历史UIP',
    `create_time`     datetime                                       DEFAULT NULL COMMENT '创建时间',
    `update_time`     datetime                                       DEFAULT NULL COMMENT '修改时间',
    `del_time`        bigint DEFAULT '0' COMMENT '删除时间戳',
    `del_flag`        tinyint DEFAULT NULL COMMENT '删除标识 0：未删除 1：已删除',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_unique_full-short-url` (`full_short_url`,`del_time`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_link_7`
(
    `id`              bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `domain`          varchar(128)                                   DEFAULT NULL COMMENT '域名',
    `short_uri`       varchar(8) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '短链接',
    `full_short_url`  varchar(128)                                   DEFAULT NULL COMMENT '完整短链接',
    `origin_url`      varchar(1024)                                  DEFAULT NULL COMMENT '原始链接',
    `click_num`       int DEFAULT '0' COMMENT '点击量',
    `gid`             varchar(32)                                    DEFAULT 'default' COMMENT '分组标识',
    `favicon`         varchar(256)                                   DEFAULT NULL COMMENT '网站图标',
    `enable_status`   tinyint DEFAULT NULL COMMENT '启用标识 0：启用 1：未启用',
    `created_type`    tinyint DEFAULT NULL COMMENT '创建类型 0：接口创建 1：控制台创建',
    `valid_date_type` tinyint DEFAULT NULL COMMENT '有效期类型 0：永久有效 1：自定义',
    `valid_date`      datetime                                       DEFAULT NULL COMMENT '有效期',
    `describe`        varchar(1024)                                  DEFAULT NULL COMMENT '描述',
    `total_pv`        int DEFAULT NULL COMMENT '历史PV',
    `total_uv`        int DEFAULT NULL COMMENT '历史UV',
    `total_uip`       int DEFAULT NULL COMMENT '历史UIP',
    `create_time`     datetime                                       DEFAULT NULL COMMENT '创建时间',
    `update_time`     datetime                                       DEFAULT NULL COMMENT '修改时间',
    `del_time`        bigint DEFAULT '0' COMMENT '删除时间戳',
    `del_flag`        tinyint DEFAULT NULL COMMENT '删除标识 0：未删除 1：已删除',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_unique_full-short-url` (`full_short_url`,`del_time`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_link_8`
(
    `id`              bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `domain`          varchar(128)                                   DEFAULT NULL COMMENT '域名',
    `short_uri`       varchar(8) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '短链接',
    `full_short_url`  varchar(128)                                   DEFAULT NULL COMMENT '完整短链接',
    `origin_url`      varchar(1024)                                  DEFAULT NULL COMMENT '原始链接',
    `click_num`       int DEFAULT '0' COMMENT '点击量',
    `gid`             varchar(32)                                    DEFAULT 'default' COMMENT '分组标识',
    `favicon`         varchar(256)                                   DEFAULT NULL COMMENT '网站图标',
    `enable_status`   tinyint DEFAULT NULL COMMENT '启用标识 0：启用 1：未启用',
    `created_type`    tinyint DEFAULT NULL COMMENT '创建类型 0：接口创建 1：控制台创建',
    `valid_date_type` tinyint DEFAULT NULL COMMENT '有效期类型 0：永久有效 1：自定义',
    `valid_date`      datetime                                       DEFAULT NULL COMMENT '有效期',
    `describe`        varchar(1024)                                  DEFAULT NULL COMMENT '描述',
    `total_pv`        int DEFAULT NULL COMMENT '历史PV',
    `total_uv`        int DEFAULT NULL COMMENT '历史UV',
    `total_uip`       int DEFAULT NULL COMMENT '历史UIP',
    `create_time`     datetime                                       DEFAULT NULL COMMENT '创建时间',
    `update_time`     datetime                                       DEFAULT NULL COMMENT '修改时间',
    `del_time`        bigint DEFAULT '0' COMMENT '删除时间戳',
    `del_flag`        tinyint DEFAULT NULL COMMENT '删除标识 0：未删除 1：已删除',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_unique_full-short-url` (`full_short_url`,`del_time`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_link_9`
(
    `id`              bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `domain`          varchar(128)                                   DEFAULT NULL COMMENT '域名',
    `short_uri`       varchar(8) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '短链接',
    `full_short_url`  varchar(128)                                   DEFAULT NULL COMMENT '完整短链接',
    `origin_url`      varchar(1024)                                  DEFAULT NULL COMMENT '原始链接',
    `click_num`       int DEFAULT '0' COMMENT '点击量',
    `gid`             varchar(32)                                    DEFAULT 'default' COMMENT '分组标识',
    `favicon`         varchar(256)                                   DEFAULT NULL COMMENT '网站图标',
    `enable_status`   tinyint DEFAULT NULL COMMENT '启用标识 0：启用 1：未启用',
    `created_type`    tinyint DEFAULT NULL COMMENT '创建类型 0：接口创建 1：控制台创建',
    `valid_date_type` tinyint DEFAULT NULL COMMENT '有效期类型 0：永久有效 1：自定义',
    `valid_date`      datetime                                       DEFAULT NULL COMMENT '有效期',
    `describe`        varchar(1024)                                  DEFAULT NULL COMMENT '描述',
    `total_pv`        int DEFAULT NULL COMMENT '历史PV',
    `total_uv`        int DEFAULT NULL COMMENT '历史UV',
    `total_uip`       int DEFAULT NULL COMMENT '历史UIP',
    `create_time`     datetime                                       DEFAULT NULL COMMENT '创建时间',
    `update_time`     datetime                                       DEFAULT NULL COMMENT '修改时间',
    `del_time`        bigint DEFAULT '0' COMMENT '删除时间戳',
    `del_flag`        tinyint DEFAULT NULL COMMENT '删除标识 0：未删除 1：已删除',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_unique_full-short-url` (`full_short_url`,`del_time`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_link_goto_0`
(
    `id`             bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `gid`            varchar(32)  DEFAULT 'default' COMMENT '分组标识',
    `full_short_url` varchar(128) DEFAULT NULL COMMENT '完整短链接',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_full_short_url` (`full_short_url`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_link_goto_1`
(
    `id`             bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `gid`            varchar(32)  DEFAULT 'default' COMMENT '分组标识',
    `full_short_url` varchar(128) DEFAULT NULL COMMENT '完整短链接',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_full_short_url` (`full_short_url`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_link_goto_10`
(
    `id`             bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `gid`            varchar(32)  DEFAULT 'default' COMMENT '分组标识',
    `full_short_url` varchar(128) DEFAULT NULL COMMENT '完整短链接',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_full_short_url` (`full_short_url`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_link_goto_11`
(
    `id`             bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `gid`            varchar(32)  DEFAULT 'default' COMMENT '分组标识',
    `full_short_url` varchar(128) DEFAULT NULL COMMENT '完整短链接',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_full_short_url` (`full_short_url`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_link_goto_12`
(
    `id`             bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `gid`            varchar(32)  DEFAULT 'default' COMMENT '分组标识',
    `full_short_url` varchar(128) DEFAULT NULL COMMENT '完整短链接',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_full_short_url` (`full_short_url`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_link_goto_13`
(
    `id`             bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `gid`            varchar(32)  DEFAULT 'default' COMMENT '分组标识',
    `full_short_url` varchar(128) DEFAULT NULL COMMENT '完整短链接',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_full_short_url` (`full_short_url`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_link_goto_14`
(
    `id`             bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `gid`            varchar(32)  DEFAULT 'default' COMMENT '分组标识',
    `full_short_url` varchar(128) DEFAULT NULL COMMENT '完整短链接',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_full_short_url` (`full_short_url`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_link_goto_15`
(
    `id`             bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `gid`            varchar(32)  DEFAULT 'default' COMMENT '分组标识',
    `full_short_url` varchar(128) DEFAULT NULL COMMENT '完整短链接',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_full_short_url` (`full_short_url`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_link_goto_2`
(
    `id`             bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `gid`            varchar(32)  DEFAULT 'default' COMMENT '分组标识',
    `full_short_url` varchar(128) DEFAULT NULL COMMENT '完整短链接',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_full_short_url` (`full_short_url`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_link_goto_3`
(
    `id`             bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `gid`            varchar(32)  DEFAULT 'default' COMMENT '分组标识',
    `full_short_url` varchar(128) DEFAULT NULL COMMENT '完整短链接',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_full_short_url` (`full_short_url`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_link_goto_4`
(
    `id`             bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `gid`            varchar(32)  DEFAULT 'default' COMMENT '分组标识',
    `full_short_url` varchar(128) DEFAULT NULL COMMENT '完整短链接',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_full_short_url` (`full_short_url`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_link_goto_5`
(
    `id`             bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `gid`            varchar(32)  DEFAULT 'default' COMMENT '分组标识',
    `full_short_url` varchar(128) DEFAULT NULL COMMENT '完整短链接',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_full_short_url` (`full_short_url`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_link_goto_6`
(
    `id`             bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `gid`            varchar(32)  DEFAULT 'default' COMMENT '分组标识',
    `full_short_url` varchar(128) DEFAULT NULL COMMENT '完整短链接',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_full_short_url` (`full_short_url`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_link_goto_7`
(
    `id`             bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `gid`            varchar(32)  DEFAULT 'default' COMMENT '分组标识',
    `full_short_url` varchar(128) DEFAULT NULL COMMENT '完整短链接',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_full_short_url` (`full_short_url`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_link_goto_8`
(
    `id`             bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `gid`            varchar(32)  DEFAULT 'default' COMMENT '分组标识',
    `full_short_url` varchar(128) DEFAULT NULL COMMENT '完整短链接',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_full_short_url` (`full_short_url`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `t_link_goto_9`
(
    `id`             bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `gid`            varchar(32)  DEFAULT 'default' COMMENT '分组标识',
    `full_short_url` varchar(128) DEFAULT NULL COMMENT '完整短链接',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_full_short_url` (`full_short_url`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
