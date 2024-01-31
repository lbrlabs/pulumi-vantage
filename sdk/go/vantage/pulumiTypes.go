// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vantage

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SavedFiltersFilter struct {
	CostReportTokens []string `pulumi:"costReportTokens"`
	Title            string   `pulumi:"title"`
	Token            string   `pulumi:"token"`
	WorkspaceToken   string   `pulumi:"workspaceToken"`
}

// SavedFiltersFilterInput is an input type that accepts SavedFiltersFilterArgs and SavedFiltersFilterOutput values.
// You can construct a concrete instance of `SavedFiltersFilterInput` via:
//
//	SavedFiltersFilterArgs{...}
type SavedFiltersFilterInput interface {
	pulumi.Input

	ToSavedFiltersFilterOutput() SavedFiltersFilterOutput
	ToSavedFiltersFilterOutputWithContext(context.Context) SavedFiltersFilterOutput
}

type SavedFiltersFilterArgs struct {
	CostReportTokens pulumi.StringArrayInput `pulumi:"costReportTokens"`
	Title            pulumi.StringInput      `pulumi:"title"`
	Token            pulumi.StringInput      `pulumi:"token"`
	WorkspaceToken   pulumi.StringInput      `pulumi:"workspaceToken"`
}

func (SavedFiltersFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SavedFiltersFilter)(nil)).Elem()
}

func (i SavedFiltersFilterArgs) ToSavedFiltersFilterOutput() SavedFiltersFilterOutput {
	return i.ToSavedFiltersFilterOutputWithContext(context.Background())
}

func (i SavedFiltersFilterArgs) ToSavedFiltersFilterOutputWithContext(ctx context.Context) SavedFiltersFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SavedFiltersFilterOutput)
}

// SavedFiltersFilterArrayInput is an input type that accepts SavedFiltersFilterArray and SavedFiltersFilterArrayOutput values.
// You can construct a concrete instance of `SavedFiltersFilterArrayInput` via:
//
//	SavedFiltersFilterArray{ SavedFiltersFilterArgs{...} }
type SavedFiltersFilterArrayInput interface {
	pulumi.Input

	ToSavedFiltersFilterArrayOutput() SavedFiltersFilterArrayOutput
	ToSavedFiltersFilterArrayOutputWithContext(context.Context) SavedFiltersFilterArrayOutput
}

type SavedFiltersFilterArray []SavedFiltersFilterInput

func (SavedFiltersFilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SavedFiltersFilter)(nil)).Elem()
}

func (i SavedFiltersFilterArray) ToSavedFiltersFilterArrayOutput() SavedFiltersFilterArrayOutput {
	return i.ToSavedFiltersFilterArrayOutputWithContext(context.Background())
}

func (i SavedFiltersFilterArray) ToSavedFiltersFilterArrayOutputWithContext(ctx context.Context) SavedFiltersFilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SavedFiltersFilterArrayOutput)
}

type SavedFiltersFilterOutput struct{ *pulumi.OutputState }

func (SavedFiltersFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SavedFiltersFilter)(nil)).Elem()
}

func (o SavedFiltersFilterOutput) ToSavedFiltersFilterOutput() SavedFiltersFilterOutput {
	return o
}

func (o SavedFiltersFilterOutput) ToSavedFiltersFilterOutputWithContext(ctx context.Context) SavedFiltersFilterOutput {
	return o
}

func (o SavedFiltersFilterOutput) CostReportTokens() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SavedFiltersFilter) []string { return v.CostReportTokens }).(pulumi.StringArrayOutput)
}

func (o SavedFiltersFilterOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v SavedFiltersFilter) string { return v.Title }).(pulumi.StringOutput)
}

func (o SavedFiltersFilterOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v SavedFiltersFilter) string { return v.Token }).(pulumi.StringOutput)
}

func (o SavedFiltersFilterOutput) WorkspaceToken() pulumi.StringOutput {
	return o.ApplyT(func(v SavedFiltersFilter) string { return v.WorkspaceToken }).(pulumi.StringOutput)
}

type SavedFiltersFilterArrayOutput struct{ *pulumi.OutputState }

func (SavedFiltersFilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SavedFiltersFilter)(nil)).Elem()
}

func (o SavedFiltersFilterArrayOutput) ToSavedFiltersFilterArrayOutput() SavedFiltersFilterArrayOutput {
	return o
}

func (o SavedFiltersFilterArrayOutput) ToSavedFiltersFilterArrayOutputWithContext(ctx context.Context) SavedFiltersFilterArrayOutput {
	return o
}

func (o SavedFiltersFilterArrayOutput) Index(i pulumi.IntInput) SavedFiltersFilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SavedFiltersFilter {
		return vs[0].([]SavedFiltersFilter)[vs[1].(int)]
	}).(SavedFiltersFilterOutput)
}

type GetAccessGrantsAccessGrant struct {
	Access        string `pulumi:"access"`
	ResourceToken string `pulumi:"resourceToken"`
	TeamToken     string `pulumi:"teamToken"`
	Token         string `pulumi:"token"`
}

// GetAccessGrantsAccessGrantInput is an input type that accepts GetAccessGrantsAccessGrantArgs and GetAccessGrantsAccessGrantOutput values.
// You can construct a concrete instance of `GetAccessGrantsAccessGrantInput` via:
//
//	GetAccessGrantsAccessGrantArgs{...}
type GetAccessGrantsAccessGrantInput interface {
	pulumi.Input

	ToGetAccessGrantsAccessGrantOutput() GetAccessGrantsAccessGrantOutput
	ToGetAccessGrantsAccessGrantOutputWithContext(context.Context) GetAccessGrantsAccessGrantOutput
}

