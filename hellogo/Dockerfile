################ RUN AS ################ 
# docker run --rm -p 8080:8080 hellogo:v1

FROM golang:1.10.2-alpine3.7 AS builder

WORKDIR dist

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app .

# ============================================

FROM alpine:3.7

WORKDIR /app

COPY --from=builder /go/dist/ .

EXPOSE 8080

CMD ["./app"]
