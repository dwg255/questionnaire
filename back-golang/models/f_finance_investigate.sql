/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : localhost:3306
 Source Schema         : financial

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 10/07/2020 11:13:46
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for f_finance_investigate
-- ----------------------------
DROP TABLE IF EXISTS `f_finance_investigate`;
CREATE TABLE `f_finance_investigate`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `register_time` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '注册日期',
  `company_name` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '公司名称',
  `main_business` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '主营业务',
  `registered_capital` decimal(19, 9) NOT NULL COMMENT '注册资本',
  `paid_capital` decimal(19, 9) NOT NULL COMMENT '实收资本',
  `tax_level` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '纳税等级',
  `tax_amount` decimal(19, 9) NOT NULL COMMENT '上一年总资产（万元）',
  `assets` decimal(19, 9) NOT NULL COMMENT '上一年度总资产（单位/万元）',
  `liabilities` decimal(19, 9) NOT NULL COMMENT '上一年度总负债（单位/万元）',
  `loan` decimal(19, 9) NOT NULL COMMENT '银行贷款（单位/万元）',
  `income` decimal(19, 9) NOT NULL COMMENT '上一年度营业收入（单位/万元）',
  `profit` decimal(19, 9) NOT NULL COMMENT '上一年度净利润（单位/万元）',
  `account` decimal(19, 9) NOT NULL COMMENT '拟融资额度（单位/万元）',
  `purpose` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '资金用途',
  `guarantee` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '担保方式',
  `contact` varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '企业联系人',
  `phone` char(16) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '联系电话',
  `created` datetime(0) NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 18 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
