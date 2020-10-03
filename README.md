# api-soccer 

### Point's
| K | V
| ------ | ------ | 
| Project Structure | Package, Manual DI, Etc 
| Go Mod | Module
| MySql | Database 
| Gin | Http Handler, Routing, Endpoint
| Gorm | ORM, Relationship Model, Query, Transactional
| Viper | App Configuration 
| Testify | Test

#
### Project Structure
```
-root project
     └ app              » Application Project
     └ config           » Configuration Files (JSON)
     └ db               » Migration Script
     └ test             » Test


- app                   » Application Project 
   └ apis               » Application Layer / Application Business Rules
       └ endpoints      » Request-Response Controller  
       └ param          » Request Parameter
   └ entities           » Domain Layer
   └ injection          » Manual Dependency Injection
       └ repositories   » Repository types
       └ services       » Service types
   └ repository         » Data Layer 
   └ services           » Enterprise Business Rules
```

#
### Tools Arguments
| Arguments | Description | Ex
| ------  | ------ | ------ | 
|  -h     | Help   |  go run main.go -h
|  -start | Starting with default environment (local)   |  go run main.go -start
|  -start env | Starting selected environment |  go run main.go -start dev

| Environment | Description | Ex
| ------  | ------ | ------ | 
|  -local | Local   |  go run main.go -start
|  -dev | Development | go run main.go -start dev
|  -prod | Production | go run main.go -start prod




#
### API Endpoint's
```
GET    /api/v1/               » app/app/apis/endpoints.home                         » Home / Base informations
GET    /api/v1/teams          » app/app/apis/endpoints.(*TeamEndpoint).getAll       » Get all teams
GET    /api/v1/team/:id       » app/app/apis/endpoints.(*TeamEndpoint).getById      » Get team by id
POST   /api/v1/team/create    » app/app/apis/endpoints.(*TeamEndpoint).create       » Create Team with Players (Body JSON)
GET    /api/v1/players        » app/app/apis/endpoints.(*PlayerEndpoint).getAll     » Get all player
GET    /api/v1/players/:id    » app/app/apis/endpoints.(*PlayerEndpoint).getByTeam  » Get Player by Team ID
GET    /api/v1/player/:id     » app/app/apis/endpoints.(*PlayerEndpoint).getById    » Get Player by Player ID
```


#
### Database
DB Migration location at projectroot/db/migration/000001_init_schema.up.sql
```
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

```
