// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceWhiskyHunterWhiskyHunter string

const (
	SourceWhiskyHunterWhiskyHunterWhiskyHunter SourceWhiskyHunterWhiskyHunter = "whisky-hunter"
)

func (e SourceWhiskyHunterWhiskyHunter) ToPointer() *SourceWhiskyHunterWhiskyHunter {
	return &e
}

func (e *SourceWhiskyHunterWhiskyHunter) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "whisky-hunter":
		*e = SourceWhiskyHunterWhiskyHunter(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceWhiskyHunterWhiskyHunter: %v", v)
	}
}

type SourceWhiskyHunter struct {
	SourceType *SourceWhiskyHunterWhiskyHunter `json:"sourceType,omitempty"`
}
