/*
Navicat MySQL Data Transfer

Source Server         : localhost_3306
Source Server Version : 80016
Source Host           : localhost:3306
Source Database       : simple_demo

Target Server Type    : MYSQL
Target Server Version : 80016
File Encoding         : 65001

Date: 2022-06-13 13:08:06
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `comment`
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
  `comment_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  `video_id` bigint(20) unsigned DEFAULT NULL,
  `content` longtext NOT NULL,
  `created_at` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`comment_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of comment
-- ----------------------------

-- ----------------------------
-- Table structure for `follow`
-- ----------------------------
DROP TABLE IF EXISTS `follow`;
CREATE TABLE `follow` (
  `user_id` bigint(20) unsigned NOT NULL,
  `fan_id` bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`fan_id`),
  KEY `fk_follow_fans` (`fan_id`),
  CONSTRAINT `fk_follow_fans` FOREIGN KEY (`fan_id`) REFERENCES `user` (`userid`),
  CONSTRAINT `fk_follow_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`userid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of follow
-- ----------------------------

-- ----------------------------
-- Table structure for `like`
-- ----------------------------
DROP TABLE IF EXISTS `like`;
CREATE TABLE `like` (
  `user_id` bigint(20) unsigned NOT NULL,
  `video_id` bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`video_id`),
  KEY `fk_like_video` (`video_id`),
  CONSTRAINT `fk_like_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`userid`),
  CONSTRAINT `fk_like_video` FOREIGN KEY (`video_id`) REFERENCES `video` (`video_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of like
-- ----------------------------

-- ----------------------------
-- Table structure for `user`
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `userid` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(191) NOT NULL,
  `password` longtext NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `token` longtext NOT NULL,
  PRIMARY KEY (`userid`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('1', 'zhanglei', '87a38998227cbbc23dcad51cd7f76ab2', '2022-06-13 12:56:35.957', '2022-06-13 12:56:35.957', '0000-00-00 00:00:00.000', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyTmFtZSI6InpoYW5nbGVpIiwidXNlclBhc3N3b3JkIjoiODdhMzg5OTgyMjdjYmJjMjNkY2FkNTFjZDdmNzZhYjIiLCJleHAiOjE2NTUwOTk3OTUsImlhdCI6MTY1NTA5NjE5NSwiaXNzIjoiRG91eWluIiwic3ViIjoidXNlclRva2VuIn0.EO9ChKBfFsd9pKMkYoFoDutJn7J8h3V6yjXSRVrCQug');
INSERT INTO `user` VALUES ('2', 'devin', 'e10adc3949ba59abbe56e057f20f883e', '2022-06-13 12:56:45.256', '2022-06-13 12:56:45.256', '0000-00-00 00:00:00.000', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyTmFtZSI6ImRldmluIiwidXNlclBhc3N3b3JkIjoiZTEwYWRjMzk0OWJhNTlhYmJlNTZlMDU3ZjIwZjg4M2UiLCJleHAiOjE2NTUwOTk4MDUsImlhdCI6MTY1NTA5NjIwNSwiaXNzIjoiRG91eWluIiwic3ViIjoidXNlclRva2VuIn0.OSHmiOTOO7fnaFXKtSM27CfI1gfEy5_N8DO_PsyUc-I');
INSERT INTO `user` VALUES ('3', 'test', 'e10adc3949ba59abbe56e057f20f883e', '2022-06-13 12:56:49.292', '2022-06-13 12:56:49.292', '0000-00-00 00:00:00.000', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyTmFtZSI6InRlc3QiLCJ1c2VyUGFzc3dvcmQiOiJlMTBhZGMzOTQ5YmE1OWFiYmU1NmUwNTdmMjBmODgzZSIsImV4cCI6MTY1NTA5OTgwOSwiaWF0IjoxNjU1MDk2MjA5LCJpc3MiOiJEb3V5aW4iLCJzdWIiOiJ1c2VyVG9rZW4ifQ.ON8_TppDyzjblmR0Oel7jiqkz4NpNZ830UPcmoaMeTA');

-- ----------------------------
-- Table structure for `video`
-- ----------------------------
DROP TABLE IF EXISTS `video`;
CREATE TABLE `video` (
  `video_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned NOT NULL,
  `title` longtext NOT NULL,
  `created_at` bigint(20) DEFAULT NULL,
  `play_url` longtext,
  `cover_url` longtext,
  PRIMARY KEY (`video_id`),
  KEY `fk_user_videos` (`user_id`),
  CONSTRAINT `fk_user_videos` FOREIGN KEY (`user_id`) REFERENCES `user` (`userid`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of video
-- ----------------------------
INSERT INTO `video` VALUES ('1', '1', 'my first', '1655084877', '3a47609f70dde7424a72635f0a3ac2be.mp4', 'bear.jpg');
INSERT INTO `video` VALUES ('2', '1', 'my second', '1655090257', '7de56f312545cde64d43e790c06acc81.mp4', 'bear.jpg');
INSERT INTO `video` VALUES ('3', '1', 'my third', '1655094370', '9ddecd33ec57f06b624caa7a62e58f0a.mp4', 'bear.jpg');
INSERT INTO `video` VALUES ('4', '1', 'my fourth', '1655094488', '54edcc6ea75a5317e8d7fdd053001341.mp4', 'bear.jpg');
INSERT INTO `video` VALUES ('5', '2', 'hi one', '1655094555', '92ea485fbedc3b294ab5308a47b30934.mp4', 'bear.jpg');
INSERT INTO `video` VALUES ('6', '2', 'hi two', '1655094590', '761feeb8f2adf6ddec787091ad615091.mp4', 'bear.jpg');
INSERT INTO `video` VALUES ('7', '2', 'hi three', '1655094680', '5889f16269f5cc7e7da432e5a14ffa14.mp4', 'bear.jpg');
INSERT INTO `video` VALUES ('8', '3', 'hey man1', '1655094712', '6889aed365f7470edfff41ed3bb7a978.mp4', 'bear.jpg');
INSERT INTO `video` VALUES ('9', '3', '22222', '1655094752', '087619f1697d3708944ca2028bd53236.mp4', 'bear.jpg');
INSERT INTO `video` VALUES ('10', '3', '3333', '1655094785', '558404624fd279d368781193fb58b840.mp4', 'bear.jpg');
