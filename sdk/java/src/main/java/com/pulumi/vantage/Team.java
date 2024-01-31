// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vantage;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vantage.TeamArgs;
import com.pulumi.vantage.Utilities;
import com.pulumi.vantage.inputs.TeamState;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a Team.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vantage.Team;
 * import com.pulumi.vantage.TeamArgs;
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
 *         var demoTeam = new Team(&#34;demoTeam&#34;, TeamArgs.builder()        
 *             .description(&#34;Demo Team Description&#34;)
 *             .name(&#34;Demo Team&#34;)
 *             .userEmails(&#34;support@vantage.sh&#34;)
 *             .workspaceTokens(&#34;wrkspc_47c3254c790e9351&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="vantage:index/team:Team")
public class Team extends com.pulumi.resources.CustomResource {
    /**
     * Description of the team.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return Description of the team.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Name of the team.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return Name of the team.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Unique team identifier.
     * 
     */
    @Export(name="token", type=String.class, parameters={})
    private Output<String> token;

    /**
     * @return Unique team identifier.
     * 
     */
    public Output<String> token() {
        return this.token;
    }
    /**
     * User emails.
     * 
     */
    @Export(name="userEmails", type=List.class, parameters={String.class})
    private Output<List<String>> userEmails;

    /**
     * @return User emails.
     * 
     */
    public Output<List<String>> userEmails() {
        return this.userEmails;
    }
    /**
     * User tokens.
     * 
     */
    @Export(name="userTokens", type=List.class, parameters={String.class})
    private Output<List<String>> userTokens;

    /**
     * @return User tokens.
     * 
     */
    public Output<List<String>> userTokens() {
        return this.userTokens;
    }
    /**
     * Workspace tokens to add the team to.
     * 
     */
    @Export(name="workspaceTokens", type=List.class, parameters={String.class})
    private Output<List<String>> workspaceTokens;

    /**
     * @return Workspace tokens to add the team to.
     * 
     */
    public Output<List<String>> workspaceTokens() {
        return this.workspaceTokens;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Team(String name) {
        this(name, TeamArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Team(String name, TeamArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Team(String name, TeamArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vantage:index/team:Team", name, args == null ? TeamArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Team(String name, Output<String> id, @Nullable TeamState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vantage:index/team:Team", name, state, makeResourceOptions(options, id));
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
    public static Team get(String name, Output<String> id, @Nullable TeamState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Team(name, id, state, options);
    }
}
