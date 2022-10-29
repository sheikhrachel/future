package testutil

import (
	"encoding/json"
	"strings"
	"testing"
)

type assertionFunc func(actual interface{}, expected ...interface{}) string

func So(
	t testing.TB,
	actual interface{},
	assertion assertionFunc,
	expected ...interface{},
) {
	t.Helper()
	if msg := assertion(actual, expected...); msg != "" {
		// Sometimes msg is marshalled JSON. Detect this and extract the "Message" property.
		if strings.HasPrefix(msg, `{"Message":`) {
			var obj map[string]interface{}
			if err := json.Unmarshal([]byte(msg), &obj); err == nil {
				if m, ok := obj["Message"].(string); ok {
					msg = m
				}
			}
		}
		t.Fatal("\n" + msg)
	}
}
