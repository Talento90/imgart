# IMGART [![Build Status](https://travis-ci.org/Talento90/imgart.svg?branch=master)](https://travis-ci.org/Talento90/imgart) [![Go Report Card](https://goreportcard.com/badge/github.com/talento90/imgart)](https://goreportcard.com/report/github.com/talento90/imgart)


![logo](https://imgart.onrender.com/api/v1/images?imgSrc=https://raw.githubusercontent.com/Talento90/imgart/master/assets/gopher.png&filters=[{"id":"overlay","parameters":{"position":[25,75],"url":"https://goo.gl/UBrXeo"}},{"id":"overlay","parameters":{"position":[22,-35],"url":"https://goo.gl/aEkkDh"}}])


IMGART it's an HTTP service for image processing based on filters and profiles.

Documentation: https://imgart.onrender.com/api/v1/docs

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

## Usage

`{host}/api/v1/images?imgSrc={0}&profile={1}&filters={2}`

* *host*: Server address
* *imgSrc*: Image URL
* *profile*: Profile we want to apply
* *filters*: List of filters to process

**Example**

```https://imgart.onrender.com/api/v1/images?imgSrc=https://goo.gl/mq7yPD&profile=example&filters=[{"id":"rotate","parameters":{"angle":-90}}]```

**Result**

![result](https://imgart.onrender.com/api/v1/images?imgSrc=https://goo.gl/mq7yPD&profile=example&filters=[{"id":"rotate","parameters":{"angle":-90}}])


## Effects

The engine behind image manipulation is the fabulous library: github.com/disintegration/


**Available Effects**

| Effect     | JSON                                                                                             | Result                                                                                                                                                                                                                                                                                                                                                 |
| ---------- | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| overlay    | `{"id":"overlay","parameters":{"position":[25,75],"url":"https://goo.gl/UBrXeo","opacity":100}}` | ![overlay](https://imgart.onrender.com/api/v1/images?imgSrc=https://raw.githubusercontent.com/Talento90/imgart/master/assets/gopher.png&filters=%5B%7B%22id%22:%22overlay%22,%22parameters%22:%7B%22position%22:%5B25,75%5D,%22url%22:%22https://raw.githubusercontent.com/Talento90/imgart/master/assets/mustache.png%22,%22opacity%22:100%7D%7D%5D) |
| resize     | `{"id":"resize","parameters":{"width":25,"height":50,"filter":"linear"}}`                        | ![resize](https://imgart.onrender.com/api/v1/images?imgSrc=https://raw.githubusercontent.com/Talento90/imgart/master/assets/gopher.png&filters=[{"id":"resize","parameters":{"width":25,"height":50,"filter":"linear"}}])                                                                                                                             |
| crop       | `{"id":"crop","parameters":{"rectangle":[0,0,202,150]}}`                                         | ![crop](https://imgart.onrender.com/api/v1/images?imgSrc=https://raw.githubusercontent.com/Talento90/imgart/master/assets/gopher.png&filters=[{"id":"crop","parameters":{"rectangle":[0,0,202,150]}}])                                                                                                                                                |
| rotate     | `{"id":"rotate","parameters":{"angle":-90,"bgcolor":"transparent"}}`                             | ![rotate](https://imgart.onrender.com/api/v1/images?imgSrc=https://raw.githubusercontent.com/Talento90/imgart/master/assets/gopher.png&filters=[{"id":"rotate","parameters":{"angle":-90,"bgcolor":"transparent"}}])                                                                                                                                  |
| blur       | `{"id":"blur","parameters":{"sigma":0.9}`                                                        | ![blur](https://imgart.onrender.com/api/v1/images?imgSrc=https://raw.githubusercontent.com/Talento90/imgart/master/assets/gopher.png&filters=[{"id":"blur","parameters":{"sigma":0.9}}])                                                                                                                                                              |
| brightness | `{"id":"brightness","parameters":{"percentage":-50}}`                                            | ![brightness](https://imgart.onrender.com/api/v1/images?imgSrc=https://raw.githubusercontent.com/Talento90/imgart/master/assets/gopher.png&filters=[{"id":"brightness","parameters":{"percentage":-50}}])                                                                                                                                             |
| contrast   | `{"id":"contrast","parameters":{"percentage":100}}`                                              | ![contrast](https://imgart.onrender.com/api/v1/images?imgSrc=https://raw.githubusercontent.com/Talento90/imgart/master/assets/gopher.png&filters=[{"id":"contrast","parameters":{"percentage":100}}])                                                                                                                                                 |
| gamma      | `{"id":"gamma","parameters":{"gamma":0.2}}`                                                      | ![gamma](https://imgart.onrender.com/api/v1/images?imgSrc=https://raw.githubusercontent.com/Talento90/imgart/master/assets/gopher.png&filters=[{"id":"gamma","parameters":{"gamma":0.2}}])                                                                                                                                                            |
		
It's possible to combine multiple effects:

```json
/api/v1/images?imgSrc=https://raw.githubusercontent.com/Talento90/imgart/master/assets/gopher.png&filters=[{"id":"overlay","parameters":{"position":[25,75],"url":"https://goo.gl/UBrXeo"}},{"id":"overlay","parameters":{"position":[22,-35],"url":"https://goo.gl/aEkkDh"}}, {"id":"crop","parameters":{"rectangle":[0,0,202,150]}}]
```
![result](https://imgart.onrender.com/api/v1/images?imgSrc=https://raw.githubusercontent.com/Talento90/imgart/master/assets/gopher.png&filters=[{"id":"overlay","parameters":{"position":[25,75],"url":"https://goo.gl/UBrXeo"}},{"id":"overlay","parameters":{"position":[22,-35],"url":"https://goo.gl/aEkkDh"}},{"id":"crop","parameters":{"rectangle":[0,0,202,150]}}])


## Profiles

If you don't want to specify filters in URL, you can create a profile with all pre configured filters and then use it in query parameters `&profile={profile-id}`.


**Create Profile**
```json
POST /api/v1/profiles

{
    "id": "my-profile",
    "filters": [
        { "id": "brightness", "parameters":{ "percentage":0.3}},
        { "id": "crop", "parameters":{"rectangle":[0,0,100,200]}}
    ]
}
```

**Using profile in query parameters**

`GET /api/v1/images?imgSrc={image url}&profile=my-profile`

