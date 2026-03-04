package longestsubstringwithoutrepeatingcharacters

import (
	"fmt"
	"strings"
)

/* Given a string s, find the length of the longest substring without duplicate characters.

Example 1:
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3. Note that "bca" and "cab" are also correct answers.

Example 2:
Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Example 3:
Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.

Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

Constraints:
0 <= s.length <= 5 * 104
s consists of English letters, digits, symbols and spaces.
*/

func Init() {
	result := lengthOfLongestSubstring("pwwkew")

	fmt.Printf("Result: %v", result)
}

func lengthOfLongestSubstring(s string) int {
	chars := strings.Split(s, "")
	longestSubstring := ""

	i := 0
	substring := ""
	for j := i; j < len(chars); {
		if strings.ContainsAny(substring, chars[j]) {
			// end of the substring
			if len(substring) > len(longestSubstring) {
				longestSubstring = substring
			}

			i++
			j = i
			substring = ""

			continue
		}

		//the char is not in the substring. So it is appended
		substring = fmt.Sprintf("%s%s", substring, chars[j])
		j++
	}

	if len(substring) > len(longestSubstring) {
		longestSubstring = substring
	}

	return len(longestSubstring)
}
