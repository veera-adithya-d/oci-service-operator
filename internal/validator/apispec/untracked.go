package apispec

import "strings"

const intentionalUntrackedPrefix = "Intentionally untracked: "

var reviewedUntrackedReasons = map[string]string{}

func reviewedUntrackedReason(targetName string) string {
	return reviewedUntrackedReasons[strings.TrimSpace(targetName)]
}

func isIntentionalUntrackedReason(reason string) bool {
	return strings.HasPrefix(strings.TrimSpace(reason), intentionalUntrackedPrefix)
}

func readOnlyResponseReason(sdkType string) string {
	return intentionalUntrackedPrefix + "spec is empty and the SDK only exposes read-only response payloads via " + sdkType + ", which would make every field appear missing from desired state."
}

func responseBodyReason(sdkType string) string {
	return intentionalUntrackedPrefix + "spec is empty and the SDK only returns " + sdkType + " in the response body, not as a desired-state payload."
}

func scalarContentReason(description string) string {
	return intentionalUntrackedPrefix + "spec is empty and " + description + ", not a reusable SDK struct for desired-state validation."
}

func binaryContentReason(description string) string {
	return intentionalUntrackedPrefix + "spec is empty and " + description + ", not a reusable SDK struct for desired-state validation."
}
