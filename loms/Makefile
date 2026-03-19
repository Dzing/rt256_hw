CURDIR=$(shell pwd)
BINDIR=${CURDIR}/bin
PACKAGE=route/loms/cmd/app

build: bindir
	go build -buildvcs=false -o ${BINDIR}/loms ${PACKAGE}

bindir:
	mkdir -p ${BINDIR}

precommit: lint
	@echo "precommit finish"

lint: 
	go vet ./...