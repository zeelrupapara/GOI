FROM golang:1.19-alpine AS build
ENV CGO_ENABLED=0

WORKDIR /go/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN cp .env.docker /app/.env

RUN go build -o /app/GPAT

FROM docker.io/improwised/golang-base
COPY --from=build /app/ /app/
ENTRYPOINT ["/app/GPAT"]
