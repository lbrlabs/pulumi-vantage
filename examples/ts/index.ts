import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as vantage from "@lbrlabs/pulumi-vantage";

const config = new pulumi.Config();
const vantageSnsTopicArn = config.get("vantageSnsTopicArn") || "arn:aws:sns:us-east-1:630399649041:cost-and-usage-report-uploaded";
const curReportName = config.get("curReportName") || "VantageReport";
const compatibilityPrivateBucketAcl = config.getBoolean("compatibilityPrivateBucketAcl") || false;
const enableAutopilot = config.getBoolean("enableAutopilot") || true;
const default = vantage.getAwsProviderInfo({});
const vantageAssumeRole = Promise.all([_default, _default]).then(([_default, _default1]) => aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        actions: ["sts:AssumeRole"],
        principals: [{
            type: "AWS",
            identifiers: [_default.iamRoleArn],
        }],
        conditions: [{
            variable: "sts:ExternalId",
            test: "StringEquals",
            values: [_default1.externalId],
        }],
    }],
}));
const vantageCostAndUsageReportsBucket = new aws.s3.BucketV2("vantageCostAndUsageReportsBucket", {forceDestroy: true});
const vantageCurRetrieval = aws.iam.getPolicyDocumentOutput({
    statements: [{
        effect: "Allow",
        actions: [
            "s3:GetObject",
            "s3:GetObjectAcl",
        ],
        resources: [pulumi.interpolate`${vantageCostAndUsageReportsBucket.arn}/*`],
    }],
});
const vantageCrossAccountConnectionIamRole = new aws.iam.Role("vantageCrossAccountConnectionIamRole", {
    inlinePolicies: [
        {
            name: "VantageCostandUsageReportRetrieval",
            policy: vantageCurRetrieval.json,
        },
        {
            name: "root",
            policy: _default.then(_default => _default.rootPolicy),
        },
        {
            name: "VantageCloudWatchMetricsReadOnly",
            policy: _default.then(_default => _default.cloudwatchMetricsPolicy),
        },
        {
            name: "VantageAdditionalResourceReadOnly",
            policy: _default.then(_default => _default.additionalResourcesPolicy),
        },
    ],
    name: "vantage_cross_account_connection",
    assumeRolePolicy: vantageAssumeRole.then(vantageAssumeRole => vantageAssumeRole.json),
});
const vantageCurAccess = aws.iam.getPolicyDocumentOutput({
    statements: [
        {
            effect: "Allow",
            actions: [
                "s3:GetBucketAcl",
                "s3:GetBucketPolicy",
            ],
            principals: [{
                type: "Service",
                identifiers: ["billingreports.amazonaws.com"],
            }],
            resources: [vantageCostAndUsageReportsBucket.arn],
        },
        {
            effect: "Allow",
            actions: ["s3:PutObject"],
            principals: [{
                type: "Service",
                identifiers: ["billingreports.amazonaws.com"],
            }],
            resources: [pulumi.interpolate`${vantageCostAndUsageReportsBucket.arn}/*`],
        },
        {
            effect: "Allow",
            actions: [
                "s3:GetObject",
                "s3:GetObjectAcl",
            ],
            principals: [{
                type: "AWS",
                identifiers: [vantageCrossAccountConnectionIamRole.arn],
            }],
            resources: [pulumi.interpolate`${vantageCostAndUsageReportsBucket.arn}/*`],
        },
    ],
});
const vantageCrossAccountConnectionPolicy = new aws.iam.RolePolicyAttachment("vantageCrossAccountConnectionPolicy", {
    role: vantageCrossAccountConnectionIamRole.name,
    policyArn: "arn:aws:iam::aws:policy/job-function/ViewOnlyAccess",
}, {
    parent: vantageCrossAccountConnectionIamRole,
});
const vantageCostAndUsageReportsBucketOwnershipControls = new aws.s3.BucketOwnershipControls("vantageCostAndUsageReportsBucketOwnershipControls", {
    bucket: vantageCostAndUsageReportsBucket.id,
    rule: {
        objectOwnership: "ObjectWriter",
    },
}, {
    parent: vantageCostAndUsageReportsBucket,
});
const vantageCostAndUsageReportsBucketAcl = new aws.s3.BucketAclV2("vantageCostAndUsageReportsBucketAcl", {
    bucket: vantageCostAndUsageReportsBucket.id,
    acl: "private",
}, {
    parent: vantageCostAndUsageReportsBucket,
    dependsOn: [vantageCostAndUsageReportsBucketOwnershipControls],
});
const vantageCostAndUsageReportsLifecycleConfiguration = new aws.s3.BucketLifecycleConfigurationV2("vantageCostAndUsageReportsLifecycleConfiguration", {
    bucket: vantageCostAndUsageReportsBucket.id,
    rules: [{
        id: "remove-old-reports",
        expiration: {
            days: 200,
        },
        status: "Enabled",
    }],
}, {
    parent: vantageCostAndUsageReportsBucket,
});
const vantageCostAndUsageReportPublicAccessBlock = new aws.s3.BucketPublicAccessBlock("vantageCostAndUsageReportPublicAccessBlock", {
    bucket: vantageCostAndUsageReportsBucket.id,
    blockPublicAcls: true,
    blockPublicPolicy: true,
    restrictPublicBuckets: true,
    ignorePublicAcls: true,
}, {
    parent: vantageCostAndUsageReportsBucket,
});
const vantageCostAndUsageReportsBucketPolicy = new aws.s3.BucketPolicy("vantageCostAndUsageReportsBucketPolicy", {
    bucket: vantageCostAndUsageReportsBucket.id,
    policy: vantageCurAccess.json,
}, {
    parent: vantageCostAndUsageReportsBucket,
});
const vantageCostAndUsageReports = new aws.cur.ReportDefinition("vantageCostAndUsageReports", {
    reportName: "VantageUsageReports",
    timeUnit: "DAILY",
    format: "textORcsv",
    compression: "GZIP",
    additionalSchemaElements: ["RESOURCES"],
    s3Bucket: vantageCostAndUsageReportsBucket.id,
    s3Region: "us-east-1",
    s3Prefix: "daily-v1",
    reportVersioning: "OVERWRITE_REPORT",
    refreshClosedReports: true,
}, {
    dependsOn: [
        vantageCostAndUsageReportsBucket,
        vantageCostAndUsageReportsBucketOwnershipControls,
        vantageCostAndUsageReportsBucketAcl,
        vantageCostAndUsageReportsLifecycleConfiguration,
        vantageCostAndUsageReportPublicAccessBlock,
        vantageCostAndUsageReportsBucketPolicy,
    ],
});
const vantageCostAndUsageReportsBucketNotification = new aws.s3.BucketNotification("vantageCostAndUsageReportsBucketNotification", {
    bucket: vantageCostAndUsageReportsBucket.id,
    topics: [{
        topicArn: vantageSnsTopicArn,
        events: ["s3:ObjectCreated:*"],
        filterSuffix: ".csv.gz",
    }],
}, {
    parent: vantageCostAndUsageReportsBucket,
});
const vantageAwsProvider = new vantage.AwsProvider("VantageAwsProvider", {
    crossAccountArn: vantageCrossAccountConnectionIamRole.arn,
    bucketArn: vantageCostAndUsageReportsBucket.arn,
});
