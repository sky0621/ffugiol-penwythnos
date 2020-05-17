CREATE TABLE IF NOT EXISTS viewing_history (
  id varchar(36) NOT NULL,
  user_id varchar(36) NOT NULL REFERENCES viewer(id),
  movie_id varchar(36) NOT NULL references movie(id),
  created_at timestamp NOT NULL DEFAULT current_timestamp,
  PRIMARY KEY (id)
);
