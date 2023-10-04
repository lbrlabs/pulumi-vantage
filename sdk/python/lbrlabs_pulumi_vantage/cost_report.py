# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['CostReportArgs', 'CostReport']

@pulumi.input_type
class CostReportArgs:
    def __init__(__self__, *,
                 title: pulumi.Input[str],
                 filter: Optional[pulumi.Input[str]] = None,
                 folder_token: Optional[pulumi.Input[str]] = None,
                 saved_filter_tokens: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 workspace_token: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a CostReport resource.
        :param pulumi.Input[str] title: Title of the Cost Report
        :param pulumi.Input[str] filter: Filter query to apply to the Cost Report
        :param pulumi.Input[str] folder_token: Token of the folder this report resides in.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] saved_filter_tokens: Saved filter tokens to be applied to the report.
        :param pulumi.Input[str] workspace_token: Workspace token to add the cost report to.
        """
        pulumi.set(__self__, "title", title)
        if filter is not None:
            pulumi.set(__self__, "filter", filter)
        if folder_token is not None:
            pulumi.set(__self__, "folder_token", folder_token)
        if saved_filter_tokens is not None:
            pulumi.set(__self__, "saved_filter_tokens", saved_filter_tokens)
        if workspace_token is not None:
            pulumi.set(__self__, "workspace_token", workspace_token)

    @property
    @pulumi.getter
    def title(self) -> pulumi.Input[str]:
        """
        Title of the Cost Report
        """
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: pulumi.Input[str]):
        pulumi.set(self, "title", value)

    @property
    @pulumi.getter
    def filter(self) -> Optional[pulumi.Input[str]]:
        """
        Filter query to apply to the Cost Report
        """
        return pulumi.get(self, "filter")

    @filter.setter
    def filter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filter", value)

    @property
    @pulumi.getter(name="folderToken")
    def folder_token(self) -> Optional[pulumi.Input[str]]:
        """
        Token of the folder this report resides in.
        """
        return pulumi.get(self, "folder_token")

    @folder_token.setter
    def folder_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "folder_token", value)

    @property
    @pulumi.getter(name="savedFilterTokens")
    def saved_filter_tokens(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Saved filter tokens to be applied to the report.
        """
        return pulumi.get(self, "saved_filter_tokens")

    @saved_filter_tokens.setter
    def saved_filter_tokens(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "saved_filter_tokens", value)

    @property
    @pulumi.getter(name="workspaceToken")
    def workspace_token(self) -> Optional[pulumi.Input[str]]:
        """
        Workspace token to add the cost report to.
        """
        return pulumi.get(self, "workspace_token")

    @workspace_token.setter
    def workspace_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "workspace_token", value)


@pulumi.input_type
class _CostReportState:
    def __init__(__self__, *,
                 filter: Optional[pulumi.Input[str]] = None,
                 folder_token: Optional[pulumi.Input[str]] = None,
                 saved_filter_tokens: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 title: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 workspace_token: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering CostReport resources.
        :param pulumi.Input[str] filter: Filter query to apply to the Cost Report
        :param pulumi.Input[str] folder_token: Token of the folder this report resides in.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] saved_filter_tokens: Saved filter tokens to be applied to the report.
        :param pulumi.Input[str] title: Title of the Cost Report
        :param pulumi.Input[str] token: Unique cost report identifier
        :param pulumi.Input[str] workspace_token: Workspace token to add the cost report to.
        """
        if filter is not None:
            pulumi.set(__self__, "filter", filter)
        if folder_token is not None:
            pulumi.set(__self__, "folder_token", folder_token)
        if saved_filter_tokens is not None:
            pulumi.set(__self__, "saved_filter_tokens", saved_filter_tokens)
        if title is not None:
            pulumi.set(__self__, "title", title)
        if token is not None:
            pulumi.set(__self__, "token", token)
        if workspace_token is not None:
            pulumi.set(__self__, "workspace_token", workspace_token)

    @property
    @pulumi.getter
    def filter(self) -> Optional[pulumi.Input[str]]:
        """
        Filter query to apply to the Cost Report
        """
        return pulumi.get(self, "filter")

    @filter.setter
    def filter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filter", value)

    @property
    @pulumi.getter(name="folderToken")
    def folder_token(self) -> Optional[pulumi.Input[str]]:
        """
        Token of the folder this report resides in.
        """
        return pulumi.get(self, "folder_token")

    @folder_token.setter
    def folder_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "folder_token", value)

    @property
    @pulumi.getter(name="savedFilterTokens")
    def saved_filter_tokens(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Saved filter tokens to be applied to the report.
        """
        return pulumi.get(self, "saved_filter_tokens")

    @saved_filter_tokens.setter
    def saved_filter_tokens(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "saved_filter_tokens", value)

    @property
    @pulumi.getter
    def title(self) -> Optional[pulumi.Input[str]]:
        """
        Title of the Cost Report
        """
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "title", value)

    @property
    @pulumi.getter
    def token(self) -> Optional[pulumi.Input[str]]:
        """
        Unique cost report identifier
        """
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "token", value)

    @property
    @pulumi.getter(name="workspaceToken")
    def workspace_token(self) -> Optional[pulumi.Input[str]]:
        """
        Workspace token to add the cost report to.
        """
        return pulumi.get(self, "workspace_token")

    @workspace_token.setter
    def workspace_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "workspace_token", value)


