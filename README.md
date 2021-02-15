# Kraken v1 RESTful API

A RESTful API for Kraken v1 written in Go

## Endpoints

### Private Methods

* [SCHEME:HOST:PORT]/kraken/v1/balance


## Docker

### Build

**Notice**:
Using semantic versioning at the end, which is also tagged at github

```dockerfile
docker build -t deemount/kraken:v0.1.1 .
```

### Run

```dockerfile
docker run --publish 8686:8686 --detach --name dcgs deemount/kraken:v0.1.1  
```

### History

* create architecture, add functionalities, first setup
* first upload & initial commit

### To Do's

* github tags
* database connection
* more methods & functionality
* more documentation
