package semtag

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewVersion(t *testing.T) {
	tests := []struct {
		desc  string
		input *NewVersionInput
		want  *Version
	}{
		{
			desc:  "ok: default version",
			input: nil,
			want: &Version{
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
			want: &Version{
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

func TestVersionBumpMajor(t *testing.T) {
	tests := []struct {
		desc string
		ver  Version
		want Version
	}{
		{
			desc: "ok: bump",
			ver: Version{
				Prefix: "v",
				Major:  0,
				Minor:  1,
				Patch:  2,
			},
			want: Version{
				Prefix: "v",
				Major:  1,
				Minor:  0,
				Patch:  0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			tt.ver.BumpMajor()
			assert.Equal(tt.want, tt.ver)
		})
	}
}

func TestVersionBumpMinor(t *testing.T) {
	tests := []struct {
		desc string
		ver  Version
		want Version
	}{
		{
			desc: "ok: bump",
			ver: Version{
				Prefix: "v",
				Major:  0,
				Minor:  1,
				Patch:  2,
			},
			want: Version{
				Prefix: "v",
				Major:  0,
				Minor:  2,
				Patch:  0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			tt.ver.BumpMinor()
			assert.Equal(tt.want, tt.ver)
		})
	}
}

func TestVersionBumpPatch(t *testing.T) {
	tests := []struct {
		desc string
		ver  Version
		want Version
	}{
		{
			desc: "ok: bump",
			ver: Version{
				Prefix: "v",
				Major:  0,
				Minor:  1,
				Patch:  2,
			},
			want: Version{
				Prefix: "v",
				Major:  0,
				Minor:  1,
				Patch:  3,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			tt.ver.BumpPatch()
			assert.Equal(tt.want, tt.ver)
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
				CandidateVersion:    "rc.1",
				Suffix:              "module1",
			},
			want: "v-0.1.0-rc.1-module1",
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
				Prefix:              "v",
				PrefixWithSeparator: true,
				Major:               1,
				Minor:               2,
				Patch:               3,
				Separator:           "-",
				CandidateVersion:    "rc.1",
				Suffix:              "test",
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
