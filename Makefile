build-all:
	cd cart && GOOS=linux GOARCH=amd64 make build
	cd loms && GOOS=linux GOARCH=amd64 make build

precommit:
	cd cart && make precommit
	cd loms && make precommit