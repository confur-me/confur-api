.PHONY: all deps build

default: all
	
all: deps build

deps:
	@go get github.com/tools/godep
	@godep restore

build:
	@godep go build ./confur.go

install:
	@if [ -a ./confur ] ; \
	then \
		sudo cp ./confur /usr/local/bin; \
	else \
		echo "Run make first"; \
	fi;
