// Copyright (c) 2020 Hervé Gouchet. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package naming_test

import (
	"strings"
	"testing"

	"github.com/matryer/is"
	"github.com/rvflash/naming"
)

const (
	seps     = " _-:"
	sample   = "Little-sample varName b_42"
	repeated = "c--:-c C_-_c cccCCC   cC"
	complex  = "camelCase CONSTANT_CASE kebab-case PascalCase snake_case Train-Case"
)

func TestJoin(t *testing.T) {
	t.Parallel()

	words := strings.Fields(sample)

	for name, tc := range map[string]struct {
		// inputs
		format func(s string) string
		// outputs
		out string
	}{
		"CamelCase":     {format: naming.CamelCase, out: "littleSampleVarNameB42"},
		"ConstantCase":  {format: naming.ConstantCase, out: "LITTLE_SAMPLE_VAR_NAME_B_42"},
		"FlatCase":      {format: naming.FlatCase, out: "littlesamplevarnameb42"},
		"KebabCase":     {format: naming.KebabCase, out: "little-sample-var-name-b-42"},
		"PascalCase":    {format: naming.PascalCase, out: "LittleSampleVarNameB42"},
		"SnakeCase":     {format: naming.SnakeCase, out: "little_sample_var_name_b_42"},
		"TrainCase":     {format: naming.TrainCase, out: "Little-Sample-Var-Name-B-42"},
		"UpperFlatCase": {format: naming.UpperFlatCase, out: "LITTLESAMPLEVARNAMEB42"},
	} {
		tt := tc
		t.Run(name, func(t *testing.T) {
			is.NewRelaxed(t).Equal(tt.out, naming.Join(tt.format, words...))
		})
	}
}

func TestFields(t *testing.T) {
	t.Parallel()

	for name, tc := range map[string]struct {
		in  string
		out []string
	}{
		"Default":  {},
		"Blank":    {in: seps, out: []string{}},
		"Repeated": {in: repeated, out: []string{"c", "c", "C", "c", "ccc", "CCC", "c", "C"}},
		"OK": {
			in: complex,
			out: []string{
				"camel", "Case", "CONSTANT", "CASE", "kebab", "case",
				"Pascal", "Case", "snake", "case", "Train", "Case",
			},
		},
	} {
		tt := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			is.NewRelaxed(t).Equal(tt.out, naming.Fields(tt.in)) // mismatch fields
		})
	}
}

func TestCamelCase(t *testing.T) {
	t.Parallel()
	are := is.New(t)
	are.Equal("", naming.CamelCase(""))                           // mismatch default
	are.Equal("littleSampleVarNameB42", naming.CamelCase(sample)) // mismatch sample
}

func TestConstantCase(t *testing.T) {
	t.Parallel()
	are := is.New(t)
	are.Equal("", naming.ConstantCase(""))                                // mismatch default
	are.Equal("LITTLE_SAMPLE_VAR_NAME_B_42", naming.ConstantCase(sample)) // mismatch sample
}

func TestFlatCase(t *testing.T) {
	t.Parallel()
	are := is.New(t)
	are.Equal("", naming.FlatCase(""))                           // mismatch default
	are.Equal("littlesamplevarnameb42", naming.FlatCase(sample)) // mismatch sample
}

func TestKebabCase(t *testing.T) {
	t.Parallel()
	are := is.New(t)
	are.Equal("", naming.KebabCase(""))                                // mismatch default
	are.Equal("little-sample-var-name-b-42", naming.KebabCase(sample)) // mismatch sample
}

func TestPascalCase(t *testing.T) {
	t.Parallel()
	are := is.New(t)
	are.Equal("", naming.PascalCase(""))                           // mismatch default
	are.Equal("LittleSampleVarNameB42", naming.PascalCase(sample)) // mismatch sample
}

func TestSnakeCase(t *testing.T) {
	t.Parallel()
	are := is.New(t)
	are.Equal("", naming.SnakeCase(""))                                // mismatch default
	are.Equal("little_sample_var_name_b_42", naming.SnakeCase(sample)) // mismatch sample
}

func TestTrainCase(t *testing.T) {
	are := is.New(t)
	are.Equal("", naming.TrainCase(""))                                // mismatch default
	are.Equal("Little-Sample-Var-Name-B-42", naming.TrainCase(sample)) // mismatch sample
}

func TestUpperFlatCase(t *testing.T) {
	t.Parallel()
	are := is.New(t)
	are.Equal("", naming.UpperFlatCase(""))                           // mismatch default
	are.Equal("LITTLESAMPLEVARNAMEB42", naming.UpperFlatCase(sample)) // mismatch sample
}
