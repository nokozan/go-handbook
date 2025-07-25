package strategy

type StrategyFunc func(string) bool

type Validator struct {
	strategies []StrategyFunc
}

func NewValidator(strategies ...StrategyFunc) *Validator {
	return &Validator{strategies: strategies}
}

func (v *Validator) Validator(input string) bool {
	for _, strategy := range v.strategies {
		if !strategy(input) {
			return false
		}
	}
	return true
}

/// common reusable strategy functions

func NotEmpty(s string) bool {
	return len(s) > 0
}

func HasAtSymbol(s string) bool {
	return len(s) > 3 && contains(s, "@")
}

func MinLength(n int) StrategyFunc {
	return func(s string) bool {
		return len(s) >= n
	}
}

func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
