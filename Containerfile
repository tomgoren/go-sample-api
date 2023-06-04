FROM golang:1.19.4-alpine3.17 as build
WORKDIR /app
COPY src/* /app
RUN CGO_ENABLED=0 go build -ldflags="-extldflags=-static -s -w" -o main

FROM scratch
WORKDIR /app
COPY --from=build /app/main /app/main
ENTRYPOINT ["/app/main"]
