/*
 Navicat Premium Data Transfer

 Source Server         : 118.89.237.69
 Source Server Type    : MySQL
 Source Server Version : 50722
 Source Host           : 118.89.237.69:3306
 Source Schema         : db_hb

 Target Server Type    : MySQL
 Target Server Version : 50722
 File Encoding         : 65001

 Date: 22/04/2019 10:19:39
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for about
-- ----------------------------
DROP TABLE IF EXISTS `about`;
CREATE TABLE `about` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `body` longtext NOT NULL,
  `time` datetime NOT NULL,
  `photo` varchar(255) NOT NULL DEFAULT '',
  `title` varchar(255) NOT NULL DEFAULT '',
  `music` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of about
-- ----------------------------
BEGIN;
INSERT INTO `about` VALUES (1, '迪丽热巴（Dilraba），1992年6月3日出生于新疆乌鲁木齐市，中国内地影视女演员。2013年，迪丽热巴因主演个人首部电视剧《阿娜尔罕》而出道。2014年，她主演了奇幻剧《逆光之恋》。迪丽热巴（Dilraba），1992年6月3日出生于新疆乌鲁木齐市，中国内地影视女演员。2013年，迪丽热巴因主演个人首部电视剧《阿娜尔罕》而出道。2014年，她主演了奇幻剧《逆光之恋》。迪丽热巴（Dilraba），1992年6月3日出生于新疆乌鲁木齐市，中国内地影视女演员。2013年，迪丽热巴因主演个人首部电视剧《阿娜尔罕》而出道。2014年，她主演了奇幻剧《逆光之恋》。迪丽热巴（Dilraba），1992年6月3日出生于新疆乌鲁木齐市，中国内地影视女演员。2013年，迪丽热巴因主演个人首部电视剧《阿娜尔罕》而出道。2014年，她主演了奇幻剧《逆光之恋》。迪丽热巴（Dilraba），1992年6月3日出生于新疆乌鲁木齐市，中国内地影视女演员。2013年，迪丽热巴因主演个人首部电视剧《阿娜尔罕》而出道。2014年，她主演了奇幻剧《逆光之恋》。迪丽热巴（Dilraba），1992年6月3日出生于新疆乌鲁木齐市，中国内地影视女演员。2013年，迪丽热巴因主演个人首部电视剧《阿娜尔罕》而出道。2014年，她主演了奇幻剧《逆光之恋》。迪丽热巴（Dilraba），1992年6月3日出生于新疆乌鲁木齐市，中国内地影视女演员。2013年，迪丽热巴因主演个人首部电视剧《阿娜尔罕》而出道。2014年，她主演了奇幻剧《逆光之恋》。', '2018-06-26 14:12:33', '', '迪丽热巴（Dilraba），1992年6月3日出生于新疆乌鲁木齐市，中国内地影视女演员。2013年，迪丽热巴因主演个人首部电视剧《阿娜尔罕》而出道。2014年，她主演了奇幻剧《逆光之恋》。', '3778678');
COMMIT;

-- ----------------------------
-- Table structure for album
-- ----------------------------
DROP TABLE IF EXISTS `album`;
CREATE TABLE `album` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL DEFAULT '',
  `content` longtext NOT NULL,
  `photo_path` varchar(255) NOT NULL DEFAULT '',
  `time` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of album
-- ----------------------------
BEGIN;
INSERT INTO `album` VALUES (1, '个人博客模板《simple》', '好咖啡要和朋友一起品尝，好“模板”也要和同样喜欢它的人一起分享。', '', '2018-06-26 19:02:18');
INSERT INTO `album` VALUES (2, '个人博客模板《simple》', '好咖啡要和朋友一起品尝，好“模板”也要和同样喜欢它的人一起分享。', '', '2018-06-26 19:02:40');
INSERT INTO `album` VALUES (3, '个人博客模板《simple》', '好咖啡要和朋友一起品尝，好“模板”也要和同样喜欢它的人一起分享。', '', '2018-06-26 19:03:03');
COMMIT;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL DEFAULT '',
  `body` longtext NOT NULL,
  `time` datetime NOT NULL,
  `category_id` int(11) NOT NULL,
  `readers_id` int(11) NOT NULL,
  `cover` varchar(255) NOT NULL DEFAULT '',
  `key` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `readers_id` (`readers_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of article
-- ----------------------------
BEGIN;
INSERT INTO `article` VALUES (1, '个人博客，属于我的小世界！', '本文很长，记录了我博客建站初到现在的过程，还有我从毕业到现在的一个状态，感谢您的阅读，如果你还是学生，也许你能从此文中，找到我们曾经相似的地方。如果你已经工作，有自己的博客，我想，你并没有忘记当初建立个人博客的初衷吧！\n\n我的个人博客已经建站有8年的时间了，对它的热爱，一直都是只增未减。回想大学读书的那几年，那会儿非常流行QQ空间，我们寝室的室友还经常邀约去学校的网吧做自己的空间。系里有个男生，空间做得非常漂亮，什么悬浮，开场动画，音乐播放器，我们女生羡慕得不得了。还邀约他跟我们一起去通宵弄空间，网上可以找到很多免费的flash资源，还有音乐，那也是第一次接触js，知道在浏览器输入一个地址，修改一下数据，就能调用一些背景出来。然后把自己QQ空间弄得漂漂亮亮的，经常邀约室友来互踩。我记得08年地震，第二天晚上，我们寝室的几个人还淡定的在寝室装扮空间呢！', '2018-06-26 14:13:06', 1, 1, '', '个人博客');
INSERT INTO `article` VALUES (2, '爱情没有永远，地老天荒也走不完', '也许，爱情没有永远，地老天荒也走不完，生命终结的末端，苦短情长。站在岁月的边端，那些美丽的定格，心伤的绝恋，都被四季的掩埋，一去不返。徒剩下这荒芜的花好月圆，一路相随，流离天涯背负了谁的思念？', '2018-06-26 16:36:26', 1, 2, '', '个人博客,日记,三星');
INSERT INTO `article` VALUES (3, '女孩都有浪漫的小情怀——浪漫的求婚词', '还在为浪漫的求婚词而烦恼不知道该怎么说吗？女孩子都有着浪漫的小情怀，对于求婚更是抱着满满的浪漫期待，也希望在求婚那一天对方可以给自己一个最浪漫的求婚词。', '2018-06-26 16:38:24', 2, 3, '', 'java');
INSERT INTO `article` VALUES (4, '擦肩而过', '擦肩而过》文/清河鱼 编绘/天朝羽打开一扇窗，我不曾把你想得平常。看季节一一过往。你停留的那个地方，是否依然花儿开放？在夜里守靠着梦中的，想那仿佛前世铭刻进心肠的', '2018-06-26 16:43:54', 3, 4, '', 'go');
COMMIT;

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of category
-- ----------------------------
BEGIN;
INSERT INTO `category` VALUES (1, '心情');
INSERT INTO `category` VALUES (2, '技术');
INSERT INTO `category` VALUES (3, '日记');
COMMIT;

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `time` datetime NOT NULL,
  `articles_id` int(11) NOT NULL,
  `content` longtext NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of comment
-- ----------------------------
BEGIN;
INSERT INTO `comment` VALUES (1, '吴延博', '2018-06-27 10:39:52', 1, '来作死');
INSERT INTO `comment` VALUES (2, '吴延博', '2018-06-27 11:34:50', 2, '我来了！');
INSERT INTO `comment` VALUES (3, '老王', '2018-06-27 11:35:12', 3, '我来了！');
INSERT INTO `comment` VALUES (5, '老王', '2018-06-27 11:36:30', 1, 'ok？\r\n');
INSERT INTO `comment` VALUES (7, '托尼', '2018-06-27 11:42:00', 2, '钢铁侠！！！！！！！！！！！！！！！！！！');
INSERT INTO `comment` VALUES (16, 'dsadas', '2018-06-27 22:19:44', 1, '???');
INSERT INTO `comment` VALUES (17, 'ww', '2018-06-28 14:20:57', 1, 'ww');
INSERT INTO `comment` VALUES (18, '刘锦涛', '2018-08-01 17:31:35', 1, '我是王八蛋！！！');
INSERT INTO `comment` VALUES (19, 'u', '2019-04-22 10:09:19', 1, '？？\r\n');
COMMIT;

-- ----------------------------
-- Table structure for file
-- ----------------------------
DROP TABLE IF EXISTS `file`;
CREATE TABLE `file` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `time` datetime NOT NULL,
  `path` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for friendship_link
-- ----------------------------
DROP TABLE IF EXISTS `friendship_link`;
CREATE TABLE `friendship_link` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL DEFAULT '',
  `link` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of friendship_link
-- ----------------------------
BEGIN;
INSERT INTO `friendship_link` VALUES (1, '老辣条的风味博客', 'http://www.laolatiao.xyz');
COMMIT;

-- ----------------------------
-- Table structure for reader
-- ----------------------------
DROP TABLE IF EXISTS `reader`;
CREATE TABLE `reader` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `number` int(11) NOT NULL DEFAULT '0',
  `appreciate` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of reader
-- ----------------------------
BEGIN;
INSERT INTO `reader` VALUES (1, 2, 27);
INSERT INTO `reader` VALUES (2, 40, 341);
INSERT INTO `reader` VALUES (3, 40, 40);
INSERT INTO `reader` VALUES (4, 20, 2);
COMMIT;

-- ----------------------------
-- Table structure for recommend
-- ----------------------------
DROP TABLE IF EXISTS `recommend`;
CREATE TABLE `recommend` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL DEFAULT '',
  `link` varchar(255) NOT NULL DEFAULT '',
  `clicks` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of recommend
-- ----------------------------
BEGIN;
INSERT INTO `recommend` VALUES (1, '你是什么人便会遇上什么人', '/', 0);
INSERT INTO `recommend` VALUES (2, '帝国cms 列表页调用子栏目，没有则不显示栏目名称', '/', 1);
INSERT INTO `recommend` VALUES (3, '第二届 优秀个人博客模板比赛参选活动', '/', 2);
INSERT INTO `recommend` VALUES (4, '个人博客模板《绅士》后台管理', '/', 3);
COMMIT;

-- ----------------------------
-- Table structure for tag_cloud
-- ----------------------------
DROP TABLE IF EXISTS `tag_cloud`;
CREATE TABLE `tag_cloud` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `lable` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of tag_cloud
-- ----------------------------
BEGIN;
INSERT INTO `tag_cloud` VALUES (1, '三星');
INSERT INTO `tag_cloud` VALUES (2, '华为');
INSERT INTO `tag_cloud` VALUES (3, '小米');
INSERT INTO `tag_cloud` VALUES (4, '浪潮');
INSERT INTO `tag_cloud` VALUES (5, 'java');
INSERT INTO `tag_cloud` VALUES (6, 'go');
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(255) NOT NULL DEFAULT '',
  `password` varchar(255) NOT NULL DEFAULT '',
  `token` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_name` (`user_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

SET FOREIGN_KEY_CHECKS = 1;
