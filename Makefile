.PHONY: all deps build

default: all
	
all: deps build

deps:
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
