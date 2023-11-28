CREATE TABLE user (
  id bigint AUTO_INCREMENT PRIMARY KEY COMMENT '用户ID',
  username varchar(255) UNIQUE NOT NULL COMMENT '用户名',
  email varchar(255) UNIQUE NOT NULL COMMENT '邮箱',
  password varchar(255) NOT NULL COMMENT '密码',
  created_at datetime DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '创建时间',
  updated_at datetime DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户表';


create table project (
  id bigint auto_increment PRIMARY KEY COMMENT '项目ID',
  project_name varchar(255) NOT NULL COMMENT '项目名称',
  project_description varchar(255) COMMENT '项目描述',
  is_public tinyint DEFAULT '0' COMMENT '是否公开 0：是 1：否',
  is_deleted tinyint DEFAULT '0' COMMENT '是否删除 0：未删除 1：删除',
  created_by bigint NOT NULL COMMENT '创建人',
  created_at datetime DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '创建时间',
  updated_at datetime DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='项目表';

