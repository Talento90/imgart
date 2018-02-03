# IMGART [![Build Status](https://travis-ci.org/Talento90/imgart.svg?branch=master)](https://travis-ci.org/Talento90/imgart) [![Go Report Card](https://goreportcard.com/badge/github.com/Talento90/imgart)](https://goreportcard.com/report/github.com/Talento90/imgart)


IMGART it's an HTTP service for image processing based on filters and profiles.

### Features
- Image manipulation
- Image caching (Redis)
- Predefined Profiles (MongoDB)
- API Documentation (Swagger Specification)
- Error Handling
- Docker
- Debugging using Delve and Docker
- Makefile
- Testing
- Graceful shutdown
- Healthcheck

### Docker Support
- Dockerfile (Development)
- Dockerfile.CI (Production)
- docker-compose.yml (Run application)
- docker-compose.debug.yml (Run application with delve for debugging)


### Setup Project
- Clone repository: `git clone git@github.com:Talento90/imgart.git`
- Install dependencies: `make deps`
- Run using docker: `docker-compose up` or `make docker`
- Open application: `open http:localhost:4005`

## Effects

The engine behind image manipulation is this fabulous library: github.com/disintegration/imaging


|Effect     |JSON                     													  	|Result  	|
|-----------|-------------------------------------------------------------------------------|-----------|
|overlay    |`{"id":"overlay","parameters":{"position":[100,200],"url":"","opacity":100}}`	| 	asd  	|
|resize     |`{"id":"resize","parameters":{"width":300,"height":700,"filter":"linear"}}`  	| 	asd  	|
|crop    	|`{"id":"crop","parameters":{"rectangle":[0,0,100,200]}}`                     	| 	asd  	|
|rotate    	|`{"id":"rotate","parameters":{"angle":-90,"bgcolor":"transparent"}}`         	| 	asd  	|
|blur    	|`{"id":"blur","parameters":{"sigma":0.8}`         							  	| 	asd  	|
|brightness |`{"id":"brightness","parameters":{"percentage":0.3}}`         					| 	asd 	|
|contrast   |`{"id":"contrast","parameters":{"percentage":0.7}}`         					| 	asd  	|
|gamma    	|`{"id":"gamma","parameters":{"gamma":0.6}}`         							| 	asd  	|
		


## Profiles


If you don't want to pass filters in URL you can simple create a profile with all pre configure filters and then use it in query parameters `&profile={profile-id}`.


**Create Profile**
```json
POST /api/v1/profiles

{
    "id": "my-profile",
    "filters: [
        {"id":"brightness","parameters":{"percentage":0.3}},
        {"id":"crop","parameters":{"rectangle":[0,0,100,200]}}
    ]
}
```

**Using profile in query parameters**

`GET /api/v1/images?imgSrc=""&profile=my-profile`
