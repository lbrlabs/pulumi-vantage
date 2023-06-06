// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tls

import (
	_ "embed"
	"fmt"
	"path/filepath"
	"unicode"

	"github.com/lbrlabs/pulumi-vantage/provider/pkg/version"
	tfpfbridge "github.com/pulumi/pulumi-terraform-bridge/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/vantage-sh/terraform-provider-vantage/shim"
)

// all of the tls token components used below.
const (
	vantagePkg = "vantage"
	vantageMod = "index"
)

// vantageMember manufactures a type token for the TLS package and the given module and type.
func vantageMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(vantagePkg + ":" + mod + ":" + mem)
}

// vantageType manufactures a type token for the TLS package and the given module and type.
func vantageType(mod string, typ string) tokens.Type {
	return tokens.Type(vantageMember(mod, typ))
}

// vantageDataSource manufactures a standard resource token given a module and resource name.
// It automatically uses the TLS package and names the file by simply lower casing the data
// source's first character.
func vantageDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return vantageMember(mod+"/"+fn, res)
}

// vantageResource manufactures a standard resource token given a module and resource name.
// It automatically uses the TLS package and names the file by simply lower casing the resource's
// first character.
func vantageResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return vantageType(mod+"/"+fn, res)
}

//go:embed cmd/pulumi-resource-vantage/bridge-metadata.json
var bridgeMetadata []byte

// Provider returns additional overlaid schema and metadata associated with the tls package.
func Provider() tfpfbridge.ProviderInfo {
	info := tfbridge.ProviderInfo{
		Name:              "vantage",
		DisplayName:       "Vantage",
		Publisher:         "lbrlabs",
		Description:       "A Pulumi package to create resource in vantage.sh.",
		PluginDownloadURL: "github://api.github.com/lbrlabs",
		Keywords:          []string{"pulumi", "vantage"},
		License:           "Apache-2.0",
		Homepage:          "https://pulumi.io",
		Repository:        "https://github.com/lbrlabs/pulumi-vantage",
		Version:           version.Version,
		GitHubOrg:         "vantage-sh",
		MetadataInfo:      tfbridge.NewProviderMetadata(bridgeMetadata),
		Resources: map[string]*tfbridge.ResourceInfo{
			"vantage_aws_provider": {
				Tok: vantageResource(vantageMod, "AwsProvider"),
			},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"vantage_aws_provider_info": {Tok: vantageDataSource(vantageMod, "getAwsProviderInfo")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@lbrlabs/pulumi-vantage",
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
			},
		},
		Python: &tfbridge.PythonInfo{
			PackageName: "lbrlabs_pulumi_vantage",
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/lbrlabs/pulumi-%[1]s/sdk/", vantagePkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				vantagePkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RootNamespace: "Lbrlabs.PulumiPackage",
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
			Namespaces: map[string]string{
				"vantage": "Vantage",
			},
		},
	}

	return tfpfbridge.ProviderInfo{
		ProviderInfo: info,
		NewProvider:  shim.NewProvider,
	}
}
