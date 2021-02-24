SET NAMES utf8mb4;

/*
    帖子表
 */
DROP TABLE IF EXISTS `cw_post`;
CREATE TABLE `cw_post` (
    `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `user_id` int(11) UNSIGNED NOT NULL COMMENT '用户id',
    `topic` varchar(255) NOT NULL COMMENT '帖子标题',
    `create_time` int(11) UNSIGNED NOT NULL COMMENT '创建时间',
    `last_update` int(11) UNSIGNED NOT NULL COMMENT '最后更新时间',
    `reply_num` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '帖子评论+回复总数',
    `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态,(0正常1删除)',
    PRIMARY KEY (`id`) USING BTREE,
    KEY `ids_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT = 10000 CHARACTER SET =utf8 COLLATE = utf8_general_ci COMMENT = '帖子表' ROW_FORMAT = Compact;

/*
    评论表
 */
DROP TABLE IF EXISTS `cw_comment`;
CREATE TABLE `cw_comment` (
    `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `user_id` int(11) UNSIGNED NOT NULL COMMENT '用户id',
    `post_id` int(11) UNSIGNED NOT NULL COMMENT '帖子id',
    `content` varchar(1024) NOT NULL COMMENT '评论内容',
    `create_time` int(11) NOT NULL COMMENT '创建时间',
    `floor` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '第几楼',
    `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态,(0正常1删除)',
    `reply_num` int(11) UNSIGNED NOT NULL COMMENT '总回复数',
    PRIMARY KEY (`id`) USING BTREE,
    KEY `ids_user_id` (`user_id`),
    KEY `ids_post_id` (`post_id`)
) ENGINE=InnoDB AUTO_INCREMENT = 10000 CHARACTER SET =utf8 COLLATE = utf8_general_ci COMMENT = '评论表' ROW_FORMAT = Compact;

/*
    回复表
 */
DROP TABLE IF EXISTS `cw_reply`;
CREATE TABLE `cw_reply` (
    `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `user_id` int(11) UNSIGNED NOT NULL COMMENT '用户id',
    `post_id` int(11) UNSIGNED NOT NULL COMMENT '帖子id',
    `comment_id` int(11) UNSIGNED NOT NULL COMMENT '评论id',
    `content` varchar(1024) NOT NULL COMMENT '评论内容',
    `parent_id` int(11) UNSIGNED NOT NULL COMMENT '回复哪个的id',
    `create_time` int(11) UNSIGNED NOT NULL COMMENT '创建时间',
    `floor` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '楼中楼第几楼',
    `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态,(0正常1删除)',
    PRIMARY KEY (`id`) USING BTREE,
    KEY `ids_user_id` (`user_id`),
    KEY `ids_post_id` (`post_id`),
    KEY `ids_comment_id` (`comment_id`)
) ENGINE=InnoDB AUTO_INCREMENT = 10000 CHARACTER SET =utf8 COLLATE = utf8_general_ci COMMENT = '回复表' ROW_FORMAT = Compact;