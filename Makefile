.DEFAULT_GOAL := build

server: mongo
	@echo "🍔 Starting server..."
	@docker-compose up app

mongo:
	@echo "🥑 Starting mongo database in detached mode..."
	@docker-compose up -d mongo

stop:
	@echo "❌ Bringing down the container..."
	@docker-compose down

build: mongo server
