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
GET    /api/v1/                  --> app/app/apis/endpoints.home
GET    /api/v1/teams             --> app/app/apis/endpoints.(*TeamEndpoint).getAll
GET    /api/v1/team/:id          --> app/app/apis/endpoints.(*TeamEndpoint).getById
POST   /api/v1/team/create       --> app/app/apis/endpoints.(*TeamEndpoint).create
GET    /api/v1/players           --> app/app/apis/endpoints.(*PlayerEndpoint).getAll
GET    /api/v1/players/:id       --> app/app/apis/endpoints.(*PlayerEndpoint).getByTeam
GET    /api/v1/player/:id        --> app/app/apis/endpoints.(*PlayerEndpoint).getById
```
