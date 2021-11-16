#!/bin/bash

# Runs the docker container on port 27017:27017 with name "mongo" on dettached mode
# more details: when stopped it's deleted and it has a volume on ".cache" directory
docker run --rm -d --name mongo -p 27017:27017 -v $(pwd)/.cache:/data/db mongo:5.0.3 
