package intelowl

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type IntelOwlClientOptions struct {
	Url         string
	Token       string
	Certificate string
}

type successResponse struct {
	Data interface{}
}

type IntelOwlClient struct {
	options   *IntelOwlClientOptions
	client    *http.Client
	Tag       *Tag
	Analyzer  *Analyzer
	Connector *Connector
	Job       *Job
}

type errorResponse struct {
	Detail string
}

func (client *IntelOwlClient) makeRequest(ctx context.Context, request *http.Request, typeOfData interface{}) error {
	request = request.WithContext(ctx)

	request.Header.Set("Content-Type", "application/json")

	tokenString := fmt.Sprintf("token %s", client.options.Token)

	request.Header.Set("Authorization", tokenString)
	response, err := client.client.Do(request)

	if err != nil {
		// return nil, err
		fmt.Println("bb")
		fmt.Println(err)
		return err
	}

	defer response.Body.Close()

	statusCode := response.StatusCode
	if statusCode < http.StatusOK || statusCode >= http.StatusBadRequest {
		var errorResp errorResponse
		msgBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("aa")
			errorMessage := fmt.Sprintf("Could not convert JSON response. Status code: %d", statusCode)
			return errors.New(errorMessage)
		}
		fmt.Println("JBFJDBF")
		fmt.Println(string(msgBytes))
		json.Unmarshal(msgBytes, &errorResp)
		errorMessage := fmt.Sprintf("Error Message: %s", errorResp.Detail)
		return errors.New(errorMessage)
	}

	sucessResp := successResponse{
		Data: typeOfData,
	}
	msgBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Could not convert JSON response. Status code: %d", statusCode)
	}
	fmt.Println("IN MAKE REQUEST")
	fmt.Println(string(msgBytes))
	fmt.Println("IN MAKE REQUEST END")
	json.Unmarshal(msgBytes, &sucessResp.Data)
	// fmt.Println(sucessResp.Data)
	return nil
}

type Tag struct {
	client *IntelOwlClient
	ID     int    `json:"id"`
	Label  string `json:"label"`
	Color  string `json:"color"`
}

type ConfigType struct {
	Queue         string `json:"queue"`
	SoftTimeLimit int    `json:"soft_time_limit"`
}

type Secret struct {
	EnvironmentVariableKey string `json:"env_var_key"`
	Description            string `json:"description"`
	Required               bool   `json:"required"`
}

type Parameter struct {
	Value       string `json:"value"`
	Type        int    `json:"type"`
	Description string `json:"description"`
}

type VerificationType struct {
	Configured     bool     `json:"configured"`
	ErrorMessage   string   `json:"error_message"`
	MissingSecrets []string `json:"missing_secrets"`
}

type BaseConfigurationType struct {
	Name         string                      `json:"name"`
	PythonModule string                      `json:"python_module"`
	Disabled     bool                        `json:"disabled"`
	Description  string                      `json:"description"`
	Config       ConfigType                  `json:"config"`
	Secrets      map[string]Secret           `json:"secrets"`
	Params       map[string]Parameter        `json:"params"`
	Verification map[string]VerificationType `json:"verification"`
}

type AnalyzerConfiguration struct {
	BaseConfigurationType
	Type                  string   `json:"type"`
	ExternalService       bool     `json:"external_service"`
	LeaksInfo             bool     `json:"leaks_info"`
	DockerBased           bool     `json:"docker_based"`
	RunHash               bool     `json:"run_hash"`
	RunHashType           string   `json:"run_hash_type"`
	SupportedFiletypes    []string `json:"supported_filetypes"`
	NotSupportedFiletypes []string `json:"not_supported_filetypes"`
	ObservableSupported   []string `json:"observable_supported"`
}

type ConnectorConfiguration struct {
	BaseConfigurationType
	MaximumTlp string `json:"maximum_tlp"`
}

type Analyzer struct {
	client *IntelOwlClient
}

type Connector struct {
	client *IntelOwlClient
}

type statusResponse struct {
	Status bool `json:"status"`
}

type Report struct {
	Name                 string
	Status               string
	Report               map[string]interface{}
	Errors               []string
	ProcessTime          int
	StartTime            time.Time
	Endime               time.Time
	RuntimeConfiguration map[string]interface{}
}

