// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// CreateCompletionRequestModel - ID of the model to use. You can use the [List models](/docs/api-reference/models/list) API to see all of your available models, or see our [Model overview](/docs/models/overview) for descriptions of them.
type CreateCompletionRequestModel string

const (
	CreateCompletionRequestModelTextDavinci003 CreateCompletionRequestModel = "text-davinci-003"
	CreateCompletionRequestModelTextDavinci002 CreateCompletionRequestModel = "text-davinci-002"
	CreateCompletionRequestModelTextDavinci001 CreateCompletionRequestModel = "text-davinci-001"
	CreateCompletionRequestModelCodeDavinci002 CreateCompletionRequestModel = "code-davinci-002"
	CreateCompletionRequestModelTextCurie001   CreateCompletionRequestModel = "text-curie-001"
	CreateCompletionRequestModelTextBabbage001 CreateCompletionRequestModel = "text-babbage-001"
	CreateCompletionRequestModelTextAda001     CreateCompletionRequestModel = "text-ada-001"
)

func (e CreateCompletionRequestModel) ToPointer() *CreateCompletionRequestModel {
	return &e
}

func (e *CreateCompletionRequestModel) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "text-davinci-003":
		fallthrough
	case "text-davinci-002":
		fallthrough
	case "text-davinci-001":
		fallthrough
	case "code-davinci-002":
		fallthrough
	case "text-curie-001":
		fallthrough
	case "text-babbage-001":
		fallthrough
	case "text-ada-001":
		*e = CreateCompletionRequestModel(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateCompletionRequestModel: %v", v)
	}
}

type CreateCompletionRequestPromptType string

const (
	CreateCompletionRequestPromptTypeStr                   CreateCompletionRequestPromptType = "str"
	CreateCompletionRequestPromptTypeArrayOfstr            CreateCompletionRequestPromptType = "arrayOfstr"
	CreateCompletionRequestPromptTypeArrayOfinteger        CreateCompletionRequestPromptType = "arrayOfinteger"
	CreateCompletionRequestPromptTypeArrayOfarrayOfinteger CreateCompletionRequestPromptType = "arrayOfarrayOfinteger"
)

type CreateCompletionRequestPrompt struct {
	Str                   *string
	ArrayOfstr            []string
	ArrayOfinteger        []int64
	ArrayOfarrayOfinteger [][]int64

	Type CreateCompletionRequestPromptType
}

func CreateCreateCompletionRequestPromptStr(str string) CreateCompletionRequestPrompt {
	typ := CreateCompletionRequestPromptTypeStr

	return CreateCompletionRequestPrompt{
		Str:  &str,
		Type: typ,
	}
}

func CreateCreateCompletionRequestPromptArrayOfstr(arrayOfstr []string) CreateCompletionRequestPrompt {
	typ := CreateCompletionRequestPromptTypeArrayOfstr

	return CreateCompletionRequestPrompt{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func CreateCreateCompletionRequestPromptArrayOfinteger(arrayOfinteger []int64) CreateCompletionRequestPrompt {
	typ := CreateCompletionRequestPromptTypeArrayOfinteger

	return CreateCompletionRequestPrompt{
		ArrayOfinteger: arrayOfinteger,
		Type:           typ,
	}
}

func CreateCreateCompletionRequestPromptArrayOfarrayOfinteger(arrayOfarrayOfinteger [][]int64) CreateCompletionRequestPrompt {
	typ := CreateCompletionRequestPromptTypeArrayOfarrayOfinteger

	return CreateCompletionRequestPrompt{
		ArrayOfarrayOfinteger: arrayOfarrayOfinteger,
		Type:                  typ,
	}
}

func (u *CreateCompletionRequestPrompt) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	str := new(string)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&str); err == nil {
		u.Str = str
		u.Type = CreateCompletionRequestPromptTypeStr
		return nil
	}

	arrayOfstr := []string{}
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&arrayOfstr); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = CreateCompletionRequestPromptTypeArrayOfstr
		return nil
	}

	arrayOfinteger := []int64{}
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&arrayOfinteger); err == nil {
		u.ArrayOfinteger = arrayOfinteger
		u.Type = CreateCompletionRequestPromptTypeArrayOfinteger
		return nil
	}

	arrayOfarrayOfinteger := [][]int64{}
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&arrayOfarrayOfinteger); err == nil {
		u.ArrayOfarrayOfinteger = arrayOfarrayOfinteger
		u.Type = CreateCompletionRequestPromptTypeArrayOfarrayOfinteger
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u CreateCompletionRequestPrompt) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return json.Marshal(u.Str)
	}

	if u.ArrayOfstr != nil {
		return json.Marshal(u.ArrayOfstr)
	}

	if u.ArrayOfinteger != nil {
		return json.Marshal(u.ArrayOfinteger)
	}

	if u.ArrayOfarrayOfinteger != nil {
		return json.Marshal(u.ArrayOfarrayOfinteger)
	}

	return nil, nil
}

type CreateCompletionRequestStopType string

const (
	CreateCompletionRequestStopTypeStr        CreateCompletionRequestStopType = "str"
	CreateCompletionRequestStopTypeArrayOfstr CreateCompletionRequestStopType = "arrayOfstr"
)

type CreateCompletionRequestStop struct {
	Str        *string
	ArrayOfstr []string

	Type CreateCompletionRequestStopType
}

