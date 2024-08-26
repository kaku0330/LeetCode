package greatestcommondivisorofstrings

func GcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

	num := GCP(len(str1), len(str2))

	return str2[:num]
}

func GCP(str1 int, str2 int) int {
	for str2 != 0 {
		str1, str2 = str2, str1%str2
	}

	return str1
}
