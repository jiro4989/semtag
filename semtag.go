package semtag

import (
	"fmt"
	"strings"
)

type Version struct {
	Prefix           string
	Major            int
	Minor            int
	Patch            int
	Separator        string
	CandidateVersion string
	Suffix           string
}

type NewVersionInput struct {
	Prefix           string
	Major            int
	Minor            int
	Patch            int
	Separator        string
	CandidateVersion string
	Suffix           string
}

func NewVersion(input *NewVersionInput) *Version {
	if input == nil {
		return NewDefaultVersion()
	}

	return &Version{
		Prefix:           input.Prefix,
		Major:            input.Major,
		Minor:            input.Minor,
		Patch:            input.Patch,
		Separator:        input.Separator,
		CandidateVersion: input.CandidateVersion,
		Suffix:           input.Suffix,
	}
}

func NewDefaultVersion() *Version {
	return &Version{
		Prefix:           "v",
		Major:            0,
		Minor:            1,
		Patch:            0,
		Separator:        "",
		CandidateVersion: "",
		Suffix:           "",
	}
}

func (v *Version) BumpMajor() {
	v.Major++
	v.Minor = 0
	v.Patch = 0
}

func (v *Version) BumpMinor() {
	v.Minor++
	v.Patch = 0
}

func (v *Version) BumpPatch() {
	v.Patch++
}

func (v *Version) String() string {
	var sb strings.Builder

	if v.Prefix != "" {
		sb.WriteString(v.Prefix)
	}

	sb.WriteString(fmt.Sprint(v.Major))
	sb.WriteString(".")
	sb.WriteString(fmt.Sprint(v.Minor))
	sb.WriteString(".")
	sb.WriteString(fmt.Sprint(v.Patch))

	if v.CandidateVersion != "" {
		sb.WriteString(v.Separator)
		sb.WriteString(v.CandidateVersion)
	}

	if v.Suffix != "" {
		sb.WriteString(v.Separator)
		sb.WriteString(v.Suffix)
	}

	return sb.String()
}
