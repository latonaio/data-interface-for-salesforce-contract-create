package resources

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Contract struct
type Contract struct {
	method   string
	metadata map[string]interface{}
}

func (c *Contract) objectName() string {
	const obName = "Contract"
	return obName
}

// newContract writes that new Customer instance
func NewContract(metadata map[string]interface{}) (*Contract, error) {
	rawMethod, ok := metadata["method"]
	if !ok {
		return nil, errors.New("missing required parameters: method")
	}
	method, ok := rawMethod.(string)
	if !ok {
		return nil, errors.New("failed to convert interface{} to string")
	}
	return &Contract{
		method:   method,
		metadata: metadata,
	}, nil
}

// postMetadata mold customer post metadata
func (c *Contract) postMetadata() (map[string]interface{}, error) {
	params := map[string]interface{}{}
	paramsIF, paramsOk := c.metadata["params"]
	if paramsOk && paramsIF != nil {
		if _, ok := paramsIF.(map[string]interface{}); !ok {
			return nil, errors.New("failed to convert interface{} to map[string]interface{}")
		}
		params = paramsIF.(map[string]interface{})
	}

	accountIdIF, accountIdOk := c.metadata["account_id"]
	if !accountIdOk {
		return nil, errors.New("account_id is required")
	}
	accountId, ok := accountIdIF.(string)
	if !ok {
		return nil, errors.New("failed to convert interface{} to string")
	}
	params["AccountId"] = accountId
	params["Account__c"] = accountId
	bytes, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("JSON marshal error: %v", err)
	}
	body := string(bytes)
	identifierIF, exists := c.metadata["identifier"]
	if !exists {
		return nil, errors.New("identifier is required")
	}
	identifier, ok := identifierIF.(string)
	if !ok {
		return nil, errors.New("failed to convert interface{} to string")
	}
	return buildMetadata(c.method, c.objectName(), "", map[string]interface{}{"identifier": identifier}, nil, body), nil
}

// BuildMetadata
func (c *Contract) BuildMetadata() (map[string]interface{}, error) {
	switch c.method {
	case "post":
		return c.postMetadata()
	}
	return nil, fmt.Errorf("invalid method: %s", c.method)
}

func buildMetadata(method, object, pathParam string, data map[string]interface{}, queryParams map[string]string, body string) map[string]interface{} {
	metadata := map[string]interface{}{
		"method":         method,
		"object":         object,
		"connection_key": "contract_post",
		"metadata":       data,
	}
	if len(pathParam) > 0 {
		metadata["path_param"] = pathParam
	}
	if queryParams != nil {
		metadata["query_params"] = queryParams
	}
	if body != "" {
		metadata["body"] = body
	}
	return metadata
}
