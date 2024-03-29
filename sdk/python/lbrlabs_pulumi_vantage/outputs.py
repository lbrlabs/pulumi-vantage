# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'SavedFiltersFilterResult',
    'GetAccessGrantsAccessGrantResult',
    'GetCostReportsCostReportResult',
    'GetDashboardsDashboardResult',
    'GetFoldersFolderResult',
    'GetTeamsTeamResult',
    'GetUsersUserResult',
    'GetWorkspacesWorkspaceResult',
]

@pulumi.output_type
class SavedFiltersFilterResult(dict):
    def __init__(__self__, *,
                 cost_report_tokens: Sequence[str],
                 title: str,
                 token: str,
                 workspace_token: str):
        pulumi.set(__self__, "cost_report_tokens", cost_report_tokens)
        pulumi.set(__self__, "title", title)
        pulumi.set(__self__, "token", token)
        pulumi.set(__self__, "workspace_token", workspace_token)

    @property
    @pulumi.getter(name="costReportTokens")
    def cost_report_tokens(self) -> Sequence[str]:
        return pulumi.get(self, "cost_report_tokens")

    @property
    @pulumi.getter
    def title(self) -> str:
        return pulumi.get(self, "title")

    @property
    @pulumi.getter
    def token(self) -> str:
        return pulumi.get(self, "token")

    @property
    @pulumi.getter(name="workspaceToken")
    def workspace_token(self) -> str:
        return pulumi.get(self, "workspace_token")


@pulumi.output_type
class GetAccessGrantsAccessGrantResult(dict):
    def __init__(__self__, *,
                 access: str,
                 resource_token: str,
                 team_token: str,
                 token: str):
        pulumi.set(__self__, "access", access)
        pulumi.set(__self__, "resource_token", resource_token)
        pulumi.set(__self__, "team_token", team_token)
        pulumi.set(__self__, "token", token)

    @property
    @pulumi.getter
    def access(self) -> str:
        return pulumi.get(self, "access")

    @property
    @pulumi.getter(name="resourceToken")
    def resource_token(self) -> str:
        return pulumi.get(self, "resource_token")

    @property
    @pulumi.getter(name="teamToken")
    def team_token(self) -> str:
        return pulumi.get(self, "team_token")

    @property
    @pulumi.getter
    def token(self) -> str:
        return pulumi.get(self, "token")


@pulumi.output_type
class GetCostReportsCostReportResult(dict):
    def __init__(__self__, *,
                 filter: str,
                 folder_token: str,
                 saved_filter_tokens: Sequence[str],
                 title: str,
                 token: str,
                 workspace_token: str):
        pulumi.set(__self__, "filter", filter)
        pulumi.set(__self__, "folder_token", folder_token)
        pulumi.set(__self__, "saved_filter_tokens", saved_filter_tokens)
        pulumi.set(__self__, "title", title)
        pulumi.set(__self__, "token", token)
        pulumi.set(__self__, "workspace_token", workspace_token)

    @property
    @pulumi.getter
    def filter(self) -> str:
        return pulumi.get(self, "filter")

    @property
    @pulumi.getter(name="folderToken")
    def folder_token(self) -> str:
        return pulumi.get(self, "folder_token")

    @property
    @pulumi.getter(name="savedFilterTokens")
    def saved_filter_tokens(self) -> Sequence[str]:
        return pulumi.get(self, "saved_filter_tokens")

    @property
    @pulumi.getter
    def title(self) -> str:
        return pulumi.get(self, "title")

    @property
    @pulumi.getter
    def token(self) -> str:
        return pulumi.get(self, "token")

    @property
    @pulumi.getter(name="workspaceToken")
    def workspace_token(self) -> str:
        return pulumi.get(self, "workspace_token")


