package internal

import (
	"crypto/sha256"
	"fmt"
)

func SumPassword(password string)  string{
	
	sum := sha256.Sum256([]byte(password))
	return fmt.Sprintf("%x", sum)
}