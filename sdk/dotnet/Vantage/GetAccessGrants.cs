// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Vantage
{
    public static class GetAccessGrants
    {
        public static Task<GetAccessGrantsResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAccessGrantsResult>("vantage:index/getAccessGrants:getAccessGrants", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetAccessGrantsResult
    {
        public readonly ImmutableArray<Outputs.GetAccessGrantsAccessGrantResult> AccessGrants;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetAccessGrantsResult(
            ImmutableArray<Outputs.GetAccessGrantsAccessGrantResult> accessGrants,

            string id)
        {
            AccessGrants = accessGrants;
            Id = id;
        }
    }
}
