package util

import "regexp"

func HtmlDetect(input string) (boolFound bool, matchStr string) {
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	boolFound = re.MatchString(input)
	if boolFound {
		matchStr = re.FindString(input)
	} else {
		matchStr = ""
	}
	return
}

func LatexDetect(input string) (boolFound bool, matchStr string) {
	re, _ := regexp.Compile("\\$(.+?)\\$")
	boolFound = re.MatchString(input)
	if boolFound {
		matchStr = re.FindString(input)
	} else {
		matchStr = ""
	}
	return
}

func MultiDetect(input string) (boolFound bool, matchStr string) {
	re1, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	re2, _ := regexp.Compile("\\$(.+?)\\$")

	boolFound = re1.MatchString(input)
	if boolFound {
		matchStr = re1.FindString(input)
		return
	}

	boolFound = re2.MatchString(input)
	if boolFound {
		matchStr = re2.FindString(input)
		return
	}

	return
}
