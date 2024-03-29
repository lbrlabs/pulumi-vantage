// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vantage;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DashboardArgs extends com.pulumi.resources.ResourceArgs {

    public static final DashboardArgs Empty = new DashboardArgs();

    /**
     * Determines how to group costs in the Dashboard.
     * 
     */
    @Import(name="dateBin")
    private @Nullable Output<String> dateBin;

    /**
     * @return Determines how to group costs in the Dashboard.
     * 
     */
    public Optional<Output<String>> dateBin() {
        return Optional.ofNullable(this.dateBin);
    }

    /**
     * Determines the date range in the Dashboard. Guaranteed to be set to &#39;custom&#39; if &#39;start*date&#39; and &#39;end*date&#39; are set.
     * 
     */
    @Import(name="dateInterval")
    private @Nullable Output<String> dateInterval;

    /**
     * @return Determines the date range in the Dashboard. Guaranteed to be set to &#39;custom&#39; if &#39;start*date&#39; and &#39;end*date&#39; are set.
     * 
     */
    public Optional<Output<String>> dateInterval() {
        return Optional.ofNullable(this.dateInterval);
    }

    /**
     * The end date for the date range for CostReports in the Dashboard. ISO 8601 Formatted. Overwrites &#39;date_interval&#39; if set.
     * 
     */
    @Import(name="endDate")
    private @Nullable Output<String> endDate;

    /**
     * @return The end date for the date range for CostReports in the Dashboard. ISO 8601 Formatted. Overwrites &#39;date_interval&#39; if set.
     * 
     */
    public Optional<Output<String>> endDate() {
        return Optional.ofNullable(this.endDate);
    }

    /**
     * The start date for the date range for CostReports in the Dashboard. ISO 8601 Formatted. Overwrites &#39;date_interval&#39; if set.
     * 
     */
    @Import(name="startDate")
    private @Nullable Output<String> startDate;

    /**
     * @return The start date for the date range for CostReports in the Dashboard. ISO 8601 Formatted. Overwrites &#39;date_interval&#39; if set.
     * 
     */
    public Optional<Output<String>> startDate() {
        return Optional.ofNullable(this.startDate);
    }

    /**
     * Title of the dashboard
     * 
     */
    @Import(name="title", required=true)
    private Output<String> title;

    /**
     * @return Title of the dashboard
     * 
     */
    public Output<String> title() {
        return this.title;
    }

    /**
     * Tokens for widgets present in the dashboard. Currently only cost report tokens are supported.
     * 
     */
    @Import(name="widgetTokens", required=true)
    private Output<List<String>> widgetTokens;

    /**
     * @return Tokens for widgets present in the dashboard. Currently only cost report tokens are supported.
     * 
     */
    public Output<List<String>> widgetTokens() {
        return this.widgetTokens;
    }

    /**
     * The token for the Workspace the Dashboard is a part of.
     * 
     */
    @Import(name="workspaceToken")
    private @Nullable Output<String> workspaceToken;

    /**
     * @return The token for the Workspace the Dashboard is a part of.
     * 
     */
    public Optional<Output<String>> workspaceToken() {
        return Optional.ofNullable(this.workspaceToken);
    }

    private DashboardArgs() {}

    private DashboardArgs(DashboardArgs $) {
        this.dateBin = $.dateBin;
        this.dateInterval = $.dateInterval;
        this.endDate = $.endDate;
        this.startDate = $.startDate;
        this.title = $.title;
        this.widgetTokens = $.widgetTokens;
        this.workspaceToken = $.workspaceToken;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DashboardArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DashboardArgs $;

        public Builder() {
            $ = new DashboardArgs();
        }

        public Builder(DashboardArgs defaults) {
            $ = new DashboardArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param dateBin Determines how to group costs in the Dashboard.
         * 
         * @return builder
         * 
         */
        public Builder dateBin(@Nullable Output<String> dateBin) {
            $.dateBin = dateBin;
            return this;
        }

        /**
         * @param dateBin Determines how to group costs in the Dashboard.
         * 
         * @return builder
         * 
         */
        public Builder dateBin(String dateBin) {
            return dateBin(Output.of(dateBin));
        }

        /**
         * @param dateInterval Determines the date range in the Dashboard. Guaranteed to be set to &#39;custom&#39; if &#39;start*date&#39; and &#39;end*date&#39; are set.
         * 
         * @return builder
         * 
         */
        public Builder dateInterval(@Nullable Output<String> dateInterval) {
            $.dateInterval = dateInterval;
            return this;
        }

        /**
         * @param dateInterval Determines the date range in the Dashboard. Guaranteed to be set to &#39;custom&#39; if &#39;start*date&#39; and &#39;end*date&#39; are set.
         * 
         * @return builder
         * 
         */
        public Builder dateInterval(String dateInterval) {
            return dateInterval(Output.of(dateInterval));
        }

        /**
         * @param endDate The end date for the date range for CostReports in the Dashboard. ISO 8601 Formatted. Overwrites &#39;date_interval&#39; if set.
         * 
         * @return builder
         * 
         */
        public Builder endDate(@Nullable Output<String> endDate) {
            $.endDate = endDate;
            return this;
        }

        /**
         * @param endDate The end date for the date range for CostReports in the Dashboard. ISO 8601 Formatted. Overwrites &#39;date_interval&#39; if set.
         * 
         * @return builder
         * 
         */
        public Builder endDate(String endDate) {
            return endDate(Output.of(endDate));
        }

        /**
         * @param startDate The start date for the date range for CostReports in the Dashboard. ISO 8601 Formatted. Overwrites &#39;date_interval&#39; if set.
         * 
         * @return builder
         * 
         */
        public Builder startDate(@Nullable Output<String> startDate) {
            $.startDate = startDate;
            return this;
        }

        /**
         * @param startDate The start date for the date range for CostReports in the Dashboard. ISO 8601 Formatted. Overwrites &#39;date_interval&#39; if set.
         * 
         * @return builder
         * 
         */
        public Builder startDate(String startDate) {
            return startDate(Output.of(startDate));
        }

        /**
         * @param title Title of the dashboard
         * 
         * @return builder
         * 
         */
        public Builder title(Output<String> title) {
            $.title = title;
            return this;
        }

        /**
         * @param title Title of the dashboard
         * 
         * @return builder
         * 
         */
        public Builder title(String title) {
            return title(Output.of(title));
        }

        /**
         * @param widgetTokens Tokens for widgets present in the dashboard. Currently only cost report tokens are supported.
         * 
         * @return builder
         * 
         */
        public Builder widgetTokens(Output<List<String>> widgetTokens) {
            $.widgetTokens = widgetTokens;
            return this;
        }

        /**
         * @param widgetTokens Tokens for widgets present in the dashboard. Currently only cost report tokens are supported.
         * 
         * @return builder
         * 
         */
        public Builder widgetTokens(List<String> widgetTokens) {
            return widgetTokens(Output.of(widgetTokens));
        }

        /**
         * @param widgetTokens Tokens for widgets present in the dashboard. Currently only cost report tokens are supported.
         * 
         * @return builder
         * 
         */
        public Builder widgetTokens(String... widgetTokens) {
            return widgetTokens(List.of(widgetTokens));
        }

        /**
         * @param workspaceToken The token for the Workspace the Dashboard is a part of.
         * 
         * @return builder
         * 
         */
        public Builder workspaceToken(@Nullable Output<String> workspaceToken) {
            $.workspaceToken = workspaceToken;
            return this;
        }

        /**
         * @param workspaceToken The token for the Workspace the Dashboard is a part of.
         * 
         * @return builder
         * 
         */
        public Builder workspaceToken(String workspaceToken) {
            return workspaceToken(Output.of(workspaceToken));
        }

        public DashboardArgs build() {
            $.title = Objects.requireNonNull($.title, "expected parameter 'title' to be non-null");
            $.widgetTokens = Objects.requireNonNull($.widgetTokens, "expected parameter 'widgetTokens' to be non-null");
            return $;
        }
    }

}
