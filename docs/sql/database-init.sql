create table if not exists article
(
    id            int unsigned auto_increment
        primary key,
    title         varchar(255) default '' null,
    user_id       varchar(255) default '' not null,
    tags          varchar(255) default '' null,
    cover         varchar(255) default '' null,
    content       longtext                null,
    trans_content longtext                null,
    html_id       int                     null,
    trans_html_id int                     null,
    ctime         datetime                null,
    mtime         datetime                null on update CURRENT_TIMESTAMP,
    abstract      varchar(255) default '' null
);

create table if not exists file
(
    id        int auto_increment comment 'Primary Key'
        primary key,
    create_at datetime     null comment 'Create Time',
    update_at datetime     null comment 'Update Time',
    file_key  varchar(255) null,
    format    varchar(255) null,
    user_id   varchar(255) null,
    size      bigint       null
);

create table if not exists share
(
    id        int auto_increment
        primary key,
    platform  varchar(20)  not null,
    user_id   int          null,
    ip        varchar(50)  null,
    index_key varchar(100) null,
    create_at datetime     null,
    constraint share_pk
        unique (index_key)
);

create table if not exists video_subtitle
(
    video_id    int         not null,
    subtitle_id int         null,
    user_id     int         not null,
    language    varchar(20) not null
);

create table if not exists video_task
(
    id          int          not null
        primary key,
    resource_id varchar(255) null,
    output      varchar(255) null,
    status      int          null,
    create_at   timestamp    null,
    update_at   timestamp    null,
    check (`status` in (-(1), 0, 1))
);

