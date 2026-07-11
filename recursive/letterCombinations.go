// https://leetcode.com/problems/letter-combinations-of-a-phone-number/
package main

// 用数组映射，索引为数字字符-'0'
var digitMap = [10]string{
	"",    // 0
	"",    // 1
	"abc", // 2
	"def", // 3
	"ghi", // 4
	"jkl", // 5
	"mno", // 6
	"pqrs",// 7
	"tuv", // 8
	"wxyz",// 9
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil // 如果更符合你场景，也可返回 []string{}
	}

	// 预估容量，减少扩容
	total := 1
	lengths := make([]int, 0, len(digits))
	for i := 0; i < len(digits); i++ {
		d := digits[i] - '0'
		if d > 9 { // 非数字字符
			continue
		}
		ls := len(digitMap[d])
		if ls == 0 { // 跳过 0/1 或者无效
			continue
		}
		lengths = append(lengths, ls)
		total *= ls
	}
	if total == 1 && len(lengths) == 0 {
		return nil // 全是无效字符
	}

	// 迭代构造
	result := make([]string, 1, total)
	result[0] = ""

	for i := 0; i < len(digits); i++ {
		d := digits[i] - '0'
		if d > 9 || len(digitMap[d]) == 0 {
			continue // 跳过 0/1 或无效
		}
		letters := digitMap[d]
		// 依据当前 letters 数量预分配下一层容量
		next := make([]string, 0, len(result)*len(letters))
		for _, prefix := range result {
			for j := 0; j < len(letters); j++ {
				// letters[j:j+1] 避免中间 string 转换
				next = append(next, prefix+letters[j:j+1])
			}
		}
		result = next
	}

	return result
}
