
-- +migrate Up
CREATE TABLE assignment (
  id INT(11) AUTO_INCREMENT NOT NULL,
  request_id INT(11)  NOT NULL,
  public_ip VARCHAR(255) NOT NULL,
  created_at DATETIME NOT NULL DEFAULT current_timestamp,
  updated_at DATETIME NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
  deleted_at DATETIME,
  PRIMARY KEY(id)
);

CREATE INDEX idx_assignment_on_request_id ON assignment(request_id);

-- +migrate Down
DROP TABLE assignment;