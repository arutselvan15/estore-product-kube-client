GO=GOOS=linux GOARCH=amd64 GO111MODULE=on go

CR_REPO_PATH=github.com/arutselvan15/estore-product-kube-client
CR_GROUP=estore
CR_VERSION=v1

init-pkg-dir:
	@echo "==> Init pkg dir..."
	mkdir -p pkg/apis/${CR_GROUP}/${CR_VERSION}
	touch pkg/apis/${CR_GROUP}/${CR_VERSION}/doc.go
	touch pkg/apis/${CR_GROUP}/${CR_VERSION}/register.go
	touch pkg/apis/${CR_GROUP}/${CR_VERSION}/types.go

mod-init:
	@echo "==> Mod Init..."
	${GO} mod init

all: clean deps fmt check test

clean:
	@echo "==> Cleaning..."
	rm -f report.json coverage.out

deps:
	@echo "==> Getting Dependencies..."
	${GO} mod tidy
	${GO} mod download

fmt:
	@echo "==> Code Formatting..."
	${GO} fmt ./...

check: fmt
	@echo "==> Code Check..."
	golangci-lint run --fast --tests

test: clean
	@echo "==> Testing..."
	CGO_ENABLED=0 ${GO} test -v -covermode=atomic -count=1 ./... -coverprofile coverage.out
	CGO_ENABLED=1 ${GO} test -race -covermode=atomic -count=1 ./... -json > report.json
	${GO} tool cover -func=coverage.out

gen-client:
	@echo "==> Generate go client..."
	cd ${GOPATH}/src/k8s.io/code-generator; ./generate-groups.sh all ${CR_REPO_PATH}/pkg/client ${CR_REPO_PATH}/pkg/apis ${CR_GROUP}:${CR_VERSION} --go-header-file ./hack/boilerplate.go.txt

rm-client:
	@echo "==> Remove go client..."
	rm -rf ${GOPATH}/src/${CR_REPO_PATH}/pkg/client ${GOPATH}/src/${CR_REPO_PATH}/pkg/apis/${CR_GROUP}/${CR_VERSION}/zz_generated.deepcopy.go