type GetAccessGrantsAccessGrantArgs struct {
	Access        pulumi.StringInput `pulumi:"access"`
	ResourceToken pulumi.StringInput `pulumi:"resourceToken"`
	TeamToken     pulumi.StringInput `pulumi:"teamToken"`
	Token         pulumi.StringInput `pulumi:"token"`
}

func (GetAccessGrantsAccessGrantArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAccessGrantsAccessGrant)(nil)).Elem()
}

func (i GetAccessGrantsAccessGrantArgs) ToGetAccessGrantsAccessGrantOutput() GetAccessGrantsAccessGrantOutput {
	return i.ToGetAccessGrantsAccessGrantOutputWithContext(context.Background())
}

func (i GetAccessGrantsAccessGrantArgs) ToGetAccessGrantsAccessGrantOutputWithContext(ctx context.Context) GetAccessGrantsAccessGrantOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetAccessGrantsAccessGrantOutput)
}

// GetAccessGrantsAccessGrantArrayInput is an input type that accepts GetAccessGrantsAccessGrantArray and GetAccessGrantsAccessGrantArrayOutput values.
// You can construct a concrete instance of `GetAccessGrantsAccessGrantArrayInput` via:
//
//	GetAccessGrantsAccessGrantArray{ GetAccessGrantsAccessGrantArgs{...} }
type GetAccessGrantsAccessGrantArrayInput interface {
	pulumi.Input

	ToGetAccessGrantsAccessGrantArrayOutput() GetAccessGrantsAccessGrantArrayOutput
	ToGetAccessGrantsAccessGrantArrayOutputWithContext(context.Context) GetAccessGrantsAccessGrantArrayOutput
}

type GetAccessGrantsAccessGrantArray []GetAccessGrantsAccessGrantInput

func (GetAccessGrantsAccessGrantArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetAccessGrantsAccessGrant)(nil)).Elem()
}

func (i GetAccessGrantsAccessGrantArray) ToGetAccessGrantsAccessGrantArrayOutput() GetAccessGrantsAccessGrantArrayOutput {
	return i.ToGetAccessGrantsAccessGrantArrayOutputWithContext(context.Background())
}

func (i GetAccessGrantsAccessGrantArray) ToGetAccessGrantsAccessGrantArrayOutputWithContext(ctx context.Context) GetAccessGrantsAccessGrantArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetAccessGrantsAccessGrantArrayOutput)
}

type GetAccessGrantsAccessGrantOutput struct{ *pulumi.OutputState }

func (GetAccessGrantsAccessGrantOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAccessGrantsAccessGrant)(nil)).Elem()
}

func (o GetAccessGrantsAccessGrantOutput) ToGetAccessGrantsAccessGrantOutput() GetAccessGrantsAccessGrantOutput {
	return o
}

func (o GetAccessGrantsAccessGrantOutput) ToGetAccessGrantsAccessGrantOutputWithContext(ctx context.Context) GetAccessGrantsAccessGrantOutput {
	return o
}

func (o GetAccessGrantsAccessGrantOutput) Access() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccessGrantsAccessGrant) string { return v.Access }).(pulumi.StringOutput)
}

func (o GetAccessGrantsAccessGrantOutput) ResourceToken() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccessGrantsAccessGrant) string { return v.ResourceToken }).(pulumi.StringOutput)
}

func (o GetAccessGrantsAccessGrantOutput) TeamToken() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccessGrantsAccessGrant) string { return v.TeamToken }).(pulumi.StringOutput)
}

func (o GetAccessGrantsAccessGrantOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccessGrantsAccessGrant) string { return v.Token }).(pulumi.StringOutput)
}

type GetAccessGrantsAccessGrantArrayOutput struct{ *pulumi.OutputState }

func (GetAccessGrantsAccessGrantArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetAccessGrantsAccessGrant)(nil)).Elem()
}

func (o GetAccessGrantsAccessGrantArrayOutput) ToGetAccessGrantsAccessGrantArrayOutput() GetAccessGrantsAccessGrantArrayOutput {
	return o
}

func (o GetAccessGrantsAccessGrantArrayOutput) ToGetAccessGrantsAccessGrantArrayOutputWithContext(ctx context.Context) GetAccessGrantsAccessGrantArrayOutput {
	return o
}

func (o GetAccessGrantsAccessGrantArrayOutput) Index(i pulumi.IntInput) GetAccessGrantsAccessGrantOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetAccessGrantsAccessGrant {
		return vs[0].([]GetAccessGrantsAccessGrant)[vs[1].(int)]
	}).(GetAccessGrantsAccessGrantOutput)
}

type GetCostReportsCostReport struct {
	Filter            string   `pulumi:"filter"`
	FolderToken       string   `pulumi:"folderToken"`
	SavedFilterTokens []string `pulumi:"savedFilterTokens"`
	Title             string   `pulumi:"title"`
	Token             string   `pulumi:"token"`
	WorkspaceToken    string   `pulumi:"workspaceToken"`
}

