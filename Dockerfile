# base image
FROM golang:1.16-alpine

# work dir to house image and run the rest of our Docker commands inside the directory
WORKDIR /app

# copy go.mod file to workdir (/app)
COPY go.mod ./
# install go modules
# RUN go mod download

# copy source code into the image
COPY *.go ./

# build app which create a Go binary with the same name in the root of the Docker image
RUN go build -o /web_health

# run the image
CMD [ "/web_health" ]