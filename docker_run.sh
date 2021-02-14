#!/bin/bash

echo "[###] BUILDING"
docker build -t forum --progress=plain .

echo "[###] REMOVE OLD PROCESS"
docker container stop $(docker ps -q)

echo "[###] RUN ..."
docker run -p 8000:8000 --rm -d forum

# docker run -v /tmp/forum-db:/src/forum.db

echo "[###] DONE!"  