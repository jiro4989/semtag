package main

import (
	"io"
	"strconv"
	"strings"
)

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
	v.prefix = buf.String()

	n, err := readNumber(r)
	if err != nil {
		return nil, err
	}
	v.major = n

	n, err = readNumber(r)
	if err != nil {
		return nil, err
	}
	v.minor = n

	n, err = readNumber(r)
	if err != nil {
		return nil, err
	}
	v.patch = n
	r.UnreadRune()

	ch, _, err := r.ReadRune()
	if err == io.EOF {
		return &v, nil
	}
	if err != nil {
		return nil, err
	}
	v.separator = string(ch)

	s, err := readPart(r, v.separator)
	if err == io.EOF {
		return &v, nil
	}
	if err != nil {
		return nil, err
	}
	v.candidateVersion = s

	s, err = readPart(r, v.separator)
	if err == io.EOF {
		return &v, nil
	}
	if err != nil {
		return nil, err
	}
	v.suffix = s

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
