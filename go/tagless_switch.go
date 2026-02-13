// Grade Calculator: Use a "tagless" switch (no variable after the keyword)
// to check a score variable. If $>90$ print "A", $>80$ "B", etc.

package main

func getGrade(score int) string {
	// Note: No variable after the word 'switch'
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	case score >= 60:
		return "D"
	default:
		return "F"
	}
}
