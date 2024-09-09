# Localstack

## Running

```shell
docker-compose up -d
```

1. sign-in https://app.localstack.cloud/
2. dashboard: https://app.localstack.cloud/dashboard

## AWS CLI configuration

```shell
sudo apt update
sudo apt install python3
python3 --version

sudo apt install build-essential zlib1g-dev libncurses5-dev libgdbm-dev libnss3-dev libssl-dev libreadline-dev libffi-dev wget
sudo apt install -y python3-pip
pip3 --version

pip install awscli-local
awslocal --version
```

### awslocal

```shell
awslocal iam create-user --user-name localstack-user
awslocal iam create-access-key --user-name localstack-user
awslocal sts get-session-token

awslocal sqs create-queue --queue-name order-payment-queue
awslocal sns create-topic --name order-payment-events
awslocal sqs delete-queue --queue-url https://localhost.localstack.cloud:4566/000000000000/order-payment-queue
awslocal sns subscribe --topic-arn arn:aws:sns:us-east-1:000000000000:order-payment-events --protocol sqs --notification-endpoint arn:aws:sqs:us-east-1:000000000000:order-payment-queue

awslocal sqs list-queues
awslocal sqs receive-message --queue-url http://sqs.us-east-1.localhost.localstack.cloud:4566/000000000000/order-payment-queue
awslocal sqs purge-queue --queue-url http://sqs.us-east-1.localhost.localstack.cloud:4566/000000000000/order-payment-queue
```

### aws

```shell
aws --endpoint-url=http://localhost:4566 sns create-topic --name order-payment-events --region us-east-1 --profile localstack
aws --endpoint-url=http://localhost:4566 sqs create-queue --queue-name order-payment-queue --region us-east-1 --profile localstack
aws --endpoint-url=http://localhost:4566 sns subscribe --topic-arn arn:aws:sns:us-east-1:000000000000:order-payment-events --protocol sqs --notification-endpoint arn:aws:sqs:us-east-1:000000000000:order-payment-queue
```

## References

- https://docs.localstack.cloud/overview
- https://docs.localstack.cloud/user-guide/aws/sts/
- https://medium.com/@valdemarjuniorr/como-configurar-localstack-para-mapear-servicos-da-aws-c8c25e6363b4
- https://www.zup.com.br/blog/localstack-simule-ambientes-aws-localmente