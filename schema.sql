create table if not exists keys (
  id integer primary key,
  key text not null
);
create table if not exists actions (
  id integer primary key,
  name text not null unique,
  workdir text not null,
  command text not null,
  recover_command text not null
);
create table if not exists logs (
  id integer primary key,
  action integer not null,
  at text not null,
  command text not null,
  output text not null,
  exit_code integer not null,
  FOREIGN KEY(action) REFERENCES actions(id)
);
