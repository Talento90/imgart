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
|overlay    |`{"id":"overlay","parameters":{"position":[25,75],"url":"https://raw.githubusercontent.com/Talento90/imgart/master/assets/mustache.png","opacity":100}}`|![overlay](https://imgart.herokuapp.com/api/v1/images?imgSrc=https://raw.githubusercontent.com/Talento90/imgart/master/assets/gopher.png&filters=%5B%7B%22id%22:%22overlay%22,%22parameters%22:%7B%22position%22:%5B25,75%5D,%22url%22:%22https://raw.githubusercontent.com/Talento90/imgart/master/assets/mustache.png%22,%22opacity%22:100%7D%7D%5D)|
|resize     |`{"id":"resize","parameters":{"width":25,"height":50,"filter":"linear"}}`  	|![resize](https://imgart.herokuapp.com/api/v1/images?imgSrc=https://raw.githubusercontent.com/Talento90/imgart/master/assets/gopher.png&filters=[{"id":"resize","parameters":{"width":25,"height":50,"filter":"linear"}}])|
|crop    	|`{"id":"crop","parameters":{"rectangle":[0,0,202,150]}}`                     	|![crop](https://imgart.herokuapp.com/api/v1/images?imgSrc=https://raw.githubusercontent.com/Talento90/imgart/master/assets/gopher.png&filters=[{%22id%22:%22crop%22,%22parameters%22:{%22rectangle%22:[0,0,202,150]}}])|
|rotate    	|`{"id":"rotate","parameters":{"angle":-90,"bgcolor":"transparent"}}`         	|![rotate](https://imgart.herokuapp.com/api/v1/images?imgSrc=https://raw.githubusercontent.com/Talento90/imgart/master/assets/gopher.png&filters=[{"id":"rotate","parameters":{"angle":-90,"bgcolor":"transparent"}}])|
|blur    	|`{"id":"blur","parameters":{"sigma":0.9}`         							  	|![blur](https://imgart.herokuapp.com/api/v1/images?imgSrc=https://raw.githubusercontent.com/Talento90/imgart/master/assets/gopher.png&filters=[{%22id%22:%22blur%22,%22parameters%22:{%22sigma%22:0.9}}])|
|brightness |`{"id":"brightness","parameters":{"percentage":-50}}`         					|![brightness](https://imgart.herokuapp.com/api/v1/images?imgSrc=https://raw.githubusercontent.com/Talento90/imgart/master/assets/gopher.png&filters=[{"id":"brightness","parameters":{"percentage":-50}}])|
|contrast   |`{"id":"contrast","parameters":{"percentage":100}}`         					|![contrast](https://imgart.herokuapp.com/api/v1/images?imgSrc=https://raw.githubusercontent.com/Talento90/imgart/master/assets/gopher.png&filters=[{%22id%22:%22contrast%22,%22parameters%22:{%22percentage%22:100}}])|
|gamma    	|`{"id":"gamma","parameters":{"gamma":0.2}}`         							|![gamma](https://imgart.herokuapp.com/api/v1/images?imgSrc=https://raw.githubusercontent.com/Talento90/imgart/master/assets/gopher.png&filters=[{%22id%22:%22gamma%22,%22parameters%22:{%22gamma%22:0.2}}])|
		

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

`GET /api/v1/images?imgSrc={image url}&profile=my-profile`
