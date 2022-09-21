FROM golang:1.19.1-alpine3.16 as builder

RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -o main .

FROM alpine:latest
RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/main /app/
CMD ["/app/main"]