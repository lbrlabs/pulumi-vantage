// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getFolders(opts?: pulumi.InvokeOptions): Promise<GetFoldersResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vantage:index/getFolders:getFolders", {
    }, opts);
}

/**
 * A collection of values returned by getFolders.
 */
export interface GetFoldersResult {
    readonly folders: outputs.GetFoldersFolder[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
