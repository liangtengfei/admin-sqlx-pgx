-- ----------------------------
-- 1、系统用户表
-- ----------------------------
-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS ago_user_id_seq;

-- Table Definition
CREATE TABLE "public"."ago_user"
(
    "id"          int8         NOT NULL DEFAULT nextval('ago_user_id_seq'::regclass),
    "dept_id"     int8         NOT NULL DEFAULT 0,
    "user_name"   varchar(200) NOT NULL DEFAULT '':: character varying,
    "real_name"   varchar(200) NOT NULL DEFAULT '':: character varying,
    "mobile"      varchar(100) NOT NULL DEFAULT '':: character varying,
    "email"       varchar(100) NOT NULL DEFAULT '':: character varying,
    "password"    varchar(200) NOT NULL DEFAULT '':: character varying,
    "sex"         varchar      NOT NULL DEFAULT '0':: character varying,
    "avatar"      varchar      NOT NULL DEFAULT '':: character varying,
    "posts"       varchar      NOT NULL DEFAULT '':: character varying,
    "status"      varchar(10)  NOT NULL DEFAULT '0':: character varying,
    "del_flag"    varchar(10)  NOT NULL DEFAULT 'N':: character varying,
    "create_time" timestamp    NOT NULL DEFAULT now(),
    "update_time" timestamp    NOT NULL DEFAULT '0001-01-01 00:00:00':: timestamp without time zone,
    "create_by"   varchar      NOT NULL DEFAULT '':: character varying,
    "update_by"   varchar      NOT NULL DEFAULT '':: character varying,
    "remark"      varchar      NOT NULL DEFAULT '':: character varying,
    PRIMARY KEY ("id")
);

-- Column Comment
COMMENT ON COLUMN "public"."ago_user"."id" IS '唯一标识';
COMMENT ON COLUMN "public"."ago_user"."dept_id" IS '部门编号';
COMMENT ON COLUMN "public"."ago_user"."user_name" IS '登录名称';
COMMENT ON COLUMN "public"."ago_user"."real_name" IS '真实姓名';
COMMENT ON COLUMN "public"."ago_user"."mobile" IS '手机号码';
COMMENT ON COLUMN "public"."ago_user"."email" IS '用户邮箱';
COMMENT ON COLUMN "public"."ago_user"."password" IS '用户密码';
COMMENT ON COLUMN "public"."ago_user"."sex" IS '用户性别（0男 1女 2未知）';
COMMENT ON COLUMN "public"."ago_user"."avatar" IS '用户头像';
COMMENT ON COLUMN "public"."ago_user"."posts" IS '岗位编号数组';
COMMENT ON COLUMN "public"."ago_user"."status" IS '状态（默认0）';
COMMENT ON COLUMN "public"."ago_user"."del_flag" IS '删除标记';
COMMENT ON COLUMN "public"."ago_user"."create_time" IS '创建时间';
COMMENT ON COLUMN "public"."ago_user"."update_time" IS '更新时间';
COMMENT ON COLUMN "public"."ago_user"."create_by" IS '创建人员';
COMMENT ON COLUMN "public"."ago_user"."update_by" IS '更新人员';
COMMENT ON COLUMN "public"."ago_user"."remark" IS '备注';
COMMENT ON TABLE "public"."ago_user" IS '系统用户表';

-- 唯一索引
CREATE UNIQUE INDEX IF NOT EXISTS ago_user_username_unique ON "public"."ago_user"(user_name);

-- 初始化系统管理员
INSERT INTO "public"."ago_user" ("id", "dept_id", "user_name", "real_name", "mobile", "email", "password", "sex",
                                 "avatar", "posts", "status", "del_flag", "create_time", "update_time", "create_by",
                                 "update_by", "remark")
VALUES (1, 100, 'admin', '小伙纸', '18530030305', '', '$2a$10$gxepnF43.K4eugBqdyxC7OZKoYiQOcK6s4m2eZGSeSa8KRuhn/h9q',
        '0', '', '', '0', 'N', '2022-05-25 11:35:33', '0001-01-01 00:00:00', '系统初始化', '', '');

-- ----------------------------
-- 2、系统部门表
-- ----------------------------
-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS ago_dept_id_seq;

