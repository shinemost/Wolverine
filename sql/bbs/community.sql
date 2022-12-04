create table `community`
(
    `id`             int(11)                                 not null auto_increment,
    -- 板块id
    `community_id`   int(10) unsigned                        not null,
    -- 板块名称
    `community_name` varchar(128) collate utf8mb4_general_ci not null,
    -- 板块介绍
    `introduction`   varchar(256) collate utf8mb4_general_ci not null,
    `create_time`    timestamp                               not null default current_timestamp,
    `update_time`    timestamp                               not null default current_timestamp on update current_timestamp,
    primary key (`id`),
    unique key `idx_community_id` (`community_id`),
    unique key `idx_community_name` (`community_name`)
) engine = InnoDB
  default charset = utf8mb4
  collate = utf8mb4_general_ci;