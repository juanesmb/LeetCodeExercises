package palindromenumber

import (
	"fmt"
	"math"
)

/*Given an integer x, return true if x is a palindrome, and false otherwise.

Example 1:
Input: x = 121 64546
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.

Example 2:
Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.

Example 3:
Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

Constraints:
-231 <= x <= 231 - 1

Follow up: Could you solve it without converting the integer to a string?*/

func Init() {
	result := isPalindrome(64546)

	fmt.Println(result)
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x == 0 {
		return true
	}

	modulePalindrome := x

	result := math.Log10(float64(x))

	power := int(result)
	c := power / 2

	divisor := int(math.Pow(10, float64(power)))
	modulus := 10

	for i := 0; i <= c; i++ {
		leftDigit := (x / divisor) % 10
		rightDigit := modulePalindrome % modulus
		if leftDigit != rightDigit {
			return false
		}

		divisor = divisor / 10
		modulePalindrome = modulePalindrome / 10
	}

	return true
}

/*func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x == 0 {
		return true
	}

	str := strconv.Itoa(x)

	fmt.Println(str)

	i := 0
	j := len(str) - 1

	for i <= j {
		if str[i] != str[j] {
			return false
		}

		i++
		j--
	}

	return true
}*/
