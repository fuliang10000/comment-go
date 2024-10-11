CREATE TABLE `comment_theme` (
                                 `id` BIGINT UNSIGNED NOT NULL auto_increment COMMENT '主键ID',
                                 `userId` BIGINT NOT NULL DEFAULT 0 COMMENT '用户ID',
                                 `entityId` BIGINT NOT NULL DEFAULT 0 COMMENT '主题ID',
                                 `type` TINYINT UNSIGNED NOT NULL DEFAULT 1 COMMENT '主题类型 1:文章，2:视频',
                                 `count` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '主题跟评论数量',
                                 `createdAt` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
                                 `updatedAt` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改时间',
                                 `deletedAt` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
                                 PRIMARY KEY ( `id` ),
                                 KEY `idx_userId` ( `userId` ),
                                 KEY `idx_entityId_type` ( `entityId`, `type` )
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '评论主题表';

CREATE TABLE `comment_index` (
                                 `id` BIGINT UNSIGNED NOT NULL auto_increment COMMENT '主键ID',
                                 `userId` BIGINT NOT NULL DEFAULT 0 COMMENT '用户ID',
                                 `entityId` BIGINT NOT NULL DEFAULT 0 COMMENT '主题ID',
                                 `type` TINYINT UNSIGNED NOT NULL DEFAULT 1 COMMENT '主题类型 1:文章，2:视频',
                                 `root` BIGINT NOT NULL DEFAULT 0 COMMENT '跟评论ID，0表示跟评论',
                                 `parentId` BIGINT NOT NULL DEFAULT 0 COMMENT '父级评论ID',
                                 `count` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '主题跟评论数量',
                                 `createdAt` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
                                 `updatedAt` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改时间',
                                 `deletedAt` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
                                 PRIMARY KEY ( `id` ),
                                 KEY `idx_userId` ( `userId` ),
                                 KEY `idx_entityId_type` ( `entityId`, `type` )
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '评论索引表';

CREATE TABLE `comment_content` (
                                   `id` BIGINT UNSIGNED NOT NULL auto_increment COMMENT '主键ID',
                                   `indexId` BIGINT NOT NULL DEFAULT 0 COMMENT '索引表ID',
                                   `content` TEXT DEFAULT NULL COMMENT '评论内容',
                                   `createdAt` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
                                   `updatedAt` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改时间',
                                   PRIMARY KEY ( `id` ),
                                   KEY `idx_indexId` ( `indexId` )
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '评论内容表';