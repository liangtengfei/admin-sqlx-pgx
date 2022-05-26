-- ----------------------------
-- 7、系统配置表
-- ----------------------------
-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS ago_config_id_seq;

-- Table Definition
CREATE TABLE "public"."ago_config"
(
    "id"           int8        NOT NULL DEFAULT nextval('ago_config_id_seq'::regclass),
    "config_name"  varchar     NOT NULL DEFAULT ''::character varying,
    "config_key"   varchar     NOT NULL DEFAULT ''::character varying,
    "config_value" varchar     NOT NULL DEFAULT ''::character varying,
    "status"       varchar(10) NOT NULL DEFAULT '0':: character varying,
    "del_flag"     varchar(10) NOT NULL DEFAULT 'N':: character varying,
    "create_time"  timestamp   NOT NULL DEFAULT now(),
    "update_time"  timestamp   NOT NULL DEFAULT '0001-01-01 00:00:00':: timestamp without time zone,
    "create_by"    varchar     NOT NULL DEFAULT '':: character varying,
    "update_by"    varchar     NOT NULL DEFAULT '':: character varying,
    "remark"       varchar     NOT NULL DEFAULT '':: character varying,
    PRIMARY KEY ("id")
);

-- Column Comment
COMMENT ON COLUMN "public"."ago_config"."id" IS '唯一标识';
COMMENT ON COLUMN "public"."ago_config"."config_name" IS '配置名称';
COMMENT ON COLUMN "public"."ago_config"."config_key" IS '配置标识';
COMMENT ON COLUMN "public"."ago_config"."config_value" IS '配置内容';
COMMENT ON COLUMN "public"."ago_config"."status" IS '状态（默认0）';
COMMENT ON COLUMN "public"."ago_config"."del_flag" IS '删除标记';
COMMENT ON COLUMN "public"."ago_config"."create_time" IS '创建时间';
COMMENT ON COLUMN "public"."ago_config"."update_time" IS '更新时间';
COMMENT ON COLUMN "public"."ago_config"."create_by" IS '创建人员';
COMMENT ON COLUMN "public"."ago_config"."update_by" IS '更新人员';
COMMENT ON COLUMN "public"."ago_config"."remark" IS '备注';

-- ----------------------------
-- 8、通知公告表
-- ----------------------------
-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS ago_notice_id_seq;

-- Table Definition
CREATE TABLE "public"."ago_notice"
(
    "id"             int8        NOT NULL DEFAULT nextval('ago_notice_id_seq'::regclass),
    "notice_title"   varchar     NOT NULL DEFAULT ''::character varying,
    "notice_type"    varchar(1)  NOT NULL DEFAULT '1'::character varying,
    "notice_content" text        NOT NULL DEFAULT ''::text,
    "status"         varchar(10) NOT NULL DEFAULT '0':: character varying,
    "del_flag"       varchar(10) NOT NULL DEFAULT 'N':: character varying,
    "create_time"    timestamp   NOT NULL DEFAULT now(),
    "update_time"    timestamp   NOT NULL DEFAULT '0001-01-01 00:00:00':: timestamp without time zone,
    "create_by"      varchar     NOT NULL DEFAULT '':: character varying,
    "update_by"      varchar     NOT NULL DEFAULT '':: character varying,
    "remark"         varchar     NOT NULL DEFAULT '':: character varying,
    PRIMARY KEY ("id")
);

-- Column Comment
COMMENT ON COLUMN "public"."ago_notice"."id" IS '唯一标识';
COMMENT ON COLUMN "public"."ago_notice"."notice_type" IS '公告类型（1通知 2公告）';
COMMENT ON COLUMN "public"."ago_notice"."notice_content" IS '公告内容';
COMMENT ON COLUMN "public"."ago_notice"."status" IS '状态（默认0）';
COMMENT ON COLUMN "public"."ago_notice"."del_flag" IS '删除标记';
COMMENT ON COLUMN "public"."ago_notice"."create_time" IS '创建时间';
COMMENT ON COLUMN "public"."ago_notice"."update_time" IS '更新时间';
COMMENT ON COLUMN "public"."ago_notice"."create_by" IS '创建人员';
COMMENT ON COLUMN "public"."ago_notice"."update_by" IS '更新人员';
COMMENT ON COLUMN "public"."ago_notice"."remark" IS '备注';