type Job struct {
	ID                       int        `json:"id"`
	Tags                     []Tag      `json:"tags"`
	AnalyzerReports          []Report   `json:"analyzer_reports,omitempty"`
	ConnectorReports         []Report   `json:"connector_reports,omitempty"`
	Source                   string     `json:"source"`
	IsSample                 bool       `json:"is_sample"`
	Md5                      string     `json:"md5"`
	ObservableName           string     `json:"observable_name"`
	ObservableClassification string     `json:"observable_classification"`
	FileName                 string     `json:"file_name"`
	FileMimetype             string     `json:"file_mimetype"`
	Status                   string     `json:"status"`
	AnalyzersRequested       []string   `json:"analyzers_requested,omitempty" `
	ConnectorsRequested      []string   `json:"connectors_requested,omitempty"`
	AnalyzersToExecute       []string   `json:"analyzers_to_execute,omitempty"`
	ConnectorsToExecute      []string   `json:"connectors_to_execute,omitempty"`
	ReceivedRequestTime      *time.Time `json:"received_request_time,omitempty"`
	FinishedAnalysisTime     *time.Time `json:"finished_analysis_time,omitempty"`
	Tlp                      string     `json:"tlp,omitempty"`
	Errors                   []string   `json:"errors,omitempty"`
	ProcessTime              float64    `json:"process_time,omitempty"`
	NoOfAnalyzersExecuted    string     `json:"no_of_analyzers_executed,omitempty"`
	NoOfConnectorsExecuted   string     `json:"no_of_connectors_executed,omitempty"`
	client                   *IntelOwlClient
}

func MakeNewIntelOwlClient(options *IntelOwlClientOptions, httpClient *http.Client) IntelOwlClient {
	if httpClient == nil {
		httpClient = &http.Client{Timeout: time.Duration(10) * time.Second}
	}
	client := IntelOwlClient{
		options: options,
		client:  httpClient,
	}
	client.Tag = &Tag{client: &client}
	client.Analyzer = &Analyzer{client: &client}
	client.Connector = &Connector{client: &client}
	client.Job = &Job{client: &client}
	return client
}

type BasicAnalysisParams struct {
	Source               string                 `json:"source"`
	Tlp                  string                 `json:"tlp"`
	CheckExistingOrForce string                 `json:"check_existing_or_force"`
	RuntimeConfiguration map[string]interface{} `json:"runtime_configuration"`
	AnalyzersRequested   []string               `json:"analyzers_requested"`
	ConnectorsRequested  []string               `json:"connectors_requested"`
	TagsLabels           []string               `json:"tags_labels"`
}

func (basicAnalysisParams *BasicAnalysisParams) GetTLPList() []string {
	return []string{
		"WHITE",
		"GREEN",
		"AMBER",
		"RED",
	}
}

func contains(value string, array []string) bool {
	for _, elem := range array {
		if value == elem {
			return true
		}
	}
	return false
}

type ObservableAnalysisParams struct {
	BasicAnalysisParams
	ObservableName           string `json:"observable_name"`
	ObservableClassification string `json:"classification"`
}

type FileAnalysisParams struct {
	BasicAnalysisParams
}

type AnalysisResponse struct {
	JobID             int      `json:"job_id"`
	Status            string   `json:"status"`
	Warnings          []string `json:"warnings"`
	AnalyzersRunning  []string `json:"analyzers_running"`
	ConnectorsRunning []string `json:"connectors_running"`
}

func (client *IntelOwlClient) CreateObservableAnalysis(ctx context.Context, params *ObservableAnalysisParams) (*AnalysisResponse, error) {
	requestUrl := fmt.Sprintf("%s/api/analyze_observable", client.options.Url)

	tlpList := params.GetTLPList()

	found := contains(params.Tlp, tlpList)

	if params.Tlp == "" || !found {
		params.Tlp = tlpList[0]
	}

	jsonData, _ := json.Marshal(params)

	fmt.Println("JSON DATA")
	fmt.Println(string(jsonData))

	request, err := http.NewRequest("POST", requestUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("???")
		return nil, err
	}

	analysisResponse := AnalysisResponse{}
	if err := client.makeRequest(ctx, request, &analysisResponse); err != nil {
		fmt.Println(">>>>")
		return nil, err
	}
	return &analysisResponse, nil

}

