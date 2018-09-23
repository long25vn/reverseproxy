# reverse proxy

## install 

```
./install.sh
```

## method.json

```
[
  {
    "service": "auth",
    "method": "login",
    "ispublic": true
  },
  {
    "service": "blog",
    "method": "createpost",
    "ispublic": false
  }
]
```

## Format

```
http://domain/api/service/method   
```

## Login

```
 curl -X POST http://0.0.0.0/api/auth/login   
```

```
{
    "access_token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwicm9sZSI6MSwiaWF0IjoxNTE2MjM5MDIyfQ.gO9vEwksOuFttxM1vE7xifrqoZ2PMh9dt_lGH2p2zvY",
    "expires_in":21600,
    "token_type":"Bearer"
}⏎
```

## Create POST (accept role = 1 )

### case 1: Wrong url

```
curl -X POST http://0.0.0.0:80/api/blog   

URL is Invalid!
```

### case 2: no token or a wrong token

```
curl -X POST http://0.0.0.0:80/api/blog/createpost

Unauthorized!
```

### case 3: token (role = 2)

```
curl http://0.0.0.0:80/api/blog/createpost -X POST -H "Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwicm9sZSI6MiwiaWF0IjoxNTE2MjM5MDIyfQ.u5-jKQxdwvLiE6uOn2ZcKE52ehtd9SfC-HW7olPICqg"

Unauthorized!
```

### case 4: token (role = 1)

```
curl http://0.0.0.0:80/api/blog/createpost -X POST -H "Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwicm9sZSI6MSwiaWF0IjoxNTE2MjM5MDIyfQ.gO9vEwksOuFttxM1vE7xifrqoZ2PMh9dt_lGH2p2zvY"

{"message":"Tao Bai Viet Thanh Cong!"}⏎
```
