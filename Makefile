.DEFAULT_GOAL := build

server-debug: mongo
	@echo "🍉 Starting server with debug mode..."
	@docker-compose up app debug-app

server: mongo
	@echo "🍔 Starting server without debug mode..."
	@docker-compose up app

mongo:
	@echo "🥑 Starting mongo database in detached mode..."
	@docker-compose up -d mongo

stop:
	@echo "❌ Bringing down the container..."
	@docker-compose down

build: mongo server
