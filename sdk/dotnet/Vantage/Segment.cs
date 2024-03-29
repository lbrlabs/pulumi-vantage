// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Vantage
{
    /// <summary>
    /// Manages a Segment.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Vantage = Lbrlabs.PulumiPackage.Vantage;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var demoSegment = new Vantage.Segment("demoSegment", new()
    ///     {
    ///         Description = "This is still a demo segment",
    ///         Filter = "(costs.provider = 'aws' AND tags.name = NULL)",
    ///         Priority = 50,
    ///         Title = "Demo Segment",
    ///         TrackUnallocated = false,
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [VantageResourceType("vantage:index/segment:Segment")]
    public partial class Segment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the Segment.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The filter query language to apply to the Segment. Additional documentation available at https://docs.vantage.sh/vql.
        /// </summary>
        [Output("filter")]
        public Output<string> Filter { get; private set; } = null!;

        /// <summary>
        /// The token of the parent Segment this new Segment belongs to. Determines the Workspace the segment is assigned to.
        /// </summary>
        [Output("parentSegmentToken")]
        public Output<string> ParentSegmentToken { get; private set; } = null!;

        /// <summary>
        /// The priority of the Segment.
        /// </summary>
        [Output("priority")]
        public Output<int> Priority { get; private set; } = null!;

        /// <summary>
        /// The title of the Segment.
        /// </summary>
        [Output("title")]
        public Output<string> Title { get; private set; } = null!;

        /// <summary>
        /// Unique segment identifier
        /// </summary>
        [Output("token")]
        public Output<string> Token { get; private set; } = null!;

        /// <summary>
        /// Whether or not to track unallocated resources in this Segment.
        /// </summary>
        [Output("trackUnallocated")]
        public Output<bool> TrackUnallocated { get; private set; } = null!;

        /// <summary>
        /// Workspace token to add the segment to.
        /// </summary>
        [Output("workspaceToken")]
        public Output<string> WorkspaceToken { get; private set; } = null!;


        /// <summary>
        /// Create a Segment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Segment(string name, SegmentArgs args, CustomResourceOptions? options = null)
            : base("vantage:index/segment:Segment", name, args ?? new SegmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Segment(string name, Input<string> id, SegmentState? state = null, CustomResourceOptions? options = null)
            : base("vantage:index/segment:Segment", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/lbrlabs",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Segment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Segment Get(string name, Input<string> id, SegmentState? state = null, CustomResourceOptions? options = null)
        {
            return new Segment(name, id, state, options);
        }
    }

    public sealed class SegmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the Segment.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The filter query language to apply to the Segment. Additional documentation available at https://docs.vantage.sh/vql.
        /// </summary>
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        /// <summary>
        /// The token of the parent Segment this new Segment belongs to. Determines the Workspace the segment is assigned to.
        /// </summary>
        [Input("parentSegmentToken")]
        public Input<string>? ParentSegmentToken { get; set; }

        /// <summary>
        /// The priority of the Segment.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// The title of the Segment.
        /// </summary>
        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        /// <summary>
        /// Whether or not to track unallocated resources in this Segment.
        /// </summary>
        [Input("trackUnallocated")]
        public Input<bool>? TrackUnallocated { get; set; }

        /// <summary>
        /// Workspace token to add the segment to.
        /// </summary>
        [Input("workspaceToken")]
        public Input<string>? WorkspaceToken { get; set; }

        public SegmentArgs()
        {
        }
        public static new SegmentArgs Empty => new SegmentArgs();
    }

    public sealed class SegmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the Segment.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The filter query language to apply to the Segment. Additional documentation available at https://docs.vantage.sh/vql.
        /// </summary>
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        /// <summary>
        /// The token of the parent Segment this new Segment belongs to. Determines the Workspace the segment is assigned to.
        /// </summary>
        [Input("parentSegmentToken")]
        public Input<string>? ParentSegmentToken { get; set; }

        /// <summary>
        /// The priority of the Segment.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// The title of the Segment.
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        /// <summary>
        /// Unique segment identifier
        /// </summary>
        [Input("token")]
        public Input<string>? Token { get; set; }

        /// <summary>
        /// Whether or not to track unallocated resources in this Segment.
        /// </summary>
        [Input("trackUnallocated")]
        public Input<bool>? TrackUnallocated { get; set; }

        /// <summary>
        /// Workspace token to add the segment to.
        /// </summary>
        [Input("workspaceToken")]
        public Input<string>? WorkspaceToken { get; set; }

        public SegmentState()
        {
        }
        public static new SegmentState Empty => new SegmentState();
    }
}
