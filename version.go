package semtag

import (
	"sort"

	"github.com/hashicorp/go-version"
)

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
	return vers[len(vers)-1], nil
}
