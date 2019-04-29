Go Simple Features : 
-----------------
 - crud rest api
 - redis caching
 - jwt auth
 - middleware with gorilla
 - working with gorm

INSTALLATION : 
-----
 - brew install dep
 - cd $GOPATH/src/github.com/me/go_jwt_redis
 - dep ensure

USAGE : 
-----
###  USER ENDPOINTS 
 - **POST** http://localhost:8081/users/create (*create users data*)
 - **DELETE** http://localhost:8081/users/delete/{id} (*delete users data*)

### GET TOKEN

```
    curl -X POST --user username:password http://localhost:8081/login
```

###  MAHASISWA ENDPOINTS 
    * Note : always use token as authorization bearer each request *
 - **GET** http://localhost:8081/mahasiswas (*list all mahasiswa data*)
 - **POST** http://localhost:8081/mahasiswas/create (*insert mahasiswa data*)
 - **PATCH** http://localhost:8081/mahasiswas/update/{id} (*update mahasiswa data*)
 - **DELETE** http://localhost:8081/mahasiswas/delete/{id} (*delete mahasiswa data*)


TODO
-----
 - Update Documentation