// GetCostReportsCostReportInput is an input type that accepts GetCostReportsCostReportArgs and GetCostReportsCostReportOutput values.
// You can construct a concrete instance of `GetCostReportsCostReportInput` via:
//
//	GetCostReportsCostReportArgs{...}
type GetCostReportsCostReportInput interface {
	pulumi.Input

	ToGetCostReportsCostReportOutput() GetCostReportsCostReportOutput
	ToGetCostReportsCostReportOutputWithContext(context.Context) GetCostReportsCostReportOutput
}

type GetCostReportsCostReportArgs struct {
	Filter            pulumi.StringInput      `pulumi:"filter"`
	FolderToken       pulumi.StringInput      `pulumi:"folderToken"`
	SavedFilterTokens pulumi.StringArrayInput `pulumi:"savedFilterTokens"`
	Title             pulumi.StringInput      `pulumi:"title"`
	Token             pulumi.StringInput      `pulumi:"token"`
	WorkspaceToken    pulumi.StringInput      `pulumi:"workspaceToken"`
}

func (GetCostReportsCostReportArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCostReportsCostReport)(nil)).Elem()
}

func (i GetCostReportsCostReportArgs) ToGetCostReportsCostReportOutput() GetCostReportsCostReportOutput {
	return i.ToGetCostReportsCostReportOutputWithContext(context.Background())
}

func (i GetCostReportsCostReportArgs) ToGetCostReportsCostReportOutputWithContext(ctx context.Context) GetCostReportsCostReportOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetCostReportsCostReportOutput)
}

// GetCostReportsCostReportArrayInput is an input type that accepts GetCostReportsCostReportArray and GetCostReportsCostReportArrayOutput values.
// You can construct a concrete instance of `GetCostReportsCostReportArrayInput` via:
//
//	GetCostReportsCostReportArray{ GetCostReportsCostReportArgs{...} }
type GetCostReportsCostReportArrayInput interface {
	pulumi.Input

	ToGetCostReportsCostReportArrayOutput() GetCostReportsCostReportArrayOutput
	ToGetCostReportsCostReportArrayOutputWithContext(context.Context) GetCostReportsCostReportArrayOutput
}

type GetCostReportsCostReportArray []GetCostReportsCostReportInput

func (GetCostReportsCostReportArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetCostReportsCostReport)(nil)).Elem()
}

func (i GetCostReportsCostReportArray) ToGetCostReportsCostReportArrayOutput() GetCostReportsCostReportArrayOutput {
	return i.ToGetCostReportsCostReportArrayOutputWithContext(context.Background())
}

func (i GetCostReportsCostReportArray) ToGetCostReportsCostReportArrayOutputWithContext(ctx context.Context) GetCostReportsCostReportArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetCostReportsCostReportArrayOutput)
}

type GetCostReportsCostReportOutput struct{ *pulumi.OutputState }

func (GetCostReportsCostReportOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCostReportsCostReport)(nil)).Elem()
}

func (o GetCostReportsCostReportOutput) ToGetCostReportsCostReportOutput() GetCostReportsCostReportOutput {
	return o
}

func (o GetCostReportsCostReportOutput) ToGetCostReportsCostReportOutputWithContext(ctx context.Context) GetCostReportsCostReportOutput {
	return o
}

func (o GetCostReportsCostReportOutput) Filter() pulumi.StringOutput {
	return o.ApplyT(func(v GetCostReportsCostReport) string { return v.Filter }).(pulumi.StringOutput)
}

func (o GetCostReportsCostReportOutput) FolderToken() pulumi.StringOutput {
	return o.ApplyT(func(v GetCostReportsCostReport) string { return v.FolderToken }).(pulumi.StringOutput)
}

func (o GetCostReportsCostReportOutput) SavedFilterTokens() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetCostReportsCostReport) []string { return v.SavedFilterTokens }).(pulumi.StringArrayOutput)
}

func (o GetCostReportsCostReportOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v GetCostReportsCostReport) string { return v.Title }).(pulumi.StringOutput)
}

func (o GetCostReportsCostReportOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v GetCostReportsCostReport) string { return v.Token }).(pulumi.StringOutput)
}

func (o GetCostReportsCostReportOutput) WorkspaceToken() pulumi.StringOutput {
	return o.ApplyT(func(v GetCostReportsCostReport) string { return v.WorkspaceToken }).(pulumi.StringOutput)
}

type GetCostReportsCostReportArrayOutput struct{ *pulumi.OutputState }

func (GetCostReportsCostReportArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetCostReportsCostReport)(nil)).Elem()
}

func (o GetCostReportsCostReportArrayOutput) ToGetCostReportsCostReportArrayOutput() GetCostReportsCostReportArrayOutput {
	return o
}

func (o GetCostReportsCostReportArrayOutput) ToGetCostReportsCostReportArrayOutputWithContext(ctx context.Context) GetCostReportsCostReportArrayOutput {
	return o
}

func (o GetCostReportsCostReportArrayOutput) Index(i pulumi.IntInput) GetCostReportsCostReportOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetCostReportsCostReport {
		return vs[0].([]GetCostReportsCostReport)[vs[1].(int)]
	}).(GetCostReportsCostReportOutput)
}

type GetDashboardsDashboard struct {
	DateBin        string   `pulumi:"dateBin"`
	DateInterval   string   `pulumi:"dateInterval"`
	EndDate        string   `pulumi:"endDate"`
	StartDate      string   `pulumi:"startDate"`
	Title          string   `pulumi:"title"`
	Token          string   `pulumi:"token"`
	WidgetTokens   []string `pulumi:"widgetTokens"`
	WorkspaceToken string   `pulumi:"workspaceToken"`
}

