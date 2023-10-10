FROM golang:1.21.2-alpine3.18 as build
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o ./api main.go

FROM alpine
WORKDIR /app
COPY --from=build /app/api .
CMD ["./api"]