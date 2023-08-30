# syntax=docker/dockerfile:1

#STAGE 1: Build
FROM golang:1.18 as build

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY . .

ARG GITHUB_TOKEN
#RUN git config --global url."https://${GITHUB_TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/"
RUN go mod download


RUN export CGO_ENABLED=0 && go build -o main

#STAGE 2: Deployment
FROM alpine:latest

USER nobody:nobody
COPY --from=build /app/main ./

CMD [ "./main", "-addr", "0.0.0.0" ]
