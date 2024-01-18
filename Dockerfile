FROM golang:1.20 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o CC-project .

FROM alpine:latest
RUN apk --no-cache add tzdata # add time zones to alpine
WORKDIR /app
COPY --from=builder /app/CC-project .

CMD ["./CC-project"]