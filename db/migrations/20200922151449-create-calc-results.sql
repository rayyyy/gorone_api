
-- +migrate Up
CREATE TABLE calc_results (
  id INT(11) AUTO_INCREMENT NOT NULL,
  key_name VARCHAR(255) NOT NULL,
  result INT(11) NOT NULL,
  created_at DATETIME NOT NULL DEFAULT current_timestamp,
  updated_at DATETIME NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
  deleted_at DATETIME,
  PRIMARY KEY(id)
);

-- +migrate Down
DROP TABLE calc_results;