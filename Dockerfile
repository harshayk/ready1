# golang image where workspace (GOPATH) configured at /go.
FROM golang:latest

# Copy the local package files to the container’s workspace.
ADD . /go/src/github.com/harshayk/harsha

# Build the golang-docker command inside the container.
RUN go install github.com/harshayk/harsha

# Run the golang-docker command when the container starts.
ENTRYPOINT /go/bin/harsha

# http server listens on port 8080.
EXPOSE 8080
