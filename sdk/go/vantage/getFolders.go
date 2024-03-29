// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vantage

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetFolders(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetFoldersResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetFoldersResult
	err := ctx.Invoke("vantage:index/getFolders:getFolders", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getFolders.
type GetFoldersResult struct {
	Folders []GetFoldersFolder `pulumi:"folders"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
