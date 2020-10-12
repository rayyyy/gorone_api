
-- +migrate Up
CREATE TABLE requests (
  id INT(11) AUTO_INCREMENT NOT NULL,
  redis_tag_name VARCHAR(255) NOT NULL,
  job_type VARCHAR(255) NOT NULL,
  desired_count INT(11) NOT NULL,
  body JSON,
  created_at DATETIME NOT NULL DEFAULT current_timestamp,
  updated_at DATETIME NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
  done_at DATETIME,
  deleted_at DATETIME,
  PRIMARY KEY(id)
);

-- +migrate Down
DROP TABLE requests;