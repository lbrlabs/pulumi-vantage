// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vantage.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AccessGrantState extends com.pulumi.resources.ResourceArgs {

    public static final AccessGrantState Empty = new AccessGrantState();

    /**
     * Access level of the grant. Must be either `allowed` or `denied`.
     * 
     */
    @Import(name="access")
    private @Nullable Output<String> access;

    /**
     * @return Access level of the grant. Must be either `allowed` or `denied`.
     * 
     */
    public Optional<Output<String>> access() {
        return Optional.ofNullable(this.access);
    }

    /**
     * Token of the resource being granted.
     * 
     */
    @Import(name="resourceToken")
    private @Nullable Output<String> resourceToken;

    /**
     * @return Token of the resource being granted.
     * 
     */
    public Optional<Output<String>> resourceToken() {
        return Optional.ofNullable(this.resourceToken);
    }

    /**
     * Token of the team being granted.
     * 
     */
    @Import(name="teamToken")
    private @Nullable Output<String> teamToken;

    /**
     * @return Token of the team being granted.
     * 
     */
    public Optional<Output<String>> teamToken() {
        return Optional.ofNullable(this.teamToken);
    }

    /**
     * Token of the access grant.
     * 
     */
    @Import(name="token")
    private @Nullable Output<String> token;

    /**
     * @return Token of the access grant.
     * 
     */
    public Optional<Output<String>> token() {
        return Optional.ofNullable(this.token);
    }

    private AccessGrantState() {}

    private AccessGrantState(AccessGrantState $) {
        this.access = $.access;
        this.resourceToken = $.resourceToken;
        this.teamToken = $.teamToken;
        this.token = $.token;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AccessGrantState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AccessGrantState $;

        public Builder() {
            $ = new AccessGrantState();
        }

        public Builder(AccessGrantState defaults) {
            $ = new AccessGrantState(Objects.requireNonNull(defaults));
        }

        /**
         * @param access Access level of the grant. Must be either `allowed` or `denied`.
         * 
         * @return builder
         * 
         */
        public Builder access(@Nullable Output<String> access) {
            $.access = access;
            return this;
        }

        /**
         * @param access Access level of the grant. Must be either `allowed` or `denied`.
         * 
         * @return builder
         * 
         */
        public Builder access(String access) {
            return access(Output.of(access));
        }

        /**
         * @param resourceToken Token of the resource being granted.
         * 
         * @return builder
         * 
         */
        public Builder resourceToken(@Nullable Output<String> resourceToken) {
            $.resourceToken = resourceToken;
            return this;
        }

        /**
         * @param resourceToken Token of the resource being granted.
         * 
         * @return builder
         * 
         */
        public Builder resourceToken(String resourceToken) {
            return resourceToken(Output.of(resourceToken));
        }

        /**
         * @param teamToken Token of the team being granted.
         * 
         * @return builder
         * 
         */
        public Builder teamToken(@Nullable Output<String> teamToken) {
            $.teamToken = teamToken;
            return this;
        }

        /**
         * @param teamToken Token of the team being granted.
         * 
         * @return builder
         * 
         */
        public Builder teamToken(String teamToken) {
            return teamToken(Output.of(teamToken));
        }

        /**
         * @param token Token of the access grant.
         * 
         * @return builder
         * 
         */
        public Builder token(@Nullable Output<String> token) {
            $.token = token;
            return this;
        }

        /**
         * @param token Token of the access grant.
         * 
         * @return builder
         * 
         */
        public Builder token(String token) {
            return token(Output.of(token));
        }

        public AccessGrantState build() {
            return $;
        }
    }

}