// GetDashboardsDashboardInput is an input type that accepts GetDashboardsDashboardArgs and GetDashboardsDashboardOutput values.
// You can construct a concrete instance of `GetDashboardsDashboardInput` via:
//
//	GetDashboardsDashboardArgs{...}
type GetDashboardsDashboardInput interface {
	pulumi.Input

	ToGetDashboardsDashboardOutput() GetDashboardsDashboardOutput
	ToGetDashboardsDashboardOutputWithContext(context.Context) GetDashboardsDashboardOutput
}

type GetDashboardsDashboardArgs struct {
	DateBin        pulumi.StringInput      `pulumi:"dateBin"`
	DateInterval   pulumi.StringInput      `pulumi:"dateInterval"`
	EndDate        pulumi.StringInput      `pulumi:"endDate"`
	StartDate      pulumi.StringInput      `pulumi:"startDate"`
	Title          pulumi.StringInput      `pulumi:"title"`
	Token          pulumi.StringInput      `pulumi:"token"`
	WidgetTokens   pulumi.StringArrayInput `pulumi:"widgetTokens"`
	WorkspaceToken pulumi.StringInput      `pulumi:"workspaceToken"`
}

func (GetDashboardsDashboardArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDashboardsDashboard)(nil)).Elem()
}

func (i GetDashboardsDashboardArgs) ToGetDashboardsDashboardOutput() GetDashboardsDashboardOutput {
	return i.ToGetDashboardsDashboardOutputWithContext(context.Background())
}

func (i GetDashboardsDashboardArgs) ToGetDashboardsDashboardOutputWithContext(ctx context.Context) GetDashboardsDashboardOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetDashboardsDashboardOutput)
}

// GetDashboardsDashboardArrayInput is an input type that accepts GetDashboardsDashboardArray and GetDashboardsDashboardArrayOutput values.
// You can construct a concrete instance of `GetDashboardsDashboardArrayInput` via:
//
//	GetDashboardsDashboardArray{ GetDashboardsDashboardArgs{...} }
type GetDashboardsDashboardArrayInput interface {
	pulumi.Input

	ToGetDashboardsDashboardArrayOutput() GetDashboardsDashboardArrayOutput
	ToGetDashboardsDashboardArrayOutputWithContext(context.Context) GetDashboardsDashboardArrayOutput
}

type GetDashboardsDashboardArray []GetDashboardsDashboardInput

func (GetDashboardsDashboardArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetDashboardsDashboard)(nil)).Elem()
}

func (i GetDashboardsDashboardArray) ToGetDashboardsDashboardArrayOutput() GetDashboardsDashboardArrayOutput {
	return i.ToGetDashboardsDashboardArrayOutputWithContext(context.Background())
}

func (i GetDashboardsDashboardArray) ToGetDashboardsDashboardArrayOutputWithContext(ctx context.Context) GetDashboardsDashboardArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetDashboardsDashboardArrayOutput)
}

type GetDashboardsDashboardOutput struct{ *pulumi.OutputState }

func (GetDashboardsDashboardOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDashboardsDashboard)(nil)).Elem()
}

func (o GetDashboardsDashboardOutput) ToGetDashboardsDashboardOutput() GetDashboardsDashboardOutput {
	return o
}

func (o GetDashboardsDashboardOutput) ToGetDashboardsDashboardOutputWithContext(ctx context.Context) GetDashboardsDashboardOutput {
	return o
}

func (o GetDashboardsDashboardOutput) DateBin() pulumi.StringOutput {
	return o.ApplyT(func(v GetDashboardsDashboard) string { return v.DateBin }).(pulumi.StringOutput)
}

func (o GetDashboardsDashboardOutput) DateInterval() pulumi.StringOutput {
	return o.ApplyT(func(v GetDashboardsDashboard) string { return v.DateInterval }).(pulumi.StringOutput)
}

func (o GetDashboardsDashboardOutput) EndDate() pulumi.StringOutput {
	return o.ApplyT(func(v GetDashboardsDashboard) string { return v.EndDate }).(pulumi.StringOutput)
}

func (o GetDashboardsDashboardOutput) StartDate() pulumi.StringOutput {
	return o.ApplyT(func(v GetDashboardsDashboard) string { return v.StartDate }).(pulumi.StringOutput)
}

func (o GetDashboardsDashboardOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v GetDashboardsDashboard) string { return v.Title }).(pulumi.StringOutput)
}

func (o GetDashboardsDashboardOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v GetDashboardsDashboard) string { return v.Token }).(pulumi.StringOutput)
}

func (o GetDashboardsDashboardOutput) WidgetTokens() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDashboardsDashboard) []string { return v.WidgetTokens }).(pulumi.StringArrayOutput)
}

func (o GetDashboardsDashboardOutput) WorkspaceToken() pulumi.StringOutput {
	return o.ApplyT(func(v GetDashboardsDashboard) string { return v.WorkspaceToken }).(pulumi.StringOutput)
}

type GetDashboardsDashboardArrayOutput struct{ *pulumi.OutputState }

func (GetDashboardsDashboardArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetDashboardsDashboard)(nil)).Elem()
}

