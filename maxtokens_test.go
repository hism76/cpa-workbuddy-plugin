package main

import (
	"encoding/json"
	"testing"
)

func TestNormalizeMaxTokensInPlace(t *testing.T) {
	// 1. Default to 64000 when neither max_tokens nor max_completion_tokens is provided
	body1 := []byte(`{"model":"hy4-preview","messages":[]}`)
	out1 := prepareUpstreamBody(body1, nil, nil, "hy4-preview")
	var obj1 map[string]any
	if err := json.Unmarshal(out1, &obj1); err != nil {
		t.Fatal(err)
	}
	if obj1["max_tokens"] != float64(64000) || obj1["max_completion_tokens"] != float64(64000) {
		t.Fatalf("default want 64000, got max_tokens=%v, max_completion_tokens=%v", obj1["max_tokens"], obj1["max_completion_tokens"])
	}

	// 2. max_tokens specified
	body2 := []byte(`{"model":"hy4-preview","max_tokens":2048,"messages":[]}`)
	out2 := prepareUpstreamBody(body2, nil, nil, "hy4-preview")
	var obj2 map[string]any
	if err := json.Unmarshal(out2, &obj2); err != nil {
		t.Fatal(err)
	}
	if obj2["max_tokens"] != float64(2048) || obj2["max_completion_tokens"] != float64(2048) {
		t.Fatalf("max_tokens want 2048, got max_tokens=%v, max_completion_tokens=%v", obj2["max_tokens"], obj2["max_completion_tokens"])
	}

	// 3. max_completion_tokens overrides max_tokens (priority)
	body3 := []byte(`{"model":"hy4-preview","max_tokens":2048,"max_completion_tokens":8192,"messages":[]}`)
	out3 := prepareUpstreamBody(body3, nil, nil, "hy4-preview")
	var obj3 map[string]any
	if err := json.Unmarshal(out3, &obj3); err != nil {
		t.Fatal(err)
	}
	if obj3["max_tokens"] != float64(8192) || obj3["max_completion_tokens"] != float64(8192) {
		t.Fatalf("max_completion_tokens priority want 8192, got max_tokens=%v, max_completion_tokens=%v", obj3["max_tokens"], obj3["max_completion_tokens"])
	}
}
