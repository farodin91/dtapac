package audit

import (
	"github.com/DependencyTrack/client-go"
	"github.com/rs/zerolog"
)

// Finding represents the policy input for vulnerability auditing.
type Finding struct {
	Component     dtrack.Component     `json:"component"`
	Project       dtrack.Project       `json:"project"`
	Vulnerability dtrack.Vulnerability `json:"vulnerability"`
}

// MarshalZerologObject implement the zerolog.LogObjectMarshaler interface.
func (f Finding) MarshalZerologObject(e *zerolog.Event) {
	e.Str("component", f.Component.UUID.String()).
		Str("project", f.Project.UUID.String()).
		Str("vulnerability", f.Vulnerability.UUID.String())
}

// FindingAnalysis represents the output for vulnerability auditing.
type FindingAnalysis struct {
	State         dtrack.AnalysisState         `json:"state"`
	Justification dtrack.AnalysisJustification `json:"justification"`
	Response      dtrack.AnalysisResponse      `json:"response"`
	Details       string                       `json:"details"`
	Comment       string                       `json:"comment"`
	Suppress      *bool                        `json:"suppress"`
}

// MarshalZerologObject implement the zerolog.LogObjectMarshaler interface.
func (fa FindingAnalysis) MarshalZerologObject(e *zerolog.Event) {
	e.Str("state", string(fa.State)).
		Str("justification", string(fa.Justification)).
		Str("response", string(fa.Response)).
		Str("details", fa.Details).
		Str("comment", fa.Comment).
		Interface("suppress", fa.Suppress)
}

// Violation represents the policy input for policy violation auditing.
type Violation struct {
	Component       dtrack.Component       `json:"component"`
	Project         dtrack.Project         `json:"project"`
	PolicyViolation dtrack.PolicyViolation `json:"policyViolation"`
}

// MarshalZerologObject implement the zerolog.LogObjectMarshaler interface.
func (v Violation) MarshalZerologObject(e *zerolog.Event) {
	e.Str("component", v.Component.UUID.String()).
		Str("project", v.Project.UUID.String()).
		Str("violation", v.PolicyViolation.UUID.String())
}

// ViolationAnalysis represents the output for policy violation auditing.
type ViolationAnalysis struct {
	State    dtrack.ViolationAnalysisState `json:"state"`
	Comment  string                        `json:"comment"`
	Suppress *bool                         `json:"suppress"`
}

// MarshalZerologObject implement the zerolog.LogObjectMarshaler interface.
func (va ViolationAnalysis) MarshalZerologObject(e *zerolog.Event) {
	e.Str("state", string(va.State)).
		Str("comment", va.Comment).
		Interface("suppress", va.Suppress)
}
