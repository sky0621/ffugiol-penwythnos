
CREATE TABLE IF NOT EXISTS work (
  id varchar(36) NOT NULL,
  name varchar(256) NOT NULL,
  price bigint,
  create_user varchar(256) DEFAULT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  update_user varchar(256) DEFAULT NULL,
  updated_at timestamp NULL DEFAULT NULL,
  deleted_at timestamp NULL DEFAULT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS work_holder (
  id varchar(36) NOT NULL,
  first_name varchar(256) NOT NULL,
  last_name varchar(256) NOT NULL,
  nickname varchar(256) DEFAULT NULL,
  create_user varchar(256) DEFAULT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  update_user varchar(256) DEFAULT NULL,
  updated_at timestamp NULL DEFAULT NULL,
  deleted_at timestamp NULL DEFAULT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS work_holder_relation (
  work_id varchar(36) NOT NULL,
  work_holder_id varchar(36) NOT NULL,
  create_user varchar(256) DEFAULT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  update_user varchar(256) DEFAULT NULL,
  updated_at timestamp NULL DEFAULT NULL,
  deleted_at timestamp NULL DEFAULT NULL,
  PRIMARY KEY (work_id, work_holder_id)
);

CREATE TABLE IF NOT EXISTS organization (
  id varchar(36) NOT NULL,
  name varchar(256) NOT NULL,
  upper_organization_id varchar(36) NOT NULL,
  create_user varchar(256) DEFAULT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  update_user varchar(256) DEFAULT NULL,
  updated_at timestamp NULL DEFAULT NULL,
  deleted_at timestamp NULL DEFAULT NULL,
  PRIMARY KEY (id)
);
