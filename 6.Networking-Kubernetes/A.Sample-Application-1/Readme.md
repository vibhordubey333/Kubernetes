### Architecture

1. Application Architecture
![image](https://user-images.githubusercontent.com/22407855/209440389-73e65dca-a9af-47c4-8d4b-35228b006e15.png)

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
