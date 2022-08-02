# Dockerfile References: https://docs.docker.com/engine/reference/builder/

FROM golang:latest

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apt-get update && apt-get -y upgrade && \
    apt-get -y install bash git

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Expose port 8067 to the outside world
EXPOSE 8067

# Run the executable
CMD ["./start.sh"]