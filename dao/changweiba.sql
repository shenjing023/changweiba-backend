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
    `password` varchar(20) NOT NULL COMMENT '密码',
    `avatar` varchar NOT NULL COMMENT '头像',
    `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态',
    `score` int(11) NOT NULL DEFAULT 0 COMMENT '分数',
    `banned_reason` tinyint(1) NOT NULL DEFAULT 0 COMMENT '被封原因',
    `create_time` int(11) NOT NULL DEFAULT 0 COMMENT '创建时间',
    `last_update` int(11) NOT NULL DEFAULT 0 COMMENT '最后更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET =utf8 COLLATE = utf8_general_ci COMMENT = '用户表' ROW_FORMAT = Compact;

SET FOREIGN_KEY_CHECKS = 1;