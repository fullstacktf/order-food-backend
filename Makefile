.DEFAULT_GOAL := build

permissions: 
	@sudo chmod +x ./.scripts/stop_mongo.sh ./.scripts/run_server.sh ./.scripts/build_mongo.sh

stop: permissions
	@echo "🍉 Stopping mongo database container..."
	@./.scripts/stop_mongo.sh

server: permissions 
	@echo "🍔 Starting server"
	@./.scripts/run_server.sh

mongo: permissions 
	@echo "🥑 Starting mongo database..."
	@./.scripts/build_mongo.sh