#!/bin/bash

docker run --rm -it --name node-docker -v $PWD:/home/app -w /home/app -e "PORT=9000" -p 8080:9000 -u node node:latest /bin/bash
