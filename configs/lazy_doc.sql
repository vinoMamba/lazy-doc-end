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


-- 目录表
CREATE TABLE `directories` (
  `id` bigint NOT NULL  COMMENT '目录ID', 
  `parent_id` bigint NOT NULL COMMENT '父目录ID',
  `dir_name` varchar(64) NOT NULL COMMENT '目录名称',
  `is_deleted` tinyint DEFAULT '0' COMMENT '是否删除 0: 否，1: 是',
  `is_public` tinyint DEFAULT '0' COMMENT '状态 0: 公开，1: 私密',
  `created_by` bigint NOT NULL COMMENT '创建人ID',
  `updated_by` bigint NOT NULL COMMENT '修改人ID',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMary KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='目录信息';
