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

type Tag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func listTaskDefinitions(ecssvc *ecs.ECS, specifiedname string) []string {
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
	taskdefs := []string{}
	for _, taskdef := range taskDefList.TaskDefinitionArns {
		if strings.Contains(*taskdef, specifiedname) {
			taskdefs = append(taskdefs, *taskdef)
		}
	}
	return taskdefs
}

func listServices(ecssvc *ecs.ECS, cluster string) []string {
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
	services := []string{}
	for _, service := range serviceList.ServiceArns {
		services = append(services, *service)
	}
	return services
}

func listClusters(ecssvc *ecs.ECS) []string {
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
	clusters := []string{}
	for _, cluster := range clusterList.ClusterArns {
		clusters = append(clusters, *cluster)
	}

	return clusters

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

	// fmt.Println("--- Task Definition Arns")
	// taskDefList := listTaskDefinitions(ecssvc, specifiedname)
	// for _, taskdef := range taskDefList {
	// 	fmt.Println(taskdef)
	// }

	fmt.Println("--- Service Arns")

	var cluster string = "rf-sandbox-infratest02-ecsc"
	serviceList := listServices(ecssvc, cluster)
	for _, service := range serviceList {
		fmt.Println(service)
	}

	fmt.Println("--- Cluster Arns")

	clusterList := listClusters(ecssvc)
	for _, cluster := range clusterList {
		fmt.Println(cluster)
	}

	fmt.Println("--- Task Definitions, Tags")

	var specifiedname string = "infratest02"
	taskDefList := listTaskDefinitions(ecssvc, specifiedname)
	for _, taskdef := range taskDefList {
		tasks := describeTaskDefinition(ecssvc, taskdef)
		for _, tag := range tasks.Tags {
			fmt.Println(taskdef, *tag.Value)
		}
	}
}
