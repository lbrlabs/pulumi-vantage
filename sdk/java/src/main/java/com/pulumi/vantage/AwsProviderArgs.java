// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vantage;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AwsProviderArgs extends com.pulumi.resources.ResourceArgs {

    public static final AwsProviderArgs Empty = new AwsProviderArgs();

    /**
     * Bucket ARN for where CUR data is being stored.
     * 
     */
    @Import(name="bucketArn")
    private @Nullable Output<String> bucketArn;

    /**
     * @return Bucket ARN for where CUR data is being stored.
     * 
     */
    public Optional<Output<String>> bucketArn() {
        return Optional.ofNullable(this.bucketArn);
    }

    /**
     * ARN to use for cross account access.
     * 
     */
    @Import(name="crossAccountArn", required=true)
    private Output<String> crossAccountArn;

    /**
     * @return ARN to use for cross account access.
     * 
     */
    public Output<String> crossAccountArn() {
        return this.crossAccountArn;
    }

    private AwsProviderArgs() {}

    private AwsProviderArgs(AwsProviderArgs $) {
        this.bucketArn = $.bucketArn;
        this.crossAccountArn = $.crossAccountArn;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AwsProviderArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AwsProviderArgs $;

        public Builder() {
            $ = new AwsProviderArgs();
        }

        public Builder(AwsProviderArgs defaults) {
            $ = new AwsProviderArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param bucketArn Bucket ARN for where CUR data is being stored.
         * 
         * @return builder
         * 
         */
        public Builder bucketArn(@Nullable Output<String> bucketArn) {
            $.bucketArn = bucketArn;
            return this;
        }

        /**
         * @param bucketArn Bucket ARN for where CUR data is being stored.
         * 
         * @return builder
         * 
         */
        public Builder bucketArn(String bucketArn) {
            return bucketArn(Output.of(bucketArn));
        }

        /**
         * @param crossAccountArn ARN to use for cross account access.
         * 
         * @return builder
         * 
         */
        public Builder crossAccountArn(Output<String> crossAccountArn) {
            $.crossAccountArn = crossAccountArn;
            return this;
        }

        /**
         * @param crossAccountArn ARN to use for cross account access.
         * 
         * @return builder
         * 
         */
        public Builder crossAccountArn(String crossAccountArn) {
            return crossAccountArn(Output.of(crossAccountArn));
        }

        public AwsProviderArgs build() {
            $.crossAccountArn = Objects.requireNonNull($.crossAccountArn, "expected parameter 'crossAccountArn' to be non-null");
            return $;
        }
    }

}
