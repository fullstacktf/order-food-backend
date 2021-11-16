.DEFAULT_GOAL := build

stop: 
	@echo "🍉 Stopping mongo database container..."
	@./.scripts/stop_mongo.sh

server:  
	@echo "🍔 Starting server..."
	@./.scripts/run_server.sh

mongo:  
	@echo "🥑 Starting mongo database..."
	@./.scripts/build_mongo.sh