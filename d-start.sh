echo Building image...
docker image build -f Dockerfile -t ascii-art-web .
if [ $? -eq 0 ]
then
    echo -e '\nList of all images:'
    docker image ls
    echo -e '\nRunning image inside the container...'
    docker run -p 8080:8080 -d --name ascii-art-web-container ascii-art-web
    docker ps -a
    echo -e '\nYou can stop and remove all running containers and perform a cleanup by running d-end.sh script.'
    echo -e '\nGetting into container... Type ls -l to see the contents and exit to quit.'
    docker exec -it ascii-art-web-container /bin/bash
else
    echo Image building failed.
    exit 1
fi