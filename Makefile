.PHONY: build-images
build-images:
	docker build -t ghcr.io/keyval-dev/odigos/autoscaler:$(TAG) . --build-arg SERVICE_NAME=autoscaler
	docker build -t ghcr.io/keyval-dev/odigos/scheduler:$(TAG) . --build-arg SERVICE_NAME=scheduler
	docker build -t ghcr.io/keyval-dev/odigos/lang-detector:$(TAG)  -f langDetector/Dockerfile . --build-arg SERVICE_NAME=langDetector
	docker build -t ghcr.io/keyval-dev/odigos/ui:$(TAG) ui/ -f ui/Dockerfile
	docker build -t ghcr.io/keyval-dev/odigos/visioncart:$(TAG) visioncart/ -f visioncart/Dockerfile
	docker build -t ghcr.io/keyval-dev/odigos/init:$(TAG) init/ -f init/Dockerfile
	docker build -t ghcr.io/keyval-dev/odigos/instrumentor:$(TAG) . --build-arg SERVICE_NAME=instrumentor

.PHONY: push-images
push-images:
	docker push ghcr.io/keyval-dev/odigos/autoscaler:$(TAG)
	docker push ghcr.io/keyval-dev/odigos/scheduler:$(TAG)
	docker push ghcr.io/keyval-dev/odigos/lang-detector:$(TAG)
	docker push ghcr.io/keyval-dev/odigos/ui:$(TAG)
	docker push ghcr.io/keyval-dev/odigos/instrumentor:$(TAG)
	docker push ghcr.io/keyval-dev/odigos/init:$(TAG)
	docker push ghcr.io/keyval-dev/odigos/visioncart:$(TAG)

.PHONY: load-to-kind
load-to-kind:
	kind load docker-image ghcr.io/keyval-dev/odigos/autoscaler:$(TAG)
	kind load docker-image ghcr.io/keyval-dev/odigos/scheduler:$(TAG)
	kind load docker-image ghcr.io/keyval-dev/odigos/lang-detector:$(TAG)
	kind load docker-image ghcr.io/keyval-dev/odigos/ui:$(TAG)
	kind load docker-image ghcr.io/keyval-dev/odigos/visioncart:$(TAG)
	kind load docker-image ghcr.io/keyval-dev/odigos/init:$(TAG)
	kind load docker-image ghcr.io/keyval-dev/odigos/instrumentor:$(TAG)