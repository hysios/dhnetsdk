.PHONY: build	

dev:
	@docker run -it -v $(shell pwd):/go/src dahua-netsdk:dev bash

build:
	@go build -o bin/dahuanet-linux ./example

sync:
	@scp bin/dahuanet-linux devserver2:~/dahuasdk
	
dbuild:
	@docker build -t dahua-netsdk:dev .
