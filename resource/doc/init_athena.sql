create database `athena`;

use `athena`;

-- table of tag
create table `tag` (
    `id` int not null auto_increment comment '标签ID',
    `name` varchar(64) not null comment '标签名称',
    `created_at` timestamp not null default current_timestamp comment '创建时间',
    `updated_at` timestamp not null default current_timestamp on update current_timestamp comment '修改时间',
    `deleted_at` timestamp default null comment '删除时间',
    primary key (`id`)
);

-- table of article
create table `article` (
    `id` int not null auto_increment comment '文章ID',  
    `title` varchar(128) not null comment '文章标题',
    `content` text not null comment '文章内容',
    `tag_id` int not null comment '文章所属标签ID',
    `created_at` timestamp not null default current_timestamp comment '创建时间',
    `updated_at` timestamp not null default current_timestamp on update current_timestamp comment '修改时间',
    `deleted_at` timestamp default null comment '删除时间',
    primary key (`id`)
);

-- table of category
create table `category` (
    `id` int not null auto_increment comment '分类ID',
    `name` varchar(64) not null comment '分类名称',
    `created_at` timestamp not null default current_timestamp comment '创建时间',
    `updated_at` timestamp not null default current_timestamp on update current_timestamp comment '修改时间',
    `deleted_at` timestamp default null comment '删除时间',
)

-- table of role
create table `role` (
    `id` int not null auto_increment comment '角色ID',
    `name` varchar(64) not null comment '角色名称',
    `created_at` timestamp not null default current_timestamp comment '创建时间',
    `updated_at` timestamp not null default current_timestamp on update current_timestamp comment '修改时间',
    `deleted_at` timestamp default null comment '删除时间',
);

-- table of menu
create table `menu` (
    `id` int not null auto_increment comment '菜单ID',
    `name` varchar(64) not null comment '菜单名称',
    `url` varchar(128) not null comment '菜单URL',
    `created_at` timestamp not null default current_timestamp comment '创建时间',
    `updated_at` timestamp not null default current_timestamp on update current_timestamp comment '修改时间',
    `deleted_at` timestamp default null comment '删除时间',
);