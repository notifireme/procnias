
FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR /app
ADD . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -o "procnias" -a -ldflags '-extldflags "-static"' ./cmd/procnias/main.go

FROM gcr.io/distroless/static
COPY --from=builder /app/procnias /procnias
COPY ./config.yml /config.yml
CMD ["/procnias"]
