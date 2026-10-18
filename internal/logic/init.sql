/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `mms_members`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `mms_members` (
  `id` char(36) COLLATE utf8mb4_bin NOT NULL COMMENT 'UUID',
  `created_at` timestamp NOT NULL COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL COMMENT 'Update Time | 修改日期',
  `status` tinyint unsigned DEFAULT '1' COMMENT 'Status 1: normal 2: ban | 状态 1 正常 2 禁用',
  `username` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'Member''s login name | 登录名',
  `password` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'Password | 密码',
  `nickname` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'Nickname | 昵称',
  `mobile` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'Mobile number | 手机号',
  `email` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'Email | 邮箱号',
  `avatar` varchar(512) COLLATE utf8mb4_bin DEFAULT '' COMMENT 'Avatar | 头像路径',
  `wechat_open_id` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'Wechat Open ID | 微信 Open ID',
  `expired_at` datetime DEFAULT NULL COMMENT 'Member expired time | 会员到期时间',
  `rank_id` bigint unsigned DEFAULT NULL COMMENT 'Member Rank ID | 会员等级ID',
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `nickname` (`nickname`),
  UNIQUE KEY `member_username_email` (`username`,`email`),
  UNIQUE KEY `member_wechat_open_id` (`wechat_open_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `mms_ranks`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `mms_ranks` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL COMMENT 'Update Time | 修改日期',
  `name` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'Rank name | 等级名称',
  `code` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'Rank code | 等级码',
  `description` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'Rank description | 等级描述',
  `remark` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'Remark | 备注',
  PRIMARY KEY (`id`),
  UNIQUE KEY `memberrank_code` (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `mms_oauth_providers`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `mms_oauth_providers` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL COMMENT 'Update Time | 修改日期',
  `name` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The provider''s name | 提供商名称',
  `client_id` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The client id | 客户端 id',
  `client_secret` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The client secret | 客户端密钥',
  `redirect_url` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The redirect url | 跳转地址',
  `scopes` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The scopes | 权限范围',
  `auth_url` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The auth url of the provider | 认证地址',
  `token_url` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The token url of the provider | 获取 token地址',
  `auth_style` bigint unsigned NOT NULL COMMENT 'The auth style, 0: auto detect 1: third party log in 2: log in with username and password',
  `info_url` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The URL to request user information by token | 用户信息请求地址',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `mms_tokens`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `mms_tokens` (
  `id` char(36) COLLATE utf8mb4_bin NOT NULL COMMENT 'UUID',
  `created_at` timestamp NOT NULL COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL COMMENT 'Update Time | 修改日期',
  `status` tinyint unsigned DEFAULT '1' COMMENT 'Status 1: normal 2: ban | 状态 1 正常 2 禁用',
  `uuid` char(36) COLLATE utf8mb4_bin NOT NULL COMMENT ' User''s UUID | 用户的UUID',
  `token` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'Token string | Token 字符串',
  `username` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT 'unknown' COMMENT 'Username | 用户名',
  `source` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'Log in source such as GitHub | Token 来源 （本地为core, 第三方如github等）',
  `expired_at` timestamp NOT NULL COMMENT ' Expire time | 过期时间',
  PRIMARY KEY (`id`),
  KEY `mms_token_uuid` (`uuid`),
  KEY `mms_token_expired_at` (`expired_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Seed data: member ranks
--

INSERT IGNORE INTO `mms_ranks` VALUES
(1,'2026-10-15 06:49:53','2026-10-15 06:49:53','memberRank.normal','001','普通会员 | Normal Member','普通会员 | Normal Member'),
(2,'2026-10-15 06:49:53','2026-10-15 06:49:53','memberRank.vip','002','VIP','VIP');

--
-- Seed data: oauth providers
--

INSERT IGNORE INTO `mms_oauth_providers` VALUES
(1,'2026-10-15 06:49:53','2026-10-15 06:49:53','google','your client id','your client secret','http://localhost:3100/oauth/login/callback','email openid','https://accounts.google.com/o/oauth2/auth','https://oauth2.googleapis.com/token',1,'https://www.googleapis.com/oauth2/v2/userinfo?access_token=TOKEN'),
(2,'2026-10-15 06:49:53','2026-10-15 06:49:53','github','your client id','your client secret','http://localhost:3100/oauth/login/callback','email openid','https://github.com/login/oauth/authorize','https://github.com/login/oauth/access_token',2,'https://api.github.com/user');

--
-- Seed data: apis (member service REST endpoints)
--

INSERT IGNORE INTO `sys_apis` VALUES
(1001,'2026-10-15 06:49:53','2026-10-15 06:49:53','/mms-api/init/database','apiDesc.initDatabase','base','Mms','GET',1),
(1002,'2026-10-15 06:49:53','2026-10-15 06:49:53','/mms-api/captcha','apiDesc.getCaptcha','captcha','Mms','GET',1),
(1003,'2026-10-15 06:49:53','2026-10-15 06:49:53','/mms-api/captcha/email','apiDesc.getEmailCaptcha','captcha','Mms','POST',1),
(1004,'2026-10-15 06:49:53','2026-10-15 06:49:53','/mms-api/captcha/sms','apiDesc.getSmsCaptcha','captcha','Mms','POST',1),
(1005,'2026-10-15 06:49:53','2026-10-15 06:49:53','/mms-api/member/create','apiDesc.createMember','member','Mms','POST',0),
(1006,'2026-10-15 06:49:53','2026-10-15 06:49:53','/mms-api/member/update','apiDesc.updateMember','member','Mms','POST',0),
(1007,'2026-10-15 06:49:53','2026-10-15 06:49:53','/mms-api/member/delete','apiDesc.deleteMember','member','Mms','POST',0),
(1008,'2026-10-15 06:49:53','2026-10-15 06:49:53','/mms-api/member/list','apiDesc.getMemberList','member','Mms','POST',0),
(1009,'2026-10-15 06:49:53','2026-10-15 06:49:53','/mms-api/member','apiDesc.getMemberById','member','Mms','POST',0),
(1010,'2026-10-15 06:49:53','2026-10-15 06:49:53','/mms-api/member/profile','apiDesc.modifyProfile','member','Mms','POST',0),
(1011,'2026-10-15 06:49:53','2026-10-15 06:49:53','/mms-api/member/logout','apiDesc.logout','member','Mms','GET',0),
(1012,'2026-10-15 06:49:53','2026-10-15 06:49:53','/mms-api/member/bind/wechat','apiDesc.bindWechat','member','Mms','POST',0),
(1013,'2026-10-15 06:49:53','2026-10-15 06:49:53','/mms-api/member/login','apiDesc.login','member','Mms','POST',1),
(1014,'2026-10-15 06:49:53','2026-10-15 06:49:53','/mms-api/member/login_by_email','apiDesc.loginByEmail','member','Mms','POST',1),
(1015,'2026-10-15 06:49:53','2026-10-15 06:49:53','/mms-api/member/login_by_sms','apiDesc.loginBySms','member','Mms','POST',1),
(1016,'2026-10-15 06:49:53','2026-10-15 06:49:53','/mms-api/member/register','apiDesc.register','member','Mms','POST',1),
(1017,'2026-10-15 06:49:53','2026-10-15 06:49:53','/mms-api/member/register_by_email','apiDesc.registerByEmail','member','Mms','POST',1),
(1018,'2026-10-15 06:49:53','2026-10-15 06:49:53','/mms-api/member/register_by_sms','apiDesc.registerBySms','member','Mms','POST',1),
(1019,'2026-10-15 06:49:53','2026-10-15 06:49:53','/mms-api/member/reset_password_by_email','apiDesc.resetPasswordByEmail','member','Mms','POST',1),
(1020,'2026-10-15 06:49:53','2026-10-15 06:49:53','/mms-api/member/reset_password_by_sms','apiDesc.resetPasswordBySms','member','Mms','POST',1),
(1021,'2026-10-15 06:49:53','2026-10-15 06:49:53','/mms-api/member/vip','apiDesc.testVip','member','Mms','GET',0),
(1022,'2026-10-15 06:49:53','2026-10-15 06:49:53','/mms-api/member_rank/create','apiDesc.createMemberRank','member_rank','Mms','POST',0),
(1023,'2026-10-15 06:49:53','2026-10-15 06:49:53','/mms-api/member_rank/update','apiDesc.updateMemberRank','member_rank','Mms','POST',0),
(1024,'2026-10-15 06:49:53','2026-10-15 06:49:53','/mms-api/member_rank/delete','apiDesc.deleteMemberRank','member_rank','Mms','POST',0),
(1025,'2026-10-15 06:49:53','2026-10-15 06:49:53','/mms-api/member_rank/list','apiDesc.getMemberRankList','member_rank','Mms','POST',0),
(1026,'2026-10-15 06:49:53','2026-10-15 06:49:53','/mms-api/member_rank','apiDesc.getMemberRankById','member_rank','Mms','POST',0),
(1027,'2026-10-15 06:49:53','2026-10-15 06:49:53','/mms-api/oauth/login','apiDesc.oauthLogin','oauth_provider','Mms','POST',1),
(1028,'2026-10-15 06:49:53','2026-10-15 06:49:53','/mms-api/oauth/login/callback','apiDesc.oauthCallback','oauth_provider','Mms','GET',1),
(1029,'2026-10-15 06:49:53','2026-10-15 06:49:53','/mms-api/oauth/login/wechat/mini_program','apiDesc.wechatMiniProgramLogin','oauth_provider','Mms','POST',1),
(1030,'2026-10-15 06:49:53','2026-10-15 06:49:53','/mms-api/oauth_provider/create','apiDesc.createOauthProvider','oauth_provider','Mms','POST',0),
(1031,'2026-10-15 06:49:53','2026-10-15 06:49:53','/mms-api/oauth_provider/update','apiDesc.updateOauthProvider','oauth_provider','Mms','POST',0),
(1032,'2026-10-15 06:49:53','2026-10-15 06:49:53','/mms-api/oauth_provider/delete','apiDesc.deleteOauthProvider','oauth_provider','Mms','POST',0),
(1033,'2026-10-15 06:49:53','2026-10-15 06:49:53','/mms-api/oauth_provider/list','apiDesc.getOauthProviderList','oauth_provider','Mms','POST',0),
(1034,'2026-10-15 06:49:53','2026-10-15 06:49:53','/mms-api/oauth_provider','apiDesc.getOauthProviderById','oauth_provider','Mms','POST',0),
(1035,'2026-10-15 06:49:53','2026-10-15 06:49:53','/mms-api/token/create','apiDesc.createToken','token','Mms','POST',0),
(1036,'2026-10-15 06:49:53','2026-10-15 06:49:53','/mms-api/token/update','apiDesc.updateToken','token','Mms','POST',0),
(1037,'2026-10-15 06:49:53','2026-10-15 06:49:53','/mms-api/token/delete','apiDesc.deleteToken','token','Mms','POST',0),
(1038,'2026-10-15 06:49:53','2026-10-15 06:49:53','/mms-api/token/list','apiDesc.getTokenList','token','Mms','POST',0),
(1039,'2026-10-15 06:49:53','2026-10-15 06:49:53','/mms-api/token','apiDesc.getTokenById','token','Mms','POST',0),
(1040,'2026-10-15 06:49:53','2026-10-15 06:49:53','/mms-api/token/logout','apiDesc.forceLoggingOut','token','Mms','POST',0);

--
-- Seed data: casbin policies for role 001
--

INSERT IGNORE INTO `casbin_rules` VALUES
(6001,'p','001','/mms-api/init/database','GET','','',''),
(6002,'p','001','/mms-api/captcha','GET','','',''),
(6003,'p','001','/mms-api/captcha/email','POST','','',''),
(6004,'p','001','/mms-api/captcha/sms','POST','','',''),
(6005,'p','001','/mms-api/member/create','POST','','',''),
(6006,'p','001','/mms-api/member/update','POST','','',''),
(6007,'p','001','/mms-api/member/delete','POST','','',''),
(6008,'p','001','/mms-api/member/list','POST','','',''),
(6009,'p','001','/mms-api/member','POST','','',''),
(6010,'p','001','/mms-api/member/profile','POST','','',''),
(6011,'p','001','/mms-api/member/logout','GET','','',''),
(6012,'p','001','/mms-api/member/bind/wechat','POST','','',''),
(6013,'p','001','/mms-api/member/login','POST','','',''),
(6014,'p','001','/mms-api/member/login_by_email','POST','','',''),
(6015,'p','001','/mms-api/member/login_by_sms','POST','','',''),
(6016,'p','001','/mms-api/member/register','POST','','',''),
(6017,'p','001','/mms-api/member/register_by_email','POST','','',''),
(6018,'p','001','/mms-api/member/register_by_sms','POST','','',''),
(6019,'p','001','/mms-api/member/reset_password_by_email','POST','','',''),
(6020,'p','001','/mms-api/member/reset_password_by_sms','POST','','',''),
(6021,'p','001','/mms-api/member/vip','GET','','',''),
(6022,'p','001','/mms-api/member_rank/create','POST','','',''),
(6023,'p','001','/mms-api/member_rank/update','POST','','',''),
(6024,'p','001','/mms-api/member_rank/delete','POST','','',''),
(6025,'p','001','/mms-api/member_rank/list','POST','','',''),
(6026,'p','001','/mms-api/member_rank','POST','','',''),
(6027,'p','001','/mms-api/oauth/login','POST','','',''),
(6028,'p','001','/mms-api/oauth/login/callback','GET','','',''),
(6029,'p','001','/mms-api/oauth/login/wechat/mini_program','POST','','',''),
(6030,'p','001','/mms-api/oauth_provider/create','POST','','',''),
(6031,'p','001','/mms-api/oauth_provider/update','POST','','',''),
(6032,'p','001','/mms-api/oauth_provider/delete','POST','','',''),
(6033,'p','001','/mms-api/oauth_provider/list','POST','','',''),
(6034,'p','001','/mms-api/oauth_provider','POST','','',''),
(6035,'p','001','/mms-api/token/create','POST','','',''),
(6036,'p','001','/mms-api/token/update','POST','','',''),
(6037,'p','001','/mms-api/token/delete','POST','','',''),
(6038,'p','001','/mms-api/token/list','POST','','',''),
(6039,'p','001','/mms-api/token','POST','','',''),
(6040,'p','001','/mms-api/token/logout','POST','','','');

--
-- Seed data: member management menus
--

INSERT IGNORE INTO `sys_menus` VALUES
(101,'2026-10-15 06:49:53','2026-10-15 06:49:53',1,1,0,'/member_dir','MemberManagementDirectory','','LAYOUT',0,'Mms',NULL,'route.memberManagement','ic:round-person-outline',0,0,0,0,'',0,0,0,20,'',1000000),
(102,'2026-10-15 06:49:53','2026-10-15 06:49:53',1,2,1,'/member','MemberManagement','','/mms/member/index',0,'Mms',NULL,'route.memberManagement','ic:round-person-outline',0,0,0,0,'',0,0,0,20,'',101),
(103,'2026-10-15 06:49:53','2026-10-15 06:49:53',2,2,1,'/member_rank','MemberRankManagement','','/mms/memberRank/index',0,'Mms',NULL,'route.memberRankManagement','ant-design:crown-outlined',0,0,0,0,'',0,0,0,20,'',101),
(104,'2026-10-15 06:49:53','2026-10-15 06:49:53',3,2,1,'/member_token','MemberTokenManagement','','/mms/token/index',0,'Mms',NULL,'route.tokenManagement','ant-design:lock-outlined',0,0,0,0,'',0,0,0,20,'',101),
(105,'2026-10-15 06:49:53','2026-10-15 06:49:53',3,2,1,'/member_oauth','MemberOauthManagement','','/mms/oauth/index',0,'Mms',NULL,'route.oauthManagement','ant-design:unlock-filled',0,0,0,0,'',0,0,0,20,'',101);

-- role 1 is authorized for the member management menus. The menu ids are resolved
-- by path because sys_menus.path is unique and core may already have seeded them.
INSERT IGNORE INTO `role_menus` (`role_id`,`menu_id`)
SELECT 1, `id` FROM `sys_menus`
WHERE `path` IN ('/member_dir','/member','/member_rank','/member_token','/member_oauth');

/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
