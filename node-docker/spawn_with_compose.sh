#!/bin/bash

docker-compose up --build

docker-compose down

docker-compose run --rm --service-ports node_dev_env