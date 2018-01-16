# gorpo
Service to modify images.

## Third Party Dependencies
- https://github.com/golang/dep
- https://github.com/julienschmidt/httprouter
- https://github.com/disintegration/imaging

## Quality Tools
- https://goreportcard.com/
- https://travis-ci.org/

## Run Project
- docker-compose up

## Project Stucture

https://peter.bourgon.org/go-best-practices-2016/#repository-structure

- cmd
- pkg
    - gorpo
        | image
        | effect
        | profile
        - image
            - service
            - logging
        - profile
            - service
            - logging
    - config
    - effect
    - repository
    - log
    - httpapi
    - errors

## Server

- Redis
- Effects Hash (effects order can have a different order but the result is the same...)
- Tests
- Panic Handler
- Makefile
- Context (implement context mechanism)
- Graceful Shutdown
- Metrics
- HealthCheck


Errors:

- duplicate profile ids
- validate filter parameters

## Commands
go list -f '{{ join .Imports "\n" }}' package_path