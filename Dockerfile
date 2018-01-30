FROM golang
LABEL maintainer "Marco Talento <marcotalento90@gmail.com>"

# We need to add ${GOPATH}/bin to PATH to have access dlv
ENV PATH ${GOPATH}/bin:$PATH

# Setting working directory
WORKDIR ${GOPATH}/src/github.com/talento90/gorpo

# Copy source code
COPY . .

# Get delve debugger
RUN go get github.com/derekparker/delve/cmd/dlv

# Build our application
RUN go build -o gorpoapi cmd/gorpoapi/main.go

# Expose server and debug port
EXPOSE 4005 2345

# Execute our application
CMD ["/gorpoapi"]