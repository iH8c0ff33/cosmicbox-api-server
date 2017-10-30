package version

import (
	"github.com/coreos/go-semver/semver"
)

var (
	VersionMajor int64
	VersionMinor int64 = 1
	VersionPatch int64 = 0
	VersionPre string
	VersionDev string
)

var Version = semver.Version{
	Major: VersionMajor,
	Minor: VersionMinor,
	Patch: VersionPatch,
	PreRelease: semver.PreRelease(VersionPre),
	Metadata: VersionDev,
}