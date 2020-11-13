docker image build -f Dockerfile -t "gt" .

echo "----------------------------------------------------------------"

docker images

echo "----------------------------------------------------------------"

docker container run -p 9090:8181 --detach --name gtservice gt

echo "----------------------------------------------------------------"

docker ps -a

echo "----------------------------------------------------------------"

docker exec -it gtservice /bin/bash
