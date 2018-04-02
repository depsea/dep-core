create database depsea;

use depsea;

-- role
create table role(
    id serial,
    name varchar(10) not null,
    description varchar(128) not null,
    create_time int not null,
    update_time int not null
);

-- user
create table _user(
    id serial,
    username varchar(30) not null,
    password varchar(128) not null,
    email varchar(128),
    website varchar(128),
    role_id int not null,
    create_time int not null,
    update_time int not null
);

-- tag
create table tag(
    id serial,
    name varchar(10) not null,
    description varchar(128) not null,
    create_time int not null,
    update_time int not null
);

-- article
create table article(
    id serial,
    title varchar(30) not null,
    description varchar(128) not null,
    body text not null,
    create_time int not null,
    update_time int not null,
    publish_time int not null,
    status int not null
);

