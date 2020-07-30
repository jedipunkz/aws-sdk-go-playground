package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/inspector"
)

const (
	templateArn = "arn:aws:inspector:ap-northeast-1:513123329229:target/0-8wfvNcpK/template/0-R2L5fYYr"
)

// func listAssessmentRuns(svc *inspector.Inspector) (runList *inspector.ListAssessmentRunsOutput) {
func listAssessmentRuns(svc *inspector.Inspector) (listRuns *inspector.ListAssessmentRunsOutput) {
	input := &inspector.ListAssessmentRunsInput{
		AssessmentTemplateArns: []*string{
			aws.String(templateArn),
		},
		MaxResults: aws.Int64(123),
	}

	result, err := svc.ListAssessmentRuns(input)
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
	}

	return result
}

func describeAssessmentRunsArns(svc *inspector.Inspector, listRuns *inspector.ListAssessmentRunsOutput) {
	for _, v := range listRuns.AssessmentRunArns {
		input := &inspector.DescribeAssessmentRunsInput{
			AssessmentRunArns: []*string{
				aws.String(*v),
			},
		}

		result, err := svc.DescribeAssessmentRuns(input)
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
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				fmt.Println(err.Error())
			}
			return
		}
		fmt.Println(result.AssessmentRuns)
	}
}

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
	// fmt.Println(list.FindingArns)
	// fmt.Println(*list.NextToken)
	// for i, v := range list.FindingArns {
	// 	fmt.Println(i)
	// 	fmt.Println(*v)
	// }
	return list
}

func describeFindings(svc *inspector.Inspector, list *inspector.ListFindingsOutput) []float64 {
	var resp []float64

	for _, v := range list.FindingArns {
		input := &inspector.DescribeFindingsInput{
			FindingArns: []*string{
				aws.String(*v),
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
		fmt.Println(*result.Findings[0].NumericSeverity)
		// fmt.Println(*result.Findings[0].Title)
		// if i == 1 {
		// 	fmt.Println(result.Findings.NumericSeverity)
		// 	return resp
		// }
		// severity := result.Findings[0].NumericSeverity
		resp = append(resp, *result.Findings[0].NumericSeverity)
		// fmt.Println(i)
		// fmt.Println(result.Findings[i])
		// fmt.Println(result.Findings)
		// fmt.Println(result.Findings)
	}

	return resp
}

func main() {
	svc := inspector.New(session.New(&aws.Config{
		Region: aws.String("ap-northeast-1"),
		// Credentials: credentials.NewSharedCredentials("", "sandbox"),
	}))
	// svc := inspector.New(session.New())

	listRuns := listAssessmentRuns(svc)
	// for _, v := range listRuns.AssessmentRunArns {
	// 	fmt.Println(*v)
	// }
	describeAssessmentRunsArns(svc, listRuns)

	// prd
	// list := listFindings(svc, "arn:aws:inspector:ap-northeast-1:513123329229:target/0-8wfvNcpK/template/0-R2L5fYYr/run/0-4zTKz1hJ")
	list := listFindings(svc, "arn:aws:inspector:ap-northeast-1:513123329229:target/0-8wfvNcpK/template/0-R2L5fYYr/run/0-I2ckLRxC")
	// sandbox
	// list := listFindings(svc, "arn:aws:inspector:ap-northeast-1:395127550274:target/0-uFsCrU0Z/template/0-xiZgiy9t/run/0-TM9I4yeZ")
	resp := describeFindings(svc, list)
	fmt.Sprintln(resp)

	// describeFindings(svc, list)
}
