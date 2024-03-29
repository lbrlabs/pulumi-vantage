// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Vantage.Outputs
{

    [OutputType]
    public sealed class GetUsersUserResult
    {
        public readonly string Email;
        public readonly string Name;
        public readonly string Role;
        public readonly string Token;

        [OutputConstructor]
        private GetUsersUserResult(
            string email,

            string name,

            string role,

            string token)
        {
            Email = email;
            Name = name;
            Role = role;
            Token = token;
        }
    }
}
