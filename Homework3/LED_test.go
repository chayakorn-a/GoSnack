package Homework3

import "testing"

func PrintError(t *testing.T, Actual, Expected [3][3]string) {
	t.Error("LED of ", 0, " is incorrect. Actual is \n",
		Actual[0][0], Actual[0][1], Actual[0][2], "\n",
		Actual[1][0], Actual[1][1], Actual[1][2], "\n",
		Actual[2][0], Actual[2][1], Actual[2][2], "\n",
		"Expected is \n",
		Expected[0][0], Expected[0][1], Expected[0][2], "\n",
		Expected[1][0], Expected[1][1], Expected[1][2], "\n",
		Expected[2][0], Expected[2][1], Expected[2][2], "\n")
}

func TestDigitZero(t *testing.T) {
	// Arrange:
	Expected := [3][3]string{{" ", "_", " "},
		{"|", " ", "|"},
		{"|", "_", "|"}}
	// Action
	Actual := CreateLedDigit(0)

	// Assertion
	if Actual != Expected {
		PrintError(t, Actual, Expected)
	}
}

func TestDigitOne(t *testing.T) {
	// Arrange:
	Expected := [3][3]string{{" ", " ", " "},
		{" ", " ", "|"},
		{" ", " ", "|"}}
	// Action
	Actual := CreateLedDigit(1)

	// Assertion
	if Actual != Expected {
		PrintError(t, Actual, Expected)
	}
}

func TestDigitTwo(t *testing.T) {
	// Arrange:
	Expected := [3][3]string{{" ", "_", " "},
		{" ", "_", "|"},
		{"|", "_", " "}}
	// Action
	Actual := CreateLedDigit(2)

	// Assertion
	if Actual != Expected {
		PrintError(t, Actual, Expected)
	}
}

func TestDigitThree(t *testing.T) {
	// Arrange:
	Expected := [3][3]string{{" ", "_", " "},
		{" ", "_", "|"},
		{" ", "_", "|"}}
	// Action
	Actual := CreateLedDigit(3)

	// Assertion
	if Actual != Expected {
		PrintError(t, Actual, Expected)
	}
}

func TestDigitFour(t *testing.T) {
	// Arrange:
	Expected := [3][3]string{{" ", " ", " "},
		{"|", "_", "|"},
		{" ", " ", "|"}}
	// Action
	Actual := CreateLedDigit(4)

	// Assertion
	if Actual != Expected {
		PrintError(t, Actual, Expected)
	}
}

func TestDigitFive(t *testing.T) {
	// Arrange:
	Expected := [3][3]string{{" ", "_", " "},
		{"|", "_", " "},
		{" ", "_", "|"}}
	// Action
	Actual := CreateLedDigit(5)

	// Assertion
	if Actual != Expected {
		PrintError(t, Actual, Expected)
	}
}

func TestDigitSix(t *testing.T) {
	// Arrange:
	Expected := [3][3]string{{" ", "_", " "},
		{"|", "_", " "},
		{"|", "_", "|"}}
	// Action
	Actual := CreateLedDigit(6)

	// Assertion
	if Actual != Expected {
		PrintError(t, Actual, Expected)
	}
}

func TestDigitSeven(t *testing.T) {
	// Arrange:
	Expected := [3][3]string{{" ", "_", " "},
		{" ", " ", "|"},
		{" ", " ", "|"}}
	// Action
	Actual := CreateLedDigit(7)

	// Assertion
	if Actual != Expected {
		PrintError(t, Actual, Expected)
	}
}

func TestDigitEight(t *testing.T) {
	// Arrange:
	Expected := [3][3]string{{" ", "_", " "},
		{"|", "_", "|"},
		{"|", "_", "|"}}
	// Action
	Actual := CreateLedDigit(8)

	// Assertion
	if Actual != Expected {
		PrintError(t, Actual, Expected)
	}
}

