package Homework3

import "testing"

var LedList = [10][3][3]string{
	{ // 0
		{" ", "_", " "},
		{"|", " ", "|"},
		{"|", "_", "|"},
	},
	{ // 1
		{" ", " ", " "},
		{" ", " ", "|"},
		{" ", " ", "|"},
	},
	{ // 2
		{" ", "_", " "},
		{" ", "_", "|"},
		{"|", "_", " "},
	},
	{ // 3
		{" ", "_", " "},
		{" ", "_", "|"},
		{" ", "_", "|"},
	},
	{ // 4
		{" ", " ", " "},
		{"|", "_", "|"},
		{" ", " ", "|"},
	},
	{ // 5
		{" ", "_", " "},
		{"|", "_", " "},
		{" ", "_", "|"},
	},
	{ // 6
		{" ", "_", " "},
		{"|", "_", " "},
		{"|", "_", "|"},
	},
	{ // 7
		{" ", "_", " "},
		{" ", " ", "|"},
		{" ", " ", "|"},
	},
	{ // 8
		{" ", "_", " "},
		{"|", "_", "|"},
		{"|", "_", "|"},
	},
	{ // 9
		{" ", "_", " "},
		{"|", "_", "|"},
		{" ", " ", "|"},
	},
}

func TestPrintZerotoNine(t *testing.T) {
	// Arrange: need grid 3*3
	var Expected [3][3]string

	for i := 0; i < 10; i++ {
		Expected = LedList[i]
		// Action
		Actual := CreateLedDigit(i)

		// Assertion
		if Expected != Actual {
			t.Error("LED of ", i, " is incorrect. Actual is \n",
				Actual[0][0], Actual[0][1], Actual[0][2], "\n",
				Actual[1][0], Actual[1][1], Actual[1][2], "\n",
				Actual[2][0], Actual[2][1], Actual[2][2], "\n",
				"Expected is \n",
				Expected[0][0], Expected[0][1], Expected[0][2], "\n",
				Expected[1][0], Expected[1][1], Expected[1][2], "\n",
				Expected[2][0], Expected[2][1], Expected[2][2], "\n")
		}
	}
}
