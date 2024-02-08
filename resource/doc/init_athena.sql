create database `athena`;

use `athena`;

-- table of user
create table `user` (
    `id` int not null auto_increment comment '用户ID',
    `passport` varchar(64) not null comment '用户账户',
    `password` varchar(128) not null comment '用户密码',
    `email` varchar(128) not null comment '用户邮箱',
    `signature` varchar(256) not null default '' comment '用户个人简介',
    `created_at` timestamp not null default current_timestamp comment '创建时间',
    `updated_at` timestamp not null default current_timestamp on update current_timestamp comment '修改时间',
    `deleted_at` timestamp default null comment '删除时间',
    primary key (`id`)
);

-- table of tag
create table `tag` (
    `id` int not null auto_increment comment '标签ID',
    `name` varchar(64) not null comment '标签名称',
    `author_id` int not null comment '标签作者ID',
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
    `image` varchar(128) default '' comment '文章图片封面',
    `author_id` int not null comment '文章作者ID',
    `created_at` timestamp not null default current_timestamp comment '创建时间',
    `updated_at` timestamp not null default current_timestamp on update current_timestamp comment '修改时间',
    `deleted_at` timestamp default null comment '删除时间',
    primary key (`id`)
);

-- table of article_tag
create table `article_tag` (
    `id` int not null auto_increment comment '文章标签关联ID',
    `article_id` int not null comment '文章ID',
    `tag_id` int not null comment '标签ID',
    `created_at` timestamp not null default current_timestamp comment '创建时间',
    `updated_at` timestamp not null default current_timestamp on update current_timestamp comment '修改时间',
    `deleted_at` timestamp default null comment '删除时间',
    primary key (`id`)
);

-- table of category
create table `category` (
    `id` int not null auto_increment comment '分类ID',
    `name` varchar(64) not null comment '分类名称',
    `author_id` int not null comment '分类作者ID',
    `created_at` timestamp not null default current_timestamp comment '创建时间',
    `updated_at` timestamp not null default current_timestamp on update current_timestamp comment '修改时间',
    `deleted_at` timestamp default null comment '删除时间',
    primary key (`id`)
);

-- table of role
create table `role` (
    `id` int not null auto_increment comment '角色ID',
    `name` varchar(64) not null comment '角色名称',
    `created_at` timestamp not null default current_timestamp comment '创建时间',
    `updated_at` timestamp not null default current_timestamp on update current_timestamp comment '修改时间',
    `deleted_at` timestamp default null comment '删除时间',
    primary key (`id`)
);

-- table of menu
create table `menu` (
    `id` int not null auto_increment comment '菜单ID',
    `name` varchar(64) not null comment '菜单名称',
    `url` varchar(128) not null comment '菜单URL',
    `created_at` timestamp not null default current_timestamp comment '创建时间',
    `updated_at` timestamp not null default current_timestamp on update current_timestamp comment '修改时间',
    `deleted_at` timestamp default null comment '删除时间',
    primary key (`id`)
);