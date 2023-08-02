// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"openai/internal/sdk/pkg/models/shared"
)

type RetrieveFineTuneRequest struct {
	// The ID of the fine-tune job
	//
	FineTuneID string `pathParam:"style=simple,explode=false,name=fine_tune_id"`
}

type RetrieveFineTuneResponse struct {
	ContentType string
	// OK
	FineTune    *shared.FineTune
	StatusCode  int
	RawResponse *http.Response
}
