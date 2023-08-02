// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// CreateEditRequestModel2 - ID of the model to use. You can use the `text-davinci-edit-001` or `code-davinci-edit-001` model with this endpoint.
type CreateEditRequestModel2 string

const (
	CreateEditRequestModel2TextDavinciEdit001 CreateEditRequestModel2 = "text-davinci-edit-001"
	CreateEditRequestModel2CodeDavinciEdit001 CreateEditRequestModel2 = "code-davinci-edit-001"
)

func (e CreateEditRequestModel2) ToPointer() *CreateEditRequestModel2 {
	return &e
}

func (e *CreateEditRequestModel2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "text-davinci-edit-001":
		fallthrough
	case "code-davinci-edit-001":
		*e = CreateEditRequestModel2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateEditRequestModel2: %v", v)
	}
}

type CreateEditRequest struct {
	// The input text to use as a starting point for the edit.
	Input *string `json:"input,omitempty"`
	// The instruction that tells the model how to edit the prompt.
	Instruction string `json:"instruction"`
	// ID of the model to use. You can use the `text-davinci-edit-001` or `code-davinci-edit-001` model with this endpoint.
	Model interface{} `json:"model"`
	// How many edits to generate for the input and instruction.
	N *int64 `json:"n,omitempty"`
	// What sampling temperature to use, between 0 and 2. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic.
	//
	// We generally recommend altering this or `top_p` but not both.
	//
	Temperature *float64 `json:"temperature,omitempty"`
	// An alternative to sampling with temperature, called nucleus sampling, where the model considers the results of the tokens with top_p probability mass. So 0.1 means only the tokens comprising the top 10% probability mass are considered.
	//
	// We generally recommend altering this or `temperature` but not both.
	//
	TopP *float64 `json:"top_p,omitempty"`
}
