FROM quay.io/fedora/fedora as builder

RUN yum install -y golang-1.22.7

WORKDIR /app

COPY . .

RUN go build main.go 

FROM quay.io/fedora/fedora

RUN yum install -y golang-1.22.7


COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
