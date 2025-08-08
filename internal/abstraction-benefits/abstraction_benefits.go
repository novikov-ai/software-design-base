package abstraction_benefits

// ===| Top Level Abstraction |===
type Constrainter interface {
	ValidationNeeded() bool
	Validate() bool
}

// ===| Concrete Constrainter |===
type WordsLimiter struct {}

func NewWordsLimiter() WordsLimiter {
	return WordsLimiter{}
}

func (wl WordsLimiter) ValidationNeeded() bool {
	return true
}

func (wl WordsLimiter) Validate() bool {
	return false
}

// ===| Constrainter Coordinator |===
type Validator struct {
	cons []Constrainter
}

func (v *Validator) Validate() bool {
	for _, c := range v.cons {
		if !c.ValidationNeeded() {
			continue
		}

		if !c.Validate() {
			return false
		}
	}

	return true
}
