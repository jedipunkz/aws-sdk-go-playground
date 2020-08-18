package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/guregu/dynamo"
)

// EcsService is struct for ECS Service
type EcsService struct {
	svc               *ecs.ECS      `dynamo:"-"` // Ignored
	Name              string        `dynamo:"name"`
	ClusterArn        string        `dynamo:"cluster_arn"`
	TaskDefinitionArn string        `dynamo:"task_definition_arn"`
	ServiceArn        string        `dynamo:"service_arn"`
	UpdateAt          time.Time     `dynamo:"update_at"`
	LoadBalancer      *LoadBalancer `dynamo:"-"` // Ignored
}

// LoadBalancer is struct for ALB
type LoadBalancer struct {
	svc             *elbv2.ELBV2 `dynamo:"-"` // Ignored
	Name            string       `dynamo:"name"`
	Elbv2Arn        string       `dynamo:"elbv2_arn"`
	TargetGroupArn  string       `dynamo:"target_group_arn"`
	ListenerArn     string       `dynamo:"listener_arn"`
	ListenerRuleArn string       `dynamo:"listener_rule_arn"`
	DNSName         string       `dynamo:"dns_name"`
}

func main() {
	db := dynamo.New(session.New(), &aws.Config{Region: aws.String("ap-northeast-1")})
	table := db.Table("ecs_service")

	// テーブルの全ての Items を検索
	var results []EcsService
	// 取得した結果は構造体の要素に入るだけで表示はされないし戻り値もなし
	err := table.Scan().All(&results)
	if err != nil {
		fmt.Println(err)
	}

	// 検索結果全て表示
	fmt.Println(&results)
	fmt.Println("---")
	// スライスを range で回して一個ずつ表示
	for i := range results {
		fmt.Println(results[i].ClusterArn)
		fmt.Println(results[i].Name)
		fmt.Println(results[i].ServiceArn)
		fmt.Println(results[i].TaskDefinitionArn)
		fmt.Println(results[i].UpdateAt)
	}

}
