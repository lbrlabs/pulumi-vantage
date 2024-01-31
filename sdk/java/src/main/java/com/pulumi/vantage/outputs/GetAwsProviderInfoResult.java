// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vantage.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetAwsProviderInfoResult {
    /**
     * @return The policy that allows Vantage to list and describe resources from your AWS account.
     * 
     */
    private String additionalResourcesPolicy;
    /**
     * @return The policy that allows Vantage to manage autopilot.
     * 
     */
    private String autopilotPolicy;
    /**
     * @return The policy that allows Vantage to retrieve cloudwatch metrics from your AWS account.
     * 
     */
    private String cloudwatchMetricsPolicy;
    /**
     * @return The Vantage external ID to authenticate your account.
     * 
     */
    private String externalId;
    /**
     * @return The IAM role that Vantage assumes into your account.
     * 
     */
    private String iamRoleArn;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The policy that allows Vantage to manage autopilot.
     * 
     */
    private String rootPolicy;

    private GetAwsProviderInfoResult() {}
    /**
     * @return The policy that allows Vantage to list and describe resources from your AWS account.
     * 
     */
    public String additionalResourcesPolicy() {
        return this.additionalResourcesPolicy;
    }
    /**
     * @return The policy that allows Vantage to manage autopilot.
     * 
     */
    public String autopilotPolicy() {
        return this.autopilotPolicy;
    }
    /**
     * @return The policy that allows Vantage to retrieve cloudwatch metrics from your AWS account.
     * 
     */
    public String cloudwatchMetricsPolicy() {
        return this.cloudwatchMetricsPolicy;
    }
    /**
     * @return The Vantage external ID to authenticate your account.
     * 
     */
    public String externalId() {
        return this.externalId;
    }
    /**
     * @return The IAM role that Vantage assumes into your account.
     * 
     */
    public String iamRoleArn() {
        return this.iamRoleArn;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The policy that allows Vantage to manage autopilot.
     * 
     */
    public String rootPolicy() {
        return this.rootPolicy;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAwsProviderInfoResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String additionalResourcesPolicy;
        private String autopilotPolicy;
        private String cloudwatchMetricsPolicy;
        private String externalId;
        private String iamRoleArn;
        private String id;
        private String rootPolicy;
        public Builder() {}
        public Builder(GetAwsProviderInfoResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.additionalResourcesPolicy = defaults.additionalResourcesPolicy;
    	      this.autopilotPolicy = defaults.autopilotPolicy;
    	      this.cloudwatchMetricsPolicy = defaults.cloudwatchMetricsPolicy;
    	      this.externalId = defaults.externalId;
    	      this.iamRoleArn = defaults.iamRoleArn;
    	      this.id = defaults.id;
    	      this.rootPolicy = defaults.rootPolicy;
        }

        @CustomType.Setter
        public Builder additionalResourcesPolicy(String additionalResourcesPolicy) {
            this.additionalResourcesPolicy = Objects.requireNonNull(additionalResourcesPolicy);
            return this;
        }
        @CustomType.Setter
        public Builder autopilotPolicy(String autopilotPolicy) {
            this.autopilotPolicy = Objects.requireNonNull(autopilotPolicy);
            return this;
        }
        @CustomType.Setter
        public Builder cloudwatchMetricsPolicy(String cloudwatchMetricsPolicy) {
            this.cloudwatchMetricsPolicy = Objects.requireNonNull(cloudwatchMetricsPolicy);
            return this;
        }
        @CustomType.Setter
        public Builder externalId(String externalId) {
            this.externalId = Objects.requireNonNull(externalId);
            return this;
        }
        @CustomType.Setter
        public Builder iamRoleArn(String iamRoleArn) {
            this.iamRoleArn = Objects.requireNonNull(iamRoleArn);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder rootPolicy(String rootPolicy) {
            this.rootPolicy = Objects.requireNonNull(rootPolicy);
            return this;
        }
        public GetAwsProviderInfoResult build() {
            final var o = new GetAwsProviderInfoResult();
            o.additionalResourcesPolicy = additionalResourcesPolicy;
            o.autopilotPolicy = autopilotPolicy;
            o.cloudwatchMetricsPolicy = cloudwatchMetricsPolicy;
            o.externalId = externalId;
            o.iamRoleArn = iamRoleArn;
            o.id = id;
            o.rootPolicy = rootPolicy;
            return o;
        }
    }
}
