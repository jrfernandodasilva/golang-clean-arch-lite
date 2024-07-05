
# Golang Clean Architecture Lite

[![License](https://img.shields.io/github/license/jrfernandodasilva/golang-clean-arch-lite.svg)](LICENSE)
[![Written in Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)](https://golang.org/)
[![GitHub release](https://img.shields.io/github/v/release/jrfernandodasilva/golang-clean-arch-lite.svg)](https://github.com/jrfernandodasilva/golang-clean-arch-lite/releases)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)
[![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white)](https://www.docker.com/)
[![Contributors](https://img.shields.io/github/contributors/jrfernandodasilva/golang-clean-arch-lite.svg)](https://github.com/jrfernandodasilva/golang-clean-arch-lite/graphs/contributors)
[![Last Update](https://img.shields.io/github/last-commit/jrfernandodasilva/golang-clean-arch-lite.svg)](https://github.com/jrfernandodasilva/golang-clean-arch-lite/commits/main)

# Run the project

### with golang
```bash
make run
```

### with docker
```bash
make docker_run
```

### with docker compose + hot reload
```bash
docker-compose up
```

# Build the project

### with golang
```bash
make build
```

### with docker
```bash
make docker_build
```

### with docker compose
```bash
docker-compose build
```

# Clean the project

```bash
make clean
```

## Structure files
```md
├── infra
│   └── persistence
│       └── memory_user_repository.go
├── domain
│   ├── entity
│   │   └── user.go
│   ├── repository
│   │   └── user_repository.go
│   └── usecase
│       ├── create_user.go
│       ├── get_user.go
│       └── list_users.go
├── build
│   └── local
│       └── start-container.sh
├── go.sum
├── go.mod
├── main.go
├── Makefile
├── Dockerfile
├── docker-compose.yaml
├── LICENSE
└── README.md
```