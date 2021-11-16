.DEFAULT_GOAL := build

mongo-stop: 
	@echo "🍉 Stopping mongo database container..."
	@./.scripts/stop_mongo.sh

server-start:  
	@echo "🍔 Starting server..."
	@./.scripts/run_server.sh

mongo-start:  mongo-stop
	@echo "🥑 Starting mongo database..."
	@./.scripts/build_mongo.sh

build: mongo-start server-start