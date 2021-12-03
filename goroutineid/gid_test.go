package goroutineid

import (
	"testing"
)

func TestGetGID(t *testing.T) {
	gid := GetGID()

	t.Logf("goroutin id:%d\n", gid)
	if gid > 1<<32 {
		t.Errorf("grountine id:%d > 1<<32,maybe err!", gid)
	}
}
