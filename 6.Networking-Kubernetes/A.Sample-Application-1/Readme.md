### Architecture

1. Application Architecture
![image](https://user-images.githubusercontent.com/22407855/209424911-16721900-5cad-482f-8006-4280febe5a5b.png)

### Run

1. To run `docker-compose up -d --build`
2. To remove the containers `docker-compose down`
3. To build users-api `docker build -t kub-demo-users .`
4. Edit deployment file to fix `ErrImmagePull` -> `kubectl edit deployment users-deployment`

### APIs
1. Auth token
```
POST: http://192.168.1.6:32163/login
Payload: 
{
    "email": "test@gmail.com",
    "password": "testers"
}
```

2. Signup
```
POST: http://192.168.1.6:32163/signup
Payload: 
{
    "email": "test@gmail.com",
    "password": "testers"
}
```