build:
	docker build -t micro --rm .

run:
	docker run -p 8080:8080 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns \
		-idt --rm \
		--name micro \
		micro api --handler=api