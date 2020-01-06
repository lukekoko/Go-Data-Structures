package linkedList

import "testing"

func TestList_append(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		l    *List
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.l.append(tt.args.val)
		})
	}
}
