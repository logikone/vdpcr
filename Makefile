DOCKER_IMAGE ?= logikone/vdpcr
DOCKER_TAG ?= latest

docker:
	docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) .

push:
	docker push $(DOCKER_IMAGE):$(DOCKER_TAG)