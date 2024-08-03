# Go Budget App

This is the budget app BE

# Usage

1. Install Dependencies
```bash
make install
```

2. Run
```bash
make run
```

# Endpoints
Try any of the below

`Transaction` endpoint `localhost:3000/api/v1/transaction`
1. list all transactions `GET localhost:3000/api/v1/transaction`
2. get one transaction by ID `GET localhost:3000/api/v1/transaction/Wn8EnRQ3koTGPQJl1nzyCer66RMWXXSBdPG9a`
3. create transaction `POST localhost:3000/api/v1/transaction`

`User` endpoint `localhost:3000/api/v1/user`
1. list all users `GET localhost:3000/api/v1/user`

# Dependencies

1. Go-Chi is the HTTP server framework used ([official website](https://go-chi.io/#/README), [github](https://github.com/go-chi/chi))
2. mongo-go-driver is the official database driver by MongoDB ([github](https://github.com/mongodb/mongo-go-driver))
3. 