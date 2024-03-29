// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getCostReports(opts?: pulumi.InvokeOptions): Promise<GetCostReportsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vantage:index/getCostReports:getCostReports", {
    }, opts);
}

/**
 * A collection of values returned by getCostReports.
 */
export interface GetCostReportsResult {
    readonly costReports: outputs.GetCostReportsCostReport[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
