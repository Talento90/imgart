# gorpo
Service to modify images.

## Third Party Dependencies
dependecy manager

- https://github.com/golang/dep

http router

- https://github.com/julienschmidt/httprouter

image processing

- github.com/anthonynsimon/bild


## Quality Tools

- https://goreportcard.com/
- https://travis-ci.org/

## Run Project

- docker-compose up

## Project Stucture

https://peter.bourgon.org/go-best-practices-2016/#repository-structure

root folder - 

adapters - standard library wrappers (ex: logger)
config - application configurations (ex: read ENV_VARS, file, arguments)


downloaders
effects
http-api
repositories

domain/core/root - contains our application domain
    - no external/3th party dependencies!!! 
    - *only* standard library dependencies
    - interfaces everywhere
    - managers that handles application logic (dependencies must be interfaces)






## Best Practises

- packages should be singular or at least follow a convention
- package name should be representative of their content just by reading the name


# Missing Features

* Context (implement context mechanism)
* Graceful Shutdown
* Metrics
