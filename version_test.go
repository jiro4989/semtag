package semtag

import (
	"testing"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/stretchr/testify/assert"
)

func TestVersions(t *testing.T) {
	r1, err := git.Init(memory.NewStorage(), memfs.New())
	assert.NoError(t, err)
	r1.CreateTag("v0.2.0", plumbing.NewHash("v0.2.0"), nil)
	r1.CreateTag("v0.3.0", plumbing.NewHash("v0.3.0"), nil)
	r1.CreateTag("v0.1.0", plumbing.NewHash("v0.1.0"), nil)

	tests := []struct {
		desc    string
		input   *VersionsInput
		want    []*Version
		wantErr bool
	}{
		{
			desc: "ok: default version",
			input: &VersionsInput{
				Repository: r1,
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			got, err := Versions(tt.input)
			if tt.wantErr {
				assert.Error(err)
				assert.Nil(got)
				return
			}

			assert.NoError(err)
			assert.Equal(tt.want, got)
		})
	}
}
