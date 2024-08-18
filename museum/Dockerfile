FROM golang:1.23 AS build
WORKDIR /go/src
COPY internal/infra/api ./internal/infra/api
COPY cmd/museum/main.go ./cmd/museum/main.go
COPY go.sum .
COPY go.mod .

ENV CGO_ENABLED=0

RUN go build -o api .

FROM scratch AS runtime
COPY --from=build /go/src/api ./
EXPOSE 8080/tcp
ENTRYPOINT ["./api"]
