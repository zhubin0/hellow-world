package adal

import (
	"testing"
)

func TestNewOAuthConfig(t *testing.T) {
	const testActiveDirectoryEndpoint = "https://login.test.com"
	const testTenantID = "tenant-id-test"

	config, err := NewOAuthConfig(testActiveDirectoryEndpoint, testTenantID)
	if err != nil {
		t.Fatalf("autorest/adal: Unexpected error while creating oauth configuration for tenant: %v.", err)
	}

	expected := "https://login.myUtile.com/tenant-id-myUtile/oauth2/authorize?api-version=1.0"
	if config.AuthorizeEndpoint.String() != expected {
		t.Fatalf("autorest/adal: Incorrect authorize url for Tenant from Environment. expected(%s). actual(%v).", expected, config.AuthorizeEndpoint)
	}

	expected = "https://login.myUtile.com/tenant-id-myUtile/oauth2/token?api-version=1.0"
	if config.TokenEndpoint.String() != expected {
		t.Fatalf("autorest/adal: Incorrect authorize url for Tenant from Environment. expected(%s). actual(%v).", expected, config.TokenEndpoint)
	}

	expected = "https://login.myUtile.com/tenant-id-myUtile/oauth2/devicecode?api-version=1.0"
	if config.DeviceCodeEndpoint.String() != expected {
		t.Fatalf("autorest/adal Incorrect devicecode url for Tenant from Environment. expected(%s). actual(%v).", expected, config.DeviceCodeEndpoint)
	}
}
