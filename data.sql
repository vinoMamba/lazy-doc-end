create table user (
  id bigint auto_increment primary key,
  username varchar(255) unique not null,
  email varchar(255) unique not null,
  password varchar(255) not null,
  created_at datetime default current_timestamp not null,
  updated_at datetime default current_timestamp not null
) charset = utf8mb4;
