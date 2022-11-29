package semtag

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Version struct {
	Prefix              string
	PrefixWithSeparator bool
	Major               int
	Minor               int
	Patch               int
	Separator           string
	Suffix1             string
	Suffix2             string
}

type NewVersionInput struct {
	Prefix              string
	PrefixWithSeparator bool
	Major               int
	Minor               int
	Patch               int
	Separator           string
	Suffix1             string
	Suffix2             string
}

func NewVersion(input *NewVersionInput) *Version {
	if input == nil {
		return NewDefaultVersion()
	}

	return &Version{
		Prefix:              input.Prefix,
		PrefixWithSeparator: input.PrefixWithSeparator,
		Major:               input.Major,
		Minor:               input.Minor,
		Patch:               input.Patch,
		Separator:           input.Separator,
		Suffix1:             input.Suffix1,
		Suffix2:             input.Suffix2,
	}
}

func NewDefaultVersion() *Version {
	return &Version{
		Prefix:              "v",
		PrefixWithSeparator: false,
		Major:               0,
		Minor:               1,
		Patch:               0,
		Separator:           "",
		Suffix1:             "",
		Suffix2:             "",
	}
}

func (v *Version) BumpMajor() {
	v.Major++
	v.Minor = 0
	v.Patch = 0
}

func (v *Version) BumpMinor() {
	v.Minor++
	v.Patch = 0
}

func (v *Version) BumpPatch() {
	v.Patch++
}

func (v *Version) String() string {
	var sb strings.Builder

	if v.Prefix != "" {
		sb.WriteString(v.Prefix)
		if v.PrefixWithSeparator {
			sb.WriteString(v.Separator)
		}
	}

	sb.WriteString(fmt.Sprint(v.Major))
	sb.WriteString(".")
	sb.WriteString(fmt.Sprint(v.Minor))
	sb.WriteString(".")
	sb.WriteString(fmt.Sprint(v.Patch))

	if v.Suffix1 != "" {
		sb.WriteString(v.Separator)
		sb.WriteString(v.Suffix1)
	}

	if v.Suffix2 != "" {
		sb.WriteString(v.Separator)
		sb.WriteString(v.Suffix2)
	}
	return sb.String()
}

type ParseInput struct {
	Tag string
}

func Parse(input *ParseInput) (*Version, error) {
	var v Version

	var buf strings.Builder
	r := strings.NewReader(input.Tag)
	for {
		ch, _, err := r.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if '0' <= ch && ch <= '9' {
			r.UnreadRune()
			break
		}
		buf.WriteRune(ch)
	}
	v.Prefix = buf.String()

	n, err := readNumber(r)
	if err != nil {
		return nil, err
	}
	v.Major = n

	n, err = readNumber(r)
	if err != nil {
		return nil, err
	}
	v.Minor = n

	n, err = readNumber(r)
	if err != nil {
		return nil, err
	}
	v.Patch = n
	r.UnreadRune()

	ch, _, err := r.ReadRune()
	if err == io.EOF {
		return &v, nil
	}
	if err != nil {
		return nil, err
	}
	v.Separator = string(ch)
	if strings.HasSuffix(v.Prefix, v.Separator) {
		v.Prefix = v.Prefix[:len(v.Prefix)-1]
		v.PrefixWithSeparator = true
	}

	s, err := readPart(r, v.Separator)
	if err == io.EOF {
		return &v, nil
	}
	if err != nil {
		return nil, err
	}
	v.Suffix1 = s

	s, err = readPart(r, v.Separator)
	if err == io.EOF {
		return &v, nil
	}
	if err != nil {
		return nil, err
	}
	v.Suffix2 = s

	return &v, nil
}

func readNumber(r *strings.Reader) (int, error) {
	var buf strings.Builder
	for {
		ch, _, err := r.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, err
		}

		if ch < '0' || '9' < ch {
			break
		}
		buf.WriteRune(ch)
	}
	n, err := strconv.Atoi(buf.String())
	if err != nil {
		return 0, err
	}
	return n, nil
}

func readPart(r *strings.Reader, sep string) (string, error) {
	var buf strings.Builder
	for {
		ch, _, err := r.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}

		if string(ch) == sep {
			break
		}
		buf.WriteRune(ch)
	}
	return buf.String(), nil
}
