// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vantage;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vantage.CostReportArgs;
import com.pulumi.vantage.Utilities;
import com.pulumi.vantage.inputs.CostReportState;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a CostReport.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vantage.CostReport;
 * import com.pulumi.vantage.CostReportArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var demoReport = new CostReport(&#34;demoReport&#34;, CostReportArgs.builder()        
 *             .filter(&#34;costs.provider = &#39;aws&#39;&#34;)
 *             .folderToken(&#34;fldr_3555785cd0409118&#34;)
 *             .savedFilterTokens(            
 *                 &#34;svd_fltr_e844a2ccace05933&#34;,
 *                 &#34;svd_fltr_1b4b80a380ef4ba2&#34;)
 *             .title(&#34;Demo Report&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="vantage:index/costReport:CostReport")
public class CostReport extends com.pulumi.resources.CustomResource {
    /**
     * Filter query to apply to the Cost Report
     * 
     */
    @Export(name="filter", type=String.class, parameters={})
    private Output</* @Nullable */ String> filter;

    /**
     * @return Filter query to apply to the Cost Report
     * 
     */
    public Output<Optional<String>> filter() {
        return Codegen.optional(this.filter);
    }
    /**
     * Token of the folder this report resides in.
     * 
     */
    @Export(name="folderToken", type=String.class, parameters={})
    private Output<String> folderToken;

    /**
     * @return Token of the folder this report resides in.
     * 
     */
    public Output<String> folderToken() {
        return this.folderToken;
    }
    /**
     * Saved filter tokens to be applied to the report.
     * 
     */
    @Export(name="savedFilterTokens", type=List.class, parameters={String.class})
    private Output<List<String>> savedFilterTokens;

    /**
     * @return Saved filter tokens to be applied to the report.
     * 
     */
    public Output<List<String>> savedFilterTokens() {
        return this.savedFilterTokens;
    }
    /**
     * Title of the Cost Report
     * 
     */
    @Export(name="title", type=String.class, parameters={})
    private Output<String> title;

    /**
     * @return Title of the Cost Report
     * 
     */
    public Output<String> title() {
        return this.title;
    }
    /**
     * Unique cost report identifier
     * 
     */
    @Export(name="token", type=String.class, parameters={})
    private Output<String> token;

    /**
     * @return Unique cost report identifier
     * 
     */
    public Output<String> token() {
        return this.token;
    }
    /**
     * Workspace token to add the cost report to.
     * 
     */
    @Export(name="workspaceToken", type=String.class, parameters={})
    private Output<String> workspaceToken;

    /**
     * @return Workspace token to add the cost report to.
     * 
     */
    public Output<String> workspaceToken() {
        return this.workspaceToken;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CostReport(String name) {
        this(name, CostReportArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CostReport(String name, CostReportArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CostReport(String name, CostReportArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vantage:index/costReport:CostReport", name, args == null ? CostReportArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private CostReport(String name, Output<String> id, @Nullable CostReportState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vantage:index/costReport:CostReport", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static CostReport get(String name, Output<String> id, @Nullable CostReportState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CostReport(name, id, state, options);
    }
}
