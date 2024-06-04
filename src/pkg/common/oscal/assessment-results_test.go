package oscal_test

import (
	"testing"
	"time"

	"github.com/defenseunicorns/go-oscal/src/pkg/uuid"
	oscalTypes_1_1_2 "github.com/defenseunicorns/go-oscal/src/types/oscal-1-1-2"
	"github.com/defenseunicorns/lula/src/pkg/common/oscal"
	"github.com/defenseunicorns/lula/src/pkg/message"
)

// Create re-usable findings and observations
// use those in tests to generate test assessment results
var findingMapPass = map[string]oscalTypes_1_1_2.Finding{
	"ID-1": {
		Target: oscalTypes_1_1_2.FindingTarget{
			TargetId: "ID-1",
			Status: oscalTypes_1_1_2.ObjectiveStatus{
				State: "satisfied",
			},
		},
	},
}

var findingMapFail = map[string]oscalTypes_1_1_2.Finding{
	"ID-1": {
		Target: oscalTypes_1_1_2.FindingTarget{
			TargetId: "ID-1",
			Status: oscalTypes_1_1_2.ObjectiveStatus{
				State: "not-satisfied",
			},
		},
	},
}

var observations = []oscalTypes_1_1_2.Observation{
	{
		Collected:   time.Now(),
		Methods:     []string{"TEST"},
		UUID:        uuid.NewUUID(),
		Description: "test description",
	},
	{
		Collected:   time.Now(),
		Methods:     []string{"TEST"},
		UUID:        uuid.NewUUID(),
		Description: "test description",
	},
}

func TestIdentifyResults(t *testing.T) {
	t.Parallel()

	t.Run("Handle valid assessment containing a single result", func(t *testing.T) {

		assessment, err := oscal.GenerateAssessmentResults(findingMapPass, observations)
		if err != nil {
			t.Fatalf("error generating assessment results: %v", err)
		}

		// key name does not matter here
		var assessmentMap = map[string]*oscalTypes_1_1_2.AssessmentResults{
			"valid.yaml": assessment,
		}

		_, err = oscal.IdentifyResults(assessmentMap)
		if err == nil {
			t.Fatalf("Expected error for inability to identify multiple results : %v", err)
		}
	})

	t.Run("Handle multiple valid assessment containing a single result", func(t *testing.T) {

		assessment, err := oscal.GenerateAssessmentResults(findingMapPass, observations)
		if err != nil {
			t.Fatalf("error generating assessment results: %v", err)
		}

		assessment2, err := oscal.GenerateAssessmentResults(findingMapFail, observations)
		if err != nil {
			t.Fatalf("error generating assessment results: %v", err)
		}

		// key name does not matter here
		var assessmentMap = map[string]*oscalTypes_1_1_2.AssessmentResults{
			"valid.yaml": assessment,
			"other.yaml": assessment2,
		}

		resultMap, err := oscal.IdentifyResults(assessmentMap)
		if err != nil {
			t.Fatalf("Expected error for inability to identify multiple results : %v", err)
		}

		if resultMap["threshold"] == nil || resultMap["latest"] == nil {
			t.Fatalf("Expected results to be identified")
		}
	})
}

// Given two results - evaluate for passing
func TestEvaluateResultsPassing(t *testing.T) {
	message.NoProgress = true

	mockThresholdResult := oscalTypes_1_1_2.Result{
		Findings: &[]oscalTypes_1_1_2.Finding{
			findingMapPass["ID-1"],
		},
	}

	mockEvaluationResult := oscalTypes_1_1_2.Result{
		Findings: &[]oscalTypes_1_1_2.Finding{
			findingMapPass["ID-1"],
		},
	}

	status, _, err := oscal.EvaluateResults(&mockThresholdResult, &mockEvaluationResult)
	if err != nil {
		t.Fatal(err)
	}

	// If status is false - then something went wrong
	if !status {
		t.Fatal("error - evaluation failed")
	}

}

