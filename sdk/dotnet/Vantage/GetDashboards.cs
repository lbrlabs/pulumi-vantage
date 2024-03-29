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
    public static class GetDashboards
    {
        public static Task<GetDashboardsResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDashboardsResult>("vantage:index/getDashboards:getDashboards", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetDashboardsResult
    {
        public readonly ImmutableArray<Outputs.GetDashboardsDashboardResult> Dashboards;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetDashboardsResult(
            ImmutableArray<Outputs.GetDashboardsDashboardResult> dashboards,

            string id)
        {
            Dashboards = dashboards;
            Id = id;
        }
    }
}
