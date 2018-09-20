#!/bin/bash
# docker run -p 9090:9090 -v /promotheus-data  prom/prometheus --config.file=/prometheus-data/promconfig.yml

docker run -d -p 9090:9090 \
      -v  $(pwd)/prometheus-data:/prometheus-data \
       prom/prometheus --config.file=/prometheus-data/prometheus.yml

# /Users/henrik/golang/src/votingpower-tracker/prometheus-test
# -v /Users/henrik/golang/src/votingpower-tracker/prometheus-test/