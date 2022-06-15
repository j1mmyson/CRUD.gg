# CRUD.gg

![Main Page](https://github.com/j1mmyson/j1mmyson.github.io/blob/main/assets/img/posts/devlog/login.png?raw=true)

`CRUD.gg` is a side project designed to practice developing web and go.  
I implemented `log-in`, `log-out`, `sign-up` and `something else`.. 



## Used Stacks

- #### Back-End: Golang (v 1.16.3)

- #### DataBase: Mysql

- #### Deployment: AWS EC2, AWS RDS

- #### Front-End: Go templates, java script

## Getting Started

1. `git clone https://github.com/j1mmyson/Go_CRUD.git`

2. create `account.go`

   ``` go
   package main
   
   const (
       host     = "<your DB server's address>" // ex) dbname.blahblah.us=east-2.rds.amazonaws.com
       database = "<database name>" // ex) gocrud
       user     = "<user name>" // ex) admin
       password = "<password>" // ex) qwe123
   )
   ```

3. `go mod tidy`

4. `go build -o server`

5. run binary file

## To Do

- [ ] Find password (Give temporary password)
- [ ] Change password
- [x] Handle input exceptions
- [ ] Prevent duplicate log-in
