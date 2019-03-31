package version

import (
	"github.com/coreos/go-semver/semver"
)

var (
	VersionMajor int64 = 1
	VersionMinor int64 = 2
	VersionPatch int64 = 1
	VersionPre   string
	VersionDev   string
)

var Version = semver.Version{
	Major:      VersionMajor,
	Minor:      VersionMinor,
	Patch:      VersionPatch,
	PreRelease: semver.PreRelease(VersionPre),
	Metadata:   VersionDev,
}
