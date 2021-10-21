package sqs

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	q := getMockSQSClient()
	queueURL := "https://queue.amazonaws.com/80398EXAMPLE/MyQueue"
	{
		_, err := q.SendMessage(&sqs.SendMessageInput{
			MessageBody: aws.String("Hello, World!"),
			QueueUrl:    &queueURL,
		})
		assert.Nil(t, err)
	}
	{
		_, err := q.ReceiveMessage(&sqs.ReceiveMessageInput{
			QueueUrl: &queueURL,
		})
		assert.Nil(t, err)
	}
}
