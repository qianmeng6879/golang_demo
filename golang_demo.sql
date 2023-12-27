drop database if exists golang_demo;

create database golang_demo character set utf8mb4;

use golang_demo;


create table t_user(
    id bigint auto_increment,
    username varchar(20) not null,
    password varchar(255) not null,
    email varchar(100) null,
    avatar varchar(255),
    is_admin bit(1) default 0,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime,
    primary key(id),
    index (deleted_at)
);

insert into t_user(username, password,created_at) values('zero','admin', now());