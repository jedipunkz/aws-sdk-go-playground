package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/inspector"
)

const (
	templateArn = "arn:aws:inspector:ap-northeast-1:513123329229:target/0-8wfvNcpK/template/0-R2L5fYYr"
	region      = "ap-northeast-1"
)

// Run is struct for AssessmentRunArn
type Run struct {
	Date time.Time
	Arn  string
}

// Runs is struct for AssessmentRunArns
type Runs []Run

// Finding is struct NumericSeverity and Description
type Finding struct {
	Severity float64
	Desc     string
}

// Findings is struct for Findings
type Findings []Finding

func listAssessmentRuns(svc *inspector.Inspector, templateArn string) *inspector.ListAssessmentRunsOutput {
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
			fmt.Println(err.Error())
		}
	}
	return result
}

func describeAssessmentRuns(svc *inspector.Inspector, assessmentRuns *inspector.ListAssessmentRunsOutput) string {
	var runs Runs

	for _, v := range assessmentRuns.AssessmentRunArns {
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
				fmt.Println(err.Error())
			}
			return ""
		}
		run := Run{
			Date: *result.AssessmentRuns[0].StartedAt,
			Arn:  *result.AssessmentRuns[0].Arn,
		}

		runs = append(runs, run)
	}
	arn := getMostRecentAssessmentRun(runs)
	return arn
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
	return list
}

func describeFindings(svc *inspector.Inspector, list *inspector.ListFindingsOutput) Findings {
	var findings Findings

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
		finding := Finding{
			Severity: *result.Findings[0].NumericSeverity,
			Desc:     *result.Findings[0].Description,
		}

		findings = append(findings, finding)
	}

	return findings
}

func getMostRecentAssessmentRun(runs Runs) string {
	var max int = 0
	var arn string = ""
	for _, v := range runs {
		if int(v.Date.YearDay()) > max {
			max = int(v.Date.YearDay())
			arn = v.Arn
		}
	}
	return arn
}

func handleNotify(result Findings) {
	for _, v := range result {
		if int(v.Severity) >= 9 {
			fmt.Println(v.Desc)
			fmt.Println("---")
		}
	}
}

func main() {
	svc := inspector.New(session.New(&aws.Config{
		Region: aws.String(region),
	}))

	listTemplateArns := listAssessmentRuns(svc, templateArn)
	arn := describeAssessmentRuns(svc, listTemplateArns)
	listFindings := listFindings(svc, arn)
	result := describeFindings(svc, listFindings)
	handleNotify(result)
}
