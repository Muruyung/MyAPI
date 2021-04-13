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

### PUT Method
POST Method is a method used to run update queries against user data
#### PUT Method used in Postman:
![put_method_postman](https://raw.githubusercontent.com/Muruyung/MyAPI/master/dokumentasi%20API/Postman-Put%20User.png)
#### PUT Method result in PostgreSQL:
![put_method_navicat](https://github.com/Muruyung/MyAPI/blob/master/dokumentasi%20API/Navicat-Put%20User.png?raw=true)

### DELETE Method
POST Method is a method used to run delete queries against user data
#### DELETE Method used in Postman:
![delete_method_postman](https://raw.githubusercontent.com/Muruyung/MyAPI/master/dokumentasi%20API/Postman-Delete%20User.png)
#### DELETE Method result in PostgreSQL:
![delete_method_navicat](https://github.com/Muruyung/MyAPI/blob/master/dokumentasi%20API/Navicat-Delete%20User.png?raw=true)

### LOGIN Method
LOGIN Method is a method used to generate token and set online againts user data
#### LOGIN Method used in Postman:
![login_method_postman](https://github.com/Muruyung/MyAPI/blob/master/dokumentasi%20API/Postman-Login%20User%20(Get%20Token).png?raw=true)
#### LOGIN Method result in PostgreSQL (Set Online User):
![login_method_navicat](https://github.com/Muruyung/MyAPI/blob/master/dokumentasi%20API/Navicat-Set%20Online%20User.png?raw=true)
#### LOGIN Method result in PostgreSQL (Save Token):
![login_method_navicat](https://github.com/Muruyung/MyAPI/blob/master/dokumentasi%20API/Navicat-Save%20Token.png?raw=true)

### LOGOUT Method
LOGOUT Method is a method used to revoke token againts and set offline user data
#### LOGIN Method used in Postman:
![logout_method_postman](https://github.com/Muruyung/MyAPI/blob/master/dokumentasi%20API/Postman-Logout%20User%20(Set%20Offline).png?raw=true)
#### LOGOUT Method result in PostgreSQL (Set Offline User):
![logout_method_navicat](https://github.com/Muruyung/MyAPI/blob/master/dokumentasi%20API/Navicat-Set%20Offline%20User.png?raw=true)
#### LOGOUT Method result in PostgreSQL (Delete Token):
![logout_method_navicat](https://github.com/Muruyung/MyAPI/blob/master/dokumentasi%20API/Navicat-Revoke%20Token.png?raw=true)

### SIGNUP Method
SIGNUP Method is like POST Method
#### SIGNUP Method used in Postman:
![signup_method_postman](https://github.com/Muruyung/MyAPI/blob/master/dokumentasi%20API/Postman-Signup.png?raw=true)
#### SIGNUP Method result in PostgreSQL:
![signup_method_navicat](https://github.com/Muruyung/MyAPI/blob/master/dokumentasi%20API/Navicat-Signup%20User.png?raw=true)
#### SIGNUP Method result in PostgreSQL (Create Token Dataset):
![signup_method_navicat](https://github.com/Muruyung/MyAPI/blob/master/dokumentasi%20API/Navicat-Signup%20User%20(Create%20Token%20Dataset).png?raw=true)
