FROM golang:1.19.4-alpine3.17 as build
WORKDIR /app
COPY src/* /app
RUN go build main.go

FROM scratch
WORKDIR /app
COPY --from=build /app/main /app/main
RUN chmod 755 main
ENTRYPOINT ["/app/main"]
