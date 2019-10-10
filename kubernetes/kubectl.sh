#!/usr/bin/env bash

kubectl create -f db-service.yaml,db-deployment.yaml,billboardapi-service.yaml,billboardapi-deployment.yaml,billboardapi-claim0-persistentvolumeclaim.yaml

