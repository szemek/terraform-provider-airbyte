// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DestinationLangchainCreateRequest struct {
	Configuration DestinationLangchain `json:"configuration"`
	// The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided.
	DefinitionID *string `json:"definitionId,omitempty"`
	// Name of the destination e.g. dev-mysql-instance.
	Name        string `json:"name"`
	WorkspaceID string `json:"workspaceId"`
}

func (o *DestinationLangchainCreateRequest) GetConfiguration() DestinationLangchain {
	if o == nil {
		return DestinationLangchain{}
	}
	return o.Configuration
}

func (o *DestinationLangchainCreateRequest) GetDefinitionID() *string {
	if o == nil {
		return nil
	}
	return o.DefinitionID
}

func (o *DestinationLangchainCreateRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *DestinationLangchainCreateRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
