// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vantage.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetAccessGrantsAccessGrant {
    private String access;
    private String resourceToken;
    private String teamToken;
    private String token;

    private GetAccessGrantsAccessGrant() {}
    public String access() {
        return this.access;
    }
    public String resourceToken() {
        return this.resourceToken;
    }
    public String teamToken() {
        return this.teamToken;
    }
    public String token() {
        return this.token;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAccessGrantsAccessGrant defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String access;
        private String resourceToken;
        private String teamToken;
        private String token;
        public Builder() {}
        public Builder(GetAccessGrantsAccessGrant defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.access = defaults.access;
    	      this.resourceToken = defaults.resourceToken;
    	      this.teamToken = defaults.teamToken;
    	      this.token = defaults.token;
        }

        @CustomType.Setter
        public Builder access(String access) {
            this.access = Objects.requireNonNull(access);
            return this;
        }
        @CustomType.Setter
        public Builder resourceToken(String resourceToken) {
            this.resourceToken = Objects.requireNonNull(resourceToken);
            return this;
        }
        @CustomType.Setter
        public Builder teamToken(String teamToken) {
            this.teamToken = Objects.requireNonNull(teamToken);
            return this;
        }
        @CustomType.Setter
        public Builder token(String token) {
            this.token = Objects.requireNonNull(token);
            return this;
        }
        public GetAccessGrantsAccessGrant build() {
            final var o = new GetAccessGrantsAccessGrant();
            o.access = access;
            o.resourceToken = resourceToken;
            o.teamToken = teamToken;
            o.token = token;
            return o;
        }
    }
}
