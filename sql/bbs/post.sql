drop table if exists `post`;
create table `post`
(
    `id`           bigint(20)                               not null auto_increment,
    `post_id`      bigint(20)                               not null comment '帖子id',
    `title`        varchar(128) collate utf8mb4_general_ci  not null comment '标题',
    `content`      varchar(8192) collate utf8mb4_general_ci not null comment '内容',
    `author_id`    bigint(20)                               not null comment '作者id',
    `community_id` bigint(20)                               not null default '1' comment '板块id',
    `status`       tinyint(4)                               not null default '1' comment '帖子状态',
    `create_time`  timestamp                                not null default current_timestamp comment '创建时间',
    `update_time`  timestamp                                not null default current_timestamp on update current_timestamp comment '更新时间',
    primary key (`id`),
    -- 唯一索引
    unique key `idx_post_id` (`post_id`),
    -- 普通索引
    key `idx_author_id` (`author_id`),
    key `idx_community_id` (`community_id`)
) engine = InnoDB
  default charset = utf8mb4
  collate = utf8mb4_general_ci;