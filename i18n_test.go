package i18n

import (
	"testing"
	"fmt"
)

func TestLoad(t *testing.T) {
	fmt.Println(Load("./langs"))

	fmt.Println(T("key1"))

	SetDefault("zh_CN")

	fmt.Println(T("key1"))
}