-- ----------------------------
-- 9、字典类型表
-- ----------------------------
-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS ago_dict_type_id_seq;

-- Table Definition
CREATE TABLE "public"."ago_dict_type"
(
    "id"          int8        NOT NULL DEFAULT nextval('ago_dict_type_id_seq'::regclass),
    "dict_name"   varchar     NOT NULL DEFAULT ''::character varying,
    "dict_type"   varchar     NOT NULL DEFAULT ''::character varying,
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
COMMENT ON COLUMN "public"."ago_dict_type"."id" IS '字典主键';
COMMENT ON COLUMN "public"."ago_dict_type"."dict_name" IS '字典类型名称';
COMMENT ON COLUMN "public"."ago_dict_type"."dict_type" IS '字典类型编码';
COMMENT ON COLUMN "public"."ago_dict_type"."status" IS '状态（默认0）';
COMMENT ON COLUMN "public"."ago_dict_type"."del_flag" IS '删除标记';
COMMENT ON COLUMN "public"."ago_dict_type"."create_time" IS '创建时间';
COMMENT ON COLUMN "public"."ago_dict_type"."update_time" IS '更新时间';
COMMENT ON COLUMN "public"."ago_dict_type"."create_by" IS '创建人员';
COMMENT ON COLUMN "public"."ago_dict_type"."update_by" IS '更新人员';
COMMENT ON COLUMN "public"."ago_dict_type"."remark" IS '备注';

-- ----------------------------
-- 10、字典数据表
-- ----------------------------
-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS ago_dict_data_id_seq;

-- Table Definition
CREATE TABLE "public"."ago_dict_data"
(
    "id"          int8        NOT NULL DEFAULT nextval('ago_dict_data_id_seq'::regclass),
    "dict_label"  varchar     NOT NULL DEFAULT ''::character varying,
    "dict_value"  varchar     NOT NULL DEFAULT ''::character varying,
    "dict_type"   varchar     NOT NULL DEFAULT ''::character varying,
    "list_class"  varchar     NOT NULL DEFAULT ''::character varying,
    "css_class"   varchar     NOT NULL DEFAULT ''::character varying,
    "status"      varchar(10) NOT NULL DEFAULT '0':: character varying,
    "del_flag"    varchar(10) NOT NULL DEFAULT 'N':: character varying,
    "create_time" timestamp   NOT NULL DEFAULT now(),
    "update_time" timestamp   NOT NULL DEFAULT '0001-01-01 00:00:00':: timestamp without time zone,
    "create_by"   varchar     NOT NULL DEFAULT '':: character varying,
    "update_by"   varchar     NOT NULL DEFAULT '':: character varying,
    "remark"      varchar     NOT NULL DEFAULT '':: character varying,
    "order_num"   int4        NOT NULL DEFAULT 0,
    PRIMARY KEY ("id")
);

-- Column Comment
COMMENT ON COLUMN "public"."ago_dict_data"."id" IS '唯一标识';
COMMENT ON COLUMN "public"."ago_dict_data"."dict_label" IS '字典标签';
COMMENT ON COLUMN "public"."ago_dict_data"."dict_value" IS '字典键值';
COMMENT ON COLUMN "public"."ago_dict_data"."dict_type" IS '字典类型编码';
COMMENT ON COLUMN "public"."ago_dict_data"."list_class" IS '表格回显样式';
COMMENT ON COLUMN "public"."ago_dict_data"."css_class" IS '样式属性（其他样式扩展）';
COMMENT ON COLUMN "public"."ago_dict_data"."order_num" IS '排序编号';
COMMENT ON COLUMN "public"."ago_dict_data"."status" IS '状态（默认0）';
COMMENT ON COLUMN "public"."ago_dict_data"."del_flag" IS '删除标记';
COMMENT ON COLUMN "public"."ago_dict_data"."create_time" IS '创建时间';
COMMENT ON COLUMN "public"."ago_dict_data"."update_time" IS '更新时间';
COMMENT ON COLUMN "public"."ago_dict_data"."create_by" IS '创建人员';
COMMENT ON COLUMN "public"."ago_dict_data"."update_by" IS '更新人员';
COMMENT ON COLUMN "public"."ago_dict_data"."remark" IS '备注';

-- ----------------------------
-- 11、操作日志表
-- ----------------------------
-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS ago_operation_log_id_seq;

-- Table Definition
CREATE TABLE "public"."ago_operation_log"
(
    "id"               int8       NOT NULL DEFAULT nextval('ago_operation_log_id_seq'::regclass),
    "business_type"    varchar    NOT NULL DEFAULT ''::character varying,
    "business_title"   varchar    NOT NULL DEFAULT ''::character varying,
    "invoke_method"    varchar    NOT NULL DEFAULT ''::character varying,
    "request_method"   varchar    NOT NULL DEFAULT ''::character varying,
    "request_url"      varchar    NOT NULL DEFAULT ''::character varying,
    "client_type"      varchar(1) NOT NULL DEFAULT 0,
    "client_ip"        varchar    NOT NULL DEFAULT ''::character varying,
    "client_location"  varchar    NOT NULL DEFAULT ''::character varying,
    "client_param"     text       NOT NULL DEFAULT ''::text,
    "operation_type"   varchar(1) NOT NULL DEFAULT 0,
    "operation_result" text       NOT NULL DEFAULT ''::text,
    "error_msg"        text       NOT NULL DEFAULT ''::text,
    "status"           varchar(1) NOT NULL DEFAULT ''::character varying,
    "create_by"        varchar    NOT NULL DEFAULT ''::character varying,
    "create_time"      timestamp  NOT NULL DEFAULT now(),
    "remark"           varchar    NOT NULL DEFAULT ''::character varying,
    "dept_name"        varchar    NOT NULL DEFAULT ''::character varying,
    PRIMARY KEY ("id")
);

-- Column Comment
COMMENT ON COLUMN "public"."ago_operation_log"."id" IS '唯一标识';
COMMENT ON COLUMN "public"."ago_operation_log"."business_type" IS '业务类型';
COMMENT ON COLUMN "public"."ago_operation_log"."business_title" IS '业务内容';
COMMENT ON COLUMN "public"."ago_operation_log"."invoke_method" IS '方法名称';
COMMENT ON COLUMN "public"."ago_operation_log"."request_method" IS '请求方式';
COMMENT ON COLUMN "public"."ago_operation_log"."request_url" IS '请求URL';
COMMENT ON COLUMN "public"."ago_operation_log"."client_type" IS '操作类别（0其它 1后台用户 2手机端用户）';
COMMENT ON COLUMN "public"."ago_operation_log"."client_ip" IS '主机地址';
COMMENT ON COLUMN "public"."ago_operation_log"."client_location" IS '操作地点';
COMMENT ON COLUMN "public"."ago_operation_log"."client_param" IS '请求参数';
COMMENT ON COLUMN "public"."ago_operation_log"."operation_type" IS '操作类型（0其它 1新增 2修改 3删除）';
COMMENT ON COLUMN "public"."ago_operation_log"."operation_result" IS '返回参数';
COMMENT ON COLUMN "public"."ago_operation_log"."error_msg" IS '错误消息';
COMMENT ON COLUMN "public"."ago_operation_log"."status" IS '状态（0正常 1异常）';
COMMENT ON COLUMN "public"."ago_operation_log"."create_by" IS '操作人员';
COMMENT ON COLUMN "public"."ago_operation_log"."create_time" IS '创建时间';
COMMENT ON COLUMN "public"."ago_operation_log"."remark" IS '备注';

-- ----------------------------
-- 12、用户角色关联表
-- ----------------------------
-- Table Definition
CREATE TABLE "public"."ago_user_role"
(
    "user_id" int8 NOT NULL,
    "role_id" int8 NOT NULL,
    PRIMARY KEY ("user_id", "role_id")
);
-- ----------------------------
-- 13、角色部门关联表
-- ----------------------------
-- Table Definition
CREATE TABLE "public"."ago_role_dept"
(
    "role_id" int8 NOT NULL,
    "dept_id" int8 NOT NULL,
    PRIMARY KEY ("role_id", "dept_id")
);
-- ----------------------------
-- 14、角色菜单关联表
-- ----------------------------
-- Table Definition
CREATE TABLE "public"."ago_role_menu"
(
    "role_id" int8 NOT NULL,
    "menu_id" int8 NOT NULL,
    PRIMARY KEY ("role_id", "menu_id")
);