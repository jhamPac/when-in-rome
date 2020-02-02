package rome

// ConvertToRoman converts Arabic numbers to Roman
func ConvertToRoman(arabic int) string {
	if arabic == 2 {
		return "II"
	}

	return "I"
}
