create table user (
  id bigint auto_increment primary key,
  username varchar(255) unique not null,
  email varchar(255) unique not null,
  password varchar(255) not null,
  created_at datetime default current_timestamp not null,
  updated_at datetime default current_timestamp not null
) charset = utf8mb4;

create table project(
  id bigint auto_increment primary key,
  project_name varchar(255) not null,
  project_description varchar(255) ,
  is_public boolean default true,
  is_deleted boolean default false,
  created_by bigint not null,
  created_at datetime default current_timestamp not null,
  updated_at datetime default current_timestamp not null
) charset = utf8mb4;
