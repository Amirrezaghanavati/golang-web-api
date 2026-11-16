package common

import (
	"crypto/rand"
	"errors"
	"math/big"
	"unicode"
	"unicode/utf8"
)

const (
	passwordUppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	passwordLowercase = "abcdefghijklmnopqrstuvwxyz"
	passwordDigits    = "0123456789"
	passwordSymbols   = "!@#$%^&*()-_=+[]{}:,.?"
	passwordAlphabet  = passwordUppercase + passwordLowercase + passwordDigits + passwordSymbols
)

// PasswordRules specifies the allowed lengths and minimum counts for each character class.
type PasswordRules struct {
	MinLength      int
	MaxLength      int
	GenerateLength int
	MinUppercase   int
	MinLowercase   int
	MinDigits      int
	MinSymbols     int
}

type passwordCharacterClass struct {
	alphabet string
	minimum  int
}

func (r PasswordRules) characterClasses() [4]passwordCharacterClass {
	return [4]passwordCharacterClass{
		{passwordUppercase, r.MinUppercase},
		{passwordLowercase, r.MinLowercase},
		{passwordDigits, r.MinDigits},
		{passwordSymbols, r.MinSymbols},
	}
}

// PasswordPolicy generates and checks passwords using rules supplied by the caller.
type PasswordPolicy struct {
	rules PasswordRules
}

// NewPasswordPolicy checks the rules and returns a policy safe to share between callers.
func NewPasswordPolicy(rules PasswordRules) (*PasswordPolicy, error) {
	if rules.MinLength < 1 || rules.MaxLength < rules.MinLength {
		return nil, errors.New("password length range is invalid")
	}
	if rules.GenerateLength < rules.MinLength || rules.GenerateLength > rules.MaxLength {
		return nil, errors.New("generated password length must be within the allowed range")
	}

	remaining := rules.GenerateLength
	for _, class := range rules.characterClasses() {
		if class.minimum < 0 {
			return nil, errors.New("password character class minimums must be nonnegative")
		}
		if class.minimum > remaining {
			return nil, errors.New("password character class minimums exceed the generated password length")
		}
		remaining -= class.minimum
	}

	return &PasswordPolicy{rules: rules}, nil
}

// CheckPassword reports whether the password satisfies the configured rules.
func (p *PasswordPolicy) CheckPassword(password string) bool {
	if p == nil || p.rules.MinLength == 0 || !utf8.ValidString(password) {
		return false
	}
	length := utf8.RuneCountInString(password)
	if length < p.rules.MinLength || length > p.rules.MaxLength {
		return false
	}

	var uppercase, lowercase, digits, symbols int
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			uppercase++
		case unicode.IsLower(char):
			lowercase++
		case unicode.IsDigit(char):
			digits++
		case unicode.IsPunct(char), unicode.IsSymbol(char):
			symbols++
		}
	}
	return uppercase >= p.rules.MinUppercase &&
		lowercase >= p.rules.MinLowercase &&
		digits >= p.rules.MinDigits &&
		symbols >= p.rules.MinSymbols
}

// GeneratePassword creates a cryptographically random password that satisfies the policy.
func (p *PasswordPolicy) GeneratePassword() (string, error) {
	if p == nil || p.rules.GenerateLength == 0 {
		return "", errors.New("password policy is unconfigured")
	}

	password := make([]byte, p.rules.GenerateLength)
	offset := 0
	for _, class := range p.rules.characterClasses() {
		end := offset + class.minimum
		if err := fillRandomCharacters(password[offset:end], class.alphabet); err != nil {
			return "", err
		}
		offset = end
	}
	if err := fillRandomCharacters(password[offset:], passwordAlphabet); err != nil {
		return "", err
	}

	// Shuffle the required characters so their positions are random too.
	for i := len(password) - 1; i > 0; i-- {
		j, err := randomPasswordIndex(i + 1)
		if err != nil {
			return "", err
		}
		password[i], password[j] = password[j], password[i]
	}
	return string(password), nil
}

func fillRandomCharacters(destination []byte, alphabet string) error {
	for i := range destination {
		index, err := randomPasswordIndex(len(alphabet))
		if err != nil {
			return err
		}
		destination[i] = alphabet[index]
	}
	return nil
}

func randomPasswordIndex(length int) (int, error) {
	index, err := rand.Int(rand.Reader, big.NewInt(int64(length)))
	if err != nil {
		return 0, err
	}
	return int(index.Int64()), nil
}
