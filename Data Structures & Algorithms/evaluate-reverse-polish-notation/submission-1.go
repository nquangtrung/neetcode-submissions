func isOp(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/"
}

func evaluate(tokens []string) (int, []string) {
	token, tokens := tokens[len(tokens) - 1], tokens[:len(tokens) - 1]
	if isOp(token) {
		right, tokens := evaluate(tokens)
		left, tokens := evaluate(tokens)

		// fmt.Printf("evaluating %d %s %d\n", left, token, right)
		switch (token) {
		case "+": return left + right, tokens
		case "-": return left - right, tokens
		case "*": return left * right, tokens
		case "/": return left / right, tokens
		default: return 0, tokens
		}

	} else {
		i, _ := strconv.Atoi(token) 
		return i, tokens
	}
}

func evalRPN(tokens []string) int {
	result, _ := evaluate(tokens)
	return result
}
