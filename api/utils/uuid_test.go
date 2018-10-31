package utils

import (
	"reflect"
	"testing"
)

func TestNewUUID(t *testing.T) {
	t.Run("test uuid", func(t *testing.T) {
		got,_ := NewUUID()

		want := ""

		if !reflect.DeepEqual(got, want) {
			t.Errorf("failed got : %v ; want : %v", got, want)
		}

	})
}
