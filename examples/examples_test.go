// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package examples

import (
	"os"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func getToken(t *testing.T) string {
	name := os.Getenv("VANTAGE_API_TOKEN")
	if name == "" {
		t.Skipf("Skipping test due to missing VANTAGE_API_TOKEN environment variable")
	}

	return name
}

func getCwd(t *testing.T) string {
	cwd, err := os.Getwd()
	if err != nil {
		t.FailNow()
	}

	return cwd
}

func getBaseOptions(t *testing.T) integration.ProgramTestOptions {
	token := getToken(t)
	return integration.ProgramTestOptions{
		Config: map[string]string{
			"vantage_api_token":  token,
		},
	}
}