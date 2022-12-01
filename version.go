package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/hashicorp/go-version"
)

type Version struct {
	prefix           string
	major            int
	minor            int
	patch            int
	separator        string
	candidateVersion string
	suffix           string
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
		prefix:           input.Prefix,
		major:            input.Major,
		minor:            input.Minor,
		patch:            input.Patch,
		separator:        input.Separator,
		candidateVersion: input.CandidateVersion,
		suffix:           input.Suffix,
	}
}

func NewDefaultVersion() *Version {
	return &Version{
		prefix:           "v",
		major:            0,
		minor:            1,
		patch:            0,
		separator:        "",
		candidateVersion: "",
		suffix:           "",
	}
}

func (v *Version) BumpMajor() {
	v.major++
	v.minor = 0
	v.patch = 0
}

func (v *Version) BumpMinor() {
	v.minor++
	v.patch = 0
}

func (v *Version) BumpPatch() {
	v.patch++
}

func (v *Version) String() string {
	var sb strings.Builder

	if v.prefix != "" {
		sb.WriteString(v.prefix)
	}

	sb.WriteString(fmt.Sprint(v.major))
	sb.WriteString(".")
	sb.WriteString(fmt.Sprint(v.minor))
	sb.WriteString(".")
	sb.WriteString(fmt.Sprint(v.patch))

	if v.candidateVersion != "" {
		sb.WriteString(v.separator)
		sb.WriteString(v.candidateVersion)
	}

	if v.suffix != "" {
		sb.WriteString(v.separator)
		sb.WriteString(v.suffix)
	}

	return sb.String()
}

type VersionsInput struct {
	Tagger Tagger
}

func Versions(input *VersionsInput) ([]*Version, error) {
	tags, err := input.Tagger.Tags(".")
	if err != nil {
		return nil, err
	}

	verMap := make(map[string]*Version)
	vers := make([]*version.Version, 0)
	for _, tag := range tags {
		pi := &ParseInput{
			Tag: tag,
		}
		ver, err := Parse(pi)
		if err != nil {
			return nil, err
		}
		v, _ := version.NewVersion(tag)
		verMap[v.String()] = ver
		vers = append(vers, v)
	}

	sort.Sort(version.Collection(vers))

	sortedVers := make([]*Version, len(vers))
	for i, v := range vers {
		sortedVers[i] = verMap[v.String()]
	}

	return sortedVers, nil
}

type LatestVersionInput struct {
	Tagger Tagger
}

func LatestVersion(input *LatestVersionInput) (*Version, error) {
	vi := &VersionsInput{
		Tagger: input.Tagger,
	}
	vers, err := Versions(vi)
	if err != nil {
		return nil, err
	}
	if len(vers) < 1 {
		return nil, errors.New("no tags")
	}
	return vers[len(vers)-1], nil
}