-- Table Definition
CREATE TABLE "public"."ago_dept"
(
    "id"          int8         NOT NULL DEFAULT nextval('ago_dept_id_seq'::regclass),
    "parent_id"   int8         NOT NULL DEFAULT 0,
    "dept_name"   varchar(200) NOT NULL DEFAULT '':: character varying,
    "dept_code"   varchar(200) NOT NULL DEFAULT '':: character varying,
    "ancestors"   varchar(200) NOT NULL DEFAULT '':: character varying,
    "order_num"   int4         NOT NULL DEFAULT 0,
    "status"      varchar(10)  NOT NULL DEFAULT '0':: character varying,
    "del_flag"    varchar(10)  NOT NULL DEFAULT 'N':: character varying,
    "create_time" timestamp    NOT NULL DEFAULT now(),
    "update_time" timestamp    NOT NULL DEFAULT '0001-01-01 00:00:00':: timestamp without time zone,
    "create_by"   varchar      NOT NULL DEFAULT '':: character varying,
    "update_by"   varchar      NOT NULL DEFAULT '':: character varying,
    "remark"      varchar      NOT NULL DEFAULT '':: character varying,
    PRIMARY KEY ("id")
);

-- Column Comment
COMMENT ON COLUMN "public"."ago_dept"."id" IS '唯一标识';
COMMENT ON COLUMN "public"."ago_dept"."parent_id" IS '父部门id';
COMMENT ON COLUMN "public"."ago_dept"."dept_name" IS '部门名称';
COMMENT ON COLUMN "public"."ago_dept"."dept_code" IS '部门编码';
COMMENT ON COLUMN "public"."ago_dept"."ancestors" IS '祖级列表';
COMMENT ON COLUMN "public"."ago_dept"."order_num" IS '显示顺序';
COMMENT ON COLUMN "public"."ago_dept"."status" IS '状态（默认0）';
COMMENT ON COLUMN "public"."ago_dept"."del_flag" IS '删除标记';
COMMENT ON COLUMN "public"."ago_dept"."create_time" IS '创建时间';
COMMENT ON COLUMN "public"."ago_dept"."update_time" IS '更新时间';
COMMENT ON COLUMN "public"."ago_dept"."create_by" IS '创建人员';
COMMENT ON COLUMN "public"."ago_dept"."update_by" IS '更新人员';
COMMENT ON COLUMN "public"."ago_dept"."remark" IS '备注';
COMMENT ON TABLE "public"."ago_dept" IS '系统部门表';

INSERT INTO "public"."ago_dept" ("id", "parent_id", "dept_name", "ancestors", "order_num", "status", "del_flag",
                                 "create_time", "update_time", "create_by", "update_by", "remark", "dept_code")
VALUES (100, 0, 'GOG有限公司', '0', 0, '0', 'N', '2022-05-25 11:38:22', '0001-01-01 00:00:00', '系统初始化', '', '系统初始化', '');
INSERT INTO "public"."ago_dept" ("id", "parent_id", "dept_name", "ancestors", "order_num", "status", "del_flag",
                                 "create_time", "update_time", "create_by", "update_by", "remark", "dept_code")
VALUES (101, 100, '塞尔达部', '0,100', 1, '0', 'N', '2022-05-25 11:38:22', '0001-01-01 00:00:00', '系统初始化', '', '系统初始化', '');
INSERT INTO "public"."ago_dept" ("id", "parent_id", "dept_name", "ancestors", "order_num", "status", "del_flag",
                                 "create_time", "update_time", "create_by", "update_by", "remark", "dept_code")
VALUES (102, 100, '林克部', '0,100', 2, '0', 'N', '2022-05-25 11:38:22', '0001-01-01 00:00:00', '系统初始化', '', '系统初始化', '');

-- ----------------------------
-- 3、系统角色表
-- ----------------------------
-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS ago_role_id_seq;

-- Table Definition
CREATE TABLE "public"."ago_role"
(
    "id"          int8        NOT NULL DEFAULT nextval('ago_role_id_seq'::regclass),
    "role_name"   varchar     NOT NULL DEFAULT '':: character varying,
    "role_key"    varchar     NOT NULL DEFAULT '':: character varying,
    "order_num"   int4        NOT NULL DEFAULT 0,
    "data_scope"  varchar     NOT NULL DEFAULT '1'::bpchar,
    "status"      varchar(10) NOT NULL DEFAULT '0':: character varying,
    "del_flag"    varchar(10) NOT NULL DEFAULT 'N':: character varying,
    "create_time" timestamp   NOT NULL DEFAULT now(),
    "update_time" timestamp   NOT NULL DEFAULT '0001-01-01 00:00:00':: timestamp without time zone,
    "create_by"   varchar     NOT NULL DEFAULT '':: character varying,
    "update_by"   varchar     NOT NULL DEFAULT '':: character varying,
    "remark"      varchar     NOT NULL DEFAULT '':: character varying,
    PRIMARY KEY ("id")
);

