
CREATE TABLE IF NOT EXISTS `todo` (
  id varchar(36) NOT NULL,
  text varchar(256) NOT NULL,
  done bool NOT NULL,
  user_id varchar(36) NOT NULL,
  create_user varchar(256) DEFAULT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  update_user varchar(256) DEFAULT NULL,
  updated_at timestamp NULL DEFAULT NULL,
  delete_user varchar(256) DEFAULT NULL,
  deleted_at timestamp NULL DEFAULT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE utf8_bin;

CREATE TABLE IF NOT EXISTS `user` (
  id varchar(36) NOT NULL,
  name varchar(256) NOT NULL,
  create_user varchar(256) DEFAULT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  update_user varchar(256) DEFAULT NULL,
  updated_at timestamp NULL DEFAULT NULL,
  delete_user varchar(256) DEFAULT NULL,
  deleted_at timestamp NULL DEFAULT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE utf8_bin;
