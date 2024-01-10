-- 用户表
CREATE TABLE `users` (
  `id` bigint NOT NULL COMMENT '用户ID',
  `username` varchar(64) DEFAULT NULL COMMENT '用户名',
  `email` varchar(64) NOT NULL COMMENT '邮箱',
  `password` varchar(64) DEFAULT NULL COMMENT '密码',
  `is_deleted` tinyint DEFAULT '0' COMMENT '是否删除 0: 否，1: 是',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户信息';


-- 项目表
CREATE TABLE `projects` (
  `id` bigint NOT NULL  COMMENT '项目ID', 
  `project_name` varchar(64) DEFAULT NULL COMMENT '项目名称',
  `project_desc` varchar(255) DEFAULT NULL COMMENT '项目描述',
  `status` tinyint DEFAULT '0' COMMENT '项目状态 0: 私密 1: 公开',
  `is_deleted` tinyint DEFAULT '0' COMMENT '是否删除 0: 否，1: 是',
  `created_by` bigint NOT NULL COMMENT '创建人ID',
  `updated_by` bigint NOT NULL COMMENT '修改人ID',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMary KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='项目信息';

-- 分组
CREATE TABLE `groups` (
  `id` bigint NOT NULL  COMMENT '项目ID', 
  `group_name` varchar(64) DEFAULT NULL COMMENT '分组名称',
  `is_deleted` tinyint DEFAULT '0' COMMENT '是否删除 0: 否，1: 是',
  `created_by` bigint NOT NULL COMMENT '创建人ID',
  `updated_by` bigint NOT NULL COMMENT '修改人ID',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMary KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='分组信息';

-- 中间表

CREATE TABLE `groups_projects`(

  `id` bigint NOT NULL  COMMENT 'ID', 
  `group_id` bigint NOT NULL  COMMENT '项目ID', 
  

  PRIMary KEY (`id`)
)
