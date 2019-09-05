package Homework3

// make LED for only single digit number
func CreateLedDigit(num int) [3][3]string {
	switch {
	case num == 0:
		return [3][3]string{
			{" ", "_", " "},
			{"|", " ", "|"},
			{"|", "_", "|"},
		}
	case num == 1:
		return [3][3]string{
			{" ", " ", " "},
			{" ", " ", "|"},
			{" ", " ", "|"},
		}
	case num == 2:
		return [3][3]string{
			{" ", "_", " "},
			{" ", "_", "|"},
			{"|", "_", " "},
		}
	case num == 3:
		return [3][3]string{
			{" ", "_", " "},
			{" ", "_", "|"},
			{" ", "_", "|"},
		}
	case num == 4:
		return [3][3]string{
			{" ", " ", " "},
			{"|", "_", "|"},
			{" ", " ", "|"},
		}
	case num == 5:
		return [3][3]string{
			{" ", "_", " "},
			{"|", "_", " "},
			{" ", "_", "|"},
		}
	case num == 6:
		return [3][3]string{
			{" ", "_", " "},
			{"|", "_", " "},
			{"|", "_", "|"},
		}
	case num == 7:
		return [3][3]string{
			{" ", "_", " "},
			{" ", " ", "|"},
			{" ", " ", "|"},
		}
	case num == 8:
		return [3][3]string{
			{" ", "_", " "},
			{"|", "_", "|"},
			{"|", "_", "|"},
		}
	case num == 9:
		return [3][3]string{
			{" ", "_", " "},
			{"|", "_", "|"},
			{" ", " ", "|"},
		}
	default:
		return [3][3]string{}
	}
}

// make LED support max 3 digits number
// if exceed, print out 000
func CreateLed(num int) [3][3][3]string {
	var hundred, ten, digit int
	ret := [3][3][3]string{}

	if num > 999 {
		hundred = 0
		ten = 0
		digit = 0
	} else {
		hundred = num / 100
		ten = (num % 100) / 10
		digit = num % 10
	}

	ret[0] = CreateLedDigit(hundred)
	ret[1] = CreateLedDigit(ten)
	ret[2] = CreateLedDigit(digit)

	return ret
}
