SET NAMES utf8mb4;

# 转储表 casbin_rule
# ------------------------------------------------------------

DROP TABLE IF EXISTS `casbin_rule`;

CREATE TABLE `casbin_rule` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) DEFAULT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `casbin_rule` WRITE;
/*!40000 ALTER TABLE `casbin_rule` DISABLE KEYS */;

INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`)
VALUES
	(18,'p','1','manage:user:list','','','',''),
	(19,'p','1','manage:user:add','','','',''),
	(20,'p','1','manage:role:getAll','','','',''),
	(21,'p','1','manage:user:delete','','','',''),
	(22,'p','1','manage:user:edit','','','',''),
	(23,'p','1','manage:role:getAll','','','',''),
	(24,'p','1','manage:role:list','','','',''),
	(25,'p','1','manage:role:add','','','',''),
	(26,'p','1','manage:role:edit','','','',''),
	(27,'p','1','manage:role:getMenus','','','',''),
	(28,'p','1','manage:role:setMenus','','','',''),
	(29,'p','1','manage:menu:getAllPages','','','',''),
	(30,'p','1','manage:menu:tree','','','',''),
	(31,'p','1','manage:role:delete','','','',''),
	(32,'p','1','manage:menu:list','','','',''),
	(33,'p','1','manage:menu:add','','','',''),
	(34,'p','1','manage:menu:edit','','','',''),
	(35,'p','1','manage:menu:delete','','','',''),
	(36,'p','2','manage:user:list','','','',''),
	(37,'p','2','manage:role:list','','','',''),
	(38,'p','2','manage:menu:list','','','','');

/*!40000 ALTER TABLE `casbin_rule` ENABLE KEYS */;
UNLOCK TABLES;

# 转储表 manage_menu
# ------------------------------------------------------------

DROP TABLE IF EXISTS `manage_menu`;

CREATE TABLE `manage_menu` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `create_by` bigint DEFAULT NULL,
  `update_by` bigint DEFAULT NULL,
  `status` varchar(1) NOT NULL,
  `parent_id` bigint NOT NULL,
  `menu_type` varchar(1) NOT NULL,
  `menu_name` varchar(50) NOT NULL,
  `hide_in_menu` tinyint(1) NOT NULL,
  `active_menu` varchar(50) DEFAULT NULL,
  `order` bigint NOT NULL,
  `route_name` varchar(255) NOT NULL,
  `route_path` varchar(255) NOT NULL,
  `component` varchar(255) NOT NULL,
  `icon` varchar(255) NOT NULL,
  `icon_type` varchar(1) NOT NULL,
  `i18n_key` varchar(255) NOT NULL,
  `keep_alive` tinyint(1) NOT NULL,
  `href` longtext,
  `multi_tab` tinyint(1) DEFAULT NULL,
  `fixed_index_in_tab` bigint NOT NULL,
  `query` longtext,
  `permissions` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `constant` tinyint(1) NOT NULL,
  `button_code` longtext,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `manage_menu` WRITE;
/*!40000 ALTER TABLE `manage_menu` DISABLE KEYS */;

INSERT INTO `manage_menu` (`id`, `create_time`, `update_time`, `create_by`, `update_by`, `status`, `parent_id`, `menu_type`, `menu_name`, `hide_in_menu`, `active_menu`, `order`, `route_name`, `route_path`, `component`, `icon`, `icon_type`, `i18n_key`, `keep_alive`, `href`, `multi_tab`, `fixed_index_in_tab`, `query`, `permissions`, `constant`, `button_code`)
VALUES
	(1,'2024-11-10 21:12:35','2024-11-10 21:12:35',NULL,NULL,'1',0,'2','关于',0,'',6,'about','/about','layout.base$view.about','fluent:book-information-24-regular','1','route.about',0,'',0,0,'[]','[]',0,NULL),
	(2,'2024-11-10 21:43:39','2024-11-10 21:43:39',NULL,NULL,'1',0,'2','首页',0,'',1,'home','/home','layout.base$view.home','mdi:monitor-dashboard','1','route.home',0,'',0,0,'[]','[]',0,NULL),
	(3,'2024-11-11 09:19:18','2024-11-11 09:19:18',NULL,NULL,'1',0,'1','系统管理',0,'',5,'manage','/manage','layout.base','carbon:cloud-service-management','1','route.manage',0,'',0,0,'[]','[]',0,NULL),
	(4,'2024-11-11 09:20:26','2024-11-11 09:20:26',NULL,NULL,'1',3,'2','用户管理',0,'',1,'manage_user','/manage/user','view.manage_user','ic:round-manage-accounts','1','route.manage_user',0,'',0,0,'[]','[]',0,''),
	(7,'2024-11-11 10:29:03','2024-11-11 10:29:03',NULL,NULL,'1',3,'2','角色管理',0,'',2,'manage_role','/manage/role','view.manage_role','carbon:user-role','1','route.manage_role',0,'',0,0,'[]','[]',0,NULL),
	(8,'2024-11-11 10:29:33','2024-11-11 10:29:33',NULL,NULL,'1',3,'2','菜单管理',0,'',5,'manage_menu','/manage/menu','view.manage_menu','material-symbols:route','1','route.manage_menu',0,'',0,0,'[]','[]',0,NULL),
	(13,'2024-11-25 05:23:25','2024-11-25 05:23:25',NULL,NULL,'1',0,'2','403',1,'',0,'403','/403','layout.blank$view.403','no-permission','2','route.403',0,'',0,0,'[]','[]',1,NULL),
	(14,'2024-11-25 05:24:40','2024-11-25 05:24:40',NULL,NULL,'1',0,'2','404',1,'',0,'404','/404','layout.blank$view.404','not-found','2','route.404',0,'',0,0,'[]','[]',1,NULL),
	(15,'2024-11-25 05:25:14','2024-11-25 05:25:14',NULL,NULL,'1',0,'2','500',1,'',0,'500','/500','layout.blank$view.500','service-error','2','route.500',0,'',0,0,'[]','[]',1,NULL),
	(17,'2024-11-25 05:52:32','2024-11-25 05:52:32',NULL,NULL,'1',0,'2','login',1,'',0,'login','/login/:module(pwd-login|code-login|register|reset-pwd|bind-wechat)?','layout.blank$view.login','no-icon','2','route.login',0,'',0,0,'[]','[]',1,NULL),
	(20,'2024-11-26 15:17:16','2024-11-26 15:17:16',NULL,NULL,'1',4,'3','用户列表',0,'',1,'','','','','1','button.manage.user.list',0,'',0,0,'[]','[{\"code\":\"manage:user:list\",\"desc\":\"用户列表接口\"}]',0,'manage:user:list'),
	(21,'2024-11-26 16:04:56','2024-11-26 16:04:56',NULL,NULL,'1',4,'3','创建用户',0,'',0,'','','','','1','button.manage.user.add',0,'',0,0,'[]','[{\"code\":\"manage:user:add\",\"desc\":\"新增用户\"},{\"code\":\"manage:role:getAll\",\"desc\":\"获取所有角色\"}]',0,'manage:user:add'),
	(22,'2024-11-28 11:13:13','2024-11-28 11:13:13',NULL,NULL,'1',4,'3','删除用户',0,'',2,'','','','','1','button.manage.user.delete',0,'',0,0,'[]','[{\"code\":\"manage:user:delete\",\"desc\":\"删除用户\"}]',0,'manage:user:delete'),
	(23,'2024-11-28 11:16:09','2024-11-28 11:16:09',NULL,NULL,'1',4,'3','编辑用户',0,'',3,'','','','','1','button.manage.user.edit',0,'',0,0,'[]','[{\"code\":\"manage:user:edit\",\"desc\":\"编辑用户\"},{\"code\":\"manage:role:getAll\",\"desc\":\"获取所有角色\"}]',0,'manage:user:edit'),
	(24,'2024-11-28 11:23:32','2024-11-28 11:23:32',NULL,NULL,'1',7,'3','角色列表',0,'',0,'','','','','1','button.manage.role.list',0,'',0,0,'[]','[{\"code\":\"manage:role:list\",\"desc\":\"角色列表\"}]',0,'manage:role:list'),
	(25,'2024-11-28 11:25:25','2024-11-28 11:25:25',NULL,NULL,'1',7,'3','新增角色',0,'',1,'','','','','1','button.manage.role.add',0,'',0,0,'[]','[{\"code\":\"manage:role:add\",\"desc\":\"新增角色\"}]',0,'manage:role:add'),
	(26,'2024-11-28 11:26:03','2024-11-28 11:26:03',NULL,NULL,'1',7,'3','编辑角色',0,'',2,'','','','','1','button.manage.role.edit',0,'',0,0,'[]','[{\"code\":\"manage:role:edit\",\"desc\":\"编辑角色\"},{\"code\":\"manage:role:getMenus\",\"desc\":\"获取菜单权限\"},{\"code\":\"manage:role:setMenus\",\"desc\":\"设置菜单权限\"},{\"code\":\"manage:menu:getAllPages\",\"desc\":\"获取菜单页面\"},{\"code\":\"manage:menu:tree\",\"desc\":\"获取菜单树\"}]',0,'manage:role:edit'),
	(27,'2024-11-28 11:27:10','2024-11-28 11:27:10',NULL,NULL,'1',7,'3','删除用户',0,'',3,'','','','','1','button.manage.role.delete',0,'',0,0,'[]','[{\"code\":\"manage:role:delete\",\"desc\":\"删除用户\"}]',0,'manage:role:delete'),
	(28,'2024-11-28 12:14:07','2024-11-28 12:14:07',NULL,NULL,'1',8,'3','菜单列表',0,'',0,'','','','','1','button.manage.menu.list',0,'',0,0,'[]','[{\"code\":\"manage:menu:list\",\"desc\":\"菜单列表\"}]',0,'manage:menu:list'),
	(29,'2024-11-28 12:22:58','2024-11-28 12:22:58',NULL,NULL,'1',8,'3','添加下一级菜单',0,'',1,'','','','','1','button.manage.menu.add',0,'',0,0,'[]','[{\"code\":\"manage:menu:add\",\"desc\":\"添加下一级菜单\"}]',0,'manage:menu:add'),
	(30,'2024-11-28 12:23:59','2024-11-28 12:23:59',NULL,NULL,'1',8,'3','编辑菜单',0,'',2,'','','','','1','button.manage.menu.edit',0,'',0,0,'[]','[{\"code\":\"manage:menu:edit\",\"desc\":\"编辑菜单\"}]',0,'manage:menu:edit'),
	(31,'2024-11-28 12:24:38','2024-11-28 12:24:38',NULL,NULL,'1',8,'3','删除菜单',0,'',3,'','','','','1','button.manage.menu.delete',0,'',0,0,'[]','[{\"code\":\"manage:menu:delete\",\"desc\":\"删除菜单\"}]',0,'manage:menu:delete');

/*!40000 ALTER TABLE `manage_menu` ENABLE KEYS */;
UNLOCK TABLES;


# 转储表 manage_role
# ------------------------------------------------------------

DROP TABLE IF EXISTS `manage_role`;

CREATE TABLE `manage_role` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `create_by` bigint DEFAULT NULL,
  `update_by` bigint DEFAULT NULL,
  `name` varchar(50) NOT NULL,
  `status` varchar(1) NOT NULL,
  `code` varchar(255) NOT NULL,
  `desc` longtext NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `manage_role` WRITE;
/*!40000 ALTER TABLE `manage_role` DISABLE KEYS */;

INSERT INTO `manage_role` (`id`, `create_time`, `update_time`, `create_by`, `update_by`, `name`, `status`, `code`, `desc`)
VALUES
	(1,'2024-10-29 22:23:50','2024-10-29 22:23:50',1,NULL,'超级管理员','1','R_SUPER','超级管理员'),
	(2,'2024-11-28 12:46:13','2024-11-28 12:46:13',1,NULL,'测试用户','1','R_TEST','测试用户');

/*!40000 ALTER TABLE `manage_role` ENABLE KEYS */;
UNLOCK TABLES;


# 转储表 manage_role_menu
# ------------------------------------------------------------

DROP TABLE IF EXISTS `manage_role_menu`;

CREATE TABLE `manage_role_menu` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `create_by` bigint DEFAULT NULL,
  `update_by` bigint DEFAULT NULL,
  `role_id` bigint NOT NULL,
  `menu_id` bigint NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `manage_role_menu` WRITE;
/*!40000 ALTER TABLE `manage_role_menu` DISABLE KEYS */;

INSERT INTO `manage_role_menu` (`id`, `create_time`, `update_time`, `create_by`, `update_by`, `role_id`, `menu_id`)
VALUES
	(215,'2024-11-28 12:34:07','2024-11-28 12:34:07',NULL,NULL,1,1),
	(216,'2024-11-28 12:34:07','2024-11-28 12:34:07',NULL,NULL,1,8),
	(217,'2024-11-28 12:34:07','2024-11-28 12:34:07',NULL,NULL,1,7),
	(218,'2024-11-28 12:34:07','2024-11-28 12:34:07',NULL,NULL,1,4),
	(219,'2024-11-28 12:34:07','2024-11-28 12:34:07',NULL,NULL,1,20),
	(220,'2024-11-28 12:34:07','2024-11-28 12:34:07',NULL,NULL,1,2),
	(221,'2024-11-28 12:34:07','2024-11-28 12:34:07',NULL,NULL,1,3),
	(222,'2024-11-28 12:34:07','2024-11-28 12:34:07',NULL,NULL,1,21),
	(223,'2024-11-28 12:34:07','2024-11-28 12:34:07',NULL,NULL,1,22),
	(224,'2024-11-28 12:34:07','2024-11-28 12:34:07',NULL,NULL,1,23),
	(225,'2024-11-28 12:34:07','2024-11-28 12:34:07',NULL,NULL,1,24),
	(226,'2024-11-28 12:34:07','2024-11-28 12:34:07',NULL,NULL,1,25),
	(227,'2024-11-28 12:34:07','2024-11-28 12:34:07',NULL,NULL,1,26),
	(228,'2024-11-28 12:34:07','2024-11-28 12:34:07',NULL,NULL,1,27),
	(229,'2024-11-28 12:34:07','2024-11-28 12:34:07',NULL,NULL,1,28),
	(230,'2024-11-28 12:34:07','2024-11-28 12:34:07',NULL,NULL,1,29),
	(231,'2024-11-28 12:34:07','2024-11-28 12:34:07',NULL,NULL,1,30),
	(232,'2024-11-28 12:34:07','2024-11-28 12:34:07',NULL,NULL,1,31),
	(241,'2024-11-28 12:52:14','2024-11-28 12:52:14',NULL,NULL,2,2),
	(242,'2024-11-28 12:52:14','2024-11-28 12:52:14',NULL,NULL,2,1),
	(243,'2024-11-28 12:52:14','2024-11-28 12:52:14',NULL,NULL,2,3),
	(244,'2024-11-28 12:52:14','2024-11-28 12:52:14',NULL,NULL,2,4),
	(245,'2024-11-28 12:52:14','2024-11-28 12:52:14',NULL,NULL,2,7),
	(246,'2024-11-28 12:52:14','2024-11-28 12:52:14',NULL,NULL,2,8),
	(247,'2024-11-28 12:52:14','2024-11-28 12:52:14',NULL,NULL,2,20),
	(248,'2024-11-28 12:52:14','2024-11-28 12:52:14',NULL,NULL,2,24),
	(249,'2024-11-28 12:52:14','2024-11-28 12:52:14',NULL,NULL,2,28);

/*!40000 ALTER TABLE `manage_role_menu` ENABLE KEYS */;
UNLOCK TABLES;


# 转储表 manage_user
# ------------------------------------------------------------

DROP TABLE IF EXISTS `manage_user`;

CREATE TABLE `manage_user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `create_by` bigint DEFAULT NULL,
  `update_by` bigint DEFAULT NULL,
  `username` varchar(30) NOT NULL,
  `password` varchar(100) NOT NULL,
  `nickname` varchar(30) NOT NULL,
  `gender` varchar(1) NOT NULL,
  `phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `status` varchar(1) NOT NULL,
  `email` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_manage_user_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `manage_user` WRITE;
/*!40000 ALTER TABLE `manage_user` DISABLE KEYS */;

INSERT INTO `manage_user` (`id`, `create_time`, `update_time`, `create_by`, `update_by`, `username`, `password`, `nickname`, `gender`, `phone`, `status`, `email`)
VALUES
	(1,'2024-10-24 09:45:00','2024-10-31 09:40:13',NULL,NULL,'jzero','123456','jzero','1','','1','jaron@jaronnie.com'),
	(3,'2024-11-28 12:50:29','2024-11-28 12:50:29',NULL,NULL,'test','123456','测试用户','1','1234567890','1','test@qq.com');

/*!40000 ALTER TABLE `manage_user` ENABLE KEYS */;
UNLOCK TABLES;


# 转储表 manage_user_role
# ------------------------------------------------------------

DROP TABLE IF EXISTS `manage_user_role`;

CREATE TABLE `manage_user_role` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `create_by` bigint DEFAULT NULL,
  `update_by` bigint DEFAULT NULL,
  `user_id` bigint NOT NULL,
  `role_id` bigint NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `manage_user_role` WRITE;
/*!40000 ALTER TABLE `manage_user_role` DISABLE KEYS */;

INSERT INTO `manage_user_role` (`id`, `create_time`, `update_time`, `create_by`, `update_by`, `user_id`, `role_id`)
VALUES
	(1,'2024-10-31 09:40:13','2024-10-31 09:40:13',NULL,NULL,1,1),
	(4,'2024-11-28 12:49:54','2024-11-28 12:49:54',NULL,NULL,2,2),
	(5,'2024-11-28 12:50:29','2024-11-28 12:50:29',NULL,NULL,3,2);

/*!40000 ALTER TABLE `manage_user_role` ENABLE KEYS */;
UNLOCK TABLES;
