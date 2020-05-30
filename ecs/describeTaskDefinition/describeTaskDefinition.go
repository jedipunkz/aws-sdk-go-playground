package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
)

type Tag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func describeTaskDefinition(ecssvc *ecs.ECS, taskdefinition string) *ecs.DescribeTaskDefinitionOutput {
	input := &ecs.DescribeTaskDefinitionInput{
		TaskDefinition: aws.String(taskdefinition),
		Include: []*string{
			aws.String("TAGS"),
		},
	}

	result, err := ecssvc.DescribeTaskDefinition(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ecs.ErrCodeServerException:
				fmt.Println(ecs.ErrCodeServerException, aerr.Error())
			case ecs.ErrCodeClientException:
				fmt.Println(ecs.ErrCodeClientException, aerr.Error())
			case ecs.ErrCodeInvalidParameterException:
				fmt.Println(ecs.ErrCodeInvalidParameterException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
	}
	return result
}

func main() {
	ecssvc := ecs.New(session.New(&aws.Config{
		Region:      aws.String("ap-northeast-1"),
		Credentials: credentials.NewSharedCredentials("", "sandbox"),
	}))

	var taskdefinition string = "arn:aws:ecs:ap-northeast-1:395127550274:task-definition/rf-sandbox-infratest02-ecstd:6"
	task := describeTaskDefinition(ecssvc, taskdefinition)
	for _, tag := range task.Tags {
		fmt.Println(*tag.Value)
	}
}
