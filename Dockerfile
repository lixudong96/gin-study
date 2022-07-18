# syntax=docker/dockerfile:1
FROM golang:1.16-alpine

# NV GOPROXY hhtps://goproxy.cn,direct
WORKDIR /gin-study
COPY . ./
RUN go mod download
RUN go build -o ./gin-study

EXPOSE 3000
ENTRYPOINT ["./gin-study"]
