// Copyright (c) 2020 Hervé Gouchet. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

// Package naming provides methods to parse and build compound names for variables types,
// functions, classes or other structures in source code.
package naming

import (
	"strings"
	"unicode"
)

const (
	kebab = '-'
	snake = '_'
	space = ' '
	blank = ""
)

// Fields splits the string s around each instance of one or more consecutive white space characters,
// underscore or hyphen. Its also break title words.
// It returns a slice of substrings of s or an empty slice if s contains only white space, underscore or hyphen.
func Fields(s string) []string {
	if len(s) == 0 {
		return nil
	}
	var (
		buf = strings.Builder{}
		cut = unicode.IsUpper
		del = func(r rune) bool {
			return r == kebab || r == snake || unicode.IsSpace(r)
		}
		d, wd, c, wc bool
	)
	for _, r := range s {
		d, c = del(r), cut(r)
		if (c && !wc) || (d && !wd) {
			_, _ = buf.WriteRune(space)
		}
		if !d {
			_, _ = buf.WriteRune(r)
		}
		wd, wc = d, c
	}
	return strings.Fields(buf.String())
}

// CamelCase applies the Fields method on s and returns a string with all the words capitalized after the first one.
func CamelCase(s string) string {
	return strings.Join(mapping(Fields(s), func(k int, v string) string {
		if k > 0 {
			return strings.Title(v)
		}
		return v
	}), blank)
}

// ConstantCase applies the Fields method on s and returns a string with all the words capitalized and
// joined with an underscore.
func ConstantCase(s string) string {
	return strings.Join(mapping(Fields(s), func(k int, v string) string {
		return strings.ToTitle(v)
	}), string(snake))
}

// FlatCase applies the Fields method on s and concatenates any words to return only one in lowercase.
func FlatCase(s string) string {
	return strings.Join(mapping(Fields(s), func(k int, v string) string {
		return strings.ToLower(v)
	}), blank)
}

// KebabCase applies the Fields method on s and returns a string with a hyphen between each word
// and all of them are lowercase.
func KebabCase(s string) string {
	return strings.Join(mapping(Fields(s), func(k int, v string) string {
		return strings.ToLower(v)
	}), string(kebab))
}

// PascalCase applies the Fields method on s and returns a string with all the words capitalized.
func PascalCase(s string) string {
	return strings.Join(mapping(Fields(s), func(k int, v string) string {
		return strings.Title(v)
	}), blank)
}

// SnakeCase applies the Fields method on s and returns a string with an underscore between each word
// and all of them are lowercase.
func SnakeCase(s string) string {
	return strings.Join(mapping(Fields(s), func(k int, v string) string {
		return strings.ToLower(v)
	}), string(snake))
}

// TrainCase applies the Fields method on s and returns a string with all the words capitalized
// and separated with a hyphen.
func TrainCase(s string) string {
	return strings.Join(mapping(Fields(s), func(k int, v string) string {
		return strings.Title(v)
	}), string(kebab))
}

// UpperFlatCase applies the Fields method on s and concatenates any words to return only one in uppercase.
func UpperFlatCase(s string) string {
	return strings.Join(mapping(Fields(s), func(k int, v string) string {
		return strings.ToTitle(v)
	}), blank)
}

func mapping(a []string, f func(k int, v string) string) []string {
	if len(a) == 0 {
		return nil
	}
	for k, v := range a {
		a[k] = f(k, v)
	}
	return a
}
