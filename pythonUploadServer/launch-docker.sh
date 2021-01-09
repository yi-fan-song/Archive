#!/bin/bash

docker run \
	--rm -it \
	--name node-docker \
	-v $PWD:/home/app \
	-w /home/app \
	-p 8000:8000 \
	python:2-alpine python httpServerWithUpload.py 8000