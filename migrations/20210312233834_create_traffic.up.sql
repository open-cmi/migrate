CREATE TABLE IF NOT EXISTS user_traffic (
  email varchar(64) unique NOT NULL default '',
  upload INTEGER NOT NULL default 0,
  download INTEGER NOT NULL default 0
);
