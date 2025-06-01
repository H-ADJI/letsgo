all: db dev
db:
	@docker container start test-mysql
dev:
	@air

