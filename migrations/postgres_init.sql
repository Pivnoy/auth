CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create type status as enum (
    'approved',
    'denied',
    'pending',
    'frozen'
    );

create table secret_question (
    id uuid primary key default uuid_generate_v4(),
    question varchar(255) not null
);

create table "user" (
    id uuid primary key default uuid_generate_v4(),
    email varchar(50) not null unique,
    phone varchar(12) not null unique,
    password varchar(70) not null,
    status status not null,
    secret_question_id uuid references secret_question not null,
    secret_question_answer varchar(255) not null
);

create table admin (
    id uuid primary key default uuid_generate_v4(),
    login varchar(50) not null unique,
    password varchar(70) not null
);

insert into secret_question (id, question) values ('09e4bde9-4472-4ab0-99fd-04a5cac0a4f1', 'how to kakat?');
insert into secret_question (id, question) values ('09e4bde9-4472-4ab0-99fd-04a5cac0a4f2','how to become compilator?');
insert into secret_question (id, question) values ('09e4bde9-4472-4ab0-99fd-04a5cac0a4f3','how to hard coding?');

insert into "user" (email, phone, password, status, secret_question_id, secret_question_answer) VALUES ('alex@mail.ru', '+79212281488','$2a$14$.lAgcS4UiAiJKiD9kbvlHObowOCdV5mEmreUvG/HewCVlm0heqNN6', 'approved', '09e4bde9-4472-4ab0-99fd-04a5cac0a4f2', '$2a$14$ZQrlXRnKomR4QaMU77NnyO4FGnimYZb3pcxY/O9isq5xsKLH3ZNGa'); -- kutsenko
insert into "user" (email, phone, password, status, secret_question_id, secret_question_answer) VALUES ('roma@mail.ru', '+76969696969','$2a$14$Q1NW5.b/kWan0Y.m1Zjy5uApp9PtuTKDOdmppL.5V8/j/AczOLlHq', 'approved', '09e4bde9-4472-4ab0-99fd-04a5cac0a4f1', '$2a$14$lIDfKVDfsJN3VYKa08Hzyek5i0V1Gb8w2RalasSt291JxmteGxzp6'); -- chach