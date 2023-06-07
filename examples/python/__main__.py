import pulumi
import lbrlabs_pulumi_vantage as vantage
import pulumi_aws as aws
import pulumi_vantage as vantage

config = pulumi.Config()
vantage_sns_topic_arn = config.get("vantageSnsTopicArn")
if vantage_sns_topic_arn is None:
    vantage_sns_topic_arn = "arn:aws:sns:us-east-1:630399649041:cost-and-usage-report-uploaded"
cur_report_name = config.get("curReportName")
if cur_report_name is None:
    cur_report_name = "VantageReport"
compatibility_private_bucket_acl = config.get_bool("compatibilityPrivateBucketAcl")
if compatibility_private_bucket_acl is None:
    compatibility_private_bucket_acl = False
enable_autopilot = config.get_bool("enableAutopilot")
if enable_autopilot is None:
    enable_autopilot = True
default = vantage.get_aws_provider_info()
vantage_assume_role = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    actions=["sts:AssumeRole"],
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="AWS",
        identifiers=[default.iam_role_arn],
    )],
    conditions=[aws.iam.GetPolicyDocumentStatementConditionArgs(
        variable="sts:ExternalId",
        test="StringEquals",
        values=[default.external_id],
    )],
)])
vantage_cost_and_usage_reports_bucket = aws.s3.BucketV2("vantageCostAndUsageReportsBucket", force_destroy=True)
vantage_cur_retrieval = aws.iam.get_policy_document_output(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    actions=[
        "s3:GetObject",
        "s3:GetObjectAcl",
    ],
    resources=[vantage_cost_and_usage_reports_bucket.arn.apply(lambda arn: f"{arn}/*")],
)])
vantage_cross_account_connection_iam_role = aws.iam.Role("vantageCrossAccountConnectionIamRole",
    inline_policies=[
        aws.iam.RoleInlinePolicyArgs(
            name="VantageCostandUsageReportRetrieval",
            policy=vantage_cur_retrieval.json,
        ),
        aws.iam.RoleInlinePolicyArgs(
            name="root",
            policy=default.root_policy,
        ),
        aws.iam.RoleInlinePolicyArgs(
            name="VantageCloudWatchMetricsReadOnly",
            policy=default.cloudwatch_metrics_policy,
        ),
        aws.iam.RoleInlinePolicyArgs(
            name="VantageAdditionalResourceReadOnly",
            policy=default.additional_resources_policy,
        ),
    ],
    name="vantage_cross_account_connection",
    assume_role_policy=vantage_assume_role.json)
vantage_cur_access = aws.iam.get_policy_document_output(statements=[
    aws.iam.GetPolicyDocumentStatementArgs(
        effect="Allow",
        actions=[
            "s3:GetBucketAcl",
            "s3:GetBucketPolicy",
        ],
        principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
            type="Service",
            identifiers=["billingreports.amazonaws.com"],
        )],
        resources=[vantage_cost_and_usage_reports_bucket.arn],
    ),
    aws.iam.GetPolicyDocumentStatementArgs(
        effect="Allow",
        actions=["s3:PutObject"],
        principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
            type="Service",
            identifiers=["billingreports.amazonaws.com"],
        )],
        resources=[vantage_cost_and_usage_reports_bucket.arn.apply(lambda arn: f"{arn}/*")],
    ),
    aws.iam.GetPolicyDocumentStatementArgs(
        effect="Allow",
        actions=[
            "s3:GetObject",
            "s3:GetObjectAcl",
        ],
        principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
            type="AWS",
            identifiers=[vantage_cross_account_connection_iam_role.arn],
        )],
        resources=[vantage_cost_and_usage_reports_bucket.arn.apply(lambda arn: f"{arn}/*")],
    ),
])
vantage_cross_account_connection_policy = aws.iam.RolePolicyAttachment("vantageCrossAccountConnectionPolicy",
    role=vantage_cross_account_connection_iam_role.name,
    policy_arn="arn:aws:iam::aws:policy/job-function/ViewOnlyAccess",
    opts=pulumi.ResourceOptions(parent=vantage_cross_account_connection_iam_role))
