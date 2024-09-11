FROM docker.io/golang:1.22 AS build

WORKDIR /app

COPY go.mod ./

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /http-debug-server

FROM scratch AS release

WORKDIR /

COPY --from=build /http-debug-server /http-debug-server

ENTRYPOINT ["/http-debug-server"]