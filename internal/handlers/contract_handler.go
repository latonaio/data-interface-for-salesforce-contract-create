package handlers

import (
	"errors"
	"fmt"

	"github.com/latonaio/salesforce-data-models"
	"github.com/latonaio/aion-core/pkg/log"
)

func HandleContract(metadata map[string]interface{}) error {
	contracts, err := models.MetadataToContracts(metadata)
	if err != nil {
		return fmt.Errorf("failed to convert contracts: %v", err)
	}
	dataIF, exists := metadata["metadata"]
	if !exists {
		return errors.New("metadata is required")
	}
	data, ok := dataIF.(map[string]interface{})
	if !ok {
		return errors.New("failed to convert interface to map[string]interface{}")
	}
	identifierIF, exists := data["identifier"]
	if !exists {
		return errors.New("metadata[identifier] is required")
	}
	identifier, ok := identifierIF.(string)
	if !ok {
		return errors.New("failed to convert interface to string")
	}

	for _, contract := range contracts {
		contract.Identifier = &identifier
		if contract.SfContractID == nil {
			return fmt.Errorf("SfContractID is nil: identifier: %v", identifier)
		}
		c, err := models.ContractByID(*contract.SfContractID)
		if err != nil {
			return fmt.Errorf("failed to get contract: %v", err)
		}
		if c != nil {
			log.Printf("update contract id: %s, identifier: %s\n", *contract.SfContractID, *contract.Identifier)
			if err := contract.Update(); err != nil {
				return fmt.Errorf("failed to update contract: %v", err)
			}
		} else {
			log.Printf("register contract id: %s, identifier: %s\n", *contract.SfContractID, *contract.Identifier)
			if err := contract.Register(); err != nil {
				return fmt.Errorf("failed to register contract: %v", err)
			}
		}
	}
	return nil
}
