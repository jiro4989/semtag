package main

import (
	"errors"
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
				prefix: "v",
				major:  0,
				minor:  1,
				patch:  0,
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
				prefix: "vv",
				major:  1,
				minor:  2,
				patch:  3,
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
				prefix: "v",
				major:  0,
				minor:  1,
				patch:  2,
			},
			want: Version{
				prefix: "v",
				major:  1,
				minor:  0,
				patch:  0,
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
				prefix: "v",
				major:  0,
				minor:  1,
				patch:  2,
			},
			want: Version{
				prefix: "v",
				major:  0,
				minor:  2,
				patch:  0,
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
				prefix: "v",
				major:  0,
				minor:  1,
				patch:  2,
			},
			want: Version{
				prefix: "v",
				major:  0,
				minor:  1,
				patch:  3,
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
				prefix: "v",
				major:  0,
				minor:  1,
				patch:  0,
			},
			want: "v0.1.0",
		},
		{
			desc: "ok: standard version with separator",
			ver: Version{
				prefix:    "v-",
				major:     0,
				minor:     1,
				patch:     0,
				separator: "",
			},
			want: "v-0.1.0",
		},
		{
			desc: "ok: standard version with suffix",
			ver: Version{
				prefix:           "v",
				major:            0,
				minor:            1,
				patch:            0,
				separator:        "-",
				candidateVersion: "rc.1",
				suffix:           "module1",
			},
			want: "v0.1.0-rc.1-module1",
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

func TestVersions(t *testing.T) {
	tests := []struct {
		desc    string
		input   *VersionsInput
		want    []*Version
		wantErr bool
	}{
		{
			desc: "ok: simpel tags",
			input: &VersionsInput{
				Tagger: &MockTagger{
					tags: []string{"v0.2.0", "v0.1.0", "v0.3.0"},
					err:  nil,
				},
			},
			want: []*Version{
				{
					prefix: "v",
					major:  0,
					minor:  1,
					patch:  0,
				},
				{
					prefix: "v",
					major:  0,
					minor:  2,
					patch:  0,
				},
				{
					prefix: "v",
					major:  0,
					minor:  3,
					patch:  0,
				},
			},
		},
		{
			desc: "ng: tagger returns error",
			input: &VersionsInput{
				Tagger: &MockTagger{
					tags: []string{"v0.2.0", "v0.1.0", "v0.3.0"},
					err:  errors.New("error"),
				},
			},
			want:    nil,
			wantErr: true,
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

func TestLatestVersion(t *testing.T) {
	tests := []struct {
		desc    string
		input   *LatestVersionInput
		want    *Version
		wantErr bool
	}{
		{
			desc: "ok: simpel tags",
			input: &LatestVersionInput{
				Tagger: &MockTagger{
					tags: []string{"v0.3.0", "v0.1.0", "v0.2.0"},
					err:  nil,
				},
			},
			want: &Version{
				prefix: "v",
				major:  0,
				minor:  3,
				patch:  0,
			},
		},
		{
			desc: "ng: tagger returns error",
			input: &LatestVersionInput{
				Tagger: &MockTagger{
					tags: []string{"v0.2.0", "v0.1.0", "v0.3.0"},
					err:  errors.New("error"),
				},
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			got, err := LatestVersion(tt.input)
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
