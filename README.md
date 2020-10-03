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
