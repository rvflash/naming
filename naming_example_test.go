// Copyright (c) 2020 Hervé Gouchet. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package naming_test

import (
	"fmt"

	"github.com/rvflash/naming"
)

func ExampleFields() {
	fmt.Println(naming.Fields("This isA really-sample XML_or_Whatever    data"))
	// output: [This is A really sample XML or Whatever data]
}

func ExampleCamelCase() {
	fmt.Println(naming.CamelCase("camel case data-19"))
	// output: camelCaseData19
}

func ExampleConstantCase() {
	fmt.Println(naming.ConstantCase("cgo enabled"))
	// output: CGO_ENABLED
}

func ExampleFlatCase() {
	fmt.Println(naming.FlatCase("go-toolDir"))
	// output: gotooldir
}

func ExampleKebabCase() {
	fmt.Println(naming.KebabCase("kebab-case Data_19"))
	// output: kebab-case-data-19
}

func ExamplePascalCase() {
	fmt.Println(naming.PascalCase("camel case data-19"))
	// output: CamelCaseData19
}

func ExampleSnakeCase() {
	fmt.Println(naming.SnakeCase("kebab-case Data_19"))
	// output: kebab_case_data_19
}

func ExampleTrainCase() {
	fmt.Println(naming.TrainCase("HTTP header case"))
	// output: HTTP-Header-Case
}

func ExampleUpperFlatCase() {
	fmt.Println(naming.UpperFlatCase("Go_Host arch"))
	// output: GOHOSTARCH
}
