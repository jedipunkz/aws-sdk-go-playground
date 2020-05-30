package main

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
)

func listTaskDefinitions(ecssvc *ecs.ECS) *ecs.ListTaskDefinitionsOutput {
	input := &ecs.ListTaskDefinitionsInput{}
	taskDefList, err := ecssvc.ListTaskDefinitions(input)
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
	return taskDefList
}

func listServices(ecssvc *ecs.ECS, cluster string) *ecs.ListServicesOutput {
	input := &ecs.ListServicesInput{
		Cluster: aws.String(cluster),
	}
	serviceList, err := ecssvc.ListServices(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ecs.ErrCodeServerException:
				fmt.Println(ecs.ErrCodeServerException, aerr.Error())
			case ecs.ErrCodeClientException:
				fmt.Println(ecs.ErrCodeClientException, aerr.Error())
			case ecs.ErrCodeInvalidParameterException:
				fmt.Println(ecs.ErrCodeInvalidParameterException, aerr.Error())
			case ecs.ErrCodeClusterNotFoundException:
				fmt.Println(ecs.ErrCodeClusterNotFoundException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
	}
	return serviceList
}

func listClusters(ecssvc *ecs.ECS) *ecs.ListClustersOutput {
	input := &ecs.ListClustersInput{}

	clusterList, err := ecssvc.ListClusters(input)
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
	return clusterList
}

func main() {
	ecssvc := ecs.New(session.New(&aws.Config{
		Region:      aws.String("ap-northeast-1"),
		Credentials: credentials.NewSharedCredentials("", "sandbox"),
	}))

	fmt.Println("--- Task Definition Arns")

	taskDefList := listTaskDefinitions(ecssvc)
	for _, taskdef := range taskDefList.TaskDefinitionArns {
		if strings.Contains(*taskdef, "infratest02") {
			fmt.Println(*taskdef)
		}
	}

	fmt.Println("--- Service Arns")

	var cluster string = "rf-sandbox-infratest02-ecsc"
	serviceList := listServices(ecssvc, cluster)
	for _, service := range serviceList.ServiceArns {
		fmt.Println(*service)
	}

	fmt.Println("--- Cluster Arns")

	clusterList := listClusters(ecssvc)
	for _, cluster := range clusterList.ClusterArns {
		fmt.Println(*cluster)
	}
}
