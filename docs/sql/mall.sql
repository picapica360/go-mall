/*
    mall 数据库脚本
    MySQL - 5.7+ : Database - mall
    Date: 2020-04-07
*/

CREATE DATABASE IF NOT EXISTS `mall` default character set utf8 collate utf8_general_ci;

USE mall;


/****************************************************************************************************/
/********************************************** 用户模块 ********************************************/
/****************************************************************************************************/

-- ----------------------------
-- Table ums_admin 后台用户表(管理员)
-- ----------------------------
DROP TABLE IF EXISTS `ums_admin`;
CREATE TABLE `ums_admin` (
  `id` int(10) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '主键id',
  `username` varchar(64) NOT NULL COMMENT '用户名',
  `password` varchar(64) NOT NULL COMMENT '密码',
  `icon` varchar(256) COMMENT '头像',
  `email` varchar(100) COMMENT '邮箱',
  `nickname` varchar(32) COMMENT '昵称',
  `note` varchar(256) COMMENT '备注信息',
  `enabled` tinyint(1) NOT NULL DEFAULT '1' COMMENT '帐号启用状态：0->禁用；1->启用'
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `login_at` datetime NOT NULL COMMENT '最后登录时间',
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='后台用户表';

-- ----------------------------
-- Table ums_admin_login_log 后台用户登录日志表
-- ----------------------------
DROP TABLE IF EXISTS `ums_admin_login_log`;
CREATE TABLE `ums_admin_login_log` (
  `id` bigint(10) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '主键id',
  `admin_id` bigint(10) COMMENT '用户id, ref->ums_admin',
  `ip` varchar(64) NOT NULL COMMENT '用户登录的 IP 地址, IPv4 或 IPv6',
  `address` varchar(100) COMMENT '',
  `user_agent` varchar(100) COMMENT '浏览器登录类型',
  `created_at` datetime NOT NULL COMMENT '创建时间'
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='后台用户登录日志表';

-- ----------------------------
-- Table ums_admin_permission_relation 后台用户和权限关系表
-- ----------------------------
DROP TABLE IF EXISTS `ums_admin_permission_relation`;
CREATE TABLE `ums_admin_permission_relation` (
  `id` int(10) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '主键id',
  `admin_id` int(10) NOT NULL COMMENT '用户id, ref->ums_admin',
  `permission_id` int(10) NOT NULL COMMENT '权限id, ref->ums_permission'
) ENGINE=InnoDB AUTO_INCREMENT=1 CHARSET=utf8 COMMENT='后台用户和权限关系表(除角色中定义的权限以外的加减权限)';

-- ----------------------------
-- Table ums_admin_role_relation 后台用户和角色关系表
-- ----------------------------
DROP TABLE IF EXISTS `ums_admin_role_relation`;
CREATE TABLE `ums_admin_role_relation` (
  `id` int(10) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '主键id',
  `admin_id` int(10) NOT NULL COMMENT '用户id, ref->ums_admin',
  `role_id` int(10) NOT NULL COMMENT '角色id, ref->ums_role'
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='后台用户和角色关系表';

-- ----------------------------
-- Table ums_member 会员表
-- ----------------------------
DROP TABLE IF EXISTS `ums_member`;
CREATE TABLE `ums_member` (
  `id` bigint(10) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '主键id',
  `member_level_id` int(10) NOT NULL COMMENT '会员等级, ref->ums_member_level',
  `username` varchar(64) NOT NULL COMMENT '用户名',
  `password` varchar(64) NOT NULL COMMENT '密码',
  `nickname` varchar(64) NOT NULL COMMENT '昵称',
  `phone` varchar(64) NOT NULL COMMENT '手机号码',
  `status` tinyint(1) DEFAULT 1 COMMENT '帐号启用状态:0->禁用；1->启用',
  `created_at` datetime NOT NULL COMMENT '注册时间',
  `icon` varchar(128) COMMENT '头像',
  `gender` tinyint(1) NOT NULL DEFAULT 0 COMMENT '性别：0->未知；1->男；2->女',
  `birthday` date COMMENT '生日',
  `city` varchar(64) COMMENT '所在城市',
  `job` varchar(100) COMMENT '职业',
  `signature` varchar(128) COMMENT '个性签名',
  `source_type` int(1) COMMENT '用户来源',
  `integration` int(11) NOT NULL DEFAULT 0 COMMENT '积分',
  `growth` int(11) NOT NULL DEFAULT 0 COMMENT '成长值',
  `luckey_count` int(11) NOT NULL DEFAULT 0 COMMENT '剩余抽奖次数',
  `history_integration` int(11) NOT NULL DEFAULT 0 COMMENT '历史积分数量',
  UNIQUE KEY `idx_ums_member_username` (`username`),
  UNIQUE KEY `idx_ums_member_phone` (`phone`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='会员表';

-- ----------------------------
-- Table ums_member_level 会员等级表
-- ----------------------------
DROP TABLE IF EXISTS `ums_member_level`;
CREATE TABLE `ums_member_level` (
  `id` int(10) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '主键id',
  `name` varchar(16) NOT NULL COMMENT '会员级别名称',
  `growth_point` int(11) DEFAULT NULL,
  `status` tinyint(1) COMMENT '是否为默认等级：0->不是；1->是',
  `free_freight_point` decimal(10,2) COMMENT '免运费标准',
  `comment_growth_point` int(11) COMMENT '每次评价获取的成长值',
  `priviledge_free_freight` int(1) COMMENT '是否有免邮特权',
  `priviledge_sign_in` int(1) COMMENT '是否有签到特权',
  `priviledge_comment` int(1) COMMENT '是否有评论获奖励特权',
  `priviledge_promotion` int(1) COMMENT '是否有专享活动特权',
  `priviledge_member_price` int(1) COMMENT '是否有会员价格特权',
  `priviledge_birthday` int(1) COMMENT '是否有生日特权',
  `note` varchar(200) DEFAULT NULL
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='会员等级表';

-- ----------------------------
-- Table ums_member_login_log 会员登录记录
-- ----------------------------
DROP TABLE IF EXISTS `ums_member_login_log`;
CREATE TABLE `ums_member_login_log` (
  `id` bigint(10) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '主键id',
  `member_id` bigint(10) NOT NULL COMMENT '会员 id, ref->ums_member',
  `created_at` datetime NOT NULL COMMENT '登录时间',
  `ip` varchar(64) NOT NULL COMMENT '登录的 ip',
  `city` varchar(64) COMMENT '登录时所在的城市',
  `province` varchar(64) COMMENT '登录时所作的身份',
  `login_type` int(1) NOT NULL COMMENT '登录类型：0->PC;1->android;2->ios;3->小程序'
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='会员登录记录';

-- ----------------------------
-- Table ums_member_receive_address 会员收货地址表
-- ----------------------------
DROP TABLE IF EXISTS `ums_member_receive_address`;
CREATE TABLE `ums_member_receive_address` (
  `id` bigint(10) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '主键id',
  `member_id` bigint(10) NOT NULL COMMENT '会员 id, ref->ums_member',
  `name` varchar(100) NOT NULL COMMENT '收货人名称',
  `phone` varchar(64) NOT NULL COMMENT '电话',
  `postcode` varchar(100) COMMENT '邮政编码',
  `province` varchar(20) NOT NULL COMMENT '省份/直辖市',
  `city` varchar(20) NOT NULL COMMENT '城市',
  `region` varchar(20) DEFAULT NULL COMMENT '区/县',
  `address` varchar(128) NOT NULL COMMENT '详细地址(街道)',
  `is_default` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否是默认 1默认 0否',
  `created_date` datetime NOT NULL COMMENT '创建日期'
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='会员收货地址表';

-- ----------------------------
-- Table ums_member_stats 会员信息统计
-- ----------------------------
DROP TABLE IF EXISTS `ums_member_stats`;
CREATE TABLE `ums_member_stats` (
  `id` bigint(10) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '主键id',
  `member_id` NOT NULL COMMENT '会员 id, ref->ums_member',
  `consume_amount` decimal(10,2) NOT NULL DEFAULT 0 COMMENT '累计消费金额',
  `order_count` int(11) NOT NULL DEFAULT 0 COMMENT '订单数量',
  `coupon_count` int(11) NOT NULL DEFAULT 0 COMMENT '优惠券数量',
  `comment_count` int(11) NOT NULL DEFAULT 0 COMMENT '评价数',
  `return_order_count` int(11) NOT NULL DEFAULT 0 COMMENT '退货数量',
  `login_count` int(11) NOT NULL DEFAULT 0 COMMENT '登录次数',
  `attend_count` int(11) NOT NULL DEFAULT 0 COMMENT '关注数量',
  `fans_count` int(11) NOT NULL DEFAULT 0 COMMENT '粉丝数量',
  `collect_product_count` int(11) NOT NULL DEFAULT 0,
  `collect_subject_count` int(11) NOT NULL DEFAULT 0,
  `collect_topic_count` int(11) NOT NULL DEFAULT 0,
  `collect_comment_count` int(11) NOT NULL DEFAULT 0,
  `invite_friend_count` int(11) NOT NULL DEFAULT 0,
  `recent_order_time` datetime COMMENT '最后一次下订单时间'
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='会员统计信息';


-- ----------------------------
-- Table ums_role 后台用户角色表
-- ----------------------------
DROP TABLE IF EXISTS `ums_role`;
CREATE TABLE `ums_role` (
  `id` int(10) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '主键id',
  `name` varchar(100) NOT NULL COMMENT '名称',
  `desc` varchar(256) COMMENT '描述',
  `enabled` tinyint(1) NOT NULL DEFAULT 1 COMMENT '启用状态：0->禁用；1->启用',
  `sort` int(4) NOT NULL COMMENT '排序',
  `created_at` datetime NOT NULL COMMENT '创建时间'
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='后台用户角色表';

-- ----------------------------
-- Table ums_menu 后台菜单表
-- ----------------------------
DROP TABLE IF EXISTS `ums_menu`;
CREATE TABLE ums_menu (
  `id` int(10) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '主键id',
  `pid` int(10) NOT NULL COMMENT '父级ID, 0 表示最top',
  `title` varchar(100) NOT NULL COMMENT '菜单名称',
  `level` int(4) NOT NULL COMMENT '菜单级数',
  `sort` int(4) NOT NULL COMMENT '菜单排序',
  `name` varchar(100) NOT NULL COMMENT '前端名称',
  `icon` varchar(200) COMMENT '前端图标',
  `is_hidden` tinyint(1) NOT NULL DEFAULT 0 COMMENT '前端隐藏: 0->不隐藏;1->隐藏',
  `created_at` datetime NOT NULL COMMENT '创建时间'
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='后台菜单表';

-- ----------------------------
-- Table ums_permission 后台用户权限表
-- ----------------------------
DROP TABLE IF EXISTS `ums_permission`;
CREATE TABLE `ums_permission` (
  `id` int(10) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '主键id',
  `pid` int(10) NOT NULL COMMENT '父级权限id, 0 表示最top',
  `name` varchar(100) NOT NULL COMMENT '名称',
  `value` varchar(128) COMMENT '权限值',
  `icon` varchar(128) COMMENT '图标',
  `kind` int(1) NOT NULL COMMENT '权限类型：0->目录；1->菜单；2->按钮（接口绑定权限）',
  `uri` varchar(128) COMMENT '前端资源路径',
  `enabled` tinyint(1) NOT NULL DEFAULT 1 COMMENT '启用状态；0->禁用；1->启用',
  `sort` int(4) NOT NULL DEFAULT 0 COMMENT '排序',
  `created_at` datetime NOT NULL COMMENT '创建时间'
  `updated_at` datetime NOT NULL COMMENT '更新时间'
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='后台用户权限表';

-- ----------------------------
-- Table ums_role 后台用户角色表
-- ----------------------------
DROP TABLE IF EXISTS `ums_resource`;
CREATE TABLE `ums_resource` (
  `id` int(10) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '主键id',
  `category_id` int(10) NOT NULL COMMENT '资源分类ID, ref->ums_resource_category',
  `name` varchar(100) NOT NULL COMMENT '名称',
  `url` varchar(200) NOT NULL COMMENT '资源URL',
  `desc` varchar(256) COMMENT '描述',
  `created_at` datetime NOT NULL COMMENT '创建时间'
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='后台资源表';

-- ----------------------------
-- Table ums_resource_category 资源分类表
-- ----------------------------
DROP TABLE IF EXISTS `ums_resource_category`;
CREATE TABLE `ums_resource_category` (
  `id` int(10) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '主键id',
  `name` varchar(100) NOT NULL COMMENT '分类名称',
  `sort` int(4) NOT NULL DEFAULT 0 COMMENT '排序',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8 COMMENT='资源分类表';

-- ----------------------------
-- Table ums_role_menu_relation 后台角色菜单关系表
-- ----------------------------
DROP TABLE IF EXISTS `ums_role_menu_relation`;
CREATE TABLE `ums_role_menu_relation` (
  id int(10) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '主键id',
  `role_id` int(10) NOT NULL COMMENT '角色ID, ref->ums_role',
  `menu_id` int(10) NOT NULL COMMENT '菜单ID, ref->ums_menu'
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='后台角色菜单关系表';

-- ----------------------------
-- Table ums_role_permission_relation 后台用户角色和权限关系表
-- ----------------------------
DROP TABLE IF EXISTS `ums_role_permission_relation`;
CREATE TABLE `ums_role_permission_relation` (
  `id` int(10) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '主键id',
  `role_id` gint(10) NOT NULL COMMENT '角色ID, ref->ums_role',
  `permission_id` int(10) NOT NULL COMMENT '角色ID, ref->ums_permission',
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='后台用户角色和权限关系表';

-- ----------------------------
-- Table structure for ums_role_resource_relation
-- ----------------------------
DROP TABLE IF EXISTS `ums_role_resource_relation`;
CREATE TABLE `ums_role_resource_relation` (
  `id` int(10) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '主键id',
  `role_id` int(10) NOT NULL COMMENT '角色ID, ref->ums_role',
  `resource_id` int(10) NOT NULL COMMENT '资源ID, ref->ums_resource',
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='后台角色资源关系表';



/****************************************************************************************************/
/********************************************** 商品模块 ********************************************/
/****************************************************************************************************/

-- ----------------------------
-- Table pms_product 商品信息表
-- ----------------------------
DROP TABLE IF EXISTS `pms_product`;
CREATE TABLE `pms_product` (
  `id` bigint(10) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '主键id',
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='商品信息表';
