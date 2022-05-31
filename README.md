<h1 id="gopin-top" align="center">GoPin</h1>

<div align="center">
    <a href="#"><img src="logo.png" width="120px" alt="Hive"></a>
    <p align="center">
        Demonstration Location Management API
        <br />
        <a href="#"><strong>Explore the API Docs »</strong></a>
        <br />
    </p>
    <div>
      <a href="https://github.com/vitorsalgado/gopin/actions/workflows/ci.yml">
        <img src="https://github.com/vitorsalgado/gopin/actions/workflows/ci.yml/badge.svg" alt="CI Status" />
      </a>
      <a href="#">
        <img src="https://img.shields.io/badge/go-1.18-blue" alt="Go 1.18" />
      </a>
      <a href="https://codecov.io/gh/vitorsalgado/gopin">
        <img src="https://codecov.io/gh/vitorsalgado/gopin/branch/main/graph/badge.svg?token=FFKD8C3000"/>
      </a>
      <a href="https://conventionalcommits.org">
        <img src="https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg" alt="Conventional Commits"/>
      </a>
    </div>
</div>

## Getting Started

```
make up
```

Default **PORT** is **:8080**  
API **http://localhost:8080/api/v1**  
The application entrypoint is in **cmd/app/main.go**.  
For other commands check the [Makefile](Makefile).

## Config

This application uses Environment Variables for configurations.  
For local environment, you can use a **.env** file for configurations.
Check [.env.sample](.env.sample).

## API Docs

This project uses **Swagger** for API documentation.  
With the application running, navigate to: **/docs**.

## GoDocs

```
make docs
```

Navigate to: [GoDocs](http://127.0.0.1:6060/pkg/github.com/vitorsalgado/gopin/)

## CI

This project uses Github Actions for **Continuous Integration**.  
Check [here](https://github.com/vitorsalgado/gopin/actions).  
GH Action definition [here](.github/workflows/ci.yml).

## Docker

Docker is used for several tasks like, e2e tests, development environment, database.  
The project is idealized to be as lightweight as possible. The final image is around **10mb**.  
Check the main [Dockerfile](Dockerfile).  
The [docker-compose-dev-yml](deployments/local/docker-compose-dev.yml) creates a development environment with the
**app** with **Live Reload** and **MySQL**.

## Testing

This application uses unit, integration and e2e tests for quality assurance.  
The e2e tests solution is implemented using Docker Compose. Check [here](test).

### Unit Tests

```
make test
```

### End-2-End Tests

```
make test-e2e
```

## Built With

- Go
- Go Modules
- Docker and Docker Compose
- Swagger
- MySQL

<p align="center"><a href="#gopin-top">back to top</a></p>
