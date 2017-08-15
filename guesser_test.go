package guesser

// Although it would be better separating the packages (guesser_test), given
// the simplicity of the application it is being kept together.
import (
	"testing"
)

func TestIsPrefix(t *testing.T) {

	// To account for additional test cases, simply include them in the appropriate table
	// pass and fail scenarios are split up for organization
	var passSlice = []string{"DR", "Mr", "MrS", "mS", "miss", " MR", " MISS"}
	var failSlice = []string{"SIR", "MZ", "!_-:/@ #", "1234567890", "\\", "\n", "%d"}

	for i := 0; i < len(passSlice); i++ {
		var result = isPrefix(passSlice[i])
		if result != true {
			t.Errorf("\nInput '%s' was NOT recognized as a prefix, expected true\n ", passSlice[i])
		}
	}

	for j := 0; j < len(failSlice); j++ {
		var result = isPrefix(failSlice[j])
		if result != false {
			t.Errorf("\nInput '%s' WAS recognized as a prefix, expected false\n ", failSlice[j])
		}
	}
}

func TestGuessFirstName(t *testing.T) {

	var testCases = []struct {
		input    string
		expected string
	}{
		{"", ""},
		{" ", ""},
		{"\n", ""},
		{"%d", "%d"},
		{"dr !@#$%^&*()", "!@#$%^&*()"},
		{"Mr dr", "dr"},
		{"mrdr", "mrdr"},
		{"Mrs First Last", "First"},
		{"Mr First Middle Last", "First"},
		{" Ms LeadingSpace", "LeadingSpace"},
		{"Miss  FrontTwoSpaces Lastname", "FrontTwoSpaces"},
		{"Miss BackTwoSpaces  Lastname", "BackTwoSpaces"},
		{" NoPrefixLeadingSpace", "NoPrefixLeadingSpace"},
		{"MrsNoSpace FirstName", "MrsNoSpace"},
		{"\\Mr Backslash", "\\Mr"}, // need to test the corner case of a single backslash entry via UI testing
		{"Mrs 1234567890 Numbers", "1234567890"},
		{"Miss A1ph4nuM3ric!", "A1ph4nuM3ric!"},
	}

	for _, testTable := range testCases {
		actual := GuessFirstName(testTable.input)
		if actual != testTable.expected {
			t.Errorf("\nExpected: '%s'\nReceived: '%s' \nInput:    '%s'\n ", testTable.expected, actual, testTable.input)
		}
	}
}
