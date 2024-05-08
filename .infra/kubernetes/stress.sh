#!/bin/bash

# shellcheck disable=SC2034
for i in {1..100000}; do
  curl https://4m5ipmqfdg.execute-api.us-east-1.amazonaws.com/default/health
#  curl http://localhost:8080/health
  sleep 1
done