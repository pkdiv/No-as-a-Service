# STAGE 1: BUILD STAGE

FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum* ./

RUN go mod download

COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o server .

# STAGE 2: RUNTIME STAGE

FROM alpine:latest

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

COPY --from=builder /app/server .

COPY data/ ./data

RUN chown -R appuser:appgroup /app

USER appuser

EXPOSE 8080

ENV PORT=8080

CMD ["./server"]
