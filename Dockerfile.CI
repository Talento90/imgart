# Build Stage
FROM golang:alpine AS build
LABEL maintainer "Marco Talento <marcotalento90@gmail.com>"

# Setting working directory
WORKDIR ${GOPATH}/src/github.com/talento90/imgart

# Copy source code
COPY . .

# Build our source to generate an executable file
RUN go build -o imgartapi ./cmd/imgartapi

# Run Stage
FROM alpine
LABEL maintainer "Marco Talento <marcotalento90@gmail.com>"

# Setting working directory for our applicationn
WORKDIR /app

# Copy just the binary file from the build stage
COPY --from=build go/src/github.com/talento90/imgart/imgartapi .

# Expose http port
EXPOSE 4005

# Run our executable application
CMD [ "./imgartapi" ]