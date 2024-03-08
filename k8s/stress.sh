#!/bin/bash

# shellcheck disable=SC2034
for i in {1..100000}; do
  curl http://adf599c66d5ad49d0ab6f61b0790c2be-462296566.us-west-2.elb.amazonaws.com:8080/health
#  curl http://localhost:8080/health
  sleep 1
done