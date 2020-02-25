
FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -o "procnias" -a -ldflags '-extldflags "-static"' ./cmd/procnias/main.go

FROM scratch
COPY --from=builder /app/procnias /procnias
ENTRYPOINT ["/procnias"]
