// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package listplanmodifier

import (
	"context"
	"openai/internal/planmodifiers/utils"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

// SuppressDiff returns a plan modifier that propagates a state value into the planned value, when it is Known, and the Plan Value is Unknown
func SuppressDiff() planmodifier.List {
	return suppressDiff{}
}

// suppressDiff implements the plan modifier.
type suppressDiff struct{}

// Description returns a human-readable description of the plan modifier.
func (m suppressDiff) Description(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// MarkdownDescription returns a markdown description of the plan modifier.
func (m suppressDiff) MarkdownDescription(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// PlanModifyList implements the plan modification logic.
func (m suppressDiff) PlanModifyList(ctx context.Context, req planmodifier.ListRequest, resp *planmodifier.ListResponse) {
	// Do nothing if there is a known planned value.
	if !req.PlanValue.IsUnknown() {
		return
	}

	// Do nothing if there is an unknown configuration value
	if req.ConfigValue.IsUnknown() {
		return
	}

	if utils.IsAllStateUnknown(ctx, req.State) {
		return
	}

	resp.PlanValue = req.StateValue
}
