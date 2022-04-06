package main

import (
	"context"
	"encoding/json"
	"fmt"
	"intelowl"
)

// TODO: FROM FILE OPTION TO INIT INTELOWLCLIENT

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
	// var r interface{}
	analResp, err := client.CreateObservableAnalysis(ctx, &intelowl.ObservableAnalysisParams{
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
		bytes, _ := json.Marshal(analResp)
		fmt.Println("========== ANALYZER CONFIbghbG LIST ==========")
		fmt.Println(string(bytes))
		fmt.Println("========== ANALYZER CONFIG LIST END` ==========\n")
	}
	// status, err := client.Analyzer.HealthCheck(ctx, "jhfjiwebfjd")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(status)
	// }

	// status2, err2 := client.Connector.HealthCheck(ctx, "OpenCTI")
	// if err2 != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(status2)
	// }

	// tag, err := client.Tag.Get(ctx, 1)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(tag)
	// }

	// allTags, err := client.Tag.List(ctx)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(allTags)
	// }

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

	// connectorConfigList, err := client.Connector.GetConfigs(ctx)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(connectorConfigList)
	// 	bytes, _ := json.Marshal(connectorConfigList)
	// 	fmt.Println("========== CONNECTOR CONFIG LIST ==========")
	// 	fmt.Println(string(bytes))
	// 	fmt.Println("========== CONNECTOR CONFIG LIST END` ==========\n")
	// }

	// job, err := client.Job.Get(ctx, 33)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(job)
	// 	bytes, _ := json.Marshal(job)
	// 	fmt.Println("========== JOB ==========")
	// 	fmt.Println(string(bytes))
	// 	fmt.Println("========== JOB END ==========\n")
	// }

	// jobList, err := client.Job.List(ctx)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(jobList)
	// 	bytes, _ := json.Marshal(jobList)
	// 	fmt.Println("========== JOB LIST ==========")
	// 	fmt.Println(string(bytes))
	// 	fmt.Println("========== JOB END ==========\n")
	// }

	// client := intelowl.IntelOwlClient{
	// "ad3ee3377e3c4313aa372d10ad065558",
	// "http://localhost",
	// "",
	// 	&http.Client{Timeout: time.Duration(10) * time.Second},
	// }

	// analyzerHealthCheck, err := client.AnalyzerHealthCheck(ctx, "dnsjknd")
	// if err != nil {
	// 	fmt.Println("error")
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(*analyzerHealthCheck)
	// 	bytes, _ := json.Marshal(analyzerHealthCheck)
	// 	fmt.Println("========== ANALYZER STATUS ==========")
	// 	fmt.Println(string(bytes))
	// 	fmt.Println("========== ANALYZER STATUS END ==========\n")
	// }

	// connectorHealthCheck, err := client.ConnectorHealthCheck(ctx, "OpenCTI")
	// if err != nil {
	// 	fmt.Println("error")
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(*connectorHealthCheck)
	// 	bytes, _ := json.Marshal(connectorHealthCheck)
	// 	fmt.Println("========== ANALYZER STATUS ==========")
	// 	fmt.Println(string(bytes))
	// 	fmt.Println("========== ANALYZER STATUS END ==========\n")
	// }

	// analyzerConfigList, err := client.GetAnalyzerConfigList(ctx)
	// if err != nil {
	// 	fmt.Println("funck me")
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(*analyzerConfigList)
	// bytes, _ := json.Marshal(analyzerConfigList)
	// fmt.Println("========== ANALYZER CONFIG LIST ==========")
	// fmt.Println(string(bytes))
	// fmt.Println("========== ANALYZER CONFIG LIST END` ==========\n")
	// }

	// connectorConfigList, err := client.GetConnectorConfigList(ctx)
	// if err != nil {
	// 	fmt.Println("funck me")
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(*connectorConfigList)
	// 	bytes, _ := json.Marshal(connectorConfigList)
	// 	fmt.Println("========== CONNECTOR CONFIG LIST ==========")
	// 	fmt.Println(string(bytes))
	// 	fmt.Println("========== CONNECTOR CONFIG LIST END` ==========")
	// }

	// responseJobs, err := client.GetAllJobs(ctx)
	// if err != nil {
	// 	fmt.Println("funck me")
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(*responseJobs)
	// 	bytes, err := json.Marshal(responseJobs)
	// 	if err != nil {
	// 		fmt.Println("lol")
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println("======= ALL JOBS ======")
	// 	fmt.Println(string(bytes))
	// 	fmt.Println("======= ALL JOBS END ======")
	// }

	// responseJob, err2 := client.GetJobById(ctx, 32)
	// if err2 != nil {
	// 	fmt.Println("funck me while getting job by id")
	// 	fmt.Println(err2)
	// } else {
	// 	fmt.Println(*responseJob)
	// 	bytes, err := json.Marshal(responseJob)
	// 	if err != nil {
	// 		fmt.Println("lol")
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println("======= JOB ======")
	// 	fmt.Println(string(bytes))
	// 	fmt.Println("======= JOB END ======")
	// }

	// client.GetAllTags(ctx)
	// responseLabel, err := client.GetTagById(ctx, 1)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(*responseLabel)

	// responseLabels, err2 := client.GetAllTags(ctx)
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }
	// fmt.Println(*responseLabels)
}
