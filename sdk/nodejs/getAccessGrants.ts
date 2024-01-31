// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getAccessGrants(opts?: pulumi.InvokeOptions): Promise<GetAccessGrantsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vantage:index/getAccessGrants:getAccessGrants", {
    }, opts);
}

/**
 * A collection of values returned by getAccessGrants.
 */
export interface GetAccessGrantsResult {
    readonly accessGrants: outputs.GetAccessGrantsAccessGrant[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}