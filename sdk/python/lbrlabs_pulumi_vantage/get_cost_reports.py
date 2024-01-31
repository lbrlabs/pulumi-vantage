# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetCostReportsResult',
    'AwaitableGetCostReportsResult',
    'get_cost_reports',
]

@pulumi.output_type
class GetCostReportsResult:
    """
    A collection of values returned by getCostReports.
    """
    def __init__(__self__, cost_reports=None, id=None):
        if cost_reports and not isinstance(cost_reports, list):
            raise TypeError("Expected argument 'cost_reports' to be a list")
        pulumi.set(__self__, "cost_reports", cost_reports)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter(name="costReports")
    def cost_reports(self) -> Sequence['outputs.GetCostReportsCostReportResult']:
        return pulumi.get(self, "cost_reports")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")


class AwaitableGetCostReportsResult(GetCostReportsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCostReportsResult(
            cost_reports=self.cost_reports,
            id=self.id)


def get_cost_reports(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCostReportsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vantage:index/getCostReports:getCostReports', __args__, opts=opts, typ=GetCostReportsResult).value

    return AwaitableGetCostReportsResult(
        cost_reports=__ret__.cost_reports,
        id=__ret__.id)
