FROM golang:1.19-alpine AS build

ENV CGO_ENABLED=0

WORKDIR /go/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN cp .env.docker .env

RUN go build -o bin/GPAT

ENTRYPOINT ["bin/GPAT"]

CMD ["api"]
