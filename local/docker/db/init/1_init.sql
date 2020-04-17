CREATE TABLE IF NOT EXISTS movie (
  id varchar(36) NOT NULL,
  name varchar(256) NOT NULL,
  filename varchar(256) NOT NULL,
  scale integer NOT NULL DEFAULT 0,
  PRIMARY KEY (id)
);
