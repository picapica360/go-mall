/*
    mall 数据库脚本
    MySQL - 5.7+ : Database - mall
    Date: 2020-04-07
*/

CREATE DATABASE IF NOT EXISTS `mall` default character set utf8 collate utf8_general_ci;

USE mall;


-- demo
CREATE TABLE rate_limit_rules(
    id INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '主键id',
    area INT(11) NOT NULL DEFAULT 0 COMMENT '业务类型',
    limit_type tinyint(4) NOT NULL DEFAULT 0 COMMENT '0: default, 1: strict',
    limit_scope tinyint(4) NOT NULL DEFAULT 0 COMMENT '0: local, 1: global',
    dur_sec int(11) NOT NULL DEFAULT 0 COMMENT '持续时间',
    allowed_counts int(11) NOT NULL DEFAULT 0 COMMENT '允许发送次数',
    state tinyint(4) NOT NULL DEFAULT 0 COMMENT '0:default, 1:deleted',
    `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `mtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    KEY `ix_mtime` (`mtime`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='频率规则表';
CREATE UNIQUE INDEX uk_area_limit_type_limit_scope ON rate_limit_rules (area, limit_type, limit_scope);


/***** 用户模块 *****/
-- ----------------------------
-- Table ums_admin 后台用户表
-- ----------------------------
DROP TABLE IF EXISTS ums_admin;
CREATE TABLE ums_admin (
  id bigint(20) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '主键id',
  username varchar(64),
  password varchar(64),
  icon varchar(500) COMMENT '头像',
  email varchar(100) COMMENT '邮箱',
  nick_name varchar(200) COMMENT '昵称',
  note varchar(500) DEFAULT NULL COMMENT '备注信息',
  create_at datetime COMMENT '创建时间',
  login_at datetime COMMENT '最后登录时间',
  status int(1) DEFAULT '1' COMMENT '帐号启用状态：0->禁用；1->启用',
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='后台用户表';

-- ----------------------------
-- Table ums_admin_login_log 后台用户登录日志表
-- ----------------------------
DROP TABLE IF EXISTS ums_admin_login_log;
CREATE TABLE ums_admin_login_log (
  id bigint(20) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '主键id',
  admin_id bigint(20) COMMENT '用户 ID',
  ip varchar(64) COMMENT '用户登录的 IP 地址, IPv4 或 IPv6',
  address varchar(100) COMMENT '',
  user_agent varchar(100) COMMENT '浏览器登录类型',
  create_at datetime COMMENT '创建时间',
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='后台用户登录日志表';

-- ----------------------------
-- Table ums_admin_permission_relation 后台用户和权限关系表
-- ----------------------------
DROP TABLE IF EXISTS ums_admin_permission_relation;
CREATE TABLE ums_admin_permission_relation (
  id bigint(20) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '主键id',
  admin_id bigint(20) NOT NULL COMMENT '用户 id',
  permission_id bigint(20) NOT NULL COMMENT '权限 id',
) ENGINE=InnoDB AUTO_INCREMENT=1 CHARSET=utf8 COMMENT='后台用户和权限关系表(除角色中定义的权限以外的加减权限)';

-- ----------------------------
-- Table ums_admin_role_relation 后台用户和角色关系表
-- ----------------------------
DROP TABLE IF EXISTS ums_admin_role_relation;
CREATE TABLE ums_admin_role_relation (
  id bigint(20) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '主键id',
  admin_id bigint(20) NOT NULL COMMENT '用户id',
  role_id bigint(20) NOT NULL COMMENT '角色id',
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='后台用户和角色关系表';

-- ----------------------------
-- Table ums_member 会员表
-- ----------------------------
DROP TABLE IF EXISTS ums_member;
CREATE TABLE ums_member (
  id bigint(20) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '主键id',
  member_level_id bigint(20) DEFAULT NULL,
  username varchar(64) NOT NULL COMMENT '用户名',
  password varchar(64) NOT NULL COMMENT '密码',
  nickname varchar(64) NOT NULL COMMENT '昵称',
  phone varchar(64) NOT NULL COMMENT '手机号码',
  status int(1) DEFAULT 1 COMMENT '帐号启用状态:0->禁用；1->启用',
  create_at datetime NOT NULL COMMENT '注册时间',
  icon varchar(500) COMMENT '头像',
  gender int(1) DEFAULT 0 COMMENT '性别：0->未知；1->男；2->女',
  birthday date COMMENT '生日',
  city varchar(64) COMMENT '所做城市',
  job varchar(100) COMMENT '职业',
  personalized_signature varchar(200) COMMENT '个性签名',
  source_type int(1) COMMENT '用户来源',
  integration int(11) COMMENT '积分',
  growth int(11) COMMENT '成长值',
  luckey_count int(11) COMMENT '剩余抽奖次数',
  history_integration int(11) COMMENT '历史积分数量',
  UNIQUE KEY `idx_username` (`username`),
  UNIQUE KEY `idx_phone` (`phone`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='会员表';

-- ----------------------------
-- Table ums_member_level 会员等级表
-- ----------------------------
DROP TABLE IF EXISTS ums_member_level;
CREATE TABLE ums_member_level (
  id bigint(20) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '主键id',
  name varchar(100) NOT NULL COMMENT '会员级别名称',
  growth_point int(11) DEFAULT NULL,
  default_status int(1) DEFAULT NULL COMMENT '是否为默认等级：0->不是；1->是',
  free_freight_point decimal(10,2) DEFAULT NULL COMMENT '免运费标准',
  comment_growth_point int(11) DEFAULT NULL COMMENT '每次评价获取的成长值',
  priviledge_free_freight int(1) DEFAULT NULL COMMENT '是否有免邮特权',
  priviledge_sign_in int(1) DEFAULT NULL COMMENT '是否有签到特权',
  priviledge_comment int(1) DEFAULT NULL COMMENT '是否有评论获奖励特权',
  priviledge_promotion int(1) DEFAULT NULL COMMENT '是否有专享活动特权',
  priviledge_member_price int(1) DEFAULT NULL COMMENT '是否有会员价格特权',
  priviledge_birthday int(1) DEFAULT NULL COMMENT '是否有生日特权',
  note varchar(200) DEFAULT NULL,
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='会员等级表';

-- ----------------------------
-- Table ums_member_login_log 会员登录记录
-- ----------------------------
DROP TABLE IF EXISTS ums_member_login_log;
CREATE TABLE ums_member_login_log (
  id bigint(20) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '主键id',
  member_id bigint(20) NOT NULL COMMENT '会员 id',
  create_at datetime NOT NULL COMMENT '登录时间',
  ip varchar(64) NOT NULL COMMENT '登录的 ip',
  city varchar(64) COMMENT '登录时所作的城市',
  province varchar(64) COMMENT '登录时所作的身份',
  login_type int(1) NOT NULL COMMENT '登录类型：0->PC；1->android;2->ios;3->小程序',
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='会员登录记录';

-- ----------------------------
-- Table ums_member_receive_address 会员收货地址表
-- ----------------------------
DROP TABLE IF EXISTS ums_member_receive_address;
CREATE TABLE ums_member_receive_address (
  id bigint(20) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '主键id',
  member_id bigint(20) NOT NULL COMMENT '会员 id',
  name varchar(100) NOT NULL COMMENT '收货人名称',
  phone_number varchar(16) NOT NULL COMMENT '收货人电话',
  default_status int(1) NOT NULL COMMENT '是否为默认: 0->不默认;1->默认',
  post_code varchar(100) DEFAULT NULL COMMENT '邮政编码',
  province varchar(100) NOT NULL COMMENT '省份/直辖市',
  city varchar(100) NOT NULL COMMENT '城市',
  region varchar(100) NOT NULL COMMENT '区',
  detail_address varchar(128) NOT NULL COMMENT '详细地址(街道)',
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='会员收货地址表';

-- ----------------------------
-- Table ums_member_statistics_info 会员统计信息
-- ----------------------------
DROP TABLE IF EXISTS ums_member_statistics_info;
CREATE TABLE ums_member_statistics_info (
  id bigint(20) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '主键id',
  member_id NOT NULL COMMENT '会员 id',
  consume_amount decimal(10,2) NOT NULL DEFAULT 0 COMMENT '累计消费金额',
  order_count int(11) NOT NULL DEFAULT 0 COMMENT '订单数量',
  coupon_count int(11) NOT NULL DEFAULT 0 COMMENT '优惠券数量',
  comment_count int(11) NOT NULL DEFAULT 0 COMMENT '评价数',
  return_order_count int(11) NOT NULL DEFAULT 0 COMMENT '退货数量',
  login_count int(11) NOT NULL DEFAULT 0 COMMENT '登录次数',
  attend_count int(11) NOT NULL DEFAULT 0 COMMENT '关注数量',
  fans_count int(11) NOT NULL DEFAULT 0 COMMENT '粉丝数量',
  collect_product_count int(11) NOT NULL DEFAULT 0,
  collect_subject_count int(11) NOT NULL DEFAULT 0,
  collect_topic_count int(11) NOT NULL DEFAULT 0,
  collect_comment_count int(11) NOT NULL DEFAULT 0,
  invite_friend_count int(11) NOT NULL DEFAULT 0,
  recent_order_time datetime COMMENT '最后一次下订单时间',
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='会员统计信息';

-- ----------------------------
-- Table ums_menu 后台菜单表
-- ----------------------------
DROP TABLE IF EXISTS ums_menu;
CREATE TABLE ums_menu (
  id bigint(20) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '主键id',
  parent_id bigint(20) DEFAULT NULL COMMENT '父级ID',
  create_at datetime NOT NULL COMMENT '创建时间',
  title varchar(100) NOT NULL COMMENT '菜单名称',
  level int(4) NOT NULL COMMENT '菜单级数',
  sort int(4) NOT NULL COMMENT '菜单排序',
  name varchar(100) COMMENT '前端名称',
  icon varchar(200) COMMENT '前端图标',
  hidden int(1) COMMENT '前端隐藏: 0->不隐藏;1->隐藏',
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='后台菜单表';

-- ----------------------------
-- Table ums_permission 后台用户权限表
-- ----------------------------
DROP TABLE IF EXISTS ums_permission;
CREATE TABLE ums_permission (
  id bigint(20) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '主键id',
  pid bigint(20) DEFAULT NULL COMMENT '父级权限id',
  name varchar(100) DEFAULT NULL COMMENT '名称',
  value varchar(200) DEFAULT NULL COMMENT '权限值',
  icon varchar(500) DEFAULT NULL COMMENT '图标',
  type int(1) NOT NULL COMMENT '权限类型：0->目录；1->菜单；2->按钮（接口绑定权限）',
  uri varchar(200) COMMENT '前端资源路径',
  status int(1) NOT NULL DEFAULT 1 COMMENT '启用状态；0->禁用；1->启用',
  create_at datetime NOT NULL COMMENT '创建时间',
  sort int(11) NOT NULL DEFAULT 0 COMMENT '排序',
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='后台用户权限表';

-- ----------------------------
-- Table ums_role 后台用户角色表
-- ----------------------------
DROP TABLE IF EXISTS ums_role;
CREATE TABLE ums_role (
  id bigint(20) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '主键id',
  name varchar(100) NOT NULL COMMENT '名称',
  description varchar(500) COMMENT '描述',
  admin_count int(11) NOT NULL COMMENT '后台用户数量',
  create_at datetime NOT NULL COMMENT '创建时间',
  status int(1) DEFAULT 1 COMMENT '启用状态：0->禁用；1->启用',
  sort int(11) NOT NULL DEFAULT 0,
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8 COMMENT='后台用户角色表';

-- ----------------------------
-- Table ums_role_menu_relation 后台角色菜单关系表
-- ----------------------------
DROP TABLE IF EXISTS ums_role_menu_relation;
CREATE TABLE ums_role_menu_relation (
  id bigint(20) UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '主键id',
  role_id bigint(20) DEFAULT NULL COMMENT '角色ID',
  menu_id bigint(20) DEFAULT NULL COMMENT '菜单ID',
) ENGINE=InnoDB AUTO_INCREMENT=96 DEFAULT CHARSET=utf8 COMMENT='后台角色菜单关系表';

