# Backend 
This is the folder for Backend(Go, MySQL)

##The Structure
Infinity is based on clean architecture, the project structure is like below.

.
├── Makefile
├── README.md
├── api
│   ├── models
│   │   ├── models Files
│   ├── v1
│   │   ├── handles
│   │   ├── middleware
│   │   ├── repository
│   │   ├── services
│   │   ├── responses.go
│   ├── router.go
├── configs
│   ├── config.go
│   ├── mconfig.toml
│   ├── logs.go
│   ├── Other Constant / static files
├── database
│   ├── connection
│   │   ├── connection.go
│   │   ├── connectionInterface.go
│   ├── migrations
│   │   ├── vesion_databasename.up.sql
│   │   ├── vesion_databasename.down.sql
│   ├── dbMigrate.go
├── logs
├── postman(optional)
├── vendor (GoDep)
│   ├── vendor packages

