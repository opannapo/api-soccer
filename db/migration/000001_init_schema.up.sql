CREATE TABLE IF NOT EXISTS `tbl_team`
(
    `id`         int(11)      NOT NULL AUTO_INCREMENT,
    `name`       varchar(100) NOT NULL,
    `address`    text         NOT NULL,
    `created_at` int(11)      NOT NULL COMMENT 'timestamp',
    `updated_at` int(11)      NOT NULL COMMENT 'timestamp',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4;

CREATE TABLE IF NOT EXISTS `tbl_player`
(
    `id`        int(11)      NOT NULL AUTO_INCREMENT,
    `name`      varchar(100) NOT NULL,
    `age`       int(11)      NOT NULL,
    `join_date` int(11)      NOT NULL COMMENT 'timestamp',
    `created_at` int(11)     NOT NULL COMMENT 'timestamp',
    `updated_at` int(11)     NOT NULL COMMENT 'timestamp',
    `team_id`   int(11)      NOT NULL COMMENT 'tbl_team id',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4;

CREATE INDEX tbl_player_team_id_IDX USING BTREE ON db_test_kitabisa.tbl_player (team_id);

