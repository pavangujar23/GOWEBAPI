#Login to DockerHub
docker login --username=pavangujar23

#Tag the Docker Image
docker tag gowebserver pavangujar23/gowebserver:latest

#Push the DOcker image to DockerHub
docker push pavangujar23/gowebserver:latest