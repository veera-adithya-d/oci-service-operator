package apispec

import (
	"reflect"
	"slices"
	"strings"
	"testing"

	"github.com/oracle/oci-service-operator/internal/validator/allowlist"
	"github.com/oracle/oci-service-operator/internal/validator/sdk"
)

type testWidgetSpec struct {
	Name string `json:"name,omitempty"`
}

func TestBuildReportIncludesUntrackedTargets(t *testing.T) {
	originalTargets := targets
	t.Cleanup(func() {
		targets = originalTargets
	})

	targets = []Target{
		{
			Name:       "TestWidget",
			SpecType:   reflect.TypeOf(testWidgetSpec{}),
			SDKStructs: nil,
		},
	}

	report, err := BuildReport(nil, allowlist.Allowlist{})
	if err != nil {
		t.Fatalf("BuildReport() error = %v", err)
	}
	if len(report.Structs) != 1 {
		t.Fatalf("BuildReport() report count = %d, want 1", len(report.Structs))
	}

	got := report.Structs[0]
	if got.Spec != "TestWidget" {
		t.Fatalf("report.Structs[0].Spec = %q, want %q", got.Spec, "TestWidget")
	}
	if got.TrackingStatus != TrackingStatusUntracked {
		t.Fatalf("report.Structs[0].TrackingStatus = %q, want %q", got.TrackingStatus, TrackingStatusUntracked)
	}
	if !strings.Contains(got.TrackingReason, "no mapped SDK payloads") {
		t.Fatalf("report.Structs[0].TrackingReason = %q, want unmapped reason", got.TrackingReason)
	}
	if !HasActionable(report) {
		t.Fatal("HasActionable() = false, want true for untracked target")
	}
}

func TestBuildReportMarksReviewedUntrackedTargetsAsIntentional(t *testing.T) {
	originalTargets := targets
	t.Cleanup(func() {
		targets = originalTargets
	})

	targets = []Target{
		{
			Name:       "DNSZoneContent",
			SpecType:   reflect.TypeOf(struct{}{}),
			StatusType: reflect.TypeOf(struct{}{}),
			SDKStructs: nil,
		},
	}

	report, err := BuildReport(nil, allowlist.Allowlist{})
	if err != nil {
		t.Fatalf("BuildReport() error = %v", err)
	}
	if len(report.Structs) != 1 {
		t.Fatalf("BuildReport() report count = %d, want 1", len(report.Structs))
	}

	got := report.Structs[0]
	if got.TrackingStatus != TrackingStatusUntracked {
		t.Fatalf("report.Structs[0].TrackingStatus = %q, want %q", got.TrackingStatus, TrackingStatusUntracked)
	}
	if !isIntentionalUntrackedReason(got.TrackingReason) {
		t.Fatalf("report.Structs[0].TrackingReason = %q, want intentional untracked reason", got.TrackingReason)
	}
	if HasActionable(report) {
		t.Fatal("HasActionable() = true, want false for reviewed intentional untracked target")
	}
}

type testStatusMarker struct{}

type testWidgetStatus struct {
	OsokStatus  testStatusMarker `json:"status,omitempty"`
	DisplayName string           `json:"displayName,omitempty"`
}

func TestBuildReportUsesStatusSurfaceForStatusTargets(t *testing.T) {
	originalTargets := targets
	t.Cleanup(func() {
		targets = originalTargets
	})

	targets = []Target{
		{
			Name:       "TestReadOnlyWidget",
			SpecType:   reflect.TypeOf(testWidgetSpec{}),
			StatusType: reflect.TypeOf(testWidgetStatus{}),
			SDKStructs: []string{"example.Widget"},
		},
	}

	report, err := BuildReport([]sdk.SDKStruct{
		{
			QualifiedName: "example.Widget",
			Fields: []sdk.SDKField{
				{Name: "DisplayName", JSONName: "displayName"},
			},
		},
	}, allowlist.Allowlist{})
	if err != nil {
		t.Fatalf("BuildReport() error = %v", err)
	}
	if len(report.Structs) != 1 {
		t.Fatalf("BuildReport() report count = %d, want 1", len(report.Structs))
	}

	got := report.Structs[0]
	if got.APISurface != apiSurfaceStatus {
		t.Fatalf("report.Structs[0].APISurface = %q, want %q", got.APISurface, apiSurfaceStatus)
	}
	if len(got.PresentFields) != 1 || got.PresentFields[0].FieldName != "DisplayName" {
		t.Fatalf("report.Structs[0].PresentFields = %#v, want DisplayName present", got.PresentFields)
	}
	if len(got.ExtraSpecFields) != 0 {
		t.Fatalf("report.Structs[0].ExtraSpecFields = %#v, want OsokStatus skipped", got.ExtraSpecFields)
	}
	if HasActionable(report) {
		t.Fatal("HasActionable() = true, want false for covered status target")
	}
}

