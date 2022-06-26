#### 1. API Register
@host = localhost:8080

http://{{host}}/api/user/register

//Headers
content-type: application/json

//raw

{
    "username": "bambangsantoso",
    "email" : "bambangsantoso112@gmail.com",
    "password": "bangbangsantoso"
}


#### 2. API Login
@host = localhost:8080

http://{{host}}/api/user/login

//Headers
content-type: application/json  

//raw

{
    "username": "bambangsantoso",
    "password": "bangbangsantoso"
}

#### 3. API Account Info

http://{{host}}/api/check/profile

//Headers
content-type: application/json

Authorization: Bearer Token

### 4. API Balance

http://{{host}}/api/check/wallet

//Headers
content-type: application/json

Authorization: Bearer Token








