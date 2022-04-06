package main

import (
	"context"
	"encoding/json"
	"fmt"
	"intelowl"
)

func main() {
	welcome := "Welcome"
	fmt.Println(welcome)
	client := intelowl.MakeNewIntelOwlClient(
		&intelowl.IntelOwlClientOptions{
			"http://localhost",
			"ad3ee3377e3c4313aa372d10ad065558",
			"",
		},
		nil,
	)
	ctx := context.Background()
	analyzerResponse, err := client.CreateObservableAnalysis(ctx, &intelowl.ObservableAnalysisParams{
		intelowl.BasicAnalysisParams{
			"ip",
			"",
			"",
			map[string]interface{}{},
			[]string{},
			[]string{},
			[]string{},
		},
		"8.8.8.8",
		"",
	})
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
	} else {
		bytes, _ := json.Marshal(analyzerResponse)
		fmt.Println("JOB ID")
		fmt.Println(analyzerResponse.JobID)
		fmt.Println("JOB ID END")
		fmt.Println("========== ANALYZER RESPONSE ==========")
		fmt.Println(string(bytes))
		fmt.Println("========== ANALYZER RESPONSE END ==========")
	}
	status, err := client.Analyzer.HealthCheck(ctx, "Not an analyzer")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(status)
	}

	status2, err2 := client.Connector.HealthCheck(ctx, "OpenCTI")
	if err2 != nil {
		fmt.Println(err)
	} else {
		fmt.Println(status2)
	}

	// analyzerConfigList, err := client.Analyzer.GetConfigs(ctx)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// fmt.Println(analyzerConfigList)
	// bytes, _ := json.Marshal(analyzerConfigList)
	// fmt.Println("========== ANALYZER CONFIG LIST ==========")
	// fmt.Println(string(bytes))
	// fmt.Println("========== ANALYZER CONFIG LIST END` ==========\n")
	// }

	connectorConfigList, err := client.Connector.GetConfigs(ctx)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(connectorConfigList)
		bytes, _ := json.Marshal(connectorConfigList)
		fmt.Println("========== CONNECTOR CONFIG LIST ==========")
		fmt.Println(string(bytes))
		fmt.Println("========== CONNECTOR CONFIG LIST END` ==========\n")
	}

	job, err := client.Job.Get(ctx, 33)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(job)
		bytes, _ := json.Marshal(job)
		fmt.Println("========== JOB ==========")
		fmt.Println(string(bytes))
		fmt.Println("========== JOB END ==========\n")
	}

	jobList, err := client.Job.List(ctx)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(jobList)
		bytes, _ := json.Marshal(jobList)
		fmt.Println("========== JOB LIST ==========")
		fmt.Println(string(bytes))
		fmt.Println("========== JOB END ==========\n")
	}
}
