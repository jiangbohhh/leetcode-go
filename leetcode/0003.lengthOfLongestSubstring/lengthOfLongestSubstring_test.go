package lengthOfLongestSubstring

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	testcases := []struct {
		in   string
		want int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
	}
	for _, tc := range testcases {
		length := lengthOfLongestSubstring(tc.in)
		if tc.want != length {
			t.Errorf("lengthOfLongestSubstring:%d, want:%d", length, tc.want)
		}
	}
}
