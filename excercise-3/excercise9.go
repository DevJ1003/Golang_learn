package main

import "fmt"

func main() {

	var word string
	fmt.Print("Enter a word : ")
	fmt.Scanf("%s", &word)

	FirstIdx := 0
	LastIdx := len(word) - 1

	IsPalindrome := true
	for FirstIdx < LastIdx {

		if word[FirstIdx] != word[LastIdx] {

			IsPalindrome = false
			break

		}

		FirstIdx++
		LastIdx--

	}

	if IsPalindrome {

		fmt.Printf("'%s' is a Palindrome!\n", word)

	} else {

		fmt.Printf("'%s' is not a Palindrome!\n", word)

	}

}
