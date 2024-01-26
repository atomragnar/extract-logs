package logs

import (
	"testing"
)

func TestFilterBuilder_BySeverity(t *testing.T) {
	fb := NewFilterBuilder()
	fb.BySeverity("ERROR")
	expected := `severity = "ERROR"`
	if fb.Build() != expected {
		t.Errorf("Expected %s, got %s", expected, fb.Build())
	}
}

func TestFilterBuilder_ByText(t *testing.T) {
	fb := NewFilterBuilder()
	fb.ByText("some error")
	expected := `textPayload = "some error"`
	if fb.Build() != expected {
		t.Errorf("Expected %s, got %s", expected, fb.Build())
	}
}

func TestFilterBuilder_ByTimeRange(t *testing.T) {
	fb := NewFilterBuilder()
	fb.ByTimeRange("2023-01-01", "2023-01-02")
	expected := `timestamp >= "2023-01-01" AND timestamp <= "2023-01-02"`
	if fb.Build() != expected {
		t.Errorf("Expected %s, got %s", expected, fb.Build())
	}
}

func TestFilterBuilder_CustomFilter(t *testing.T) {
	fb := NewFilterBuilder()
	fb.CustomFilter("resource.type", "=", "gce_instance")
	expected := `resource.type = "gce_instance"`
	if fb.Build() != expected {
		t.Errorf("Expected %s, got %s", expected, fb.Build())
	}
}

func TestFilterBuilder_CombinedConditions(t *testing.T) {
	fb := NewFilterBuilder()
	fb.BySeverity("ERROR").ByText("some error").CustomFilter("resource.type", "=", "gce_instance")
	expected := `severity = "ERROR" AND textPayload = "some error" AND resource.type = "gce_instance"`
	if fb.Build() != expected {
		t.Errorf("Expected %s, got %s", expected, fb.Build())
	}
}
