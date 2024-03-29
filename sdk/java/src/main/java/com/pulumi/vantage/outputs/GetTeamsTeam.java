// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vantage.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetTeamsTeam {
    private String description;
    private String name;
    private String token;
    private List<String> userEmails;
    private List<String> userTokens;
    private List<String> workspaceTokens;

    private GetTeamsTeam() {}
    public String description() {
        return this.description;
    }
    public String name() {
        return this.name;
    }
    public String token() {
        return this.token;
    }
    public List<String> userEmails() {
        return this.userEmails;
    }
    public List<String> userTokens() {
        return this.userTokens;
    }
    public List<String> workspaceTokens() {
        return this.workspaceTokens;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetTeamsTeam defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String description;
        private String name;
        private String token;
        private List<String> userEmails;
        private List<String> userTokens;
        private List<String> workspaceTokens;
        public Builder() {}
        public Builder(GetTeamsTeam defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.name = defaults.name;
    	      this.token = defaults.token;
    	      this.userEmails = defaults.userEmails;
    	      this.userTokens = defaults.userTokens;
    	      this.workspaceTokens = defaults.workspaceTokens;
        }

        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder token(String token) {
            this.token = Objects.requireNonNull(token);
            return this;
        }
        @CustomType.Setter
        public Builder userEmails(List<String> userEmails) {
            this.userEmails = Objects.requireNonNull(userEmails);
            return this;
        }
        public Builder userEmails(String... userEmails) {
            return userEmails(List.of(userEmails));
        }
        @CustomType.Setter
        public Builder userTokens(List<String> userTokens) {
            this.userTokens = Objects.requireNonNull(userTokens);
            return this;
        }
        public Builder userTokens(String... userTokens) {
            return userTokens(List.of(userTokens));
        }
        @CustomType.Setter
        public Builder workspaceTokens(List<String> workspaceTokens) {
            this.workspaceTokens = Objects.requireNonNull(workspaceTokens);
            return this;
        }
        public Builder workspaceTokens(String... workspaceTokens) {
            return workspaceTokens(List.of(workspaceTokens));
        }
        public GetTeamsTeam build() {
            final var o = new GetTeamsTeam();
            o.description = description;
            o.name = name;
            o.token = token;
            o.userEmails = userEmails;
            o.userTokens = userTokens;
            o.workspaceTokens = workspaceTokens;
            return o;
        }
    }
}