vantage_cost_and_usage_reports_bucket_ownership_controls = aws.s3.BucketOwnershipControls("vantageCostAndUsageReportsBucketOwnershipControls",
    bucket=vantage_cost_and_usage_reports_bucket.id,
    rule=aws.s3.BucketOwnershipControlsRuleArgs(
        object_ownership="ObjectWriter",
    ),
    opts=pulumi.ResourceOptions(parent=vantage_cost_and_usage_reports_bucket))
vantage_cost_and_usage_reports_bucket_acl = aws.s3.BucketAclV2("vantageCostAndUsageReportsBucketAcl",
    bucket=vantage_cost_and_usage_reports_bucket.id,
    acl="private",
    opts=pulumi.ResourceOptions(parent=vantage_cost_and_usage_reports_bucket,
        depends_on=[vantage_cost_and_usage_reports_bucket_ownership_controls]))
vantage_cost_and_usage_reports_lifecycle_configuration = aws.s3.BucketLifecycleConfigurationV2("vantageCostAndUsageReportsLifecycleConfiguration",
    bucket=vantage_cost_and_usage_reports_bucket.id,
    rules=[aws.s3.BucketLifecycleConfigurationV2RuleArgs(
        id="remove-old-reports",
        expiration=aws.s3.BucketLifecycleConfigurationV2RuleExpirationArgs(
            days=200,
        ),
        status="Enabled",
    )],
    opts=pulumi.ResourceOptions(parent=vantage_cost_and_usage_reports_bucket))
vantage_cost_and_usage_report_public_access_block = aws.s3.BucketPublicAccessBlock("vantageCostAndUsageReportPublicAccessBlock",
    bucket=vantage_cost_and_usage_reports_bucket.id,
    block_public_acls=True,
    block_public_policy=True,
    restrict_public_buckets=True,
    ignore_public_acls=True,
    opts=pulumi.ResourceOptions(parent=vantage_cost_and_usage_reports_bucket))
vantage_cost_and_usage_reports_bucket_policy = aws.s3.BucketPolicy("vantageCostAndUsageReportsBucketPolicy",
    bucket=vantage_cost_and_usage_reports_bucket.id,
    policy=vantage_cur_access.json,
    opts=pulumi.ResourceOptions(parent=vantage_cost_and_usage_reports_bucket))
vantage_cost_and_usage_reports = aws.cur.ReportDefinition("vantageCostAndUsageReports",
    report_name="VantageUsageReports",
    time_unit="DAILY",
    format="textORcsv",
    compression="GZIP",
    additional_schema_elements=["RESOURCES"],
    s3_bucket=vantage_cost_and_usage_reports_bucket.id,
    s3_region="us-east-1",
    s3_prefix="daily-v1",
    report_versioning="OVERWRITE_REPORT",
    refresh_closed_reports=True,
    opts=pulumi.ResourceOptions(depends_on=[
            vantage_cost_and_usage_reports_bucket,
            vantage_cost_and_usage_reports_bucket_ownership_controls,
            vantage_cost_and_usage_reports_bucket_acl,
            vantage_cost_and_usage_reports_lifecycle_configuration,
            vantage_cost_and_usage_report_public_access_block,
            vantage_cost_and_usage_reports_bucket_policy,
        ]))
vantage_cost_and_usage_reports_bucket_notification = aws.s3.BucketNotification("vantageCostAndUsageReportsBucketNotification",
    bucket=vantage_cost_and_usage_reports_bucket.id,
    topics=[aws.s3.BucketNotificationTopicArgs(
        topic_arn=vantage_sns_topic_arn,
        events=["s3:ObjectCreated:*"],
        filter_suffix=".csv.gz",
    )],
    opts=pulumi.ResourceOptions(parent=vantage_cost_and_usage_reports_bucket))
vantage_aws_provider = vantage.AwsProvider("VantageAwsProvider",
    cross_account_arn=vantage_cross_account_connection_iam_role.arn,
    bucket_arn=vantage_cost_and_usage_reports_bucket.arn)
