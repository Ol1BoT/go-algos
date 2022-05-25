package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
// THIS REVERSE WORKS WITH BTYES, THERE ARE SOME CHARS THAT ARE BIGGER THAN A BTYE
// func Reverse(word string) string {

// 	var aggString string

// 	for i := len(word) - 1; i >= 0; i-- {
// 		aggString = fmt.Sprint(aggString, string(word[i]))
// 	}

// 	return aggString
// }

//THIS IS A BETTER SOLUTION as it works based on Runes rather than bytes
func Reverse(word string) string {
	var res string

	for _, r := range word {
		res = string(r) + res
	}

	return res

}
