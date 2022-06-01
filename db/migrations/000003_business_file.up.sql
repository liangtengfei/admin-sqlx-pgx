-- ----------------------------
-- 文件上传功能
-- ----------------------------
-- Table Definition
CREATE TABLE "public"."ago_file"
(
    "id"          uuid      NOT NULL,
    "file_name"   varchar   NOT NULL DEFAULT '':: character varying,
    "file_path"   varchar   NOT NULL DEFAULT '':: character varying,
    "file_url"    varchar   NOT NULL DEFAULT '':: character varying,
    "file_size"   int8      NOT NULL DEFAULT 0,
    "user_id"     int8      NOT NULL,
    "mime_type"   varchar   NOT NULL DEFAULT '':: character varying,
    "create_time" timestamp NOT NULL DEFAULT now(),
    "create_by"   varchar   NOT NULL DEFAULT '':: character varying,
    "remark"      varchar   NOT NULL DEFAULT '':: character varying,
    PRIMARY KEY ("id")
);

-- Column Comment
COMMENT ON COLUMN "public"."ago_file"."id" IS '唯一标识';
COMMENT ON COLUMN "public"."ago_file"."file_name" IS '上传文件名';
COMMENT ON COLUMN "public"."ago_file"."file_path" IS '文件存储路径';
COMMENT ON COLUMN "public"."ago_file"."file_url" IS '文件访问路径';
COMMENT ON COLUMN "public"."ago_file"."file_size" IS '文件的大小';
COMMENT ON COLUMN "public"."ago_file"."user_id" IS '用户编号';
COMMENT ON COLUMN "public"."ago_file"."mime_type" IS '文件类型';
COMMENT ON COLUMN "public"."ago_file"."create_time" IS '创建时间';
COMMENT ON COLUMN "public"."ago_file"."create_by" IS '创建人员';
COMMENT ON COLUMN "public"."ago_file"."remark" IS '备注';
COMMENT ON TABLE "public"."ago_file" IS '文件信息表';