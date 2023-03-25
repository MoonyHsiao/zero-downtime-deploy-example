# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="xxxxx <xxxxx@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

COPY . .

RUN go mod init github.com/MoonyHsiao/zero-downtime-deploy-example
RUN go mod tidy 
