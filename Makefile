DOCKER_IMAGE ?= logikone/vdpcr
DOCKER_TAG ?= latest

GO_MODULES := $(shell go list ./... | grep -v vendor)

fmt:
	@go fmt $(GO_MODULES)

vet:
	@go vet $(GO_MODULES)

test:
	@go test -v $(GO_MODULES)

docker: fmt vet test
	docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) .

push: docker
	docker push $(DOCKER_IMAGE):$(DOCKER_TAG)