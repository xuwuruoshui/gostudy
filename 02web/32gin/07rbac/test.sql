/*
 Navicat Premium Data Transfer

 Source Server         : test
 Source Server Type    : MySQL
 Source Server Version : 80028
 Source Host           : 120.78.159.42:3306
 Source Schema         : test

 Target Server Type    : MySQL
 Target Server Version : 80028
 File Encoding         : 65001

 Date: 12/04/2022 22:28:08
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for perm
-- ----------------------------
DROP TABLE IF EXISTS `perm`;
CREATE TABLE `perm`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `path` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `create_at` timestamp(0) NULL DEFAULT NULL,
  `update_at` timestamp(0) NULL DEFAULT NULL,
  `delete_at` timestamp(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of perm
-- ----------------------------
INSERT INTO `perm` VALUES (1, '/test', '2022-04-12 21:01:34', '2022-04-12 21:01:37', NULL);
INSERT INTO `perm` VALUES (2, '/test2', '2022-04-12 21:01:46', '2022-04-12 21:01:47', NULL);

-- ----------------------------
-- Table structure for perm_role
-- ----------------------------
DROP TABLE IF EXISTS `perm_role`;
CREATE TABLE `perm_role`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `perm_id` int(0) NOT NULL,
  `role_id` int(0) NOT NULL,
  `create_at` timestamp(0) NOT NULL,
  `update_at` timestamp(0) NOT NULL,
  `delete_at` timestamp(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of perm_role
-- ----------------------------
INSERT INTO `perm_role` VALUES (1, 1, 1, '2022-04-12 21:02:14', '2022-04-12 21:02:16', NULL);
INSERT INTO `perm_role` VALUES (2, 1, 2, '2022-04-12 21:05:23', '2022-04-12 21:05:25', NULL);
INSERT INTO `perm_role` VALUES (3, 2, 2, '2022-04-12 21:06:50', '2022-04-12 21:06:52', NULL);

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `create_at` timestamp(0) NULL DEFAULT NULL,
  `update_at` timestamp(0) NULL DEFAULT NULL,
  `delete_at` timestamp(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `name_unique`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of role
-- ----------------------------
INSERT INTO `role` VALUES (1, 'normal', '2022-04-12 21:01:15', '2022-04-12 21:03:36', NULL);
INSERT INTO `role` VALUES (2, 'admin', '2022-04-12 21:03:31', '2022-04-12 21:04:22', NULL);

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `username` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `password` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `create_at` timestamp(0) NULL DEFAULT NULL,
  `update_at` timestamp(0) NULL DEFAULT NULL,
  `delete_at` timestamp(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `username`(`username`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 'root', 'root', '2022-04-12 20:58:50', '2022-04-12 20:58:52', NULL);
INSERT INTO `user` VALUES (2, 'aaa', 'aaa', '2022-04-12 21:03:06', '2022-04-12 21:03:08', NULL);

-- ----------------------------
-- Table structure for user_role
-- ----------------------------
DROP TABLE IF EXISTS `user_role`;
CREATE TABLE `user_role`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `user_id` int(0) NOT NULL,
  `role_id` int(0) NOT NULL,
  `create_at` timestamp(0) NOT NULL,
  `update_at` timestamp(0) NOT NULL,
  `delete_at` timestamp(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_role
-- ----------------------------
INSERT INTO `user_role` VALUES (1, 1, 1, '2022-04-12 21:02:04', '2022-04-12 21:02:06', NULL);
INSERT INTO `user_role` VALUES (2, 1, 2, '2022-04-12 21:07:36', '2022-04-12 21:07:37', NULL);
INSERT INTO `user_role` VALUES (3, 2, 1, '2022-04-12 21:08:27', '2022-04-12 21:08:30', NULL);

-- ----------------------------
-- Procedure structure for insert_tb_item
-- ----------------------------
DROP PROCEDURE IF EXISTS `insert_tb_item`;
delimiter ;;
CREATE PROCEDURE `insert_tb_item`(num int)
begin
while num <= 10000000 do
insert into tb_item values(num,concat('',num,''),round(RAND() * 100000,2),FLOOR(RAND() * 100000),FLOOR(RAND() * 10),'1','5435343235','2019-04-20 22:37:15','2019-04-20 22:37:15');
set num = num + 1;
end while;
end
;;
delimiter ;

-- ----------------------------
-- Procedure structure for test_add
-- ----------------------------
DROP PROCEDURE IF EXISTS `test_add`;
delimiter ;;
CREATE PROCEDURE `test_add`(n int)
BEGIN
DECLARE num int DEFAULT(0);
WHILE num<=n DO
SET num = num+1;
insert into test(`id`,`name`,`email`) VALUES(num,num,num);
END WHILE;
END
;;
delimiter ;

SET FOREIGN_KEY_CHECKS = 1;