-- Column Comment
COMMENT ON COLUMN "public"."ago_role"."id" IS '唯一标识';
COMMENT ON COLUMN "public"."ago_role"."role_name" IS '角色名称';
COMMENT ON COLUMN "public"."ago_role"."role_key" IS '角色标识';
COMMENT ON COLUMN "public"."ago_role"."order_num" IS '排序编号';
COMMENT ON COLUMN "public"."ago_role"."data_scope" IS '数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）';
COMMENT ON COLUMN "public"."ago_role"."status" IS '状态（默认0）';
COMMENT ON COLUMN "public"."ago_role"."del_flag" IS '删除标记';
COMMENT ON COLUMN "public"."ago_role"."create_time" IS '创建时间';
COMMENT ON COLUMN "public"."ago_role"."update_time" IS '更新时间';
COMMENT ON COLUMN "public"."ago_role"."create_by" IS '创建人员';
COMMENT ON COLUMN "public"."ago_role"."update_by" IS '更新人员';
COMMENT ON COLUMN "public"."ago_role"."remark" IS '备注';
COMMENT ON TABLE "public"."ago_role" IS '系统角色表';

INSERT INTO "public"."ago_role" ("id", "role_name", "role_key", "order_num", "data_scope", "status", "del_flag",
                                 "create_time", "update_time", "create_by", "update_by", "remark")
VALUES (1, '超级管理员', 'admin', 0, '1', '0', 'N', '2022-05-25 11:45:59.960122', '0001-01-01 00:00:00', '系统初始化', '', '');
INSERT INTO "public"."ago_role" ("id", "role_name", "role_key", "order_num", "data_scope", "status", "del_flag",
                                 "create_time", "update_time", "create_by", "update_by", "remark")
VALUES (2, '普通角色', 'common', 0, '1', '0', 'N', '2022-05-25 11:45:59.960122', '0001-01-01 00:00:00', '系统初始化', '', '');

-- ----------------------------
-- 4、系统岗位表
-- ----------------------------
-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS ago_post_id_seq;

-- Table Definition
CREATE TABLE "public"."ago_post"
(
    "id"        int8    NOT NULL DEFAULT nextval('ago_post_id_seq'::regclass),
    "post_name" varchar NOT NULL DEFAULT '':: character varying,
    "order_num" int4    NOT NULL DEFAULT 0,
    "status"      varchar(10)  NOT NULL DEFAULT '0':: character varying,
    "del_flag"    varchar(10)  NOT NULL DEFAULT 'N':: character varying,
    "create_by"   varchar      NOT NULL DEFAULT '':: character varying,
    "create_time" timestamp    NOT NULL DEFAULT now(),
    "update_by"   varchar      NOT NULL DEFAULT '':: character varying,
    "update_time" timestamp    NOT NULL DEFAULT '0001-01-01 00:00:00':: timestamp without time zone,
    "remark"      varchar      NOT NULL DEFAULT '':: character varying,
    PRIMARY KEY ("id")
);

-- Column Comment
COMMENT ON COLUMN "public"."ago_post"."id" IS '唯一标识';
COMMENT ON COLUMN "public"."ago_post"."post_name" IS '岗位名称';
COMMENT ON COLUMN "public"."ago_post"."order_num" IS '排序编号';
COMMENT ON COLUMN "public"."ago_post"."status" IS '状态（默认0）';
COMMENT ON COLUMN "public"."ago_post"."del_flag" IS '删除标记';
COMMENT ON COLUMN "public"."ago_post"."create_time" IS '创建时间';
COMMENT ON COLUMN "public"."ago_post"."update_time" IS '更新时间';
COMMENT ON COLUMN "public"."ago_post"."create_by" IS '创建人员';
COMMENT ON COLUMN "public"."ago_post"."update_by" IS '更新人员';
COMMENT ON COLUMN "public"."ago_post"."remark" IS '备注';
COMMENT ON TABLE "public"."ago_post" IS '系统岗位表';


-- ----------------------------
-- 5、菜单权限表
-- ----------------------------
-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS ago_menu_id_seq;

-- Table Definition
CREATE TABLE "public"."ago_menu"
(
    "id"          int8        NOT NULL DEFAULT nextval('ago_menu_id_seq'::regclass),
    "menu_name"   varchar     NOT NULL DEFAULT '':: character varying,
    "menu_key"    varchar     NOT NULL DEFAULT '':: character varying,
    "parent_id"   int8        NOT NULL DEFAULT 0,
    "path"        varchar     NOT NULL DEFAULT '':: character varying,
    "menu_type"   varchar     NOT NULL DEFAULT '1':: character varying,
    "is_frame"    bool        NOT NULL DEFAULT false,
    "is_visible"  bool        NOT NULL DEFAULT true,
    "icon"        varchar     NOT NULL DEFAULT '':: character varying,
    "req_method"  varchar     NOT NULL DEFAULT 'GET':: character varying,
    "status"      varchar(10) NOT NULL DEFAULT '0':: character varying,
    "del_flag"    varchar(10) NOT NULL DEFAULT 'N':: character varying,
    "create_by"   varchar     NOT NULL DEFAULT '':: character varying,
    "create_time" timestamp   NOT NULL DEFAULT now(),
    "update_by"   varchar     NOT NULL DEFAULT '':: character varying,
    "update_time" timestamp   NOT NULL DEFAULT '0001-01-01 00:00:00':: timestamp without time zone,
    "remark"      varchar     NOT NULL DEFAULT '':: character varying,
    PRIMARY KEY ("id")
);

