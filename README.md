#### 1. API Register: to register new user
@host = localhost:8080

http://{{host}}/api/user/register

content-type: application/json

//raw

{
    "username": "bambangsantoso",
    "email" : "bambangsantoso112@gmail.com",
    "password": "bangbangsantoso"
}


#### 2. API Login: for registered users to log in to app
@host = localhost:8080

http://{{host}}/api/user/register

content-type: application/json

//raw

{
    "username": "bambangsantoso",
    "password": "bangbangsantoso"
}


