package intelowl

// import (
// 	"context"
// 	"fmt"
// 	"net/http"
// )

// type ConfigType struct {
// 	Queue         string `json:"queue"`
// 	SoftTimeLimit int    `json:"soft_time_limit"`
// }

// type Secret struct {
// 	EnvironmentVariableKey string `json:"env_var_key"`
// 	Description            string `json:"description"`
// 	Required               bool   `json:"required"`
// }

// type Parameter struct {
// 	Value       string `json:"value"`
// 	Type        int    `json:"type"`
// 	Description string `json:"description"`
// }

// type VerificationType struct {
// 	Configured     bool     `json:"configured"`
// 	ErrorMessage   string   `json:"error_message"`
// 	MissingSecrets []string `json:"missing_secrets"`
// }

// type BaseConfigurationType struct {
// 	Name         string                      `json:"name"`
// 	PythonModule string                      `json:"python_module"`
// 	Disabled     bool                        `json:"disabled"`
// 	Description  string                      `json:"description"`
// 	Config       ConfigType                  `json:"config"`
// 	Secrets      map[string]Secret           `json:"secrets"`
// 	Params       map[string]Parameter        `json:"params"`
// 	Verification map[string]VerificationType `json:"verification"`
// }

// type AnalyzerConfiguration struct {
// 	BaseConfigurationType
// 	Type                  string   `json:"type"`
// 	ExternalService       bool     `json:"external_service"`
// 	LeaksInfo             bool     `json:"leaks_info"`
// 	DockerBased           bool     `json:"docker_based"`
// 	RunHash               bool     `json:"run_hash"`
// 	RunHashType           string   `json:"run_hash_type"`
// 	SupportedFiletypes    []string `json:"supported_filetypes"`
// 	NotSupportedFiletypes []string `json:"not_supported_filetypes"`
// 	ObservableSupported   []string `json:"observable_supported"`
// }

// type ConnectorConfiguration struct {
// 	BaseConfigurationType
// 	MaximumTlp string `json:"maximum_tlp"`
// }

// type ErrorResponse struct {
// 	erorr string `json:"erorr"`
// }

// func (client *IntelOwlClient) GetAnalyzerConfigList(ctx context.Context) (*map[string]AnalyzerConfiguration, error) {
// 	requestUrl := fmt.Sprintf("%s/api/get_analyzer_configs", client.BaseUrl)
// 	request, err := http.NewRequest("GET", requestUrl, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	responseAnalyzerConfigList := map[string]AnalyzerConfiguration{}
// 	if err := client.makeRequest(ctx, request, &responseAnalyzerConfigList); err != nil {
// 		return nil, err
// 	}
// 	return &responseAnalyzerConfigList, nil
// }

// func (client *IntelOwlClient) GetConnectorConfigList(ctx context.Context) (*map[string]ConnectorConfiguration, error) {
// 	requestUrl := fmt.Sprintf("%s/api/get_connector_configs", client.BaseUrl)
// 	request, err := http.NewRequest("GET", requestUrl, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	responseConfigList := map[string]ConnectorConfiguration{}
// 	if err := client.makeRequest(ctx, request, &responseConfigList); err != nil {
// 		return nil, err
// 	}
// 	return &responseConfigList, nil
// }