-- 唯一索引
CREATE UNIQUE INDEX IF NOT EXISTS ago_menu_key_unique ON "public"."ago_menu"(menu_key);

-- Column Comment
COMMENT ON COLUMN "public"."ago_menu"."id" IS '唯一标识';
COMMENT ON COLUMN "public"."ago_menu"."menu_name" IS '菜单名称';
COMMENT ON COLUMN "public"."ago_menu"."menu_key" IS '菜单标识';
COMMENT ON COLUMN "public"."ago_menu"."parent_id" IS '上级标识';
COMMENT ON COLUMN "public"."ago_menu"."path" IS '路由地址';
COMMENT ON COLUMN "public"."ago_menu"."menu_type" IS '菜单类型（D目录 M菜单 A按钮）';
COMMENT ON COLUMN "public"."ago_menu"."is_frame" IS '是否为外链（0是 1否）';
COMMENT ON COLUMN "public"."ago_menu"."is_visible" IS '菜单状态（0显示 1隐藏）';
COMMENT ON COLUMN "public"."ago_menu"."icon" IS '菜单图标';
COMMENT ON COLUMN "public"."ago_menu"."req_method" IS '请求方法';
COMMENT ON COLUMN "public"."ago_menu"."status" IS '状态（默认0）';
COMMENT ON COLUMN "public"."ago_menu"."del_flag" IS '删除标记';
COMMENT ON COLUMN "public"."ago_menu"."create_time" IS '创建时间';
COMMENT ON COLUMN "public"."ago_menu"."update_time" IS '更新时间';
COMMENT ON COLUMN "public"."ago_menu"."create_by" IS '创建人员';
COMMENT ON COLUMN "public"."ago_menu"."update_by" IS '更新人员';
COMMENT ON COLUMN "public"."ago_menu"."remark" IS '备注';
COMMENT ON TABLE "public"."ago_menu" IS '菜单权限表';

INSERT INTO "public"."ago_menu" ("id", "menu_name", "menu_key", "parent_id", "path", "menu_type", "is_frame",
                                 "is_visible", "icon", "req_method", status, del_flag, create_time, update_time,
                                 create_by, update_by, remark)
VALUES (1, '系统管理', 'system', 0, '#', 'D', false, true, '', 'GET', '0', 'N', now(), '0001-01-01 00:00:00', '系统初始化', '', '系统初始化');

-- ----------------------------
-- 6、登录会话记录
-- ----------------------------
-- Table Definition
CREATE TABLE "public"."ago_session"
(
    "id"            uuid      NOT NULL,
    "user_name"     varchar   NOT NULL DEFAULT '':: character varying,
    "real_name"     varchar   NOT NULL DEFAULT '':: character varying,
    "refresh_token" varchar   NOT NULL DEFAULT '':: character varying,
    "user_agent"    varchar   NOT NULL DEFAULT '':: character varying,
    "client_ip"     varchar   NOT NULL DEFAULT '':: character varying,
    "is_blocked"    bool      NOT NULL DEFAULT false,
    "expires_at"    timestamp NOT NULL DEFAULT '0001-01-01 00:00:00':: timestamp without time zone,
    "create_at"     timestamp NOT NULL DEFAULT now(),
    "remark"        varchar   NOT NULL DEFAULT '':: character varying,
    PRIMARY KEY ("id")
);

-- Column Comment
COMMENT ON COLUMN "public"."ago_session"."id" IS '唯一标识';
COMMENT ON COLUMN "public"."ago_session"."user_name" IS '用户名';
COMMENT ON COLUMN "public"."ago_session"."real_name" IS '真实姓名';
COMMENT ON COLUMN "public"."ago_session"."refresh_token" IS '刷新秘钥';
COMMENT ON COLUMN "public"."ago_session"."user_agent" IS '请求信息';
COMMENT ON COLUMN "public"."ago_session"."client_ip" IS '请求地址';
COMMENT ON COLUMN "public"."ago_session"."is_blocked" IS '是否阻止';
COMMENT ON COLUMN "public"."ago_session"."expires_at" IS '过期时间';
COMMENT ON COLUMN "public"."ago_session"."create_at" IS '创建时间';
COMMENT ON COLUMN "public"."ago_session"."remark" IS '备注';
COMMENT ON TABLE "public"."ago_session" IS '登录会话记录';