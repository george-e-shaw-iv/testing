.PHONY: deps
deps:
	go mod tidy
	go mod download

.PHONY: build
build: deps
	mkdir -p bin/
	go build -o bin/ ./cmd/...

.PHONY: docker-build
docker-build:
	DOCKER_BUILDKIT=1 docker build --ssh default -t github.com/george-e-shaw-iv/testing -f deployments/serverd/Dockerfile .

.PHONY: docker-run
docker-run:
	docker run -dit github.com/george-e-shaw-iv/testing

.PHONY: docker
docker: docker-build docker-run