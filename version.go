package semtag

import (
	"sort"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/hashicorp/go-version"
)

type VersionsInput struct {
	Repository *git.Repository
}

func Versions(input *VersionsInput) ([]*Version, error) {
	// r, err := git.PlainOpen(".")
	// if err != nil {
	// 	return nil, err
	// }
	r := input.Repository
	iter, err := r.Tags()
	if err != nil {
		return nil, err
	}

	verMap := make(map[string]*Version)
	vers := make([]*version.Version, 0)
	iter.ForEach(func(ref *plumbing.Reference) error {
		tag, err := r.TagObject(ref.Hash())
		if err != nil {
			return err
		}

		pi := &ParseInput{
			Tag: tag.Name,
		}
		ver, err := Parse(pi)
		if err != nil {
			return err
		}
		name := tag.Name
		verMap[name] = ver
		v, _ := version.NewVersion(name)
		vers = append(vers, v)

		return nil
	})
	switch err {
	case nil:
		// Nothing to do
	case plumbing.ErrObjectNotFound:
		return nil, nil
	default:
		return nil, err
	}

	sort.Sort(version.Collection(vers))

	sortedVers := make([]*Version, len(vers))
	for i, v := range vers {
		sortedVers[i] = verMap[v.String()]
	}

	return sortedVers, nil
}

func LatestVersion() (*Version, error) {
	return nil, nil
}
