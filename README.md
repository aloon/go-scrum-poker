# go-scrum-poker

Free Agile Estimation App

The application is made with go using websockets

The application is running in https://go-scrum-poker.fly.dev/

---

## How to run

### Docker

```bash
$ git clone git@github.com:aloon/go-scrum-poker.git
$ cd go-scrum-poker
$ docker build -t go-scrum-poker .
$ docker run -p 8080:8080 go-scrum-poker
```

### Local

```bash
$ git clone git@github.com:aloon/go-scrum-poker.git
$ cd go-scrum-poker
$ go get -u github.com/gin-gonic/gin
$ go get -u github.com/gorilla/websocket
$ go get -u github.com/gosimple/slug
$ go get -u github.com/stretchr/testify
$ go run main.go
```

---

<!-- badge -->
[![Open Source](https://badges.frapsoft.com/os/v1/open-source.svg?v=103)](https://opensource.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![PR's Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat)](http://makeapullrequest.com)  
[![Fly Deploy](https://github.com/aloon/go-scrum-poker/actions/workflows/fly.yml/badge.svg)](https://github.com/aloon/go-scrum-poker/actions/workflows/fly.yml)
[![CodeQL](https://github.com/aloon/go-scrum-poker/actions/workflows/codeql.yml/badge.svg)](https://github.com/aloon/go-scrum-poker/actions/workflows/codeql.yml)
<!-- badge -->