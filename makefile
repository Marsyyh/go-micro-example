build:
	for dir in proto/*/ ; do \
		cd $$dir && \
		make build; \
	done
	cd auth && \
	make build
	cd api && \
	make build
run:
	cd auth && \
	make run
	cd api && \
	make run
	micro api
clean:
	docker stop auth-srv
	docker rmi $(shell docker images -f dangling=true -q)