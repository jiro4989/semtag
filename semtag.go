package semtag

import (
	"fmt"
	"strings"
)

type Version struct {
	Prefix              string
	PrefixWithSeparator bool
	Major               int
	Minor               int
	Patch               int
	Separator           string
	Suffix1             string
	Suffix2             string
}

type NewVersionInput struct {
	Prefix              string
	PrefixWithSeparator bool
	Major               int
	Minor               int
	Patch               int
	Separator           string
	Suffix1             string
	Suffix2             string
}

func NewVersion(input *NewVersionInput) Version {
	if input == nil {
		return NewDefaultVersion()
	}

	return Version{
		Prefix:              input.Prefix,
		PrefixWithSeparator: input.PrefixWithSeparator,
		Major:               input.Major,
		Minor:               input.Minor,
		Patch:               input.Patch,
		Separator:           input.Separator,
		Suffix1:             input.Suffix1,
		Suffix2:             input.Suffix2,
	}
}

func NewDefaultVersion() Version {
	return Version{
		Prefix:              "v",
		PrefixWithSeparator: false,
		Major:               0,
		Minor:               1,
		Patch:               0,
		Separator:           "",
		Suffix1:             "",
		Suffix2:             "",
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
		if v.PrefixWithSeparator {
			sb.WriteString(v.Separator)
		}
	}

	sb.WriteString(fmt.Sprint(v.Major))
	sb.WriteString(".")
	sb.WriteString(fmt.Sprint(v.Minor))
	sb.WriteString(".")
	sb.WriteString(fmt.Sprint(v.Patch))

	if v.Suffix1 != "" {
		sb.WriteString(v.Separator)
		sb.WriteString(v.Suffix1)
	}

	if v.Suffix2 != "" {
		sb.WriteString(v.Separator)
		sb.WriteString(v.Suffix2)
	}
	return sb.String()
}
