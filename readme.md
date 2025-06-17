
# DEALS

Deals Technical Backend Engineer Golang




## Getting Started

### Prequisites

This Project using the following dependencies
- **Programming Language** : Go 1.23
- **Database** : Postgres
- **Package Manager** : Go Modules

### Installation
Build the application from the sources
- ``git clone https://github.com/oktopriima/deals``
- ``cd deals``
- install dependency using ``go mod tidy`` and ``go mod vendor``
- copy the environment example ``cp env-example.yaml env.yaml``
- adjust the variable with your local params
- run the migration ``go run database/migration/migration.go``
- run the seeder ``go run database/seeder/seeder.go``
- run the application ``go run main.go``

### Documentaion
The Documentaion writen using postman
- access the `docs` folder
- import `DEALS.postman_collection.json` into your postman application
- for admin access `username: admin` and `password: admin123`
- for user is increamental from 1 to 100. ex `username: user1` and `password: pass1`
- add `{{ROUTE}}` with value `http://localhost:8000` to postman local environment
- add `{{TOKEN}}` and `{{ADMIN_TOKEN}}` to access endpoint admin and user


## Authors

- [@oktopriima](https://www.github.com/oktopriima)