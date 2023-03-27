CREATE TABLE IF NOT EXISTS clients (
  id BIGSERIAL primary key,
  name TEXT not null,
  amount decimal not null,
  lock integer not null,
  created_at TIMESTAMP default now(),
  updated_at TIMESTAMP default now()
);