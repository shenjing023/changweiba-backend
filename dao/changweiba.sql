/**
    mysql changweiba
 */
 
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

/*
    用户表
 */
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
    `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户id',
    `name` varchar(20) NOT NULL UNIQUE COMMENT '名称',
    `password` varchar(32) NOT NULL COMMENT '密码',
    `avatar` varchar(255) NOT NULL COMMENT '头像',
    `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态',
    `score` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '分数',
    `role` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户的角色',
    `banned_reason` varchar(255) NOT NULL DEFAULT 0 COMMENT '被封原因',
    `create_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
    `last_update` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '最后更新时间',
    `ip` int(10) UNSIGNED NOT NULL COMMENT '该账号注册时的ip地址',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10000 CHARACTER SET =utf8 COLLATE = utf8_general_ci COMMENT = '用户表' ROW_FORMAT =
Compact;

/*
    默认头像表
 */
DROP TABLE IF EXISTS `avatar`;
CREATE TABLE `avatar` (
    `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `url` varchar(255) NOT NULL COMMENT '头像url',
    `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态,(0正常1无效)',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT = 0 CHARACTER SET =utf8 COLLATE = utf8_general_ci COMMENT = '默认头像表' ROW_FORMAT = Compact;

/*
    帖子表
 */
DROP TABLE IF EXISTS `post`;
CREATE TABLE `post` (
    `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `user_id` int(11) UNSIGNED NOT NULL COMMENT '用户id',
    `topic` varchar(255) NOT NULL COMMENT '帖子标题',
    `create_time` int(11) UNSIGNED NOT NULL COMMENT '创建时间',
    `last_update` int(11) UNSIGNED NOT NULL COMMENT '最后更新时间',
    `reply_num` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '帖子评论+回复总数',
    `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态,(0正常1删除)',
    PRIMARY KEY (`id`) USING BTREE,
    KEY `ids_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT = 0 CHARACTER SET =utf8 COLLATE = utf8_general_ci COMMENT = '帖子表' ROW_FORMAT = Compact;

/*
    评论表
 */
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
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
) ENGINE=InnoDB AUTO_INCREMENT = 0 CHARACTER SET =utf8 COLLATE = utf8_general_ci COMMENT = '评论表' ROW_FORMAT = Compact;

/*
    回复表
 */
DROP TABLE IF EXISTS `reply`;
CREATE TABLE `reply` (
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
) ENGINE=InnoDB AUTO_INCREMENT = 0 CHARACTER SET =utf8 COLLATE = utf8_general_ci COMMENT = '回复表' ROW_FORMAT = Compact;


SET FOREIGN_KEY_CHECKS = 1;