func TestBuildReportRoutesDesiredAndObservedSDKStructsToDifferentSurfaces(t *testing.T) {
	originalTargets := targets
	t.Cleanup(func() {
		targets = originalTargets
	})

	targets = []Target{
		{
			Name:       "TestWidget",
			SpecType:   reflect.TypeOf(testWidgetSpec{}),
			StatusType: reflect.TypeOf(testWidgetStatus{}),
			SDKStructs: []string{"example.CreateWidgetDetails", "example.Widget"},
		},
	}

	report, err := BuildReport([]sdk.SDKStruct{
		{
			QualifiedName: "example.CreateWidgetDetails",
			Fields: []sdk.SDKField{
				{Name: "Name", JSONName: "name"},
			},
		},
		{
			QualifiedName: "example.Widget",
			Fields: []sdk.SDKField{
				{Name: "DisplayName", JSONName: "displayName"},
			},
		},
	}, allowlist.Allowlist{})
	if err != nil {
		t.Fatalf("BuildReport() error = %v", err)
	}
	if len(report.Structs) != 2 {
		t.Fatalf("BuildReport() report count = %d, want 2", len(report.Structs))
	}

	bySDK := make(map[string]StructReport, len(report.Structs))
	for _, structReport := range report.Structs {
		bySDK[structReport.SDKStruct] = structReport
	}

	create := bySDK["example.CreateWidgetDetails"]
	if create.APISurface != apiSurfaceSpec {
		t.Fatalf("CreateWidgetDetails APISurface = %q, want %q", create.APISurface, apiSurfaceSpec)
	}
	if len(create.PresentFields) != 1 || create.PresentFields[0].FieldName != "Name" {
		t.Fatalf("CreateWidgetDetails PresentFields = %#v, want Name present", create.PresentFields)
	}

	observed := bySDK["example.Widget"]
	if observed.APISurface != apiSurfaceStatus {
		t.Fatalf("Widget APISurface = %q, want %q", observed.APISurface, apiSurfaceStatus)
	}
	if len(observed.PresentFields) != 1 || observed.PresentFields[0].FieldName != "DisplayName" {
		t.Fatalf("Widget PresentFields = %#v, want DisplayName present", observed.PresentFields)
	}
}

func TestReviewedUntrackedReasonsCoverAllEmptyRegistryTargets(t *testing.T) {
	t.Parallel()

	var got []string
	for _, target := range Targets() {
		if len(target.SDKStructs) != 0 {
			continue
		}
		reason := reviewedUntrackedReason(target.Name)
		if !isIntentionalUntrackedReason(reason) {
			got = append(got, target.Name)
		}
	}

	if len(got) != 0 {
		t.Fatalf("empty registry targets without reviewed untracked reasons: %v", got)
	}

	var extra []string
	for targetName := range reviewedUntrackedReasons {
		matched := false
		for _, target := range Targets() {
			if target.Name == targetName && len(target.SDKStructs) == 0 {
				matched = true
				break
			}
		}
		if !matched {
			extra = append(extra, targetName)
		}
	}
	slices.Sort(extra)
	if len(extra) != 0 {
		t.Fatalf("reviewed untracked reasons without matching empty registry targets: %v", extra)
	}
}
