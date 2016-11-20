package recursion

func StringPermutations(str string) []string {
	permutations := []string{}
	if len(str) == 1 {
		permutations = append(permutations, str)
	} else {
		for i := 0; i < len(str); i++ {
			character := string(str[i])
			var rest string
			rest = str[0:i]
			rest = rest + str[i+1:]
			subPermutations := StringPermutations(rest)
			for _, subPermutation := range subPermutations {
				permutations = append(permutations, character + subPermutation)
			}
		}
	}
	return permutations;
}
