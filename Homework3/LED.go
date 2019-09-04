package Homework3

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
