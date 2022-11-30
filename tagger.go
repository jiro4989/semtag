package semtag

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

type Tagger interface {
	Tags(string) ([]string, error)
}

type GitTagger struct {
	Tagger
}

func (g *GitTagger) Tags(path string) ([]string, error) {
	r, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}
	iter, err := r.Tags()
	if err != nil {
		return nil, err
	}

	tags := make([]string, 0)
	iter.ForEach(func(ref *plumbing.Reference) error {
		tag, err := r.TagObject(ref.Hash())
		if err != nil {
			return err
		}

		tags = append(tags, tag.Name)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return tags, nil
}

func NewTagger() *GitTagger {
	return &GitTagger{}
}