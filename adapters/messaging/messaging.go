package messaging

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"os"
)

func getSession() *session.Session {
	return session.Must(session.NewSession(&aws.Config{
		Endpoint: aws.String(os.Getenv("AWS_ENDPOINT")),
		Region:   aws.String(os.Getenv("AWS_REGION")),
	}))
}

func getQueueURL(sess *session.Session, queueName string) (*sqs.GetQueueUrlOutput, error) {
	sqsClient := sqs.New(sess)

	result, err := sqsClient.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: &queueName,
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func Publish(queueName string, messageBody string) error {
	sess := getSession()
	sqsClient := sqs.New(sess)

	urlRes, err := getQueueURL(sess, queueName)
	if err != nil {
		return fmt.Errorf("got an error while trying to create queue: %v\n", err)
	}

	queueUrl := urlRes.QueueUrl
	_, err = sqsClient.SendMessage(&sqs.SendMessageInput{
		QueueUrl:    queueUrl,
		MessageBody: aws.String(messageBody),
	})

	return err
}
