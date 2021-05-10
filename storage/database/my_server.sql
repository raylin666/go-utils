CREATE TABLE `srv_jwt_secret`
(
    `id`         int(11) unsigned NOT NULL AUTO_INCREMENT,
    `secret`     varchar(80) NOT NULL COMMENT '颁布标识 Secret',
    `created_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY (`secret`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='JWT 认证颁布标识表';

CREATE TABLE `srv_jwt_users`
(
    `id`         int(11) unsigned NOT NULL AUTO_INCREMENT,
    `secret`     varchar(80)  NOT NULL COMMENT '颁布标识 Secret',
    `user_id`    varchar(50)  NOT NULL COMMENT '用户标识 ID',
    `token`      varchar(255) NOT NULL COMMENT 'TOKEN 认证',
    `ttl`        int(10) unsigned NOT NULL DEFAULT '0' COMMENT '有效时长',
    `expired_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '过期时间',
    `refresh_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '刷新时间',
    `created_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    `updated_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
    `deleted_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY (`user_id`, `secret`) USING BTREE,
    UNIQUE KEY (`token`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='JWT 认证用户表';