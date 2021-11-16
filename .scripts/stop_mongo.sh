
if [ "$(docker ps -q -f name=mongo)" ]; then
    docker stop mongo  # stops the mongo container
fi