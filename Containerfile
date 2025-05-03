FROM golang:1.24.2-alpine as build
WORKDIR /app
COPY src/* /app
RUN CGO_ENABLED=0 go build -ldflags="-extldflags=-static -s -w" -o main

FROM scratch
WORKDIR /app
COPY --from=build /app/main /app/main
ENTRYPOINT ["/app/main"]
