ORG := keratin
PROJECT := authn-server
NAME := $(ORG)/$(PROJECT)
VERSION := 1.10.1
MAIN := main.go

.PHONY: clean
clean:
	rm -rf dist

init:
	which ego > /dev/null || go get github.com/benbjohnson/ego/cmd/ego
	ego server/views

# The Linux builder is a Docker container because that's the easiest way to get the toolchain for
# CGO on a MacOS host.
.PHONY: linux-builder
linux-builder:
	docker build -f Dockerfile.linux -t $(NAME)-linux-builder .

# The Linux target is built using a special Docker image, because this Makefile assumes the host
# machine is running MacOS.
dist/authn-linux64: init
	make linux-builder
	docker run --rm \
		-v $(PWD):/go/src/github.com/$(NAME) \
		-w /go/src/github.com/$(NAME) \
		$(NAME)-linux-builder \
		sh -c " \
			GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -ldflags '-extldflags -static -X conf.VERSION=$(VERSION)' -o '$@' \
		"
	bzip2 -c "$@" > dist/authn-linux64.bz2

# The Darwin target is built using the host machine. Only runs on a MacOS host.
dist/authn-macos64: init
ifeq ($(shell uname -s),Darwin)
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 go build -ldflags "-X conf.VERSION=$(VERSION)" -o "$@"
	bzip2 -c "$@" > dist/authn-macos64.bz2
endif

# The Windows target is built using Mingw-w64.
# MacOS: brew install mingw-w64
# Linux: apt install mingw-w64
dist/authn-windows64.exe: init
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -ldflags '-X conf.VERSION=$(VERSION)' -o '$@'

# The Docker target wraps the linux/amd64 binary
.PHONY: dist/docker
dist/docker: dist/authn-linux64
	docker build --tag $(NAME):latest .

# Build all distributables
.PHONY: dist
dist: dist/docker dist/authn-macos64 dist/authn-linux64 dist/authn-windows64.exe

# Run the server
.PHONY: server
server: init
	docker-compose up -d redis
	DATABASE_URL=sqlite3://localhost/dev \
		REDIS_URL=redis://127.0.0.1:8701/11 \
		go run -ldflags "-X conf.VERSION=$(VERSION)" $(MAIN) server

# Run tests
.PHONY: test
test: init
	docker-compose up -d redis mysql postgres
	TEST_REDIS_URL=redis://127.0.0.1:8701/12 \
	  TEST_MYSQL_URL=mysql://root@127.0.0.1:8702/authnservertest \
	  TEST_POSTGRES_URL=postgres://postgres@127.0.0.1:8703/postgres?sslmode=disable \
	  go test -race ./...

# Run CI tests
.PHONY: test-ci
test-ci: init
	TEST_REDIS_URL=redis://127.0.0.1/1 \
	  TEST_MYSQL_URL=mysql://root@127.0.0.1/test \
	  TEST_POSTGRES_URL=postgres://postgres@127.0.0.1/postgres?sslmode=disable \
	  go test -covermode=count -coverprofile=coverage.out ./...

# Run benchmarks
.PHONY: benchmarks
benchmarks:
	docker-compose up -d redis
	TEST_REDIS_URL=redis://127.0.0.1:8701/12 \
		go test -run=XXX -bench=. \
			github.com/keratin/authn-server/server/meta \
			github.com/keratin/authn-server/server/sessions

# Run migrations
.PHONY: migrate
migrate:
	docker-compose up -d redis
	DATABASE_URL=sqlite3://localhost/dev \
		REDIS_URL=redis://127.0.0.1:8701/11 \
		go run -ldflags "-X conf.VERSION=$(VERSION)" $(MAIN) migrate

# Cut a release of the current version.
.PHONY: release
release: dist
	docker push $(NAME):latest
	docker tag $(NAME):latest $(NAME):$(VERSION)
	docker push $(NAME):$(VERSION)
	git push
	git tag v$(VERSION)
	git push --tags
	xdg-open https://github.com/$(NAME)/releases/tag/v$(VERSION)
	xdg-open dist
