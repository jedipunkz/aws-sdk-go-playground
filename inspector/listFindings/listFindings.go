package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/inspector"
)

func listFindings(svc *inspector.Inspector, arn string) (list *inspector.ListFindingsOutput) {
	input := &inspector.ListFindingsInput{
		AssessmentRunArns: []*string{
			aws.String(arn),
		},
		MaxResults: aws.Int64(123),
	}

	list, err := svc.ListFindings(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case inspector.ErrCodeInternalException:
				fmt.Println(inspector.ErrCodeInternalException, aerr.Error())
			case inspector.ErrCodeInvalidInputException:
				fmt.Println(inspector.ErrCodeInvalidInputException, aerr.Error())
			case inspector.ErrCodeAccessDeniedException:
				fmt.Println(inspector.ErrCodeAccessDeniedException, aerr.Error())
			case inspector.ErrCodeNoSuchEntityException:
				fmt.Println(inspector.ErrCodeNoSuchEntityException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
		return
	}
	return list
}

func describeFindings(svc *inspector.Inspector, list *inspector.ListFindingsOutput) {
	for _, i := range list.FindingArns {
		input := &inspector.DescribeFindingsInput{
			FindingArns: []*string{
				aws.String(*i),
			},
		}

		result, err := svc.DescribeFindings(input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case inspector.ErrCodeInternalException:
					fmt.Println(inspector.ErrCodeInternalException, aerr.Error())
				case inspector.ErrCodeInvalidInputException:
					fmt.Println(inspector.ErrCodeInvalidInputException, aerr.Error())
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				fmt.Println(err.Error())
			}
		}
		fmt.Println(result.Findings)
	}
}

func main() {
	svc := inspector.New(session.New(&aws.Config{
		Region:      aws.String("ap-northeast-1"),
		Credentials: credentials.NewSharedCredentials("", "readyfor"),
	}))

	list := listFindings(svc, "arn:aws:inspector:ap-northeast-1:513123329229:target/0-8wfvNcpK/template/0-R2L5fYYr/run/0-4zTKz1hJ")

	describeFindings(svc, list)
}
