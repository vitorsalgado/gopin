CREATE
DATABASE IF NOT EXISTS go;

USE go;

DROP TABLE IF EXISTS `user_locations`;

CREATE TABLE user_locations
(
  id           VARCHAR(36) NOT NULL DEFAULT (uuid()),
  user_uuid    VARCHAR(36) NOT NULL,
  session_uuid VARCHAR(36) NOT NULL,
  lat          FLOAT(10)   NOT NULL,
  lng          FLOAT(10)   NOT NULL,
  `precision`  FLOAT(10)   NOT NULL,
  reported_at  TIMESTAMP   NOT NULL,
  inserted_at  TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,

  PRIMARY KEY (id),
  INDEX        idx_user_uuid (user_uuid),
  INDEX        idx_session_uuid (session_uuid)
);
