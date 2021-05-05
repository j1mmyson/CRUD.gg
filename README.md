# CRUD.gg

---

![Main Page](https://github.com/j1mmyson/j1mmyson.github.io/blob/master/assets/img/posts/devlog/login.png?raw=true)

`CRUD.gg` is a side project designed to practice developing web and go.  
I implemented `log-in`, `log-out`, `sign-up` and `something else`.. 

## Link

> AWS EC2의 public ip는 유동 ip를 부여받기 때문에 유효하지 않은 주소일 수 있습니다..  
> (The public ip of AWS EC2 may be an invalid address because it is granted a floating ip.)

<http://ec2-3-141-190-12.us-east-2.compute.amazonaws.com/>

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
- [v] Handle input exceptions
- [ ] Prevent duplicate log-in