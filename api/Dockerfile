FROM golang:1.21-alpine3.18 as builder

WORKDIR /app

ADD go.mod ./
ADD go.sum ./
RUN go mod download -x

COPY . .

RUN go build -o vitamart ./cmd/vitamart/main.go

FROM alpine:3.18

COPY --from=builder /app/vitamart .
COPY --from=builder /app/configs/config.yml ./configs/config.yml

CMD ["./vitamart"]

