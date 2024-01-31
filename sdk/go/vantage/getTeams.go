// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vantage

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetTeams(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetTeamsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetTeamsResult
	err := ctx.Invoke("vantage:index/getTeams:getTeams", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getTeams.
type GetTeamsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id    string         `pulumi:"id"`
	Teams []GetTeamsTeam `pulumi:"teams"`
}