func TestDigitNine(t *testing.T) {
	// Arrange:
	Expected := [3][3]string{{" ", "_", " "},
		{"|", "_", "|"},
		{" ", " ", "|"}}
	// Action
	Actual := CreateLedDigit(9)

	// Assertion
	if Actual != Expected {
		PrintError(t, Actual, Expected)
	}
}

func PrintError2(t *testing.T, Actual, Expected [3][3][3]string) {
	t.Error("LED of ", 0, " is incorrect. Actual is \n",
		Actual[0][0][0], Actual[0][0][1], Actual[0][0][2], "  ", Actual[1][0][0], Actual[1][0][1], Actual[1][0][2], "  ", Actual[2][0][0], Actual[2][0][1], Actual[2][0][2], "  ", "\n",
		Actual[0][1][0], Actual[0][1][1], Actual[0][1][2], "  ", Actual[1][1][0], Actual[1][1][1], Actual[1][1][2], "  ", Actual[2][1][0], Actual[2][1][1], Actual[2][1][2], "  ", "\n",
		Actual[0][2][0], Actual[0][2][1], Actual[0][2][2], "  ", Actual[1][2][0], Actual[1][2][1], Actual[1][2][2], "  ", Actual[2][2][0], Actual[2][2][1], Actual[2][2][2], "  ", "\n",
		"Expected is \n",
		Expected[0][0][0], Expected[0][0][1], Expected[0][0][2], "  ", Expected[1][0][0], Expected[1][0][1], Expected[1][0][2], "  ", Expected[2][0][0], Expected[2][0][1], Expected[2][0][2], "  ", "\n",
		Expected[0][1][0], Expected[0][1][1], Expected[0][1][2], "  ", Expected[1][1][0], Expected[1][1][1], Expected[1][1][2], "  ", Expected[2][1][0], Expected[2][1][1], Expected[2][1][2], "  ", "\n",
		Expected[0][2][0], Expected[0][2][1], Expected[0][2][2], "  ", Expected[1][2][0], Expected[1][2][1], Expected[1][2][2], "  ", Expected[2][2][0], Expected[2][2][1], Expected[2][2][2], "  ", "\n")
}

// Put 123 then expect 123
func TestCase123(t *testing.T) {
	// Arrange:
	Expected := [3][3][3]string{{{" ", " ", " "}, //1
		{" ", " ", "|"},
		{" ", " ", "|"}},
		{{" ", "_", " "}, //2
			{" ", "_", "|"},
			{"|", "_", " "}},
		{{" ", "_", " "}, //3
			{" ", "_", "|"},
			{" ", "_", "|"}}}
	// Action
	Actual := CreateLed(123)

	// Assertion
	if Actual != Expected {
		PrintError2(t, Actual, Expected)
	}
}

// Put 5 then expect 005
func TestCase005(t *testing.T) {
	// Arrange:
	Expected := [3][3][3]string{{{" ", "_", " "},
		{"|", " ", "|"},
		{"|", "_", "|"}},
		{{" ", "_", " "},
			{"|", " ", "|"},
			{"|", "_", "|"}},
		{{" ", "_", " "},
			{"|", "_", " "},
			{" ", "_", "|"}}}
	// Action
	Actual := CreateLed(5)

	// Assertion
	if Actual != Expected {
		PrintError2(t, Actual, Expected)
	}
}

// support max 3 digits but put 4, then expect 000
func TestCaseExceed1234(t *testing.T) {
	// Arrange:
	Expected := [3][3][3]string{{{" ", "_", " "},
		{"|", " ", "|"},
		{"|", "_", "|"}},
		{{" ", "_", " "},
			{"|", " ", "|"},
			{"|", "_", "|"}},
		{{" ", "_", " "},
			{"|", " ", "|"},
			{"|", "_", "|"}}}
	// Action
	Actual := CreateLed(1234)

	// Assertion
	if Actual != Expected {
		PrintError2(t, Actual, Expected)
	}
}
