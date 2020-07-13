package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/inspector"
)

type arns struct {
	arn string
}

func (a arn) String() string {
	return fmt.Sprintf("%v", a.arn)
}

func main() {
	svc := inspector.New(session.New(&aws.Config{
		Region:      aws.String("ap-northeast-1"),
		Credentials: credentials.NewSharedCredentials("", "readyfor"),
	}))
	input := &inspector.ListFindingsInput{
		AssessmentRunArns: []*string{
			aws.String("arn:aws:inspector:ap-northeast-1:513123329229:target/0-8wfvNcpK/template/0-R2L5fYYr/run/0-4zTKz1hJ"),
		},
		MaxResults: aws.Int64(123),
	}

	result, err := svc.ListFindings(input)
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
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	var arns []*arn
	fmt.Printf("%v", arns)

	// fmt.Printf("%v \n", result.FindingArns)
	// jsonBytes := []*string(result.FindingArns)
	// fmt.Printf("%+v\n", &result.FindingArns)
	// var output Output
	// json.Unmarshal(jsonBytes, &output)
	// fmt.Println(output.arn)
}
