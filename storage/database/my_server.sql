CREATE TABLE `srv_jwt_secret`
(
    `id`         int(11) unsigned NOT NULL AUTO_INCREMENT,
    `app`        varchar(36) NOT NULL COMMENT '使用应用名称',
    `key`        varchar(36) NOT NULL COMMENT '颁布标识 Key',
    `secret`     varchar(80) NOT NULL COMMENT '颁布标识 Secret',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `expired_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '过期时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY key_secret (`key`, `secret`) USING BTREE
) ENGINE=InnoDB COMMENT='JWT 认证颁布标识表';

CREATE TABLE `srv_jwt_users`
(
    `id`         int(11) unsigned NOT NULL AUTO_INCREMENT,
    `secret`     varchar(80)  NOT NULL COMMENT '颁布标识 Secret',
    `user_id`    varchar(50)  NOT NULL COMMENT '用户标识 ID',
    `token`      varchar(255) NOT NULL COMMENT 'TOKEN 认证',
    `ttl`        int(10) unsigned NOT NULL DEFAULT '0' COMMENT '有效时长',
    `expired_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '过期时间',
    `refresh_at` datetime DEFAULT NULL COMMENT '刷新时间',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY user_secret (`user_id`, `secret`) USING BTREE,
    UNIQUE KEY (`token`) USING BTREE
) ENGINE=InnoDB COMMENT='JWT 认证用户表';