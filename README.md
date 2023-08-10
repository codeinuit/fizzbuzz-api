# FizzBuzz API ![build&test](https://github.com/codeinuit/fizzbuzz-api/actions/workflows/go-build-and-test.yml/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/codeinuit/fizzbuzz-api)](https://goreportcard.com/report/github.com/codeinuit/fizzbuzz-api)

A FizzBuzz implementation with Go as an API

Built and tested with Golang 1.20 with MySQL 8.0 

## Setup and run

### Configuration

API configuration go through environnement variables

```
# API configuration
PORT=8080

# Gin configuration
# Define it only to run in debug mode
DEBUG=

# MySQL configuration
MYSQL_HOST=localhost
MYSQL_PORT=3306
MYSQL_DB=main
MYSQL_USER=user
MYSQL_PASS=password
```

### Using Linux
To run FizzBuzzAPI, you will need to run a MySQL database using version 8.

```
# build
$ make 

# run api
$ PORT=8080 \
  MYSQL_HOST=localhost \
  MYSQL_PORT=3306 \
  MYSQL_DB=main \
  MYSQL_USER=user \
  MYSQL_PASS=password \
  ./bin/api
```

### Using Docker
FizzBuzz provides a `Dockerfile` for building and running the API, but also provides a docker compose configuration to run both database and API, that also contains the environnement default configuration. In order to use it, you will need to run the following command

```
$ docker compose up --build
```



## Documentation
Swagger implementation

You can access to the Swagger documentation using the Swagger website itself [here](https://petstore.swagger.io/?url=https://path/to/file.yaml), or localy by running a Docker image with the correct configuration located at `./fizzbuzz_api.yml`

```
$ docker pull swaggerapi/swagger-ui
$ docker run -p 80:8080 -e SWAGGER_JSON=/fizzbuzz_api.yml -v ${pwd}/fizzbuzz_api.yaml:/usr/share/nginx/html/fizzbuzz_api.yaml swaggerapi/swagger-ui
```


## Library used
- [gin-gonic/gin](https://github.com/gin-gonic/gin)
- [sirupsen/logrus](https://github.com/sirupsen/logrus)
- [stretchr/testify](https://github.com/stretchr/testify)
- [go-gorm/gorm](https://github.com/go-gorm/gorm)

## Contributing
FizzBuzz is a Golang project developed as a technical test and neither does aspire to be updated nor looking for contributions. However, any comments are welcome.