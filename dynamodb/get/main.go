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
	// EcsService
	db := dynamo.New(session.New(), &aws.Config{Region: aws.String("ap-northeast-1")})
	ecsTable := db.Table("ecs_service")

	// Item を検索して表示: 要グローバルセカンダリインデックス
	var results2 EcsService
	// err := ecsTable.Get("name", "pr-fuga-test").Index("name-index").One(&results2)
	err := ecsTable.Get("name", "pr-fuga-test").One(&results2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(results2.Name)
	fmt.Println(results2.ClusterArn)
	fmt.Println(results2.ServiceArn)
	fmt.Println(results2.TaskDefinitionArn)
	fmt.Println(results2.UpdateAt)

	// テーブルの全ての Items を検索
	// var results []EcsService
	// 取得した結果は構造体の要素に入るだけで表示はされないし戻り値もなし
	// err := ecsTable.Scan().All(&results)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// 検索結果全て表示
	// fmt.Println(&results)
	// fmt.Println("---")
	// // スライスを range で回して一個ずつ表示
	// for i := range results {
	// 	fmt.Println(results[i].ClusterArn)
	// 	fmt.Println(results[i].Name)
	// 	fmt.Println(results[i].ServiceArn)
	// 	fmt.Println(results[i].TaskDefinitionArn)
	// 	fmt.Println(results[i].UpdateAt)
	// 	fmt.Println(results[i].LoadBalancer)
	// }

	// LoadBalancer
	// lbTable := db.Table("load_balancer")
	//
	// var lbresults []LoadBalancer
	// err = lbTable.Scan().All(&lbresults)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	//
	// fmt.Println(&lbresults)
	// for j := range lbresults {
	// 	fmt.Println(lbresults[j].DNSName)
	// 	fmt.Println(lbresults[j].Elbv2Arn)
	// 	fmt.Println(lbresults[j].ListenerArn)
	// 	fmt.Println(lbresults[j].ListenerRuleArn)
	// 	fmt.Println(lbresults[j].TargetGroupArn)
	// 	fmt.Println(lbresults[j].Name)
	// }
}
