// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vantage;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vantage.SavedFilterArgs;
import com.pulumi.vantage.Utilities;
import com.pulumi.vantage.inputs.SavedFilterState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Manages a SavedFilter.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vantage.SavedFilter;
 * import com.pulumi.vantage.SavedFilterArgs;
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
 *         var demoFilter = new SavedFilter(&#34;demoFilter&#34;, SavedFilterArgs.builder()        
 *             .filter(&#34;(costs.provider = &#39;aws&#39;)&#34;)
 *             .title(&#34;Demo Saved Filter&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="vantage:index/savedFilter:SavedFilter")
public class SavedFilter extends com.pulumi.resources.CustomResource {
    /**
     * VQL Query used for this saved filter.
     * 
     */
    @Export(name="filter", type=String.class, parameters={})
    private Output<String> filter;

    /**
     * @return VQL Query used for this saved filter.
     * 
     */
    public Output<String> filter() {
        return this.filter;
    }
    /**
     * Title of the Saved Filter
     * 
     */
    @Export(name="title", type=String.class, parameters={})
    private Output<String> title;

    /**
     * @return Title of the Saved Filter
     * 
     */
    public Output<String> title() {
        return this.title;
    }
    /**
     * Unique saved filter identifier
     * 
     */
    @Export(name="token", type=String.class, parameters={})
    private Output<String> token;

    /**
     * @return Unique saved filter identifier
     * 
     */
    public Output<String> token() {
        return this.token;
    }
    /**
     * Workspace token to add the saved filter into.
     * 
     */
    @Export(name="workspaceToken", type=String.class, parameters={})
    private Output<String> workspaceToken;

    /**
     * @return Workspace token to add the saved filter into.
     * 
     */
    public Output<String> workspaceToken() {
        return this.workspaceToken;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SavedFilter(String name) {
        this(name, SavedFilterArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SavedFilter(String name, SavedFilterArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SavedFilter(String name, SavedFilterArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vantage:index/savedFilter:SavedFilter", name, args == null ? SavedFilterArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SavedFilter(String name, Output<String> id, @Nullable SavedFilterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vantage:index/savedFilter:SavedFilter", name, state, makeResourceOptions(options, id));
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
    public static SavedFilter get(String name, Output<String> id, @Nullable SavedFilterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SavedFilter(name, id, state, options);
    }
}
