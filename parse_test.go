package semtag

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		desc    string
		input   *ParseInput
		want    *Version
		wantErr bool
	}{
		{
			desc: "ok: standard version",
			input: &ParseInput{
				Tag: "v1.2.3",
			},
			want: &Version{
				Prefix: "v",
				Major:  1,
				Minor:  2,
				Patch:  3,
			},
			wantErr: false,
		},
		{
			desc: "ok: 2 suffix",
			input: &ParseInput{
				Tag: "v1.2.3-rc.1-test",
			},
			want: &Version{
				Prefix:           "v",
				Major:            1,
				Minor:            2,
				Patch:            3,
				Separator:        "-",
				CandidateVersion: "rc.1",
				Suffix:           "test",
			},
			wantErr: false,
		},
		{
			desc: "ok: prefix separator",
			input: &ParseInput{
				Tag: "v-1.2.3-rc.1-test",
			},
			want: &Version{
				Prefix:           "v-",
				Major:            1,
				Minor:            2,
				Patch:            3,
				Separator:        "-",
				CandidateVersion: "rc.1",
				Suffix:           "test",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			got, err := Parse(tt.input)
			if tt.wantErr {
				assert.Error(err)
				return
			}

			assert.NoError(err)
			assert.Equal(tt.want, got)
		})
	}
}