func TestEvaluateResultsFailed(t *testing.T) {
	message.NoProgress = true
	mockThresholdResult := oscalTypes_1_1_2.Result{
		Findings: &[]oscalTypes_1_1_2.Finding{
			findingMapPass["ID-1"],
		},
	}

	mockEvaluationResult := oscalTypes_1_1_2.Result{
		Findings: &[]oscalTypes_1_1_2.Finding{
			findingMapFail["ID-1"],
		},
	}

	status, findings, err := oscal.EvaluateResults(&mockThresholdResult, &mockEvaluationResult)
	if err != nil {
		t.Fatal(err)
	}

	// If status is true - then something went wrong
	if status {
		t.Fatal("error - evaluation was successful when it should have failed")
	}

	if len(findings["no-longer-satisfied"]) != 1 {
		t.Fatal("error - expected 1 finding, got ", len(findings["no-longer-satisfied"]))
	}

}

func TestEvaluateResultsNoFindings(t *testing.T) {
	message.NoProgress = true
	mockThresholdResult := oscalTypes_1_1_2.Result{
		Findings: &[]oscalTypes_1_1_2.Finding{},
	}

	mockEvaluationResult := oscalTypes_1_1_2.Result{
		Findings: &[]oscalTypes_1_1_2.Finding{},
	}

	status, _, err := oscal.EvaluateResults(&mockThresholdResult, &mockEvaluationResult)
	if err != nil {
		t.Fatal(err)
	}

	// If status is false - then something went wrong
	if !status {
		t.Fatal("error - evaluation failed")
	}

}

func TestEvaluateResultsNoThreshold(t *testing.T) {
	message.NoProgress = true
	mockThresholdResult := oscalTypes_1_1_2.Result{}

	mockEvaluationResult := oscalTypes_1_1_2.Result{
		Findings: &[]oscalTypes_1_1_2.Finding{
			{
				Target: oscalTypes_1_1_2.FindingTarget{
					TargetId: "ID-1",
					Status: oscalTypes_1_1_2.ObjectiveStatus{
						State: "satisfied",
					},
				},
			},
		},
	}

	_, _, err := oscal.EvaluateResults(&mockThresholdResult, &mockEvaluationResult)
	if err == nil {
		t.Fatal("error - expected error, got nil")
	}
}

func TestEvaluateResultsNewFindings(t *testing.T) {
	message.NoProgress = true
	mockThresholdResult := oscalTypes_1_1_2.Result{
		Findings: &[]oscalTypes_1_1_2.Finding{
			{
				Target: oscalTypes_1_1_2.FindingTarget{
					TargetId: "ID-1",
					Status: oscalTypes_1_1_2.ObjectiveStatus{
						State: "satisfied",
					},
				},
			},
		},
	}
	// Adding two new findings
	mockEvaluationResult := oscalTypes_1_1_2.Result{
		Findings: &[]oscalTypes_1_1_2.Finding{
			{
				Target: oscalTypes_1_1_2.FindingTarget{
					TargetId: "ID-1",
					Status: oscalTypes_1_1_2.ObjectiveStatus{
						State: "satisfied",
					},
				},
			},
			{
				Target: oscalTypes_1_1_2.FindingTarget{
					TargetId: "ID-2",
					Status: oscalTypes_1_1_2.ObjectiveStatus{
						State: "satisfied",
					},
				},
			},
			{
				Target: oscalTypes_1_1_2.FindingTarget{
					TargetId: "ID-3",
					Status: oscalTypes_1_1_2.ObjectiveStatus{
						State: "not-satisfied",
					},
				},
			},
		},
	}

	status, findings, err := oscal.EvaluateResults(&mockThresholdResult, &mockEvaluationResult)
	if err != nil {
		t.Fatal(err)
	}

	// If status is false - then something went wrong
	if !status {
		t.Fatal("error - evaluation failed")
	}

	if len(findings["new-passing-findings"]) != 1 {
		t.Fatal("error - expected 1 new finding, got ", len(findings["new-passing-findings"]))
	}

}