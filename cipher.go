package main

import "fmt"

func main() {
	fmt.Println(ceazerCipher("middle-Outz", 2))
	fmt.Println(ceazerCipher("A friend in need is a friend indeed", 20))
	fmt.Println(ceazerCipher("Always-Look-on-the-Bright-Side-of-Life", 5))
	fmt.Println(ceazerCipher("A Fool and His Money Are Soon Parted.", 27))
	fmt.Println(ceazerCipher("One should not worry over things that have already happened and that cannot be changed.", 49))
}

/*
Julius Caesar protected his confidential information by encrypting it using a cipher.
Caesar's cipher shifts each letter by a number of letters.
If the shift takes you past the end of the alphabet, just rotate back to the front of the alphabet.
In the case of a rotation by 3, w, x, y and z would map to z, a, b and c.

Create a function that takes a string s (text to be encrypted) and an integer k (the rotation factor).
It should return an encrypted string.
*/
func ceazerCipher(s string, n int) string {
	var result string
	if n > 25 {
		n = n % 26
	}
	for _, value := range s {
		if value == 32 || value == 45 || value == 46 {
			result += string(value)
			continue
		}
		if value > 64 && value < 91 {
			if value+int32(n) > 64 && value+int32(n) < 91 {
				result += string(value + int32(n))
			} else {
				result += string(value + int32(n) - int32(26))
			}
		} else {
			if value+int32(n) <= 122 {
				result += string(value + int32(n))
			} else {
				result += string(value + int32(n) - int32(26))
			}
		}
	}
	return result
}
