//uc_test
package uc

import "testing"
import "fmt"

type ucTest struct {
	in, out string
}

var ucTests = []ucTest{
	ucTest{"abC", "ABC"},
	ucTest{"Adf", "ADF"},
	ucTest{"123cbd", "123CBD"},
	ucTest{"1Dbd", "1DBD1"},
}

func TestUC(t *testing.T) {
	for _, ut := range ucTests {
		uc := UpperCase(ut.in)
		fmt.Println(ut.in, ut.out, uc)
		if uc != ut.out {
			t.Errorf("UpperCase(%s) = %s,must be %s", ut.in, uc, ut.out)
		}
	}
}
