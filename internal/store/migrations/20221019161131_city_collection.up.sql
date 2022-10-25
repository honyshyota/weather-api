CREATE TABLE city (
    name varchar not null,
    country varchar not null,
    lat real,
    lon real
);

CREATE TABLE weather (
    name varchar not null,
    country varchar not null,
    lat real,
    lon real,
    temp double precision,
    date timestamp,
    data jsonb
);

CREATE TABLE users (
    uuid varchar unique,
    name varchar unique,
    email varchar unique,
    password varchar,
    city varchar
);