# Go Budget App

This is the budget app BE

# Prerequisites

1. Golang ([official download](https://go.dev/dl/), [brew](https://formulae.brew.sh/formula/go))
2. (optional) golangci-lint ([brew](https://formulae.brew.sh/formula/golangci-lint),
[official quickstart](https://golangci-lint.run/welcome/quick-start/)) 

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
2. get one transaction by ID `GET localhost:3000/api/v1/transaction/6tsn-Mmmv-fpZS-k2Bv`
3. create transaction `POST localhost:3000/api/v1/transaction`
4. update transaction `PATCH localhost:3000/api/v1/transaction/6tsn-Mmmv-fpZS-k2Bv`

`User` endpoint `localhost:3000/api/v1/user`
1. list all users `GET localhost:3000/api/v1/user`

# Dependencies

1. Go-Chi is the HTTP server framework used ([official website](https://go-chi.io/#/README),
[github](https://github.com/go-chi/chi))
3. mongo-go-driver is the official database driver by MongoDB ([github](https://github.com/mongodb/mongo-go-driver))
