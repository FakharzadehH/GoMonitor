CREATE TABLE servers(
id serial PRIMARY KEY,
address text not null,
success int not null default 0,
failure int not null default 0,
last_failure TIMESTAMP,
created_at timestamp not null default now(),
updated_at timestamp not null default now()
)