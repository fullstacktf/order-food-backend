#!/bin/bash

docker run --rm -d --name mongo -p 27017:27017 -v /src/database/mongo-data:/data/db mongo:5.0.3 # runs mongo:latest image on dettached mode

docker exec -it mongo bash # open the container's bash

# once we're inside the container, run "mongosh"