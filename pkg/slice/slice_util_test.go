package slice

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceContainElement(t *testing.T) {
	type argInts struct {
		slice []int
		ele   int
	}
	testInts := []struct {
		name string
		args argInts
		want bool
	}{
		{
			name: "[]int - true",
			args: argInts{
				slice: []int{1, 2, 3, 4, 5},
				ele:   1,
			},
			want: true,
		},
		{
			name: "[]int - false",
			args: argInts{
				slice: []int{1, 2, 3, 4, 5},
				ele:   6,
			},
			want: false,
		},
	}
	for _, tt := range testInts {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, SliceContainElement(tt.args.slice, tt.args.ele), "SliceContainElement(%v, %v)", tt.args.slice, tt.args.ele)
		})
	}
	type argStrs struct {
		slice []string
		ele   string
	}
	testStrs := []struct {
		name string
		args argStrs
		want bool
	}{
		{
			name: "[]string - true",
			args: argStrs{
				slice: []string{"A", "B", "C"},
				ele:   "A",
			},
			want: true,
		},
		{
			name: "[]string - false",
			args: argStrs{
				slice: []string{"A", "B", "C"},
				ele:   "A",
			},
			want: true,
		},
	}
	for _, tt := range testStrs {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, SliceContainElement(tt.args.slice, tt.args.ele), "SliceContainElement(%v, %v)", tt.args.slice, tt.args.ele)
		})
	}
	type testStruct struct {
		Key   string
		Value int
	}
	type argStructs struct {
		slice []testStruct
		ele   testStruct
	}
	testStructs := []struct {
		name string
		args argStructs
		want bool
	}{
		{
			name: "[]struct - true",
			args: argStructs{
				slice: []testStruct{
					{
						Key:   "A",
						Value: 1,
					},
					{
						Key:   "A",
						Value: 2,
					},
				},
				ele: testStruct{
					Key:   "A",
					Value: 2,
				},
			},
			want: true,
		},
		{
			name: "[]struct - false",
			args: argStructs{
				slice: []testStruct{
					{
						Key:   "A",
						Value: 1,
					},
					{
						Key:   "A",
						Value: 2,
					},
				},
				ele: testStruct{
					Key:   "A",
					Value: 3,
				},
			},
			want: false,
		},
	}
	for _, tt := range testStructs {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, SliceContainElement(tt.args.slice, tt.args.ele), "SliceContainElement(%v, %v)", tt.args.slice, tt.args.ele)
		})

	}
}

func TestDiffTwoSlices(t *testing.T) {
	type argInts struct {
		oldVals []int
		newVals []int
	}
	testInts := []struct {
		name        string
		args        argInts
		wantAdded   []int
		wantRemoved []int
	}{
		{
			name: "[]int",
			args: argInts{
				oldVals: []int{1, 2, 3, 4, 5},
				newVals: []int{3, 4, 5, 6},
			},
			wantAdded:   []int{6},
			wantRemoved: []int{1, 2},
		},
	}
	for _, tt := range testInts {
		t.Run(tt.name, func(t *testing.T) {
			gotAdded, gotRemoved := DiffTwoSlices(tt.args.newVals, tt.args.oldVals)
			assert.Equalf(t, tt.wantAdded, gotAdded, "DiffTwoSlices(%v, %v)", tt.args.newVals, tt.args.oldVals)
			assert.Equalf(t, tt.wantRemoved, gotRemoved, "DiffTwoSlices(%v, %v)", tt.args.newVals, tt.args.oldVals)
		})
	}
	type argStrs struct {
		oldVals []string
		newVals []string
	}
	testStrs := []struct {
		name        string
		args        argStrs
		wantAdded   []string
		wantRemoved []string
	}{
		{
			name: "[]string ",
			args: argStrs{
				oldVals: []string{"A", "B"},
				newVals: []string{"B", "C"},
			},
			wantAdded:   []string{"C"},
			wantRemoved: []string{"A"},
		},
	}
	for _, tt := range testStrs {
		t.Run(tt.name, func(t *testing.T) {
			gotAdded, gotRemoved := DiffTwoSlices(tt.args.newVals, tt.args.oldVals)
			assert.Equalf(t, tt.wantAdded, gotAdded, "DiffTwoSlices(%v, %v)", tt.args.newVals, tt.args.oldVals)
			assert.Equalf(t, tt.wantRemoved, gotRemoved, "DiffTwoSlices(%v, %v)", tt.args.newVals, tt.args.oldVals)
		})
	}
}

func TestUniqueSlice(t *testing.T) {
	type args struct {
		oldSlice []int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			args: args{
				oldSlice: []int64{1, 2, 3, 3, 3, 3, 4, 2, 5},
			},
			want: []int64{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, UniqueSlice(tt.args.oldSlice), "UniqueSlice(%v)", tt.args.oldSlice)
		})
	}
}

func TestFind(t *testing.T) {
	type args struct {
		slice         []int
		matchFunction func(int) bool
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			args: args{
				slice: []int{1, 2, 3, 4, 5},
				matchFunction: func(i int) bool {
					if i == 5 {
						return true
					}
					return false
				},
			},
			want:    5,
			wantErr: assert.NoError,
		},
		{
			args: args{
				slice: []int{1, 2, 3, 4},
				matchFunction: func(i int) bool {
					if i == 5 {
						return true
					}
					return false
				},
			},
			want:    0,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Find(tt.args.slice, tt.args.matchFunction)
			if !tt.wantErr(t, err, fmt.Sprintf("Find in(%v)", tt.args.slice)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Find(%v, %v)", tt.args.slice, tt.args.matchFunction)
		})
	}
}
