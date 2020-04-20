CREATE TABLE IF NOT EXISTS viewing_history (
  id varchar(36) NOT NULL,
  user_id varchar(36) NOT NULL,
  movie_id varchar(36) NOT NULL,
  created_at timestamp NOT NULL DEFAULT current_timestamp,
  PRIMARY KEY (id)
);
