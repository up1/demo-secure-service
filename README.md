# Secure service
* Stateless
* Stateful

## Start services


### Start Auth service

```
$docker compose up -d auth --build
```

Call API to Get JWT token
```
$curl http://127.0.0.1:8000
```

### Start gateway
```
$docker compose up -d gateway --build
```

Call protected api with token
```
$curl --location 'http://127.0.0.1:9000' --header 'Token: your-jwt-token'
```