func CreateCreateCompletionRequestStopStr(str string) CreateCompletionRequestStop {
	typ := CreateCompletionRequestStopTypeStr

	return CreateCompletionRequestStop{
		Str:  &str,
		Type: typ,
	}
}

func CreateCreateCompletionRequestStopArrayOfstr(arrayOfstr []string) CreateCompletionRequestStop {
	typ := CreateCompletionRequestStopTypeArrayOfstr

	return CreateCompletionRequestStop{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *CreateCompletionRequestStop) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	str := new(string)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&str); err == nil {
		u.Str = str
		u.Type = CreateCompletionRequestStopTypeStr
		return nil
	}

	arrayOfstr := []string{}
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&arrayOfstr); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = CreateCompletionRequestStopTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u CreateCompletionRequestStop) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return json.Marshal(u.Str)
	}

	if u.ArrayOfstr != nil {
		return json.Marshal(u.ArrayOfstr)
	}

	return nil, nil
}

type CreateCompletionRequest struct {
	// Generates `best_of` completions server-side and returns the "best" (the one with the highest log probability per token). Results cannot be streamed.
	//
	// When used with `n`, `best_of` controls the number of candidate completions and `n` specifies how many to return – `best_of` must be greater than `n`.
	//
	// **Note:** Because this parameter generates many completions, it can quickly consume your token quota. Use carefully and ensure that you have reasonable settings for `max_tokens` and `stop`.
	//
	BestOf *int64 `json:"best_of,omitempty"`
	// Echo back the prompt in addition to the completion
	//
	Echo *bool `json:"echo,omitempty"`
	// Number between -2.0 and 2.0. Positive values penalize new tokens based on their existing frequency in the text so far, decreasing the model's likelihood to repeat the same line verbatim.
	//
	// [See more information about frequency and presence penalties.](/docs/api-reference/parameter-details)
	//
	FrequencyPenalty *float64 `json:"frequency_penalty,omitempty"`
	// Modify the likelihood of specified tokens appearing in the completion.
	//
	// Accepts a json object that maps tokens (specified by their token ID in the GPT tokenizer) to an associated bias value from -100 to 100. You can use this [tokenizer tool](/tokenizer?view=bpe) (which works for both GPT-2 and GPT-3) to convert text to token IDs. Mathematically, the bias is added to the logits generated by the model prior to sampling. The exact effect will vary per model, but values between -1 and 1 should decrease or increase likelihood of selection; values like -100 or 100 should result in a ban or exclusive selection of the relevant token.
	//
	// As an example, you can pass `{"50256": -100}` to prevent the <|endoftext|> token from being generated.
	//
	LogitBias map[string]int64 `json:"logit_bias,omitempty"`
	// Include the log probabilities on the `logprobs` most likely tokens, as well the chosen tokens. For example, if `logprobs` is 5, the API will return a list of the 5 most likely tokens. The API will always return the `logprob` of the sampled token, so there may be up to `logprobs+1` elements in the response.
	//
	// The maximum value for `logprobs` is 5.
	//
	Logprobs *int64 `json:"logprobs,omitempty"`
	// The maximum number of [tokens](/tokenizer) to generate in the completion.
	//
	// The token count of your prompt plus `max_tokens` cannot exceed the model's context length. [Example Python code](https://github.com/openai/openai-cookbook/blob/main/examples/How_to_count_tokens_with_tiktoken.ipynb) for counting tokens.
	//
	MaxTokens *int64 `json:"max_tokens,omitempty"`
	// ID of the model to use. You can use the [List models](/docs/api-reference/models/list) API to see all of your available models, or see our [Model overview](/docs/models/overview) for descriptions of them.
	//
	Model CreateCompletionRequestModel `json:"model"`
	// How many completions to generate for each prompt.
	//
	// **Note:** Because this parameter generates many completions, it can quickly consume your token quota. Use carefully and ensure that you have reasonable settings for `max_tokens` and `stop`.
	//
	N *int64 `json:"n,omitempty"`
	// Number between -2.0 and 2.0. Positive values penalize new tokens based on whether they appear in the text so far, increasing the model's likelihood to talk about new topics.
	//
	// [See more information about frequency and presence penalties.](/docs/api-reference/parameter-details)
	//
	PresencePenalty *float64 `json:"presence_penalty,omitempty"`
	// The prompt(s) to generate completions for, encoded as a string, array of strings, array of tokens, or array of token arrays.
	//
	// Note that <|endoftext|> is the document separator that the model sees during training, so if a prompt is not specified the model will generate as if from the beginning of a new document.
	//
	Prompt CreateCompletionRequestPrompt `json:"prompt"`
	// Up to 4 sequences where the API will stop generating further tokens. The returned text will not contain the stop sequence.
	//
	Stop *CreateCompletionRequestStop `json:"stop,omitempty"`
	// Whether to stream back partial progress. If set, tokens will be sent as data-only [server-sent events](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events#Event_stream_format) as they become available, with the stream terminated by a `data: [DONE]` message. [Example Python code](https://github.com/openai/openai-cookbook/blob/main/examples/How_to_stream_completions.ipynb).
	//
	Stream *bool `json:"stream,omitempty"`
	// The suffix that comes after a completion of inserted text.
	Suffix *string `json:"suffix,omitempty"`
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
	// A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse. [Learn more](/docs/guides/safety-best-practices/end-user-ids).
	//
	User *string `json:"user,omitempty"`
}
