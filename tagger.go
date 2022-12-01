package main

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

type Tagger interface {
	Tags() ([]string, error)
	CreateTag(string) error
	PushTag(string) error
}

type GitTagger struct {
	Tagger
	repo *git.Repository
}

func NewTagger(path string) (*GitTagger, error) {
	r, err := git.PlainOpen(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open repository: %w", err)
	}
	t := &GitTagger{
		repo: r,
	}
	return t, nil
}

func (g *GitTagger) Tags() ([]string, error) {
	iter, err := g.repo.Tags()
	if err != nil {
		return nil, err
	}

	tags := make([]string, 0)
	iter.ForEach(func(ref *plumbing.Reference) error {
		tag, err := g.repo.TagObject(ref.Hash())
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

func (g *GitTagger) CreateTag(tag string) error {
	h, err := g.repo.Head()
	if err != nil {
		return fmt.Errorf("failed to get HEAD: %w", err)
	}
	hash := h.Hash()
	_, err = g.repo.CreateTag(tag, hash, nil)
	if err != nil {
		return fmt.Errorf("failed to create tag: %w", err)
	}
	return nil
}

func (g *GitTagger) PushTag(tag string) error {
	return nil
}
