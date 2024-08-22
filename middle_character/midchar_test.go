package middle_character

import "testing"

func TestGetMiddle(t *testing.T) {
	testStrings := []struct {
		TestString string
		Want       string
	}{
		{TestString: "test", Want: "es"},
		{TestString: "testing", Want: "t"},
		{TestString: "testings", Want: "ti"},
	}

	for _, testStruct := range testStrings {
		t.Run(testStruct.TestString, func(t *testing.T) {
			got := GetMiddle(testStruct.TestString)
			want := testStruct.Want
			if got != want {
				t.Errorf("have got %s, but want %s", got, want)
			}
		})

	}
}
