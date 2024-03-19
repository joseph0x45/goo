create table if not exists keys (
  id INTEGER primary key,
  key text not null
);

create table if not exists actions (
  id INTEGER primary key,
  name text not null unique,
  workdir text not null,
  command text not null,
  recover_command text not null
);
