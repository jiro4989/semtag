package semtag

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewVersion(t *testing.T) {
	tests := []struct {
		desc  string
		input *NewVersionInput
		want  Version
	}{
		{
			desc:  "ok: default version",
			input: nil,
			want: Version{
				Prefix: "v",
				Major:  0,
				Minor:  1,
				Patch:  0,
			},
		},
		{
			desc: "ok: input",
			input: &NewVersionInput{
				Prefix: "vv",
				Major:  1,
				Minor:  2,
				Patch:  3,
			},
			want: Version{
				Prefix: "vv",
				Major:  1,
				Minor:  2,
				Patch:  3,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			got := NewVersion(tt.input)
			assert.Equal(tt.want, got)
		})
	}
}

func TestVersionString(t *testing.T) {
	tests := []struct {
		desc string
		ver  Version
		want string
	}{
		{
			desc: "ok: standard version",
			ver: Version{
				Prefix: "v",
				Major:  0,
				Minor:  1,
				Patch:  0,
			},
			want: "v0.1.0",
		},
		{
			desc: "ok: standard version with separator",
			ver: Version{
				Prefix:              "v",
				PrefixWithSeparator: true,
				Major:               0,
				Minor:               1,
				Patch:               0,
				Separator:           "-",
			},
			want: "v-0.1.0",
		},
		{
			desc: "ok: standard version with suffix",
			ver: Version{
				Prefix:              "v",
				PrefixWithSeparator: true,
				Major:               0,
				Minor:               1,
				Patch:               0,
				Separator:           "-",
				Suffix1:             "module1",
				Suffix2:             "rc1",
			},
			want: "v-0.1.0-module1-rc1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			got := tt.ver.String()
			assert.Equal(tt.want, got)
		})
	}
}
