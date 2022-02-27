
docker run -i -d -t -p 8080:9292 -v $HOME/.aws/credentials:/appuser/.aws/credentials:ro producer-sqs:snapshot