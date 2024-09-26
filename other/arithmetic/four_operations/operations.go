package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Calculate(expression string, numbers ...any) (float64, error) {

	numberStrings := make([]string, len(numbers))

	for i, v := range numbers {
		val, err := ToStringErr(v)
		if err != nil {
			return 0, err
		}
		numberStrings[i] = val
	}

	expressions := ExpressionTidy(expression, numberStrings...)
	expression = strings.Join(expressions, "")

	//	将中缀表达式转换为后缀表达式
	postfixExpression, err := toPostfixExpression(expression)
	if err != nil {
		return 0, err
	}

	//	计算后缀表达式
	result, err := calculatePostfixExpression(postfixExpression)
	return result, err
}

type Stack struct {
	top  uint
	data [1024]string
}

func (s *Stack) IsEmpty() bool {
	return s.top == 0
}
func (s *Stack) Push(data string) {
	s.data[s.top] = data
	s.top++
}
func (s *Stack) Pop() (string, error) {
	if s.top == 0 {
		return "", fmt.Errorf("Stack is empty")
	}
	s.top--
	return s.data[s.top], nil
}
func (s *Stack) Top() (string, error) {
	if s.top == 0 {
		return "", fmt.Errorf("Stack is empty")
	}
	return s.data[s.top-1], nil
}

func ExpressionTidy(expression string, numbers ...string) []string {
	for i, v := range numbers {
		expression = strings.ReplaceAll(expression, "$v"+strconv.Itoa(i+1), v)
	}
	str1 := strings.ReplaceAll(expression, " ", "")
	str2 := strings.ReplaceAll(str1, "*", "@*@")
	str3 := strings.ReplaceAll(str2, "/", "@/@")
	str4 := strings.ReplaceAll(str3, "+", "@+@")
	str5 := strings.ReplaceAll(str4, "-", "@-@")
	str6 := strings.ReplaceAll(str5, "(", "@(@")
	str7 := strings.ReplaceAll(str6, ")", "@)@")
	str8 := strings.ReplaceAll(str7, "@@", "@")
	return strings.Split(str8, "@")
}

func ToString(s any) string {
	str, _ := ToStringErr(s)
	return str
}

func ToStringErr(s any) (string, error) {
	switch v := s.(type) {
	case int:
		return fmt.Sprintf("%d", v), nil
	case int8:
		return fmt.Sprintf("%d", v), nil
	case int16:
		return fmt.Sprintf("%d", v), nil
	case int32:
		return fmt.Sprintf("%d", v), nil
	case int64:
		return fmt.Sprintf("%d", v), nil
	case uint:
		return fmt.Sprintf("%d", v), nil
	case uint8:
		return fmt.Sprintf("%d", v), nil
	case uint16:
		return fmt.Sprintf("%d", v), nil
	case uint32:
		return fmt.Sprintf("%d", v), nil
	case uint64:
		return fmt.Sprintf("%d", v), nil
	case float32:
		return fmt.Sprintf("%f", v), nil
	case float64:
		return fmt.Sprintf("%f", v), nil
	case string:
		return v, nil
	case []byte:
		return string(v), nil
	case []rune:
		return string(v), nil
	}
	return "", fmt.Errorf("ToStringErr: %v, unkown type: %T", s, s)
}

func ToNumber(number any) float64 {
	num, _ := ToNumberErr(number)
	return num
}
func ToNumberErr(number any) (float64, error) {
	switch v := number.(type) {
	case int:
		return float64(v), nil
	case int8:
		return float64(v), nil
	case int16:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case uint:
		return float64(v), nil
	case uint8:
		return float64(v), nil
	case uint16:
		return float64(v), nil
	case uint32:
		return float64(v), nil
	case uint64:
		return float64(v), nil
	case float32:
		return float64(v), nil
	case float64:
		return v, nil
	case string:
		num, _ := strconv.ParseFloat(v, 64)
		return num, nil
	}
	return 0, fmt.Errorf("ToNumberErr: %v, unkown type: %T", number, number)
}

func simpleCalculate(operator string, num1, num2 float64) (float64, error) {
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		return num1 / num2, nil
	}
	return 0, fmt.Errorf("calculate: unkown operator: %s", operator)
}

func priority(s byte) int {
	switch s {
	case '+':
		return 1
	case '-':
		return 1
	case '*':
		return 2
	case '/':
		return 2
	}
	return 0
}

