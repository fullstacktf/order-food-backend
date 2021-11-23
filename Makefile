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

set-initial-data: mongo-start
	@echo "🥓 Filling up the mongo database..."
	@go run ./.cache/db/set-initial-data/set-data.go

build: mongo-start server-start
