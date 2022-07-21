* create user
```
curl -X POST http://localhost:9090/user/v1 -H 'Content-Type: application/json' -d '{"username": "felix", "password": "test123", "email": "abc@example.com"}'
```


* login
```
curl -X POST http://localhost:9090/login/v1 -H 'Content-Type: application/json' -d '{"username": "felix", "password": "test123"}'
```

* get user
```
curl -X GET http://localhost:9090/user/v1 -H 'Content-Type: application/json' -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTgzOTc2MTIsInVzZXJuYW1lIjoiZmVsaXgifQ.kJBs58msrWBgFxcJ92hG7NsZfUbSXY22B7mwC8AVvic'
```