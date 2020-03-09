build:
	go build -v -o "procnias.exe"  ./cmd/procnias/main.go
build_static:
	CGO_ENABLED=0 GOOS=linux go build -v -o "procnias.exe" -a -ldflags '-extldflags "-static"' ./cmd/procnias/main.go
docker:
	docker build -t procnias .
clean:
	go clean --modcache