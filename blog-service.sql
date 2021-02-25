/*
Navicat MySQL Data Transfer

Source Server         : local
Source Server Version : 50726
Source Host           : localhost:3306
Source Database       : blog-service

Target Server Type    : MYSQL
Target Server Version : 50726
File Encoding         : 65001

Date: 2021-02-25 20:42:46
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for blog_article
-- ----------------------------
DROP TABLE IF EXISTS `blog_article`;
CREATE TABLE `blog_article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `category` int(10) unsigned DEFAULT '0' COMMENT '分类id',
  `title` varchar(100) DEFAULT '' COMMENT '文章标题',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
  `cover_image_url` varchar(255) DEFAULT '' COMMENT '封面图片地址',
  `content` longtext COMMENT '文章内容',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `is_top` tinyint(1) DEFAULT '0' COMMENT '是否置顶',
  `is_hot` tinyint(1) DEFAULT '0' COMMENT '是否热门',
  `summary` varchar(255) DEFAULT '' COMMENT '概述简介',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0为未删除、1为已删除',
  `created_at` int(10) unsigned DEFAULT '0' COMMENT '新建时间',
  `modified_at` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `deleted_at` int(10) unsigned DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=44 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='文章管理';

-- ----------------------------
-- Records of blog_article
-- ----------------------------
INSERT INTO `blog_article` VALUES ('1', '1', 'JAVA', '1', 'https://s1.ax1x.com/2020/05/14/YDhBDO.jpg', '1', '1', '1', '1', '1', '1111111', '0', '1612577991', '0', null);
INSERT INTO `blog_article` VALUES ('2', '1', 'JAVA', '1', 'https://s1.ax1x.com/2020/05/14/YDhoVg.jpg', '1', '1', '1', '1', '1', '1111111', '0', '1612578450', '0', null);
INSERT INTO `blog_article` VALUES ('3', '1', 'JAVA', '1', 'https://s1.ax1x.com/2020/05/14/YD4FR1.jpg', '1', '1', '1', '1', '1', '1111111', '0', '1612578476', '0', null);
INSERT INTO `blog_article` VALUES ('4', '1', 'JAVA', '1', 'https://s1.ax1x.com/2020/05/14/YDhagx.jpg', '1', '1', '1', '1', '1', '1111111', '0', '1612578505', '0', null);
INSERT INTO `blog_article` VALUES ('5', '1', 'JAVA', '1', 'https://s1.ax1x.com/2020/05/14/YDhagx.jpg', '1', '1', '1', '1', '1', '1111111', '0', '1612578614', '0', null);
INSERT INTO `blog_article` VALUES ('6', '1', 'JAVA', '1', 'https://s1.ax1x.com/2020/05/14/YDhagx.jpg', '1', '1', '1', '1', '1', '1111111', '0', '1612578651', '0', null);
INSERT INTO `blog_article` VALUES ('7', '1', 'JAVA', '1', 'https://s1.ax1x.com/2020/05/14/YDhagx.jpg', '1', '1', '1', '1', '1', '1111111', '0', '1612579152', '0', null);
INSERT INTO `blog_article` VALUES ('8', '1', 'JAVA', '1', 'https://s1.ax1x.com/2020/05/14/YDhagx.jpg', '1', '1', '1', '1', '1', '1111111', '0', '1612579152', '0', null);
INSERT INTO `blog_article` VALUES ('9', '1', 'JAVA', '1', 'https://s1.ax1x.com/2020/05/14/YDhagx.jpg', '1', '1', '1', '1', '1', '1111111', '0', '1612579301', '0', null);
INSERT INTO `blog_article` VALUES ('10', '1', 'JAVA', '1', 'https://s1.ax1x.com/2020/05/14/YDhagx.jpg', '1', '1', '1', '1', '1', '1111111', '0', '1612579369', '0', null);
INSERT INTO `blog_article` VALUES ('11', '1', 'JAVA', '1', 'https://s1.ax1x.com/2020/05/14/YDhagx.jpg', '1', '1', '1', '1', '1', '1111111', '0', '1612579594', '0', null);
INSERT INTO `blog_article` VALUES ('12', '1', 'JAVA', '1', 'https://s1.ax1x.com/2020/05/14/YDhagx.jpg', '1', '1', '1', '1', '1', '1111111', '0', '1612580886', '0', null);
INSERT INTO `blog_article` VALUES ('13', '1', 'JAVA', '1', 'https://s1.ax1x.com/2020/05/14/YDhagx.jpg', '1', '1', '1', '1', '1', '1111111', '0', '1612580933', '0', null);
INSERT INTO `blog_article` VALUES ('14', '1', 'JAVA', '1', 'https://s1.ax1x.com/2020/05/14/YDhagx.jpg', '1', '1', '1', '1', '1', '1111111', '0', '1612580974', '0', null);
INSERT INTO `blog_article` VALUES ('15', '1', 'JAVA', '1', 'https://s1.ax1x.com/2020/05/14/YDhagx.jpg', '1', '1', '1', '1', '1', '1111111', '0', '1612581123', '0', null);
INSERT INTO `blog_article` VALUES ('16', '1', 'JAVA', '1', 'https://s1.ax1x.com/2020/05/14/YDhagx.jpg', '1', '1', '1', '1', '1', '1111111', '0', '1612581217', '0', null);
INSERT INTO `blog_article` VALUES ('17', '1', 'JAVA', '1', 'https://s1.ax1x.com/2020/05/14/YDhagx.jpg', '1', '1', '1', '1', '1', '1111111', '0', '1612581371', '0', null);
INSERT INTO `blog_article` VALUES ('18', '1', 'JAVA', '1', 'https://s1.ax1x.com/2020/05/14/YDhagx.jpg', '1', '1', '1', '1', '1', '1111111', '0', '1612581496', '0', null);
INSERT INTO `blog_article` VALUES ('19', '1', 'JAVA', '1', 'https://s1.ax1x.com/2020/05/14/YDhU81.jpg', '1', '1', '1', '1', '1', '1111111', '0', '1612581515', '0', null);
INSERT INTO `blog_article` VALUES ('20', '1', 'JAVA', '1', 'https://s1.ax1x.com/2020/05/14/YDhBDO.jpg', '1', '1', '1', '1', '1', '1111111', '0', '1612581526', '0', null);
INSERT INTO `blog_article` VALUES ('21', '1', 'JAVA', '1', 'https://s1.ax1x.com/2020/05/14/YDhoVg.jpg', '1', '1', '1', '1', '1', '1111111', '0', '1612581555', '0', null);
INSERT INTO `blog_article` VALUES ('22', '1', 'JAVA', '1', 'https://s1.ax1x.com/2020/05/14/YD4FR1.jpg', '1', '1', '1', '1', '1', '1111111', '0', '1612581563', '0', null);
INSERT INTO `blog_article` VALUES ('23', '2', '300000', '1', 'https://s1.ax1x.com/2020/05/14/YDhagx.jpg', '1', '1', '1', '1', '1', '1111111', '0', '1612581841', '0', null);
INSERT INTO `blog_article` VALUES ('24', '2', '300000', '1', 'https://s1.ax1x.com/2020/05/14/YDhU81.jpg', '1', '1', '1', '1', '1', '1111111', '0', '1612581897', '0', null);
INSERT INTO `blog_article` VALUES ('25', '2', '300000', '1', 'https://s1.ax1x.com/2020/05/14/YDhBDO.jpg', '1', '1', '1', '1', '1', '1111111', '0', '1612582056', '0', null);
INSERT INTO `blog_article` VALUES ('26', '2', '300000', '1', 'https://s1.ax1x.com/2020/05/14/YDhoVg.jpg', '1', '1', '1', '1', '1', '1111111', '0', '1612582235', '0', null);
INSERT INTO `blog_article` VALUES ('27', '2', '300000', '1', 'https://s1.ax1x.com/2020/05/14/YD4FR1.jpg', '1', '1', '1', '1', '1', '1111111', '0', '1612582251', '0', null);
INSERT INTO `blog_article` VALUES ('28', '2', '300000', '1', 'https://s1.ax1x.com/2020/05/14/YDhagx.jpg', '1', '1', '1', '1', '1', '1111111', '0', '1612582282', '0', null);
INSERT INTO `blog_article` VALUES ('29', '2', '300000', '1', 'https://s1.ax1x.com/2020/05/14/YDhU81.jpg', '1', '1', '1', '1', '1', '1111111', '0', '1612582590', '0', null);
INSERT INTO `blog_article` VALUES ('30', '2', '300000', '1', 'https://s1.ax1x.com/2020/05/14/YDhBDO.jpg', '1', '1', '1', '1', '1', '1111111', '0', '1612582610', '0', null);
INSERT INTO `blog_article` VALUES ('31', '2', '300000', '1', 'https://s1.ax1x.com/2020/05/14/YDhoVg.jpg', '1', '1', '1', '1', '1', '1111111', '0', '1612582621', '0', null);
INSERT INTO `blog_article` VALUES ('32', '2', '300000', '1', 'https://s1.ax1x.com/2020/05/14/YD4FR1.jpg', '1', '1', '1', '1', '1', '1111111', '0', '1612582636', '0', null);
INSERT INTO `blog_article` VALUES ('33', '2', '300000', '1', 'https://s1.ax1x.com/2020/05/14/YDhagx.jpg', '1', '1', '1', '1', '1', '1111111', '0', '1612582656', '0', null);
INSERT INTO `blog_article` VALUES ('34', '2', '300000', '1', 'https://s1.ax1x.com/2020/05/14/YDhU81.jpg', '1', '1', '1', '1', '1', '1111111', '0', '1612582664', '0', null);
INSERT INTO `blog_article` VALUES ('35', '2', '300000', '1', 'https://s1.ax1x.com/2020/05/14/YDhBDO.jpg', '1', '1', '1', '1', '1', '1111111', '0', '1612582674', '0', null);
INSERT INTO `blog_article` VALUES ('36', '3', 'MYSQL', '1', 'https://s1.ax1x.com/2020/05/14/YDhoVg.jpg', '1', '1', '1', '1', '1', '1111111', '0', '1612582739', '0', null);
INSERT INTO `blog_article` VALUES ('37', '3', 'MYSQL', '1', 'https://s1.ax1x.com/2020/05/14/YD4FR1.jpg', '1', '1', '1', '1', '1', '1111111', '0', '1612582773', '0', null);
INSERT INTO `blog_article` VALUES ('38', '3', 'MYSQL', '1', 'https://s1.ax1x.com/2020/05/14/YDhagx.jpg', '1', '1', '1', '1', '1', '1111111', '0', '1612582848', '0', null);
INSERT INTO `blog_article` VALUES ('39', '3', 'MYSQL', '1', 'https://s1.ax1x.com/2020/05/14/YDhU81.jpg', '1', '1', '1', '1', '1', '1111111', '0', '1612582861', '0', null);
INSERT INTO `blog_article` VALUES ('40', '2', '300000', '1', 'https://s1.ax1x.com/2020/05/14/YDhBDO.jpg', '1', '1', '1', '1', '1', '1111111', '0', '1612583130', '0', null);
INSERT INTO `blog_article` VALUES ('41', '2', '300000', '1', 'https://s1.ax1x.com/2020/05/14/YDhoVg.jpg', '1', '1', '1', '1', '1', '1111111', '0', '1612583141', '0', null);
INSERT INTO `blog_article` VALUES ('42', '2', '1', '1', 'https://s1.ax1x.com/2020/05/14/YDhoVg.jpg', '1', '1', '1', '1', '1', '1111111', '0', '1612583465', '0', null);
INSERT INTO `blog_article` VALUES ('43', '2', '1', '1', 'https://s1.ax1x.com/2020/05/14/YDhoVg.jpg', '1', '1', '1', '1', '1', '1111111', '0', '1612592474', '0', null);

-- ----------------------------
-- Table structure for blog_article_tag
-- ----------------------------
DROP TABLE IF EXISTS `blog_article_tag`;
CREATE TABLE `blog_article_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `article_id` int(11) NOT NULL COMMENT '文章ID',
  `tag_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '标签ID',
  `created_at` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_at` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_at` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0为未删除、1为已删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='文章标签关联';

-- ----------------------------
-- Records of blog_article_tag
-- ----------------------------

-- ----------------------------
-- Table structure for blog_auth
-- ----------------------------
DROP TABLE IF EXISTS `blog_auth`;
CREATE TABLE `blog_auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `app_key` varchar(20) DEFAULT '' COMMENT 'Key',
  `app_secret` varchar(50) DEFAULT '' COMMENT 'Secret',
  `created_at` int(10) unsigned DEFAULT '0' COMMENT '新建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_at` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_at` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0为未删除、1为已删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='认证管理';

-- ----------------------------
-- Records of blog_auth
-- ----------------------------

-- ----------------------------
-- Table structure for blog_category
-- ----------------------------
DROP TABLE IF EXISTS `blog_category`;
CREATE TABLE `blog_category` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(100) DEFAULT '' COMMENT '分类名称',
  `created_at` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_at` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_at` int(10) DEFAULT NULL COMMENT '删除时间',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0为未删除、1为已删除',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='分类表';

-- ----------------------------
-- Records of blog_category
-- ----------------------------
INSERT INTO `blog_category` VALUES ('1', 'JAVA', '0', '1', '0', '1', null, '0', '1');
INSERT INTO `blog_category` VALUES ('2', 'SpringBoot', '0', '1', '0', '1', null, '0', '1');
INSERT INTO `blog_category` VALUES ('3', 'MySql', '0', '1', '0', '1', null, '0', '1');
INSERT INTO `blog_category` VALUES ('4', '随笔', '0', '1', '0', '1', null, '0', '1');

-- ----------------------------
-- Table structure for blog_focus
-- ----------------------------
DROP TABLE IF EXISTS `blog_focus`;
CREATE TABLE `blog_focus` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) DEFAULT '' COMMENT '标题',
  `img` varchar(255) DEFAULT '' COMMENT '图片',
  `deleted_at` int(10) unsigned DEFAULT NULL COMMENT '删除时间',
  `modified_at` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0为未删除、1为已删除',
  `created_at` int(10) unsigned DEFAULT '0' COMMENT '新建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COMMENT='首页焦点图';

-- ----------------------------
-- Records of blog_focus
-- ----------------------------
INSERT INTO `blog_focus` VALUES ('1', 'Akina', 'https://s1.ax1x.com/2020/05/14/YDfRnU.jpg', null, '0', '0', '0', '0', '0');
INSERT INTO `blog_focus` VALUES ('2', '使用说明', 'https://s1.ax1x.com/2020/05/14/YDf4AJ.jpg', null, '0', '0', '0', '0', '0');
INSERT INTO `blog_focus` VALUES ('3', '文章归档', 'https://s1.ax1x.com/2020/05/14/YDfT91.jpg', null, '0', '0', '0', '0', '0');

-- ----------------------------
-- Table structure for blog_site
-- ----------------------------
DROP TABLE IF EXISTS `blog_site`;
CREATE TABLE `blog_site` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `avatar` varchar(255) DEFAULT '' COMMENT '头像',
  `slogan` varchar(255) DEFAULT '' COMMENT '标语',
  `name` varchar(255) DEFAULT '' COMMENT '名称',
  `domain` varchar(255) DEFAULT '' COMMENT '域名',
  `notice` varchar(255) DEFAULT '' COMMENT '提示',
  `desc` varchar(255) DEFAULT '' COMMENT '描述',
  `created_at` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_at` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_at` int(10) DEFAULT NULL COMMENT '删除时间',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0为未删除、1为已删除',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='站点信息表';

-- ----------------------------
-- Records of blog_site
-- ----------------------------
INSERT INTO `blog_site` VALUES ('1', 'https://s2.ax1x.com/2020/01/17/1SCadg.png', 'The way up is not crowded, and most chose ease.', 'ZX′blog\'', 'https://www.fengziy.cn', '本博客后端数据及架构为gin+mysql提供', '一个It技术的探索者', '0', '1', '0', '1', null, '0');
