FROM golang:latest as builder

WORKDIR /app

COPY . .

RUN go build main.go 

FROM golang:latest

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
