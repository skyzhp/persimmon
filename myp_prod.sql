/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50559
 Source Host           : localhost:3306
 Source Schema         : myp_prod

 Target Server Type    : MySQL
 Target Server Version : 50559
 File Encoding         : 65001

 Date: 07/04/2018 12:25:24
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for pit_attachments
-- ----------------------------
DROP TABLE IF EXISTS `pit_attachments`;
CREATE TABLE `pit_attachments` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `path` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `user_id` int(11) NOT NULL DEFAULT '0',
  `hash1` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `md5` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `created_at` int(10) DEFAULT NULL,
  `updated_at` int(10) DEFAULT NULL,
  `client_ip` int(10) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Table structure for pit_categorys
-- ----------------------------
DROP TABLE IF EXISTS `pit_categorys`;
CREATE TABLE `pit_categorys` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `category_name` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `category_parent` int(11) NOT NULL DEFAULT '0',
  `category_flag` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `category_description` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `created_at` int(10) DEFAULT NULL,
  `updated_at` int(10) DEFAULT NULL,
  `client_ip` int(10) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `categorys_category_flag_unique` (`category_flag`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of pit_categorys
-- ----------------------------
BEGIN;
INSERT INTO `pit_categorys` VALUES (1, 'golang', 0, 'golang', 'golang note.', 1523074802, 1523074802, 2130706433);
COMMIT;

-- ----------------------------
-- Table structure for pit_comments
-- ----------------------------
DROP TABLE IF EXISTS `pit_comments`;
CREATE TABLE `pit_comments` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` int(10) DEFAULT '0',
  `posts_id` int(11) DEFAULT NULL,
  `name` varchar(255) COLLATE utf8_unicode_ci DEFAULT '',
  `email` varchar(255) COLLATE utf8_unicode_ci DEFAULT '',
  `url` varchar(255) COLLATE utf8_unicode_ci DEFAULT '',
  `content` text COLLATE utf8_unicode_ci,
  `markdown` text COLLATE utf8_unicode_ci,
  `created_at` int(10) DEFAULT NULL,
  `updated_at` int(10) DEFAULT NULL,
  `client_ip` int(10) DEFAULT NULL,
  `status` tinyint(1) unsigned DEFAULT '1' COMMENT '1 正常 2待审核 3垃圾评论',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Table structure for pit_links
-- ----------------------------
DROP TABLE IF EXISTS `pit_links`;
CREATE TABLE `pit_links` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8_unicode_ci DEFAULT '',
  `logo` varchar(255) COLLATE utf8_unicode_ci DEFAULT '',
  `group` varchar(255) COLLATE utf8_unicode_ci DEFAULT '',
  `url` varchar(255) COLLATE utf8_unicode_ci DEFAULT '',
  `created_at` int(10) DEFAULT NULL,
  `updated_at` int(10) DEFAULT NULL,
  `client_ip` int(10) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of pit_links
-- ----------------------------
BEGIN;
INSERT INTO `pit_links` VALUES (1, 'MrCong', 'https://o75u5ooep.qnssl.com/avatar/persimmon.jpg', '', 'https://cong5.net/', 1487105565, 1487105897, 0);
COMMIT;

-- ----------------------------
-- Table structure for pit_navications
-- ----------------------------
DROP TABLE IF EXISTS `pit_navications`;
CREATE TABLE `pit_navications` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `url` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `sort` int(11) DEFAULT '0',
  `created_at` int(10) DEFAULT NULL,
  `updated_at` int(10) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Table structure for pit_options
-- ----------------------------
DROP TABLE IF EXISTS `pit_options`;
CREATE TABLE `pit_options` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `option_title` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `option_name` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `option_value` text COLLATE utf8_unicode_ci,
  `option_group` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `option_remark` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `option_status` tinyint(1) DEFAULT '2' COMMENT '0 hidden 1 base 2 extends ',
  `data_type` tinyint(1) DEFAULT '1' COMMENT '1 text  2 textarea',
  `created_at` int(10) DEFAULT NULL,
  `updated_at` int(10) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of pit_options
-- ----------------------------
BEGIN;
INSERT INTO `pit_options` VALUES (3, '网站名称', 'site_name', 'Mr柿子 - I\'am Mr柿子.', '', '', 1, 1, 1487167761, 1487546220);
INSERT INTO `pit_options` VALUES (4, '网站关键词', 'keywords', 'Mr柿子,MrCong,PHP 开发笔记,Swift 笔记,iOS 开发笔记', NULL, NULL, 1, 2, 1487167779, 1487194870);
INSERT INTO `pit_options` VALUES (5, '网站描述', 'description', 'Hi~ ，我是刘聪(@Mr柿子),一名PHPer，爱好折腾各种新技术，热爱PHP、Swift.', '', '', 1, 2, 1487167800, 1487194870);
INSERT INTO `pit_options` VALUES (6, '导航配置', 'navigations', '[{\"name\":\"前端开发\",\"url\":\"/categories/web\"},{\"name\":\"PHP开发\",\"url\":\"/categories/php\"},{\"name\":\"Linux\",\"url\":\"/categories/linux\"},{\"name\":\"iOS开发\",\"url\":\"/categories/ios\"},{\"name\":\"其他\",\"url\":\"/categories/other\"},{\"name\":\"Swift开发\",\"url\":\"/categories/swift\"},{\"name\":\"友情链接\",\"url\":\"/friends\"}]', '', '', 0, 2, 1487198194, 1522596442);
INSERT INTO `pit_options` VALUES (7, '新浪微博', 'weibo', 'http://weibo.com/701265306', '', '', 2, 1, 1487545984, 1487546153);
INSERT INTO `pit_options` VALUES (8, 'GitHub 地址', 'github', 'https://github.com/cong5', NULL, NULL, 2, 1, 1487546011, 1487546153);
INSERT INTO `pit_options` VALUES (9, 'Google Plus', 'google_plus', 'https://plus.google.com/+CongLiuForMe/', '', '', 2, 1, 1487546054, 1487546153);
INSERT INTO `pit_options` VALUES (10, '统计代码', 'analysis', '<script src=\"https://s13.cnzz.com/z_stat.php?id=1266985713&web_id=1266985713\" language=\"JavaScript\"></script>', '', '', 2, 2, 1487546593, 1509398943);
INSERT INTO `pit_options` VALUES (11, '备案号', 'icp', '桂ICP备14007691号', '', '', 2, 1, 1487547000, 1487547020);
INSERT INTO `pit_options` VALUES (13, 'XmlRpc登陆失败次数', 'xmlrpc_login_failed', '4', NULL, NULL, 2, 1, 1489086282, 1492661326);
COMMIT;

-- ----------------------------
-- Table structure for pit_posts
-- ----------------------------
DROP TABLE IF EXISTS `pit_posts`;
CREATE TABLE `pit_posts` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `flag` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `title` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `thumb` varchar(200) COLLATE utf8_unicode_ci DEFAULT NULL,
  `category_id` int(11) NOT NULL DEFAULT '0',
  `user_id` int(11) NOT NULL,
  `content` text COLLATE utf8_unicode_ci,
  `markdown` text COLLATE utf8_unicode_ci,
  `views` int(11) DEFAULT '0',
  `comments` int(11) DEFAULT '0',
  `created_at` int(10) DEFAULT NULL,
  `updated_at` int(10) DEFAULT NULL,
  `deleted_at` int(10) DEFAULT NULL,
  `client_ip` int(10) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `posts_flag_unique` (`flag`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of pit_posts
-- ----------------------------
BEGIN;
INSERT INTO `pit_posts` VALUES (1, 'persimmon-hello-world', 'Persimmon 你好！ Golang 你好', '', 1, 1, '<div class=\"post-active\"><p>Persimmon Hello!<br>Golang Hello.<br>Welcome to persimmon blog.</p>\n<h2 id=\"first-set-golang-system-env\">First Set Golang System ENV</h2>\n<pre><code><span class=\"hljs-meta\">#</span><span class=\"bash\"> <span class=\"hljs-keyword\">if</span> your not <span class=\"hljs-built_in\">set</span> GOPATH env,please setting.</span>\n<span class=\"hljs-meta\">#</span><span class=\"bash\"><span class=\"hljs-built_in\">export</span> GOPATH=<span class=\"hljs-variable\">$HOME</span>/go</span>\n<span class=\"hljs-meta\">#</span><span class=\"bash\"><span class=\"hljs-built_in\">export</span> PATH=<span class=\"hljs-variable\">$HOME</span>/bin:<span class=\"hljs-variable\">$GOPATH</span>/bin:<span class=\"hljs-variable\">$PATH</span></span>\nexport GO15VENDOREXPERIMENT=1\n</code></pre><h2 id=\"package-management\">Package Management</h2>\n<pre><code>https:<span class=\"hljs-comment\">//github.com/Masterminds/glide</span>\n</code></pre><h2 id=\"installation\">Installation</h2>\n<pre><code>go get github.com/cong5/persimmon\n<span class=\"hljs-built_in\">cd</span> <span class=\"hljs-variable\">$GOPATH</span>/src/github.com/cong5/persimmon\nglide install\n</code></pre><h2 id=\"run-app\">Run APP</h2>\n<pre><code><span class=\"hljs-built_in\">cd</span> <span class=\"hljs-variable\">$GOPATH</span> &amp;&amp; revel run github.com/cong5/persimmon\n</code></pre><h2 id=\"build\">build</h2>\n<p>Build persimmon binary exec file to current directory.</p>\n<pre><code><span class=\"hljs-built_in\">cd</span> <span class=\"hljs-variable\">$GOPATH</span>/src/github.com/cong5/persimmon\ngo build -o ./bin/persimmon github.com/cong5/persimmon/app/tmp\n</code></pre><h2 id=\"-persimmon\">运行persimmon</h2>\n<p>Argument -importPath required.</p>\n<pre><code>./bin/persimmon -importPath=github.com/cong5/persimmon -runMode=prod -port=9100\n</code></pre><p>FYI, in the meantime, after Revel generates your app, you can just do:</p>\n<p>go build -o bin/myapp import/path/to/myapp/app/tmp<br>run gdb on the binary bin/myapp.<br>Pass in flags \"-importPath\", \"-srcPath\", \"-runMode\"</p>\n</div>', 'Persimmon Hello!\nGolang Hello.\nWelcome to persimmon blog.\n\n## First Set Golang System ENV\n\n```\n# if your not set GOPATH env,please setting.\n#export GOPATH=$HOME/go\n#export PATH=$HOME/bin:$GOPATH/bin:$PATH\nexport GO15VENDOREXPERIMENT=1\n```\n\n## Package Management\n\n```\nhttps://github.com/Masterminds/glide\n```\n\n## Installation\n\n```\ngo get github.com/cong5/persimmon\ncd $GOPATH/src/github.com/cong5/persimmon\nglide install\n```\n\n## Run APP\n\n```\ncd $GOPATH && revel run github.com/cong5/persimmon\n```\n\n## build\n\nBuild persimmon binary exec file to current directory.\n\n```\ncd $GOPATH/src/github.com/cong5/persimmon\ngo build -o ./bin/persimmon github.com/cong5/persimmon/app/tmp\n```\n\n## 运行persimmon\n\nArgument -importPath required.\n\n```\n./bin/persimmon -importPath=github.com/cong5/persimmon -runMode=prod -port=9100\n```\n\nFYI, in the meantime, after Revel generates your app, you can just do:\n\ngo build -o bin/myapp import/path/to/myapp/app/tmp\nrun gdb on the binary bin/myapp.\nPass in flags \"-importPath\", \"-srcPath\", \"-runMode\"\n', 1, 0, 1523073486, 1523073486, NULL, 2130706433);
COMMIT;

-- ----------------------------
-- Table structure for pit_posts_tags
-- ----------------------------
DROP TABLE IF EXISTS `pit_posts_tags`;
CREATE TABLE `pit_posts_tags` (
  `posts_id` int(10) unsigned NOT NULL,
  `tags_id` int(10) unsigned NOT NULL,
  `created_at` int(10) DEFAULT NULL,
  `updated_at` int(10) DEFAULT NULL,
  KEY `posts_tags_post_id_index` (`posts_id`),
  KEY `posts_tags_tag_id_index` (`tags_id`),
  CONSTRAINT `posts_tags_post_id_foreign` FOREIGN KEY (`posts_id`) REFERENCES `pit_posts` (`id`) ON DELETE CASCADE,
  CONSTRAINT `posts_tags_tag_id_foreign` FOREIGN KEY (`tags_id`) REFERENCES `pit_tags` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Table structure for pit_tags
-- ----------------------------
DROP TABLE IF EXISTS `pit_tags`;
CREATE TABLE `pit_tags` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `tags_name` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `tags_flag` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `created_at` int(10) DEFAULT NULL,
  `updated_at` int(10) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Table structure for pit_users
-- ----------------------------
DROP TABLE IF EXISTS `pit_users`;
CREATE TABLE `pit_users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `email` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `password` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `avatar` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `remember_token` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
  `created_at` int(10) DEFAULT NULL,
  `updated_at` int(10) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `users_email_unique` (`email`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of pit_users
-- ----------------------------
BEGIN;
INSERT INTO `pit_users` VALUES (1, 'persimmon', 'persimmon@cong5.net', '$2a$10$x0YiIpILZrqgymyiIVNZyO2egLPQd1zlk7N1OSi5r9AI6NL4iUpRu', 'https://o75u5ooep.qnssl.com/avatar/persimmon.jpg', '', 1523073838, 1523073838);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
