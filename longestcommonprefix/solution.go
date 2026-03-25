package longestcommonprefix

import (
	"fmt"
)

/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:
Input: strs = ["flower","flow","flight"]
Output: "fl"

Example 2:
Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.


Constraints:
1 <= strs.length <= 200
0 <= strs[i].length <= 200

strs[i] consists of only lowercase English letters if it is non-empty.
*/

func Init() {
	input := []string{"ab", "a"}

	result := longestCommonPrefix(input)

	fmt.Println(result)
}

func longestCommonPrefix(strs []string) string {
	var longestPrefix = ""

	for i, char := range strs[0] {
		for _, word := range strs {
			if len(word) > i {
				if string(word[i]) != string(char) {
					return longestPrefix
				}
			} else {
				return longestPrefix
			}
		}

		longestPrefix = fmt.Sprintf("%s%s", longestPrefix, string(char))
	}

	return longestPrefix
}
