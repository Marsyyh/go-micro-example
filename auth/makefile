build:
	GOOS=linux GOARCH=amd64 go build
	docker build --rm -t auth-svc .

run:
	docker run -p 50001:50001 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns \
		-idt --rm \
		--name auth-svc\
		auth-svc