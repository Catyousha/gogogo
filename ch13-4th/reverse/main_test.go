package main

import (
	"testing"
	"unicode/utf8"
)

func FuzzR1(f *testing.F) {
	testCases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testCases {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, orig string) {
		rev, err := R1(orig)
		if err != nil {
			return
		}

		doubleRev, err := R1(rev)
		if err != nil {
			return
		}

		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}

		if utf8.ValidString(orig) && !utf8.ValidString(string(rev)) {
			t.Errorf("Reverse: invalid UTF-8 string %q", rev)
		}
	})
}

func FuzzR2(f *testing.F) {
	testCases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testCases {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, orig string) {
		rev, err := R2(orig)
		if err != nil {
			return
		}
		doubleRev, err := R2(string(rev))
		if err != nil {
			return
		}

		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}

		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse: invalid UTF-8 string %q", rev)
		}
	})
}
