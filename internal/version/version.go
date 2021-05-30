package version

// Version represents current version.
const Version = "0.2.0"

var (
	// UIVersion is build time var and represents version of the user interface
	UIVersion string
	// GitCommit is build time var and represents curret git commit hash
	GitCommit string
	// BuildDate is build time var and represents build datetime
	BuildDate string
)

// BuildInfo defines build information
type BuildInfo struct {
	Version   string `json:"version"`
	UIVersion string `json:"ui_version"`
	GitCommit string `json:"git_commit"`
	BuildDate string `json:"build_date"`
}

// GetBuildInfo returns build information
func GetBuildInfo() BuildInfo {
	return BuildInfo{
		Version,
		UIVersion,
		GitCommit,
		BuildDate,
	}
}

// GenerateBuildVersionString returns string for CLI --version output
func GenerateBuildVersionString() string {
	versionString := "Version     " + Version + "\n" +
		"UI version  " + UIVersion + "\n" +
		"Commit      " + GitCommit + "\n" +
		"Date        " + BuildDate

	return versionString
}
