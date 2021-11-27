# syntax=docker/dockerfile:experimental

##
## Build
##
FROM golang:1.17-alpine as build

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o personal-backend .

##
## Deploy
##
FROM alpine
RUN apk add --no-cache bash dumb-init

COPY --from=build ["/build/personal-backend", "/"]

EXPOSE 3000
ENV PORT=3000
ENTRYPOINT [ "/usr/bin/dumb-init", "--" ]
CMD [ "/personal-backend" ]