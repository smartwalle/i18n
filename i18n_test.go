package i18n

import (
	"testing"
	"fmt"
)

func TestLoad(t *testing.T) {
	fmt.Println(Load("./langs"))

	fmt.Println(T("name"))

	SetDefault("en")
	fmt.Println(T("gender"))

	SetDefault("cn")
	fmt.Println(T("gender"))

	fmt.Println(Exists("fr"), Exists("en"))
}
