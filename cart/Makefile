CURDIR=$(shell pwd)
BINDIR=${CURDIR}/bin
PACKAGE=atlas.chr/vaa/hw/cart/cmd/app

build: bindir
	go build -o ${BINDIR}/cart ${PACKAGE}

bindir:
	mkdir -p ${BINDIR}

precommit: lint
	@echo "precommit finish"

lint: 
	go vet ./...