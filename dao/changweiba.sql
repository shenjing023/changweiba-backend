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
    `id` int(11) UNSIGNED NOT NULL DEFAULT 10000 AUTO_INCREMENT COMMENT '用户id',
    `name` varchar(20) NOT NULL UNIQUE COMMENT '名称',
    `password` varchar(32) NOT NULL COMMENT '密码',
    `avatar` varchar(255) NOT NULL COMMENT '头像',
    `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态',
    `score` int(11) NOT NULL DEFAULT 0 COMMENT '分数',
    `role` tinyint(1) NOT NULL DEFAULT 0 COMMENT '用户的角色',
    `banned_reason` varchar(255) NOT NULL DEFAULT 0 COMMENT '被封原因',
    `create_time` int(11) NOT NULL DEFAULT 0 COMMENT '创建时间',
    `last_update` int(11) NOT NULL DEFAULT 0 COMMENT '最后更新时间',
    `ip` int(10) UNSIGNED NOT NULL COMMENT '该账号注册时的ip地址',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET =utf8 COLLATE = utf8_general_ci COMMENT = '用户表' ROW_FORMAT = Compact;

/*
    默认头像表
 */
DROP TABLE IF EXISTS `avatar`;
CREATE TABLE `avatar` (
    `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `url` varchar(255) NOT NULL COMMENT '头像url',
    `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态,(0正常1无效)',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT = 6 CHARACTER SET =utf8 COLLATE = utf8_general_ci COMMENT = '默认头像表' ROW_FORMAT = Compact;


SET FOREIGN_KEY_CHECKS = 1;