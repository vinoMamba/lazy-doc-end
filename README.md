## Lazy Doc End

go proxy

```bash
export GOPROXY=https://goproxy.cn
```

request 

```bash
curl -H "Content-Type:application/json" -X POST -d '{"username":"vino@test.com","password":"123456","confirmPassword":"123456"}' http://127.0.0.1:3000/user/register
```

#### Database structure

1. Users
