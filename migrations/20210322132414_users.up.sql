CREATE TABLE IF NOT EXISTS users (
  email varchar(64) primary key NOT NULL,
  uuid varchar(36) NOT NULL UNIQUE
);
