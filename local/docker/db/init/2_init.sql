CREATE TABLE IF NOT EXISTS viewer (
  id varchar(36) NOT NULL,
  name varchar(256) NOT NULL,
  nickname varchar(256),
  PRIMARY KEY (id)
);
