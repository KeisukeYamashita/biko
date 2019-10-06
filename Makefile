.PHONY: build
build:
	@go build -o bin/biko \
		./cmd/biko

.PHONY: dockerbuild
dockerbuild:
	docker build . -t biko

.PHONY: install
install: build
	@cp bin/biko /usr/local/bin
	@mkdir -p ${HOME}/.biko
	@cp .config/default.toml ${HOME}/.biko/config.toml