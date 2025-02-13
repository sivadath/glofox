# Glofox API

## Overview
Basic REST apis for a fitness studio to create classes and members to book classes

## Ignored Scenarios
 If a member attempts to book a class on a given date and multiple options are available (e.g., Yoga and Pilates), the system randomly assigns one of the classes.

## Prerequisites
- Go 1.23+
- Docker (optional, for containerized deployment)

## Setup
### Running Locally
```sh
make run
```

### Running with Docker
```sh
make docker-run
```

### Running Tests
```sh
make test
```

### Lint and Code Checks
```sh
make fmt
make vet
make lint
```

to run `make lint` command `golangci-lint` should be pre-installed

## API Documentation
Swagger documentation is available at (execute the code and hit this endpoint):
```
http://localhost:8080/swagger/index.html
```

## Cleanup
```sh
make clean
make docker-clean
```

