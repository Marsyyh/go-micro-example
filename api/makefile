build:
	GOOS=linux GOARCH=amd64 go build
	docker build --rm -t api-srv .

run:
	docker run -p 50002:50001 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns \
		--rm \
		--name api-srv \
		api-srv