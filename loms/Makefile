CURDIR=$(shell pwd)
BINDIR=${CURDIR}/bin
PACKAGE=atlas.chr/vaa/hw/loms/cmd/app

build: bindir
	go build -o ${BINDIR}/loms ${PACKAGE}

bindir:
	mkdir -p ${BINDIR}

precommit: lint
	@echo "precommit finish"

lint: 
	go vet ./...