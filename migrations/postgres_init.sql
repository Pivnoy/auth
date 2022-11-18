CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table "user" (
    id uuid primary key default uuid_generate_v4(),
    email varchar(50),
    password varchar(70)
);

insert into "user" (email, password) VALUES ('alex', '$2a$14$.lAgcS4UiAiJKiD9kbvlHObowOCdV5mEmreUvG/HewCVlm0heqNN6'); -- kutsenko
insert into "user" (email, password) VALUES ('roma', '$2a$14$Q1NW5.b/kWan0Y.m1Zjy5uApp9PtuTKDOdmppL.5V8/j/AczOLlHq'); -- chach