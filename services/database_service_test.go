package services

import "testing"

func TestGetPostgresDBSpecsFromURL(t *testing.T) {
	testCases := []struct {
		input, expectedUser, expectedPw, expectedHost, expectedPort, expectedDBName string
	}{
		{"protocol://user:pw@host:port/db", "user", "pw", "host", "port", "db"},
	}

	for _, tc := range testCases {
		actualHost, actualUser, actualPort, actualPw, actualDB := GetPostgresDBSpecsFromURL(tc.input)
		if actualHost != tc.expectedHost && actualUser != tc.expectedUser &&
			actualPort != tc.expectedPort && actualPw != tc.expectedPw && actualDB != tc.expectedDBName {
			t.Errorf("GetPostgresDBSpecsFromURL(%v) failed: exptected: protocol://%v:%v@%v:%v/%v, got: protocol://%v:%v@%v:%v/%v",
				tc.input, tc.expectedUser, tc.expectedPw, tc.expectedHost, tc.expectedPort, tc.expectedDBName,
				actualUser, actualPw, actualHost, actualPort, actualDB)
		}
	}
}
