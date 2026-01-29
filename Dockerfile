FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download && go mod tidy

COPY . .

RUN CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    go build -o server main.go

# runtime stage

FROM alpine:latest

WORKDIR /app

RUN adduser -D appuser
USER appuser

COPY --from=builder /app/server .

EXPOSE 8080

ENV PORT=8080

CMD [ "./server" ]