echo Stopping the container...
docker stop ascii-art-web-container
echo -e '\nRemoving the container...'
docker rm ascii-art-web-container
# Clean all the unused objects
docker system prune -f
