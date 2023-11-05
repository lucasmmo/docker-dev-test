up:
	@go mod tidy
	@docker-compose up --build -d

down:
	@rm -rf tmp/
	@docker-compose down