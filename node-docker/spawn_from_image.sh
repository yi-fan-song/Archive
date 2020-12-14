#!/bin/bash

docker run --rm -it --name node-docker -v $PWD:/home/app -p 8080:9000 node-dev
