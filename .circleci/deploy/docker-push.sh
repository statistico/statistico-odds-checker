#!/bin/bash

set -e

aws ecr get-login-password --region "$AWS_DEFAULT_REGION" | docker login --username AWS --password-stdin "$AWS_ECR_ACCOUNT_URL/statistico-odds-checker"

docker tag "statistico-odds-checker-console" "$AWS_ECR_ACCOUNT_URL/statistico-odds-checker:latest"
docker push "$AWS_ECR_ACCOUNT_URL/statistico-odds-checker:latest"
