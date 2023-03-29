CREATE TABLE IF NOT EXISTS clients (
  id BIGSERIAL primary key,
  name TEXT not null,
  amount decimal not null,
  created_at TIMESTAMP default now(),
  updated_at TIMESTAMP default now()
);

CREATE TABLE IF NOT EXISTS transations (
  sender_id bigint not null,
  receiver_id bigint not null,
  amount decimal not null,
  finished boolean not null,
  created_at timestamp default now(),
  PRIMARY KEY (sender_id, receiver_id, created_at),
  CONSTRAINT fk_sender FOREIGN KEY(sender_id) REFERENCES clients(id),
  CONSTRAINT fk_receiver FOREIGN KEY(receiver_id) REFERENCES clients(id)
);

--ALTER TABLE IF EXISTS clients DROP COLUMN lock;
--ALTER TABLE IF EXISTS transations ADD COLUMN created_at timestamp default now();
--ALTER TABLE IF EXISTS transations ADD PRIMARY KEY (sender_id, receiver_id, created_at);