@pulumi.output_type
class GetDashboardsDashboardResult(dict):
    def __init__(__self__, *,
                 date_bin: str,
                 date_interval: str,
                 end_date: str,
                 start_date: str,
                 title: str,
                 token: str,
                 widget_tokens: Sequence[str],
                 workspace_token: str):
        pulumi.set(__self__, "date_bin", date_bin)
        pulumi.set(__self__, "date_interval", date_interval)
        pulumi.set(__self__, "end_date", end_date)
        pulumi.set(__self__, "start_date", start_date)
        pulumi.set(__self__, "title", title)
        pulumi.set(__self__, "token", token)
        pulumi.set(__self__, "widget_tokens", widget_tokens)
        pulumi.set(__self__, "workspace_token", workspace_token)

    @property
    @pulumi.getter(name="dateBin")
    def date_bin(self) -> str:
        return pulumi.get(self, "date_bin")

    @property
    @pulumi.getter(name="dateInterval")
    def date_interval(self) -> str:
        return pulumi.get(self, "date_interval")

    @property
    @pulumi.getter(name="endDate")
    def end_date(self) -> str:
        return pulumi.get(self, "end_date")

    @property
    @pulumi.getter(name="startDate")
    def start_date(self) -> str:
        return pulumi.get(self, "start_date")

    @property
    @pulumi.getter
    def title(self) -> str:
        return pulumi.get(self, "title")

    @property
    @pulumi.getter
    def token(self) -> str:
        return pulumi.get(self, "token")

    @property
    @pulumi.getter(name="widgetTokens")
    def widget_tokens(self) -> Sequence[str]:
        return pulumi.get(self, "widget_tokens")

    @property
    @pulumi.getter(name="workspaceToken")
    def workspace_token(self) -> str:
        return pulumi.get(self, "workspace_token")


@pulumi.output_type
class GetFoldersFolderResult(dict):
    def __init__(__self__, *,
                 parent_folder_token: str,
                 saved_filter_tokens: Sequence[str],
                 title: str,
                 token: str,
                 workspace_token: str):
        pulumi.set(__self__, "parent_folder_token", parent_folder_token)
        pulumi.set(__self__, "saved_filter_tokens", saved_filter_tokens)
        pulumi.set(__self__, "title", title)
        pulumi.set(__self__, "token", token)
        pulumi.set(__self__, "workspace_token", workspace_token)

    @property
    @pulumi.getter(name="parentFolderToken")
    def parent_folder_token(self) -> str:
        return pulumi.get(self, "parent_folder_token")

    @property
    @pulumi.getter(name="savedFilterTokens")
    def saved_filter_tokens(self) -> Sequence[str]:
        return pulumi.get(self, "saved_filter_tokens")

    @property
    @pulumi.getter
    def title(self) -> str:
        return pulumi.get(self, "title")

    @property
    @pulumi.getter
    def token(self) -> str:
        return pulumi.get(self, "token")

    @property
    @pulumi.getter(name="workspaceToken")
    def workspace_token(self) -> str:
        return pulumi.get(self, "workspace_token")


@pulumi.output_type
class GetTeamsTeamResult(dict):
    def __init__(__self__, *,
                 description: str,
                 name: str,
                 token: str,
                 user_emails: Sequence[str],
                 user_tokens: Sequence[str],
                 workspace_tokens: Sequence[str]):
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "token", token)
        pulumi.set(__self__, "user_emails", user_emails)
        pulumi.set(__self__, "user_tokens", user_tokens)
        pulumi.set(__self__, "workspace_tokens", workspace_tokens)

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def token(self) -> str:
        return pulumi.get(self, "token")

    @property
    @pulumi.getter(name="userEmails")
    def user_emails(self) -> Sequence[str]:
        return pulumi.get(self, "user_emails")

    @property
    @pulumi.getter(name="userTokens")
    def user_tokens(self) -> Sequence[str]:
        return pulumi.get(self, "user_tokens")

    @property
    @pulumi.getter(name="workspaceTokens")
    def workspace_tokens(self) -> Sequence[str]:
        return pulumi.get(self, "workspace_tokens")


@pulumi.output_type
class GetUsersUserResult(dict):
    def __init__(__self__, *,
                 email: str,
                 name: str,
                 role: str,
                 token: str):
        pulumi.set(__self__, "email", email)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "role", role)
        pulumi.set(__self__, "token", token)

    @property
    @pulumi.getter
    def email(self) -> str:
        return pulumi.get(self, "email")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def role(self) -> str:
        return pulumi.get(self, "role")

    @property
    @pulumi.getter
    def token(self) -> str:
        return pulumi.get(self, "token")


@pulumi.output_type
class GetWorkspacesWorkspaceResult(dict):
    def __init__(__self__, *,
                 name: str,
                 token: str):
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "token", token)

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def token(self) -> str:
        return pulumi.get(self, "token")


