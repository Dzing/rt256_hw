CURDIR=$(shell pwd)
BINDIR=${CURDIR}/bin
PACKAGE=route/cart/cmd/app

build: bindir
	go build -buildvcs=false -o ${BINDIR}/cart ${PACKAGE}

bindir:
	mkdir -p ${BINDIR}

precommit: lint
	@echo "precommit finish"

lint: 
	go vet ./...