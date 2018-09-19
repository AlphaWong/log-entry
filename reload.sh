#!/bin/bash

docker build . -t log-service && docker run -it --rm -p 7000:80 log-service