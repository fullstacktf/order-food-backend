.DEFAULT_GOAL := build

stop-mongo: 
	@echo "🍉 Stopping mongo database container..."
	@./.scripts/stop_mongo.sh

start-server:  
	@echo "🍔 Starting server..."
	@./.scripts/run_server.sh

start-mongo:  
	@echo "🥑 Starting mongo database..."
	@./.scripts/build_mongo.sh

build: start-mongo start-server