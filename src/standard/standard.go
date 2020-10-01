package standard

type Runner interface {
	Run(q int)
}

type Solution interface {
	Test()
}

func Equal(actual, expected string) bool {
	return actual == expected
}