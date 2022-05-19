package match_test

import (
	"testing"

	"github.com/amirography/rose/match"
)

func TestMatch(t *testing.T) {
	t.Parallel()
	type testCase struct {
		fn               func(string, string) string
		name, swatchName string
		want             string
	}
	cases := []testCase{
		{fn: match.Get, name: "rose", swatchName: "rp", want: "#ebbcba"},
		{fn: match.Get, name: "rse", swatchName: "rp", want: ""},
	}

	for _, tc := range cases {
		got := tc.fn(tc.name, tc.swatchName)

		if tc.want != got {
			t.Errorf("want %s, got %s", tc.want, got)
		}
	}
}
