package uniqueStr_test

import (
	"github.com/ricardgo403/uniqueStr"
	"testing"
)

func TestRun(t *testing.T) {
	err := uniqueStr.Run()
	if err != nil {
		t.Fatalf("Not expected error: %v", err)
	}
}
