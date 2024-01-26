package logs

import (
	"fmt"
	"strings"
)

type FilterBuilder struct {
	conditions []string
}

func NewFilterBuilder() *FilterBuilder {
	return &FilterBuilder{}
}

func (fb *FilterBuilder) BySeverity(severity string) *FilterBuilder {
	condition := fmt.Sprintf("severity = \"%s\"", severity)
	fb.conditions = append(fb.conditions, condition)
	return fb
}

func (fb *FilterBuilder) ByText(text string) *FilterBuilder {
	condition := fmt.Sprintf("textPayload = \"%s\"", text)
	fb.conditions = append(fb.conditions, condition)
	return fb
}

func (fb *FilterBuilder) ByTimeRange(from string, to string) *FilterBuilder {
	condition := fmt.Sprintf("timestamp >= \"%s\" AND timestamp <= \"%s\"", from, to)
	fb.conditions = append(fb.conditions, condition)
	return fb
}

func (fb *FilterBuilder) CustomFilter(filter, operand, value string) *FilterBuilder {
	condition := fmt.Sprintf("%s %s \"%s\"", filter, operand, value)
	fb.conditions = append(fb.conditions, condition)
	return fb
}

func (fb *FilterBuilder) Build() string {
	return strings.Join(fb.conditions, " AND ")
}
