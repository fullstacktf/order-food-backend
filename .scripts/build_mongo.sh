#!/bin/bash

docker run --rm -d --name mongo -p 27017:27017 -v $(pwd)/.cache:/data/db mongo:5.0.3 # runs mongo:latest image on dettached mode
