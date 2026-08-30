package translation_test

import (
	"testing"

	"hello-api/translation"
)

func TestTranslate(t *testing.T) {
	// Arrange
	tt := []struct {
		Word        string
		Language    string
		Translation string
	}{
		{
			Word:        "hello",
			Language:    "German",
			Translation: "hallo",
		},
		{
			Word:        "Hello",
			Language:    "german",
			Translation: "hallo",
		},
		{
			Word:        "hello ",
			Language:    "german",
			Translation: "hallo",
		},
	}

	for _, tc := range tt {
		// Act
		res := translation.Translate(tc.Word, tc.Language)

		// Assert
		if res != tc.Translation {
			t.Errorf(`expected %q but received %q`, tc.Translation, res)
		}
	}
}
