package main

import (
	"os"
	"testing"
)

func Test(t *testing.T) {
	os.Setenv("DB_DSN", "-")
	main()
}
