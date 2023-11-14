package krn

import (
	"reflect"
	"testing"

	"gitlab.com/nzv/prc"
)

func TestSwaps(t *testing.T) {

	prc.ProcPath = "testdata"

	tests := []struct {
		desc string
		want []Swap
	}{
		{
			desc: "ok swaps",
			want: []Swap{
				{
					Filename: "/dev/sda5",
					Type:     "partition",
					Size:     2097148,
					Used:     1024,
					Priority: -2,
				},
				{
					Filename: "/dev/sdb1",
					Type:     "file",
					Size:     1048572,
					Used:     512,
					Priority: -1,
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			got, err := Swaps()
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("swaps mismatch: got %+v, want %+v", got, tc.want)
			}
		})
	}
}