class CostReport(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 filter: Optional[pulumi.Input[str]] = None,
                 folder_token: Optional[pulumi.Input[str]] = None,
                 saved_filter_tokens: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 title: Optional[pulumi.Input[str]] = None,
                 workspace_token: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a CostReport.

        ## Example Usage

        ```python
        import pulumi
        import lbrlabs_pulumi_vantage as vantage

        demo_report = vantage.CostReport("demoReport",
            filter="costs.provider = 'aws'",
            folder_token="fldr_3555785cd0409118",
            saved_filter_tokens=[
                "svd_fltr_e844a2ccace05933",
                "svd_fltr_1b4b80a380ef4ba2",
            ],
            title="Demo Report")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] filter: Filter query to apply to the Cost Report
        :param pulumi.Input[str] folder_token: Token of the folder this report resides in.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] saved_filter_tokens: Saved filter tokens to be applied to the report.
        :param pulumi.Input[str] title: Title of the Cost Report
        :param pulumi.Input[str] workspace_token: Workspace token to add the cost report to.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CostReportArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a CostReport.

        ## Example Usage

        ```python
        import pulumi
        import lbrlabs_pulumi_vantage as vantage

        demo_report = vantage.CostReport("demoReport",
            filter="costs.provider = 'aws'",
            folder_token="fldr_3555785cd0409118",
            saved_filter_tokens=[
                "svd_fltr_e844a2ccace05933",
                "svd_fltr_1b4b80a380ef4ba2",
            ],
            title="Demo Report")
        ```

        :param str resource_name: The name of the resource.
        :param CostReportArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CostReportArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 filter: Optional[pulumi.Input[str]] = None,
                 folder_token: Optional[pulumi.Input[str]] = None,
                 saved_filter_tokens: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 title: Optional[pulumi.Input[str]] = None,
                 workspace_token: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CostReportArgs.__new__(CostReportArgs)

            __props__.__dict__["filter"] = filter
            __props__.__dict__["folder_token"] = folder_token
            __props__.__dict__["saved_filter_tokens"] = saved_filter_tokens
            if title is None and not opts.urn:
                raise TypeError("Missing required property 'title'")
            __props__.__dict__["title"] = title
            __props__.__dict__["workspace_token"] = workspace_token
            __props__.__dict__["token"] = None
        super(CostReport, __self__).__init__(
            'vantage:index/costReport:CostReport',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            filter: Optional[pulumi.Input[str]] = None,
            folder_token: Optional[pulumi.Input[str]] = None,
            saved_filter_tokens: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            title: Optional[pulumi.Input[str]] = None,
            token: Optional[pulumi.Input[str]] = None,
            workspace_token: Optional[pulumi.Input[str]] = None) -> 'CostReport':
        """
        Get an existing CostReport resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] filter: Filter query to apply to the Cost Report
        :param pulumi.Input[str] folder_token: Token of the folder this report resides in.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] saved_filter_tokens: Saved filter tokens to be applied to the report.
        :param pulumi.Input[str] title: Title of the Cost Report
        :param pulumi.Input[str] token: Unique cost report identifier
        :param pulumi.Input[str] workspace_token: Workspace token to add the cost report to.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CostReportState.__new__(_CostReportState)

        __props__.__dict__["filter"] = filter
        __props__.__dict__["folder_token"] = folder_token
        __props__.__dict__["saved_filter_tokens"] = saved_filter_tokens
        __props__.__dict__["title"] = title
        __props__.__dict__["token"] = token
        __props__.__dict__["workspace_token"] = workspace_token
        return CostReport(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def filter(self) -> pulumi.Output[Optional[str]]:
        """
        Filter query to apply to the Cost Report
        """
        return pulumi.get(self, "filter")

    @property
    @pulumi.getter(name="folderToken")
    def folder_token(self) -> pulumi.Output[str]:
        """
        Token of the folder this report resides in.
        """
        return pulumi.get(self, "folder_token")

    @property
    @pulumi.getter(name="savedFilterTokens")
    def saved_filter_tokens(self) -> pulumi.Output[Sequence[str]]:
        """
        Saved filter tokens to be applied to the report.
        """
        return pulumi.get(self, "saved_filter_tokens")

    @property
    @pulumi.getter
    def title(self) -> pulumi.Output[str]:
        """
        Title of the Cost Report
        """
        return pulumi.get(self, "title")

    @property
    @pulumi.getter
    def token(self) -> pulumi.Output[str]:
        """
        Unique cost report identifier
        """
        return pulumi.get(self, "token")

    @property
    @pulumi.getter(name="workspaceToken")
    def workspace_token(self) -> pulumi.Output[str]:
        """
        Workspace token to add the cost report to.
        """
        return pulumi.get(self, "workspace_token")