func (o GetDashboardsDashboardArrayOutput) ToGetDashboardsDashboardArrayOutput() GetDashboardsDashboardArrayOutput {
	return o
}

func (o GetDashboardsDashboardArrayOutput) ToGetDashboardsDashboardArrayOutputWithContext(ctx context.Context) GetDashboardsDashboardArrayOutput {
	return o
}

func (o GetDashboardsDashboardArrayOutput) Index(i pulumi.IntInput) GetDashboardsDashboardOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetDashboardsDashboard {
		return vs[0].([]GetDashboardsDashboard)[vs[1].(int)]
	}).(GetDashboardsDashboardOutput)
}

type GetFoldersFolder struct {
	ParentFolderToken string   `pulumi:"parentFolderToken"`
	SavedFilterTokens []string `pulumi:"savedFilterTokens"`
	Title             string   `pulumi:"title"`
	Token             string   `pulumi:"token"`
	WorkspaceToken    string   `pulumi:"workspaceToken"`
}

// GetFoldersFolderInput is an input type that accepts GetFoldersFolderArgs and GetFoldersFolderOutput values.
// You can construct a concrete instance of `GetFoldersFolderInput` via:
//
//	GetFoldersFolderArgs{...}
type GetFoldersFolderInput interface {
	pulumi.Input

	ToGetFoldersFolderOutput() GetFoldersFolderOutput
	ToGetFoldersFolderOutputWithContext(context.Context) GetFoldersFolderOutput
}

type GetFoldersFolderArgs struct {
	ParentFolderToken pulumi.StringInput      `pulumi:"parentFolderToken"`
	SavedFilterTokens pulumi.StringArrayInput `pulumi:"savedFilterTokens"`
	Title             pulumi.StringInput      `pulumi:"title"`
	Token             pulumi.StringInput      `pulumi:"token"`
	WorkspaceToken    pulumi.StringInput      `pulumi:"workspaceToken"`
}

func (GetFoldersFolderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFoldersFolder)(nil)).Elem()
}

func (i GetFoldersFolderArgs) ToGetFoldersFolderOutput() GetFoldersFolderOutput {
	return i.ToGetFoldersFolderOutputWithContext(context.Background())
}

func (i GetFoldersFolderArgs) ToGetFoldersFolderOutputWithContext(ctx context.Context) GetFoldersFolderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetFoldersFolderOutput)
}

// GetFoldersFolderArrayInput is an input type that accepts GetFoldersFolderArray and GetFoldersFolderArrayOutput values.
// You can construct a concrete instance of `GetFoldersFolderArrayInput` via:
//
//	GetFoldersFolderArray{ GetFoldersFolderArgs{...} }
type GetFoldersFolderArrayInput interface {
	pulumi.Input

	ToGetFoldersFolderArrayOutput() GetFoldersFolderArrayOutput
	ToGetFoldersFolderArrayOutputWithContext(context.Context) GetFoldersFolderArrayOutput
}

type GetFoldersFolderArray []GetFoldersFolderInput

func (GetFoldersFolderArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetFoldersFolder)(nil)).Elem()
}

func (i GetFoldersFolderArray) ToGetFoldersFolderArrayOutput() GetFoldersFolderArrayOutput {
	return i.ToGetFoldersFolderArrayOutputWithContext(context.Background())
}

func (i GetFoldersFolderArray) ToGetFoldersFolderArrayOutputWithContext(ctx context.Context) GetFoldersFolderArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetFoldersFolderArrayOutput)
}

type GetFoldersFolderOutput struct{ *pulumi.OutputState }

func (GetFoldersFolderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFoldersFolder)(nil)).Elem()
}

func (o GetFoldersFolderOutput) ToGetFoldersFolderOutput() GetFoldersFolderOutput {
	return o
}

func (o GetFoldersFolderOutput) ToGetFoldersFolderOutputWithContext(ctx context.Context) GetFoldersFolderOutput {
	return o
}

func (o GetFoldersFolderOutput) ParentFolderToken() pulumi.StringOutput {
	return o.ApplyT(func(v GetFoldersFolder) string { return v.ParentFolderToken }).(pulumi.StringOutput)
}

func (o GetFoldersFolderOutput) SavedFilterTokens() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetFoldersFolder) []string { return v.SavedFilterTokens }).(pulumi.StringArrayOutput)
}

func (o GetFoldersFolderOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v GetFoldersFolder) string { return v.Title }).(pulumi.StringOutput)
}

func (o GetFoldersFolderOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v GetFoldersFolder) string { return v.Token }).(pulumi.StringOutput)
}

func (o GetFoldersFolderOutput) WorkspaceToken() pulumi.StringOutput {
	return o.ApplyT(func(v GetFoldersFolder) string { return v.WorkspaceToken }).(pulumi.StringOutput)
}

type GetFoldersFolderArrayOutput struct{ *pulumi.OutputState }

func (GetFoldersFolderArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetFoldersFolder)(nil)).Elem()
}

func (o GetFoldersFolderArrayOutput) ToGetFoldersFolderArrayOutput() GetFoldersFolderArrayOutput {
	return o
}

func (o GetFoldersFolderArrayOutput) ToGetFoldersFolderArrayOutputWithContext(ctx context.Context) GetFoldersFolderArrayOutput {
	return o
}