func (tag *Tag) Get(ctx context.Context, tagId int) (*Tag, error) {
	requestUrl := fmt.Sprintf("%s/api/tags/%d", tag.client.options.Url, tagId)
	fmt.Println(requestUrl)
	request, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return nil, err
	}
	tagResponse := Tag{}
	if err := tag.client.makeRequest(ctx, request, &tagResponse); err != nil {
		return nil, err
	}
	return &tagResponse, nil
}

func (tag *Tag) List(ctx context.Context) (*[]Tag, error) {
	requestUrl := fmt.Sprintf("%s/api/tags", tag.client.options.Url)
	fmt.Println(requestUrl)
	request, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return nil, err
	}
	tagResponse := []Tag{}
	if err := tag.client.makeRequest(ctx, request, &tagResponse); err != nil {
		return nil, err
	}
	return &tagResponse, nil
}

func (analyzer *Analyzer) GetConfigs(ctx context.Context) (*[]AnalyzerConfiguration, error) {
	requestUrl := fmt.Sprintf("%s/api/get_analyzer_configs", analyzer.client.options.Url)
	fmt.Println(requestUrl)
	request, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return nil, err
	}
	analyzerConfigurationResponse := map[string]AnalyzerConfiguration{}
	if err := analyzer.client.makeRequest(ctx, request, &analyzerConfigurationResponse); err != nil {
		return nil, err
	}
	analyzerConfigurationList := []AnalyzerConfiguration{}
	for _, analyzerConfig := range analyzerConfigurationResponse {
		analyzerConfigurationList = append(analyzerConfigurationList, analyzerConfig)
	}

	return &analyzerConfigurationList, nil
}

func (analyzer *Analyzer) HealthCheck(ctx context.Context, analyzerName string) (bool, error) {
	requestUrl := fmt.Sprintf("%s/api/analyzer/%s/healthcheck", analyzer.client.options.Url, analyzerName)
	request, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return false, err
	}
	status := statusResponse{}
	if err := analyzer.client.makeRequest(ctx, request, &status); err != nil {
		return false, err
	}
	return status.Status, nil
}

func (connector *Connector) GetConfigs(ctx context.Context) (*[]ConnectorConfiguration, error) {
	requestUrl := fmt.Sprintf("%s/api/get_analyzer_configs", connector.client.options.Url)
	fmt.Println(requestUrl)
	request, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return nil, err
	}
	connectorConfigurationResponse := map[string]ConnectorConfiguration{}
	if err := connector.client.makeRequest(ctx, request, &connectorConfigurationResponse); err != nil {
		return nil, err
	}
	connectorConfigurationList := []ConnectorConfiguration{}
	for _, connectorConfig := range connectorConfigurationResponse {
		connectorConfigurationList = append(connectorConfigurationList, connectorConfig)
	}

	return &connectorConfigurationList, nil
}

func (connector *Connector) HealthCheck(ctx context.Context, connectorName string) (bool, error) {
	requestUrl := fmt.Sprintf("%s/api/connector/%s/healthcheck", connector.client.options.Url, connectorName)
	request, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return false, err
	}
	status := statusResponse{}
	if err := connector.client.makeRequest(ctx, request, &status); err != nil {
		return false, err
	}
	return status.Status, nil
}

func (job *Job) List(ctx context.Context) (*[]Job, error) {
	requestUrl := fmt.Sprintf("%s/api/jobs", job.client.options.Url)
	request, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return nil, err
	}
	responseJobs := []Job{}
	if err := job.client.makeRequest(ctx, request, &responseJobs); err != nil {
		return nil, err
	}
	return &responseJobs, nil
}

func (job *Job) Get(ctx context.Context, id int) (*Job, error) {
	requestUrl := fmt.Sprintf("%s/api/jobs/%d", job.client.options.Url, id)
	request, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return nil, err
	}
	responseJobs := Job{}
	if err := job.client.makeRequest(ctx, request, &responseJobs); err != nil {
		return nil, err
	}
	return &responseJobs, nil
}
