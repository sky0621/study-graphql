
-- +migrate Up
CREATE TABLE customer (
  id bigserial NOT NULL,
  name varchar(64) NOT NULL,
  age int NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE todo (
  id bigserial NOT NULL,
  task varchar(64) NOT NULL,
  user_id bigint NOT NULL REFERENCES customer(id),
  PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE customer;
DROP TABLE todo;
