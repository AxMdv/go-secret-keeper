package build

import "fmt"

var (
	buildVersion string
	buildDate    string
	buildCommit  string
)

func PrintBuildInfo() {
	fmt.Printf("Build version:%s\n", formatValue(buildVersion))
	fmt.Printf("Build date:%s\n", formatValue(buildDate))
	fmt.Printf("Build commit:%s\n", formatValue(buildCommit))
}

func formatValue(buildData string) string {
	if buildData == "" {
		return "N/A"
	}
	return buildData
}
