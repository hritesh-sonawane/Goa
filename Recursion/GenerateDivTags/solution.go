package main

// O((2n)!/((n!((n+1)!)))) time | O((2n)!/((n!((n+1)!)))) space wher n: input no.
func GenerateDivTags(numberOfTags int) []string {
	matchedDivTags := []string{}
	generatedDivTagsFromPrefix(numberOfTags, numberOfTags, "", &matchedDivTags)
	return matchedDivTags
}

func generatedDivTagsFromPrefix(openingTagsNeeded int, closingTagsNeeded int, prefix string, result *[]string) {
	if openingTagsNeeded > 0 {
		newPrefix := prefix + "<div>"
		generatedDivTagsFromPrefix(openingTagsNeeded-1, closingTagsNeeded, newPrefix, result)
	}

	if openingTagsNeeded < closingTagsNeeded {
		newPrefix := prefix + "</div>"
		generatedDivTagsFromPrefix(openingTagsNeeded, closingTagsNeeded-1, newPrefix, result)
	}

	if closingTagsNeeded == 0 {
		*result = append(*result, prefix)
	}
}