func (o GetFoldersFolderArrayOutput) Index(i pulumi.IntInput) GetFoldersFolderOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetFoldersFolder {
		return vs[0].([]GetFoldersFolder)[vs[1].(int)]
	}).(GetFoldersFolderOutput)
}

type GetTeamsTeam struct {
	Description     string   `pulumi:"description"`
	Name            string   `pulumi:"name"`
	Token           string   `pulumi:"token"`
	UserEmails      []string `pulumi:"userEmails"`
	UserTokens      []string `pulumi:"userTokens"`
	WorkspaceTokens []string `pulumi:"workspaceTokens"`
}

// GetTeamsTeamInput is an input type that accepts GetTeamsTeamArgs and GetTeamsTeamOutput values.
// You can construct a concrete instance of `GetTeamsTeamInput` via:
//
//	GetTeamsTeamArgs{...}
type GetTeamsTeamInput interface {
	pulumi.Input

	ToGetTeamsTeamOutput() GetTeamsTeamOutput
	ToGetTeamsTeamOutputWithContext(context.Context) GetTeamsTeamOutput
}

type GetTeamsTeamArgs struct {
	Description     pulumi.StringInput      `pulumi:"description"`
	Name            pulumi.StringInput      `pulumi:"name"`
	Token           pulumi.StringInput      `pulumi:"token"`
	UserEmails      pulumi.StringArrayInput `pulumi:"userEmails"`
	UserTokens      pulumi.StringArrayInput `pulumi:"userTokens"`
	WorkspaceTokens pulumi.StringArrayInput `pulumi:"workspaceTokens"`
}

func (GetTeamsTeamArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTeamsTeam)(nil)).Elem()
}

func (i GetTeamsTeamArgs) ToGetTeamsTeamOutput() GetTeamsTeamOutput {
	return i.ToGetTeamsTeamOutputWithContext(context.Background())
}

func (i GetTeamsTeamArgs) ToGetTeamsTeamOutputWithContext(ctx context.Context) GetTeamsTeamOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetTeamsTeamOutput)
}

// GetTeamsTeamArrayInput is an input type that accepts GetTeamsTeamArray and GetTeamsTeamArrayOutput values.
// You can construct a concrete instance of `GetTeamsTeamArrayInput` via:
//
//	GetTeamsTeamArray{ GetTeamsTeamArgs{...} }
type GetTeamsTeamArrayInput interface {
	pulumi.Input

	ToGetTeamsTeamArrayOutput() GetTeamsTeamArrayOutput
	ToGetTeamsTeamArrayOutputWithContext(context.Context) GetTeamsTeamArrayOutput
}

type GetTeamsTeamArray []GetTeamsTeamInput

func (GetTeamsTeamArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetTeamsTeam)(nil)).Elem()
}

func (i GetTeamsTeamArray) ToGetTeamsTeamArrayOutput() GetTeamsTeamArrayOutput {
	return i.ToGetTeamsTeamArrayOutputWithContext(context.Background())
}

func (i GetTeamsTeamArray) ToGetTeamsTeamArrayOutputWithContext(ctx context.Context) GetTeamsTeamArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetTeamsTeamArrayOutput)
}

type GetTeamsTeamOutput struct{ *pulumi.OutputState }

func (GetTeamsTeamOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTeamsTeam)(nil)).Elem()
}

func (o GetTeamsTeamOutput) ToGetTeamsTeamOutput() GetTeamsTeamOutput {
	return o
}

func (o GetTeamsTeamOutput) ToGetTeamsTeamOutputWithContext(ctx context.Context) GetTeamsTeamOutput {
	return o
}

func (o GetTeamsTeamOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetTeamsTeam) string { return v.Description }).(pulumi.StringOutput)
}

func (o GetTeamsTeamOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetTeamsTeam) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetTeamsTeamOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v GetTeamsTeam) string { return v.Token }).(pulumi.StringOutput)
}

func (o GetTeamsTeamOutput) UserEmails() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetTeamsTeam) []string { return v.UserEmails }).(pulumi.StringArrayOutput)
}

func (o GetTeamsTeamOutput) UserTokens() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetTeamsTeam) []string { return v.UserTokens }).(pulumi.StringArrayOutput)
}

func (o GetTeamsTeamOutput) WorkspaceTokens() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetTeamsTeam) []string { return v.WorkspaceTokens }).(pulumi.StringArrayOutput)
}

type GetTeamsTeamArrayOutput struct{ *pulumi.OutputState }

func (GetTeamsTeamArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetTeamsTeam)(nil)).Elem()
}

func (o GetTeamsTeamArrayOutput) ToGetTeamsTeamArrayOutput() GetTeamsTeamArrayOutput {
	return o
}

func (o GetTeamsTeamArrayOutput) ToGetTeamsTeamArrayOutputWithContext(ctx context.Context) GetTeamsTeamArrayOutput {
	return o
}

func (o GetTeamsTeamArrayOutput) Index(i pulumi.IntInput) GetTeamsTeamOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetTeamsTeam {
		return vs[0].([]GetTeamsTeam)[vs[1].(int)]
	}).(GetTeamsTeamOutput)
}

type GetUsersUser struct {
	Email string `pulumi:"email"`
	Name  string `pulumi:"name"`
	Role  string `pulumi:"role"`
	Token string `pulumi:"token"`
}