// 后缀表达式求值
func calculatePostfixExpression(expression []string) (float64, error) {
	var (
		num1 string
		num2 string
		err  error
		s    Stack //	操作栈，用于存入操作数，运算符
	)

	//	从左至右扫描后缀表达式
	for i := 0; i < len(expression); i++ {
		var cur = expression[i]

		//	1. 若读取的是运算符
		if cur[0] == '+' || cur[0] == '-' || cur[0] == '*' || cur[0] == '/' {
			//	从操作栈中弹出两个数进行运算
			num1, err = s.Pop()
			if err != nil {
				return 0, err
			}
			num2, err = s.Pop()
			if err != nil {
				return 0, err
			}

			//	先弹出的数为B，后弹出的数为A
			B, _ := ToNumberErr(num1)
			A, _ := ToNumberErr(num2)
			res, err := simpleCalculate(cur, A, B)
			if err != nil {
				return 0, err
			}

			//将中间结果压栈
			s.Push(fmt.Sprintf("%0.4f", res))
		} else {
			//	1. 若读取的是操作数，直接压栈
			s.Push(cur)
		}
	}

	//	计算结束，栈顶保存最后结果
	resultStr, err := s.Top()
	if err != nil {
		return 0, err
	}
	result, err := ToNumberErr(resultStr)
	return result, err
}

// 将中缀表达式转换成后缀表达式（逆波兰式），expression格式为后缀表达式
func toPostfixExpression(expression string) (postfixExpression []string, err error) {
	var (
		opStack Stack //	运算符堆栈
		i       int
	)

LABEL:
	for i < len(expression) { //	从左至右扫描中缀表达式
		switch {
		//	1. 若读取的是操作数，则将该操作数存入后缀表达式。
		case (expression[i] >= '0' && expression[i] <= '9') || expression[i] == '.':
			var number []byte //	如数字12.3，由'1'、'2'、'.', '3'组成
			for ; i < len(expression); i++ {
				if !((expression[i] >= '0' && expression[i] <= '9') || expression[i] == '.') {
					break
				}
				number = append(number, expression[i])
			}
			postfixExpression = append(postfixExpression, string(number))

		//	2. 若读取的是运算符：
		//	(1) 该运算符为左括号"("，则直接压入运算符堆栈。
		case expression[i] == '(':
			opStack.Push(fmt.Sprintf("%c", expression[i]))
			i++

		//	(2) 该运算符为右括号")"，则输出运算符堆栈中的运算符到后缀表达式，直到遇到左括号为止。
		case expression[i] == ')':
			for !opStack.IsEmpty() {
				data, _ := opStack.Pop()
				if data[0] == '(' {
					break
				}
				postfixExpression = append(postfixExpression, data)
			}
			i++

		//	(3) 该运算符为非括号运算符:
		case expression[i] == '+' || expression[i] == '-' || expression[i] == '*' || expression[i] == '/':
			//	(a)若运算符堆栈为空,则直接压入运算符堆栈。
			if opStack.IsEmpty() {
				opStack.Push(fmt.Sprintf("%c", expression[i]))
				i++
				continue LABEL
			}

			data, _ := opStack.Top()
			//	(b)若运算符堆栈栈顶的运算符为括号，则直接压入运算符堆栈。(只可能为左括号这种情况)
			if data[0] == '(' {
				opStack.Push(fmt.Sprintf("%c", expression[i]))
				i++
				continue LABEL
			}
			//	(c)若比运算符堆栈栈顶的运算符优先级低或相等，则输出栈顶运算符到后缀表达式,直到栈为空或者找到优先级高于当前运算符。并将当前运算符压入运算符堆栈。
			if priority(expression[i]) <= priority(data[0]) {
				tmp := priority(expression[i])
				for !opStack.IsEmpty() && tmp <= priority(data[0]) {
					postfixExpression = append(postfixExpression, data)
					opStack.Pop()
					data, _ = opStack.Top()
				}
				opStack.Push(fmt.Sprintf("%c", expression[i]))
				i++
				continue LABEL
			}
			//	(d)若比运算符堆栈栈顶的运算符优先级高，则直接压入运算符堆栈。
			opStack.Push(fmt.Sprintf("%c", expression[i]))
			i++

		default:
			err = fmt.Errorf("invalid expression:%v", expression[i])
			return
		}
	}

	//	3. 扫描结束，将运算符堆栈中的运算符依次弹出，存入后缀表达式。
	for !opStack.IsEmpty() {
		data, _ := opStack.Pop()
		if data[0] == '#' {
			break
		}
		postfixExpression = append(postfixExpression, data)
	}
	return
}
