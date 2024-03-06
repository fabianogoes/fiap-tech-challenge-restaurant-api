#!/bin/bash

# shellcheck disable=SC2034
for i in {1..100000}; do
  #curl http://a1481cebca0614312b65f621fb776ec3-1097563495.us-west-2.elb.amazonaws.com:8080/health
  curl http://localhost:8080/health
  sleep 1
done