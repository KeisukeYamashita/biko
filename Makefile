.PHONY: build
build:
	go build -o bin/biko \
	./cmd/biko

.PHONY: install
install: build
	@cp bin/biko /usr/local/bin
	@mkdir $HOME/.biko
	@cp .config/config.ini $HOME/.biko/config.ini