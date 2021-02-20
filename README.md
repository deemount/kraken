# Kraken v1 RESTful API

A RESTful API for Kraken v1 written in Go.

**Information:**
To use this application, you'll need a account at https://kraken.com. After creating a account, you'll also in need of a api-key and api-secret to connect successfully to the kraken api.

***Notice***: This application is still under development.

## Endpoints

***Notice***:
There are actually four endpoints defined

### Global Methods

* [SCHEME:HOST:PORT]/kraken/v1/home

### Private Methods

* [SCHEME:HOST:PORT]/kraken/v1/balance
* [SCHEME:HOST:PORT]/kraken/v1/tradebalance
* [SCHEME:HOST:PORT]/kraken/v1/ledger

## Environment

You need to create a .env file in your root folder.

```shell
API_KRAKEN_VERSION="0"
API_KRAKEN_URL="https://api.kraken.com"
API_KRAKEN_USERAGENT="GoKrakenBot/1.0"
API_KRAKEN_KEY=""
API_KRAKEN_SECRET=""

#
API_SERVER_HOST="http://localhost"
API_SERVER_PATH_PREFIX="/kraken/v1"
API_SERVER_PORT="8686"
API_SERVER_PATH_SRC="/src"
API_SERVER_CLIENT_LIMIT="10"

#
API_DB_DRIVER="postgres"
API_DB_USER=""
API_DB_PASSWORD=""
API_DB_PORT="5432"
API_DB_HOST="127.0.0.1"
API_DB_SSLMODE=false
API_DB_SCHEMA=""
API_DB_TABLE_PREFIX=""
API_DB_NAME="kraken"
API_DB_LOGMODE=true
API_DB_SINGULARTABLE=true

#
API_SWAGGER_HOST=""
API_SWAGGER_PORT=""

```

## Docker

### Build

***Notice***:
Using semantic versioning at the end, which is also tagged at github

```dockerfile
docker build -t deemount/kraken:v0.1.1 .
```

### Run

```dockerfile
docker run --publish 8686:8686 --detach --name krkn deemount/kraken:v0.1.1  
```

## History

* added first working request endpoint method for balance
* create architecture, add functionalities, first setup
* first upload & initial commit

### To Do's

* github tags
* database connection
* more methods & functionality
* more documentation

### More

* API Documentation https://www.kraken.com/features/api
