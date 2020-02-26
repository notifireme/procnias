build:
	CGO_ENABLED=0 GOOS=linux go build -v -o "procnias.exe" -a -ldflags '-extldflags "-static"' ./cmd/procnias/main.go
docker:
	docker build -t procnias .
clean:
	go clean --modcache