DROP TABLE IF EXISTS `profile`;
CREATE TABLE `profile` (
  `userId` int(20) unsigned NOT NULL,
  `level` int(10) unsigned NOT NULL DEFAULT 1,
  `exp` int(10) unsigned NOT NULL DEFAULT 0,
  `money` int(10) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`userId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `profile` VALUES (100, 5, 10000, 100000);