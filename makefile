build:
	for dir in proto/*/ ; do \
		cd $$dir && \
		make build; \
	done
	cd auth && \
	make build
	cd api && \
	make build
	cd micro && \
	make build
run:
	cd auth && \
	make run
	cd api && \
	make run
	cd micro && \
	make run 
clean:
	docker stop auth-svc api-svc micro
	docker rmi $(shell docker images -f dangling=true -q)