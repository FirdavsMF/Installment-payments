# clone a repository 
```
git clone https://github.com/FirdavsMF/Installment-payments.git
```
## install dependency
``` 
go get
```

## start the server
```
go run cmd/main.go
```

## open  index.html
```
web/index.html
```

![](https://github.com/FirdavsMF/Installment-payments/blob/master/%D0%B4%D0%BE%D0%BA%D1%83%D0%BC%D0%B5%D0%BD%D1%82%D1%8B/Installment-payments.png?raw=true)

## registration 
```
curl --request GET http://127.0.0.1:9999/api/accounts?phone=+992000000000&email=firdavs@gmail.com&username=Firdavs&password=Firdavs
```

## log in with your login and password
```
curl --request GET http://127.0.0.1:9999/api/accounts/login?phone=+992000000000&password=Firdavs
```

## add a product
```  
curl --request POST http://127.0.0.1:9999/api/products/save?id=0&accountid=2
```

## view all product

``` curl --request GET   http://127.0.0.1:9999/api/products/All
```

## product update
```
curl --request GET http://127.0.0.1:9999/product.edit
```

## remove product
``` 
curl --request GET    http://127.0.0.1:9999/api/products/Buy?id=12&count=10
```

