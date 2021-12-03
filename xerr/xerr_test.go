package xerr

import (
	"testing"
	//"github.com/fabletang/golibs/xerr"
)

func TestXerror(t *testing.T) {
	//var xerr Xerror
	xerror := New(500, "test")
	t.Log(xerror)
	if len(xerror.Cause) < 10 {
		t.Errorf("xerr.Cause err:%s\n", xerror)
	}
	//fmt.Printf("the 64-bit hash of 'abc' is: 0x%x\n", hash64)
	//fmt.Printf("the 64-bit hash of 'abc' is: %d\n", hash64)
	// Output:
	// the 64-bit hash of 'hello' is: 0xb48be5a931380ce8
}
