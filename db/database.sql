drop type if exists e_language_kind;
create type e_language_kind as enum ('programming', 'database');

drop table if exists languages;
create table if not exists languages (
  id serial not null,
  active boolean not null default true,
  name varchar(255) not null,
  kind e_language_kind not null,
  aliases varchar(255)[] not null,
  words varchar(255)[] not null,
  identifier uuid not null default gen_random_uuid(),
  created_at timestamp not null default current_timestamp,
  updated_at timestamp not null default current_timestamp,
  primary key (id)
);