// GetUsersUserInput is an input type that accepts GetUsersUserArgs and GetUsersUserOutput values.
// You can construct a concrete instance of `GetUsersUserInput` via:
//
//	GetUsersUserArgs{...}
type GetUsersUserInput interface {
	pulumi.Input

	ToGetUsersUserOutput() GetUsersUserOutput
	ToGetUsersUserOutputWithContext(context.Context) GetUsersUserOutput
}

type GetUsersUserArgs struct {
	Email pulumi.StringInput `pulumi:"email"`
	Name  pulumi.StringInput `pulumi:"name"`
	Role  pulumi.StringInput `pulumi:"role"`
	Token pulumi.StringInput `pulumi:"token"`
}

func (GetUsersUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUsersUser)(nil)).Elem()
}

func (i GetUsersUserArgs) ToGetUsersUserOutput() GetUsersUserOutput {
	return i.ToGetUsersUserOutputWithContext(context.Background())
}

func (i GetUsersUserArgs) ToGetUsersUserOutputWithContext(ctx context.Context) GetUsersUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetUsersUserOutput)
}

// GetUsersUserArrayInput is an input type that accepts GetUsersUserArray and GetUsersUserArrayOutput values.
// You can construct a concrete instance of `GetUsersUserArrayInput` via:
//
//	GetUsersUserArray{ GetUsersUserArgs{...} }
type GetUsersUserArrayInput interface {
	pulumi.Input

	ToGetUsersUserArrayOutput() GetUsersUserArrayOutput
	ToGetUsersUserArrayOutputWithContext(context.Context) GetUsersUserArrayOutput
}

type GetUsersUserArray []GetUsersUserInput

func (GetUsersUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetUsersUser)(nil)).Elem()
}

func (i GetUsersUserArray) ToGetUsersUserArrayOutput() GetUsersUserArrayOutput {
	return i.ToGetUsersUserArrayOutputWithContext(context.Background())
}

func (i GetUsersUserArray) ToGetUsersUserArrayOutputWithContext(ctx context.Context) GetUsersUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetUsersUserArrayOutput)
}

type GetUsersUserOutput struct{ *pulumi.OutputState }

func (GetUsersUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUsersUser)(nil)).Elem()
}

func (o GetUsersUserOutput) ToGetUsersUserOutput() GetUsersUserOutput {
	return o
}

func (o GetUsersUserOutput) ToGetUsersUserOutputWithContext(ctx context.Context) GetUsersUserOutput {
	return o
}

func (o GetUsersUserOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v GetUsersUser) string { return v.Email }).(pulumi.StringOutput)
}

func (o GetUsersUserOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetUsersUser) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetUsersUserOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v GetUsersUser) string { return v.Role }).(pulumi.StringOutput)
}

func (o GetUsersUserOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v GetUsersUser) string { return v.Token }).(pulumi.StringOutput)
}

type GetUsersUserArrayOutput struct{ *pulumi.OutputState }

func (GetUsersUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetUsersUser)(nil)).Elem()
}

func (o GetUsersUserArrayOutput) ToGetUsersUserArrayOutput() GetUsersUserArrayOutput {
	return o
}

func (o GetUsersUserArrayOutput) ToGetUsersUserArrayOutputWithContext(ctx context.Context) GetUsersUserArrayOutput {
	return o
}

func (o GetUsersUserArrayOutput) Index(i pulumi.IntInput) GetUsersUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetUsersUser {
		return vs[0].([]GetUsersUser)[vs[1].(int)]
	}).(GetUsersUserOutput)
}

type GetWorkspacesWorkspace struct {
	Name  string `pulumi:"name"`
	Token string `pulumi:"token"`
}

// GetWorkspacesWorkspaceInput is an input type that accepts GetWorkspacesWorkspaceArgs and GetWorkspacesWorkspaceOutput values.
// You can construct a concrete instance of `GetWorkspacesWorkspaceInput` via:
//
//	GetWorkspacesWorkspaceArgs{...}
type GetWorkspacesWorkspaceInput interface {
	pulumi.Input

	ToGetWorkspacesWorkspaceOutput() GetWorkspacesWorkspaceOutput
	ToGetWorkspacesWorkspaceOutputWithContext(context.Context) GetWorkspacesWorkspaceOutput
}

type GetWorkspacesWorkspaceArgs struct {
	Name  pulumi.StringInput `pulumi:"name"`
	Token pulumi.StringInput `pulumi:"token"`
}

func (GetWorkspacesWorkspaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWorkspacesWorkspace)(nil)).Elem()
}

func (i GetWorkspacesWorkspaceArgs) ToGetWorkspacesWorkspaceOutput() GetWorkspacesWorkspaceOutput {
	return i.ToGetWorkspacesWorkspaceOutputWithContext(context.Background())
}

func (i GetWorkspacesWorkspaceArgs) ToGetWorkspacesWorkspaceOutputWithContext(ctx context.Context) GetWorkspacesWorkspaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetWorkspacesWorkspaceOutput)
}

// GetWorkspacesWorkspaceArrayInput is an input type that accepts GetWorkspacesWorkspaceArray and GetWorkspacesWorkspaceArrayOutput values.
// You can construct a concrete instance of `GetWorkspacesWorkspaceArrayInput` via:
//
//	GetWorkspacesWorkspaceArray{ GetWorkspacesWorkspaceArgs{...} }
type GetWorkspacesWorkspaceArrayInput interface {
	pulumi.Input

	ToGetWorkspacesWorkspaceArrayOutput() GetWorkspacesWorkspaceArrayOutput
	ToGetWorkspacesWorkspaceArrayOutputWithContext(context.Context) GetWorkspacesWorkspaceArrayOutput
}

