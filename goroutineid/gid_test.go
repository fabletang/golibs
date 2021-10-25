package goroutineid_test

import (
	"testing"

	"github.com/fabletang/golibs/goroutineid"
)

func TestGetGID(t *testing.T) {
	t.Logf("goroutin id:%d\n", goroutineid.GetGID())
}
