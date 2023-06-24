package data

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
	- unit test file naming MUST xxx_test.go
	- import library "testing"
	- unit test function naming MUST func TestBlahblahblah(t *testing.T) {
		// the content to test with
	}

	- run the unit test by
	1. go to root folder
	2. run command go test .\unit_test_folder
*/

func TestProducts(t *testing.T) {
	type args struct { // function args definition
		p *Product
	}

	tests := []struct { // test case defination
		name string
		args args
		want bool
	}{ // test cases
		{ // test case
			name: "Test product name validation",
			args: args{&Product{Price: 0.99, SKU: "aaa-bbb-ccc-ddd"}},
			want: false,
		},
		{ // test case
			name: "Test product price validation",
			args: args{&Product{Name: "Tuna fish", Price: -0.01, SKU: "aaa-bbb-ccc-ddd"}},
			want: false,
		},
		{ // test case
			name: "Test product SKU validation",
			args: args{&Product{Name: "Tuna fish", Price: 0.99, SKU: "ABS"}},
			want: false,
		},
	}
	for _, tt := range tests { // test cases execution
		t.Run(tt.name,
			func(t *testing.T) {
				got := func(p *Product) bool {
					if err := p.Validate(); err != nil {
						fmt.Printf("Validate error: %v\n", err)
						return true
					}
					return false
				}(tt.args.p)

				assert.Equal(t, tt.want, got)
			})
	}
}

/*
func TestProductValidation(t *testing.T) {
	p := &Product{}
	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
*/
