FROM golang:1.23 as build
WORKDIR /app
COPY . .
RUN go mod download && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o glofox main.go

FROM gcr.io/distroless/base-debian11
WORKDIR /app
COPY --from=build /app/glofox .
CMD ["./glofox"]
