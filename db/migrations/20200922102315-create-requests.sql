
-- +migrate Up
CREATE TABLE requests (
  id INT(11) AUTO_INCREMENT NOT NULL,
  body JSON,
  created_at DATETIME NOT NULL DEFAULT current_timestamp,
  updated_at DATETIME NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
  deleted_at DATETIME,
  PRIMARY KEY(id)
);

-- +migrate Down
DROP TABLE requests;