type GetWorkspacesWorkspaceArray []GetWorkspacesWorkspaceInput

func (GetWorkspacesWorkspaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetWorkspacesWorkspace)(nil)).Elem()
}

func (i GetWorkspacesWorkspaceArray) ToGetWorkspacesWorkspaceArrayOutput() GetWorkspacesWorkspaceArrayOutput {
	return i.ToGetWorkspacesWorkspaceArrayOutputWithContext(context.Background())
}

func (i GetWorkspacesWorkspaceArray) ToGetWorkspacesWorkspaceArrayOutputWithContext(ctx context.Context) GetWorkspacesWorkspaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetWorkspacesWorkspaceArrayOutput)
}

type GetWorkspacesWorkspaceOutput struct{ *pulumi.OutputState }

func (GetWorkspacesWorkspaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWorkspacesWorkspace)(nil)).Elem()
}

func (o GetWorkspacesWorkspaceOutput) ToGetWorkspacesWorkspaceOutput() GetWorkspacesWorkspaceOutput {
	return o
}

func (o GetWorkspacesWorkspaceOutput) ToGetWorkspacesWorkspaceOutputWithContext(ctx context.Context) GetWorkspacesWorkspaceOutput {
	return o
}

func (o GetWorkspacesWorkspaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetWorkspacesWorkspace) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetWorkspacesWorkspaceOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v GetWorkspacesWorkspace) string { return v.Token }).(pulumi.StringOutput)
}

type GetWorkspacesWorkspaceArrayOutput struct{ *pulumi.OutputState }

func (GetWorkspacesWorkspaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetWorkspacesWorkspace)(nil)).Elem()
}

func (o GetWorkspacesWorkspaceArrayOutput) ToGetWorkspacesWorkspaceArrayOutput() GetWorkspacesWorkspaceArrayOutput {
	return o
}

func (o GetWorkspacesWorkspaceArrayOutput) ToGetWorkspacesWorkspaceArrayOutputWithContext(ctx context.Context) GetWorkspacesWorkspaceArrayOutput {
	return o
}

func (o GetWorkspacesWorkspaceArrayOutput) Index(i pulumi.IntInput) GetWorkspacesWorkspaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetWorkspacesWorkspace {
		return vs[0].([]GetWorkspacesWorkspace)[vs[1].(int)]
	}).(GetWorkspacesWorkspaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SavedFiltersFilterInput)(nil)).Elem(), SavedFiltersFilterArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SavedFiltersFilterArrayInput)(nil)).Elem(), SavedFiltersFilterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetAccessGrantsAccessGrantInput)(nil)).Elem(), GetAccessGrantsAccessGrantArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetAccessGrantsAccessGrantArrayInput)(nil)).Elem(), GetAccessGrantsAccessGrantArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetCostReportsCostReportInput)(nil)).Elem(), GetCostReportsCostReportArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetCostReportsCostReportArrayInput)(nil)).Elem(), GetCostReportsCostReportArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetDashboardsDashboardInput)(nil)).Elem(), GetDashboardsDashboardArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetDashboardsDashboardArrayInput)(nil)).Elem(), GetDashboardsDashboardArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetFoldersFolderInput)(nil)).Elem(), GetFoldersFolderArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetFoldersFolderArrayInput)(nil)).Elem(), GetFoldersFolderArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetTeamsTeamInput)(nil)).Elem(), GetTeamsTeamArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetTeamsTeamArrayInput)(nil)).Elem(), GetTeamsTeamArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetUsersUserInput)(nil)).Elem(), GetUsersUserArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetUsersUserArrayInput)(nil)).Elem(), GetUsersUserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetWorkspacesWorkspaceInput)(nil)).Elem(), GetWorkspacesWorkspaceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetWorkspacesWorkspaceArrayInput)(nil)).Elem(), GetWorkspacesWorkspaceArray{})
	pulumi.RegisterOutputType(SavedFiltersFilterOutput{})
	pulumi.RegisterOutputType(SavedFiltersFilterArrayOutput{})
	pulumi.RegisterOutputType(GetAccessGrantsAccessGrantOutput{})
	pulumi.RegisterOutputType(GetAccessGrantsAccessGrantArrayOutput{})
	pulumi.RegisterOutputType(GetCostReportsCostReportOutput{})
	pulumi.RegisterOutputType(GetCostReportsCostReportArrayOutput{})
	pulumi.RegisterOutputType(GetDashboardsDashboardOutput{})
	pulumi.RegisterOutputType(GetDashboardsDashboardArrayOutput{})
	pulumi.RegisterOutputType(GetFoldersFolderOutput{})
	pulumi.RegisterOutputType(GetFoldersFolderArrayOutput{})
	pulumi.RegisterOutputType(GetTeamsTeamOutput{})
	pulumi.RegisterOutputType(GetTeamsTeamArrayOutput{})
	pulumi.RegisterOutputType(GetUsersUserOutput{})
	pulumi.RegisterOutputType(GetUsersUserArrayOutput{})
	pulumi.RegisterOutputType(GetWorkspacesWorkspaceOutput{})
	pulumi.RegisterOutputType(GetWorkspacesWorkspaceArrayOutput{})
}
