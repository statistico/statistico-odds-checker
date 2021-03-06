#!/bin/bash

set -e

aws ecr get-login --no-include-email --region $AWS_DEFAULT_REGION | bash

docker tag "statisticooddschecker_console" "$AWS_ECR_ACCOUNT_URL/statistico-odds-checker:$CIRCLE_SHA1"
docker push "$AWS_ECR_ACCOUNT_URL/statistico-odds-checker:$CIRCLE_SHA1"
