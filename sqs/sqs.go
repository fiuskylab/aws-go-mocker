package sqs

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

type mockSQS struct {
	sqsiface.SQSAPI
	messages map[string][]*sqs.Message
}

func (m *mockSQS) SendMessage(in *sqs.SendMessageInput) (*sqs.SendMessageOutput, error) {
	m.messages[*in.QueueUrl] = append(m.messages[*in.QueueUrl], &sqs.Message{
		Body: in.MessageBody,
	})
	return &sqs.SendMessageOutput{}, nil
}
func (m *mockSQS) ReceiveMessage(in *sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error) {
	if len(m.messages[*in.QueueUrl]) == 0 {
		return &sqs.ReceiveMessageOutput{}, nil
	}
	response := m.messages[*in.QueueUrl][0:1]
	m.messages[*in.QueueUrl] = m.messages[*in.QueueUrl][1:]
	return &sqs.ReceiveMessageOutput{
		Messages: response,
	}, nil
}

// Not used. This is just here for reference.
func getRealSQSClient() sqsiface.SQSAPI {
	sess := session.Must(session.NewSession())
	return sqs.New(sess)
}
func getMockSQSClient() sqsiface.SQSAPI {
	return &mockSQS{
		messages: map[string][]*sqs.Message{},
	}
}
