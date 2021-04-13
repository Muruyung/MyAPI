# MyAPI

## Description
Simple API using clean code architecture

This project has  4 Domain layer :
 * Entity Layer
 * Usecase Layer
 * Controller Layer  
 * Delivery Layer

## Tools Used:
 * [PostgreSQL](https://github.com/lib/pq)
 * [Iris Lib](https://github.com/kataras/iris)
 * [Xorm Lib](https://xorm.io/xorm)
 * [JWT](https://github.com/dgrijalva/jwt-go)
 * [bcrypt](https://golang.org/x/crypto/bcrypt)

### 📚 [Database Config](https://github.com/Muruyung/MyAPI/blob/master/config/config.go)

## Documentation API
### POST Method
POST Method is a method used to run input queries against user data
#### POST Method used in Postman:
![post_method_postman](https://raw.githubusercontent.com/Muruyung/MyAPI/master/dokumentasi%20API/Postman-Post%20User.png)
#### POST Method result in PostgreSQL:
![post_method_navicat](https://github.com/Muruyung/MyAPI/blob/master/dokumentasi%20API/Navicat-Post%20User.png?raw=true)

### GET Method
This method has 2 method variations. There are GET List and GET by ID
#### GET List used in Postman:
![get_list_method_postman](https://github.com/Muruyung/MyAPI/blob/master/dokumentasi%20API/Postman-Get%20List%20User.png?raw=true)
#### GET by ID used in Postman:
![get_by_id_postman](https://github.com/Muruyung/MyAPI/blob/master/dokumentasi%20API/Postman-Get%20by%20ID%20User.png?raw=true)
