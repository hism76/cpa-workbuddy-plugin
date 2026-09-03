package main

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/router-for-me/CLIProxyAPI/v7/sdk/pluginapi"
)

func TestParseStoredAutoDetectsGlobalFromJWT(t *testing.T) {
	issuer := "https://www.codebuddy.ai/auth/realms/copilot"
	tok := syntheticAccessToken(t, issuer)
	raw := []byte(`{"accessToken":"` + tok + `"}`)

	sa, err := parseStored(raw)
	if err != nil {
		t.Fatalf("parseStored failed: %v", err)
	}
	if sa.Auth.Domain != "www.codebuddy.ai" {
		t.Errorf("expected domain www.codebuddy.ai, got %q", sa.Auth.Domain)
	}
	if !isGlobalDomain(sa.Auth.Domain) {
		t.Errorf("expected isGlobalDomain true for %q", sa.Auth.Domain)
	}
	if got := accountRegion(sa); got != "global" {
		t.Errorf("expected accountRegion global, got %q", got)
	}
	if got := upstreamBaseFor(sa); got != "https://www.codebuddy.ai" {
		t.Errorf("expected upstreamBaseFor https://www.codebuddy.ai, got %q", got)
	}
	if got := originRefererFor(sa); got != "https://www.codebuddy.ai" {
		t.Errorf("expected originRefererFor https://www.codebuddy.ai, got %q", got)
	}
}

func TestManagementLoginStartGlobal(t *testing.T) {
	base := loadedManagementBasePath() + "/plugins/" + providerName
	reqBody := []byte(`{"region":"global"}`)
	resp := managementResponseForTest(t, pluginapi.ManagementRequest{
		Method: http.MethodPost,
		Path:   base + "/login/start",
		Body:   reqBody,
	})
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}
	var res struct {
		Success bool   `json:"success"`
		URL     string `json:"url"`
		Region  string `json:"region"`
	}
	if err := json.Unmarshal(resp.Body, &res); err != nil {
		t.Fatalf("failed to parse response: %v", err)
	}
	if !res.Success {
		t.Fatalf("expected success true, got false")
	}
	if res.Region != "global" {
		t.Errorf("expected region global, got %q", res.Region)
	}
	if res.URL == "" {
		t.Errorf("expected non-empty login URL")
	}
}
