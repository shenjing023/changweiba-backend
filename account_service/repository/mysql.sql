/*
    用户表
 */
SET NAMES utf8mb4;

DROP TABLE IF EXISTS `cw_user`;
CREATE TABLE `cw_user`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '名称',
  `password` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '头像',
  `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态',
  `score` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '分数',
  `role` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户的角色',
  `banned_reason` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '被封原因',
  `create_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `last_update` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '最后更新时间',
  `ip` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '该账号注册时的ip地址',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_name`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10000 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户表' ROW_FORMAT = Compact;

/*
  随机头像表
*/
DROP TABLE IF EXISTS `cw_avatar`;
CREATE TABLE `cw_avatar`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '头像url',
  `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态,(0正常1无效)',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10000 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '默认头像表' ROW_FORMAT = Compact;