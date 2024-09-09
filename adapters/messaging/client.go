package messaging

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/fabianogoes/fiap-challenge/domain/entities"
	"log/slog"
	"path"
)

type AWSSQSClient struct {
	config     *entities.Config
	awsSession *session.Session
	sqsClient  *sqs.SQS
}

func NewAWSSQSClient(config *entities.Config) *AWSSQSClient {
	awsSession := session.Must(session.NewSession(&aws.Config{
		Endpoint: aws.String(config.AwsEndpoint),
		Region:   aws.String(config.AwsRegion),
	}))

	sqsClient := sqs.New(awsSession)

	return &AWSSQSClient{
		config,
		awsSession,
		sqsClient,
	}
}

func (c *AWSSQSClient) Receive(sqsQueueUrl string) *sqs.ReceiveMessageOutput {
	message, err := c.sqsClient.ReceiveMessage(c.ReceiveParams(sqsQueueUrl))
	if err != nil {
		slog.Error(fmt.Sprintf("error receiving message - %v", err))
		return nil
	}

	return message
}

func (c *AWSSQSClient) Publish(queueUrl string, messageBody string) error {
	_, err := c.sqsClient.SendMessage(&sqs.SendMessageInput{
		QueueUrl:    aws.String(queueUrl),
		MessageBody: aws.String(messageBody),
	})

	return err
}

func (c *AWSSQSClient) Delete(receiptHandle *string, queueUrl string) {
	deleteParams := &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(queueUrl),
		ReceiptHandle: receiptHandle,
	}

	_, err := c.sqsClient.DeleteMessage(deleteParams)
	if err != nil {
		slog.Error(fmt.Sprintf("error deleting message - %v", err))
	}
}

func (c *AWSSQSClient) ReceiveParams(queueURL string) *sqs.ReceiveMessageInput {
	return &sqs.ReceiveMessageInput{
		MaxNumberOfMessages: aws.Int64(1),
		QueueUrl:            aws.String(queueURL),
		WaitTimeSeconds:     aws.Int64(5),
	}
}

func (c *AWSSQSClient) CreateQueue(queueName string) {
	slog.Info(fmt.Sprintf("creating queue - %v", queueName))
	input := &sqs.CreateQueueInput{
		QueueName: aws.String(queueName),
	}
	out, err := c.sqsClient.CreateQueue(input)
	if err != nil {
		slog.Error(fmt.Sprintf("error create queue - %v", err))
	}

	slog.Info(fmt.Sprintf("Created queue - %v", *out.QueueUrl))
}

func (c *AWSSQSClient) Init() {
	if c.config.Environment != "production" {
		c.CreateQueue(extractName(c.config.PaymentQueueUrl))
		c.CreateQueue(extractName(c.config.PaymentCallbackQueueUrl))
		c.CreateQueue(extractName(c.config.KitchenQueueUrl))
		c.CreateQueue(extractName(c.config.KitchenCallbackQueueUrl))

		slog.Info("Listing all queues...")
		queues, err := c.sqsClient.ListQueues(&sqs.ListQueuesInput{})
		if err != nil {
			slog.Error(fmt.Sprintf("error listing queues - %v", err))
		}

		for _, queue := range queues.QueueUrls {
			slog.Info(fmt.Sprintf("Listing queue - %v", *queue))
		}
	}
}

func extractName(url string) string {
	return path.Base(url)
}
