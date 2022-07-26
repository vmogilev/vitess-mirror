/*
Copyright 2020 The Vitess Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package textutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitDelimitedList(t *testing.T) {
	defaultList := []string{"one", "two", "three"}
	tt := []struct {
		s    string
		list []string
	}{
		{s: "one,two,three"},
		{s: "one, two, three"},
		{s: "one,two; three  "},
		{s: "one two three"},
		{s: "one,,,two,three"},
		{s: " one, ,two,  three "},
	}

	for _, tc := range tt {
		if tc.list == nil {
			tc.list = defaultList
		}
		list := SplitDelimitedList(tc.s)
		assert.Equal(t, tc.list, list)
	}
}

func TestEscapeJoin(t *testing.T) {
	elems := []string{"normal", "with space", "with,comma", "with?question"}
	s := EscapeJoin(elems, ",")
	assert.Equal(t, "normal,with+space,with%2Ccomma,with%3Fquestion", s)
}

func TestSplitUnescape(t *testing.T) {
	{
		s := ""
		elems, err := SplitUnescape(s, ",")
		assert.NoError(t, err)
		assert.Nil(t, elems)
	}
	{
		s := "normal,with+space,with%2Ccomma,with%3Fquestion"
		expected := []string{"normal", "with space", "with,comma", "with?question"}
		elems, err := SplitUnescape(s, ",")
		assert.NoError(t, err)
		assert.Equal(t, expected, elems)
	}
}

func TestSingleWordCamel(t *testing.T) {
	tt := []struct {
		word   string
		expect string
	}{
		{
			word:   "",
			expect: "",
		},
		{
			word:   "_",
			expect: "_",
		},
		{
			word:   "a",
			expect: "A",
		},
		{
			word:   "A",
			expect: "A",
		},
		{
			word:   "_A",
			expect: "_a",
		},
		{
			word:   "mysql",
			expect: "Mysql",
		},
		{
			word:   "mySQL",
			expect: "Mysql",
		},
	}
	for _, tc := range tt {
		t.Run(tc.word, func(t *testing.T) {
			camel := SingleWordCamel(tc.word)
			assert.Equal(t, tc.expect, camel)
		})
	}
}
