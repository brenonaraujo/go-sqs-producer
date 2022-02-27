
docker run -i -d -t -p 8080:9292 -v $HOME/.aws/credentials:/home/app/.aws/credentials:ro -e AWS_DEFAULT_REGION=us-east-1 producer-sqs:snapshot