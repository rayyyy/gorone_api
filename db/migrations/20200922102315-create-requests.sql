
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

CREATE INDEX idx_requests_on_redis_tag_name ON requests(redis_tag_name);
CREATE INDEX idx_requests_on_job_type ON requests(job_type);
CREATE INDEX idx_requests_on_done_at ON requests(done_at);

-- +migrate Down
DROP TABLE requests;