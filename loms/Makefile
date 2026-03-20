CURDIR=$(shell pwd)
BINDIR=${CURDIR}/bin
PACKAGE=route/loms/cmd/app

GOVER=$(shell go version | perl -nle '/(go\d\S+)/; print $$1;')
LINTVER=v1.64.8
LINTBIN=${BINDIR}/lint_${GOVER}_${LINTVER}

build: bindir
	go build -o ${BINDIR}/app/loms ${PACKAGE}

bindir:
	mkdir -p ${BINDIR}

precommit: lint
	@echo "precommit finish"

lint: install-lint
	@echo ">>> check with go vet"
	go vet ./...
	@echo ">>> check with golangci-lint"
	${LINTBIN} run

install-lint: bindir
	test -f ${LINTBIN} || \
		(GOBIN=${BINDIR} go install github.com/golangci/golangci-lint/cmd/golangci-lint@${LINTVER} && \
		mv ${BINDIR}/golangci-lint ${LINTBIN})