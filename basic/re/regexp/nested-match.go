package main

import (
	"fmt"
	"regexp"
	"strings"
)

// 复杂模式匹配通常涉及到嵌套结构或多条件组合的匹配。在Go语言中，由于`regexp`包的一些限制，某些复杂模式可能需要采用更创造性的方法来实现。以下是一个复杂模式匹配的示例：
//
// 示例：匹配嵌套的括号
// 假设我们想要匹配像`(abc(def(ghi)jkl)mno)`这样嵌套的括号结构。
//
// 由于Go的`regexp`包不支持递归匹配，我们不能直接用一个正则表达式来实现这一点。但我们可以采用分步骤的方法来处理这种复杂模式。
// 首先，可以使用一个简单的正则表达式来匹配最内层的括号内容，然后逐层向外处理。
// 在这个示例中，我们定义了一个`matchNestedParentheses`函数，它接受一个字符串并返回所有匹配的嵌套括号。我们使用了`regexp.Compile`来编译一个匹配最内层括号的正则表达式，并在循环中逐步移除已匹配的内层括号，直到没有更多匹配为止。
//
// 虽然这种方法无法一次性匹配所有嵌套层级，但它提供了一种处理此类复杂模式的有效方式。通过这样的迭代方法，我们能够处理那些在Go的`regexp`包当前能力范围之外的复杂匹配情况。

func matchNestedParentheses(input string) []string {
	re, _ := regexp.Compile(`\([^()]*\)`)
	var matches []string

	for {
		match := re.FindString(input)
		if match == "" {
			break
		}

		matches = append(matches, strings.ReplaceAll(strings.ReplaceAll(match, "@[", "("), "]@", ")"))
		input = re.ReplaceAllString(input, strings.Join([]string{"@[", match[1 : len(match)-1], "]@"}, ""))
	}

	return matches
}

func main() {
	nested := "(abc(def(ghi)jkl)mno)"
	matches := matchNestedParentheses(nested)
	fmt.Println(matches) // 输出: ["(ghi)", "(def(ghi)jkl)", "(abc(def(ghi)jkl)mno)"]

	re, _ := regexp.Compile(`\([^()]*\)`)
	normalMatches := re.FindAllString(nested, -1)
	fmt.Println(normalMatches) // 输出: ["(ghi)"]
}
