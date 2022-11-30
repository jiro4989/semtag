package semtag

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersions(t *testing.T) {
	tests := []struct {
		desc    string
		input   *VersionsInput
		want    []*Version
		wantErr bool
	}{
		{
			desc: "ok: default version",
			input: &VersionsInput{
				Tagger: &MockTagger{
					tags: []string{"v0.2.0", "v0.1.0", "v0.3.0"},
					err:  nil,
				},
			},
			want: []*Version{
				{
					Prefix: "v",
					Major:  0,
					Minor:  1,
					Patch:  0,
				},
				{
					Prefix: "v",
					Major:  0,
					Minor:  2,
					Patch:  0,
				},
				{
					Prefix: "v",
					Major:  0,
					Minor:  3,
					Patch:  0,
				},
			},
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
