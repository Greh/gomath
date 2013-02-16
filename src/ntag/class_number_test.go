// Copyright (c) 2012 Google, Inc. All Rights Reserved.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ntag

import (
	"fmt"
	"math"
	"testing"
)

func dEquals(a, b float64) bool {
	return math.Abs(a-b) < 1e-6
}

func TestRegulator(t *testing.T) {
	p := ParseIntPoly("x^2 - 3*x + 1")
	gotr := p.regulatorRealQuad()
	r := 0.481211825059603 // From Sage.
	if !dEquals(r, gotr) {
		fmt.Printf("Poly %s regulator %f but got %f\n", p, r, gotr)
		t.FailNow()
	}
}

func TestRandomClassNumbers(t *testing.T) {
	for _, testCase := range classNumberTestCases {
		p := ParseIntPoly(testCase.polyString)
		if p.Degree() > 2 {
			continue
		}
		if p.Discriminant().Sign() > 0 {
			continue
		}
		h := testCase.classNumber
		k := MakeNumberField(p)
		hGot := k.ClassNumber()
		if hGot != h {
			fmt.Printf("%s expected class number %d got %d\n", p.String(), h, hGot)
			fmt.Println("Discriminant:", p.Discriminant())
			t.FailNow()
		}
	}
}

func TestParsePolynomial(t *testing.T) {
	testCases := []string{"x^2 + 5*x - 1001", "x^2", "-1*x", "2*x + 1", "x + 1", "x", "1", "2"}
	for _, testCase := range testCases {
		p := ParseIntPoly(testCase).String()
		if p != testCase {
			fmt.Printf("Parsed %s but got %s\n", testCase, p)
			t.FailNow()
		}
	}
}

func TestClassNumber1(t *testing.T) {
	class_number_1_nums := []int{-3, -4, -7, -8, -11, -19, -43, -67, -163}
	class_number_1 := IntSetFromSlice(class_number_1_nums)
	for a := 1; a < 10; a++ {
		for b := -10; b < 10; b++ {
			for c := -10; c < 10; c++ {
				p := MakeIntPolynomial64(int64(c), int64(b), int64(a))
				if !p.IsIrreducible() {
					continue
				}
				if p.Discriminant().Sign() > 0 {
					continue
				}
				if !IsFundamentalDiscriminant(p.Discriminant()) {
					continue
				}
				k := MakeNumberField(p)

				if k.ClassNumber() == 1 && !class_number_1.Contains(int(p.Discriminant().Int64())) {
					t.Fail()
				}
			}
		}
	}
}

var classNumberTestCases = []struct {
	classNumber int
	polyString  string
}{
	{1, "x^2 - x + 1"},
	{1, "x^2 + x + 1"},
	{2, "x^2 - x + 9"},
	{1, "x^2 + x - 1"},
	{1, "x^2 + 2*x + 3"},
	{1, "x^2 - 3*x - 5"},
	{2, "x^2 + 13*x + 1"},
	{2, "x^2 + 10*x - 1"},
	{1, "x^2 - 3"},
	{1, "x^2 - 2*x + 2"},
	{1, "x^2 - x - 1"},
	{1, "x^2 + 4"},
	{1, "x^2 - x - 8"},
	{2, "x^2 - 9*x + 4"},
	{1, "x^2 + 2"},
	{1, "x^2 + x - 3"},
	{1, "x^2 - x + 2"},
	{3, "x^2 + x + 6"},
	{1, "x^2 - x - 3"},
	{1, "x^2 + 4*x - 2"},
	{2, "x^2 - x + 4"},
	{2, "x^2 + x + 13"},
	{1, "x^2 + 9*x + 6"},
	{1, "x^2 - x + 3"},
	{1, "x^2 + 2*x - 2"},
	{2, "x^2 - 12*x + 1"},
	{1, "x^2 + 5*x - 1"},
	{1, "x^2 + 1"},
	{1, "x^2 - 5*x + 5"},
	{1, "x^2 + 2*x - 1"},
	{1, "x^2 - 51*x + 1"},
	{1, "x^2 - 3*x + 1"},
	{1, "x^2 + x - 23"},
	{2, "x^2 - 20*x - 6"},
	{1, "x^2 + 3*x + 1"},
	{1, "x^2 - 18*x + 1"},
	{2, "x^2 + 2*x + 6"},
	{1, "x^2 + x + 2"},
	{1, "x^2 + 3"},
	{1, "x^2 - 4*x + 5"},
	{11, "x^2 - 72*x - 1"},
	{1, "x^2 - 5"},
	{1, "x^2 - 14*x + 1"},
	{7, "x^2 - x + 160"},
	{1, "x^2 + x + 5"},
	{1, "x^2 + 13*x + 5"},
	{4, "x^2 - 11*x - 6"},
	{1, "x^2 - 8*x + 2"},
	{1, "x^2 - 3*x - 8"},
	{1, "x^2 + 11"},
	{1, "x^2 - 6"},
	{1, "x^2 + 5*x + 7"},
	{1, "x^2 - 1523*x + 1"},
	{1, "x^2 + 17*x - 1"},
	{1, "x^2 + x - 4"},
	{1, "x^2 + 2*x + 2"},
	{2, "x^2 - 6*x - 1"},
	{1, "x^2 - 4*x + 1"},
	{1, "x^2 + 3*x - 2"},
	{2, "x^2 + x + 4"},
	{1, "x^2 + x - 18"},
	{1, "x^2 - 15*x - 4"},
	{2, "x^2 - 19*x - 1"},
	{3, "x^2 + x + 8"},
	{8, "x^2 + 210"},
	{4, "x^2 - x + 10"},
	{1, "x^2 - x - 138"},
	{1, "x^2 - 2*x - 1"},
	{1, "x^2 - x - 4"},
	{1, "x^2 + 7*x - 1"},
	{1, "x^2 + 3*x - 1"},
	{1, "x^2 - 13*x - 2"},
	{1, "x^2 - 22*x - 3"},
	{1, "x^2 + 2*x - 74"},
	{2, "x^2 + 10"},
	{3, "x^2 + 31"},
	{1, "x^2 + 10*x + 13"},
	{2, "x^2 + 6*x - 1"},
	{1, "x^2 - 7*x - 3"},
	{1, "x^2 + 2*x - 6"},
	{1, "x^2 + 22*x + 8"},
	{3, "x^2 - 17*x - 8"},
	{2, "x^2 + 2*x + 7"},
	{1, "x^2 + 6*x + 2"},
	{1, "x^2 + 3*x + 4"},
	{1, "x^2 + 4*x - 10"},
	{1, "x^2 - x - 5"},
	{2, "x^2 + 6"},
	{1, "x^2 - 3*x - 1"},
	{1, "x^2 + 2*x + 4"},
	{1, "x^2 + 31*x - 4"},
	{1, "x^2 + 10*x + 1"},
	{2, "x^2 + 9*x - 1"},
	{1, "x^2 - 2*x + 3"},
	{1, "x^2 - 13*x - 1"},
	{3, "x^2 - x + 8"},
	{1, "x^2 - 12*x - 1"},
	{1, "x^2 - 21"},
	{1, "x^2 - 5*x + 1"},
	{1, "x^2 - 2"},
	{1, "x^3 - x - 3"},
	{6, "x^3 + 10*x^2 + 1"},
	{1, "x^3 + 3*x^2 - x + 1"},
	{1, "x^3 + x^2 - x + 31"},
	{1, "x^3 - x^2 - x - 1"},
	{1, "x^3 - x^2 - 4*x + 2"},
	{19, "x^3 + 88*x^2 - x + 1"},
	{1, "x^3 - 5*x - 3"},
	{1, "x^3 + 2*x^2 - x + 1"},
	{4, "x^3 - x + 49"},
	{1, "x^3 + 4*x^2 + 6"},
	{1, "x^3 + 39*x^2 + x - 2"},
	{1, "x^3 - x + 1"},
	{2, "x^3 - 4*x^2 - 1"},
	{3, "x^3 + 9*x^2 - x + 1"},
	{1, "x^3 - x - 1"},
	{3, "x^3 - 19*x^2 - 8*x - 1"},
	{1, "x^3 - 7*x^2 + x + 1"},
	{1, "x^3 + x^2 + 1"},
	{1, "x^3 + 5*x - 13"},
	{1, "x^3 - 2*x^2 - x + 3"},
	{1, "x^3 - x^2 + 2*x + 1"},
	{1, "x^3 + 5*x - 3"},
	{1, "x^3 - x^2 - 2*x + 1"},
	{1, "x^3 - x^2 - 3*x - 2"},
	{1, "x^3 - 2*x^2 - 25*x - 2"},
	{1, "x^3 + x - 1"},
	{1, "x^3 - x^2 + x - 3"},
	{1, "x^3 + x^2 + x - 1"},
	{8, "x^3 - 32*x^2 - 2"},
	{1, "x^3 + x^2 + 2*x - 1"},
	{1, "x^3 - x^2 - 5*x + 1"},
	{1, "x^3 + 2*x + 1"},
	{1, "x^3 + x^2 + 2*x - 2"},
	{1, "x^3 + 6*x^2 + 2*x + 2"},
	{1, "x^3 - x^2 - 18*x + 11"},
	{3, "x^3 + 20*x^2 + x - 1"},
	{1, "x^3 - 16*x^2 + 2*x + 1"},
	{1, "x^3 + 2*x^2 + x + 5"},
	{1, "x^3 + 2*x^2 - 24*x + 2"},
	{1, "x^3 - 6*x^2 + x + 1"},
	{1, "x^3 + 2*x^2 - x + 3"},
	{1, "x^3 - x^2 - x - 3"},
	{1, "x^3 + x^2 - 81*x - 1"},
	{1, "x^3 + x^2 - x + 1"},
	{1, "x^3 - x^2 + x + 1"},
	{1, "x^3 - 3*x - 1"},
	{1, "x^3 + 5*x^2 + x - 2"},
	{1, "x^3 + x^2 - x + 3"},
	{1, "x^3 - 5*x^2 + x - 1"},
	{6, "x^3 - x^2 + x - 27"},
	{1, "x^3 + x^2 - 7"},
	{3, "x^3 - x^2 + 553*x - 3"},
	{1, "x^3 - x^2 - 13*x - 1"},
	{1, "x^3 + x^2 + x + 2"},
	{1, "x^3 - x^2 - x - 11"},
	{2, "x^3 - x^2 - 32*x + 1"},
	{1, "x^3 + 3*x - 1"},
	{5, "x^3 - x^2 + x - 102"},
	{4, "x^3 - 64*x + 2"},
	{1, "x^3 - 2*x + 6"},
	{1, "x^3 + 12*x^2 + x + 2"},
	{1, "x^3 + 5*x^2 - x + 4"},
	{1, "x^3 - x^2 - 1"},
	{2, "x^3 - 5*x^2 - 3*x + 176"},
	{1, "x^3 - 8*x + 1"},
	{1, "x^3 + x^2 - 2*x + 1"},
	{1, "x^3 - 4*x^2 + x - 6"},
	{2, "x^3 + x^2 + 16*x - 1"},
	{1, "x^3 + 4"},
	{3, "x^3 - 12*x^2 + 1"},
	{1, "x^3 + x^2 + x - 5"},
	{1, "x^3 + 11*x^2 + 2*x + 16"},
	{2, "x^3 + 13*x^2 + 2*x + 1"},
	{4, "x^3 - 2*x^2 + 4*x - 9"},
	{1, "x^3 - 2*x^2 + x - 10"},
	{1, "x^3 + 2*x^2 - x - 3"},
	{1, "x^3 + x^2 - 5*x - 4"},
	{1, "x^3 + x^2 + 5*x - 2"},
	{1, "x^3 - x^2 + 5"},
	{4, "x^3 + 12*x^2 + 2*x + 1"},
	{5, "x^3 - 13*x^2 + x - 1"},
	{1, "x^3 + x^2 - 1"},
	{2, "x^3 + 14*x^2 + 2*x - 1"},
	{1, "x^3 - 3*x^2 + x - 2"},
	{1, "x^3 - x^2 + x - 14"},
	{1, "x^3 + x^2 - 4*x + 7"},
	{1, "x^3 - x^2 - x + 3"},
	{1, "x^3 - x^2 + 2*x - 1"},
	{1, "x^3 + 9*x^2 - 1"},
	{2, "x^3 + x^2 - 664*x - 1"},
	{1, "x^3 + x^2 - 2*x - 1"},
	{1, "x^3 + 2*x - 7"},
	{1, "x^3 + 9*x^2 + 3"},
	{1, "x^3 - x^2 - 4*x + 1"},
	{1, "x^3 - 55*x^2 + 4"},
	{2, "x^3 - 26*x^2 - 2*x + 1"},
	{2, "x^3 - 2*x^2 + 6*x - 1"},
	{1, "x^3 + x^2 + 2*x + 1"},
	{1, "x^3 - 3*x^2 - x + 5"},
	{1, "x^4 - x^2 - 185*x - 13"},
	{1, "x^4 - 3*x^3 + 3*x^2 - x - 6"},
	{1, "x^4 - x^3 - x^2 - 2"},
	{1, "x^4 + x^3 + x^2 + 5*x + 1"},
	{1, "x^4 + 3*x^3 + x^2 + x - 2"},
	{1, "x^4 + 5*x^3 - x^2 - 5*x + 1"},
	{1, "x^4 - 2*x^3 - 1"},
	{1, "x^4 + x^3 + 2*x^2 - x + 1"},
	{1, "x^4 + 2*x^3 - 6*x^2 - 2*x + 3"},
	{1, "x^4 + 4*x^2 - x + 4"},
	{10, "x^4 + x^3 - 27*x - 1"},
	{1, "x^4 + x^3 + 3*x^2 - 2"},
	{1, "x^4 + x^3 + 1"},
	{1, "x^4 - 5*x^2 - 2"},
	{2, "x^4 + 10*x^3 - 7"},
	{1, "x^4 - x^3 - 1"},
	{1, "x^4 - 3*x^3 - x^2 - 4"},
	{2, "x^4 + 4*x^2 + 1"},
	{1, "x^4 + x^3 - x - 4"},
	{1, "x^4 + x^3 + 6*x^2 - 3*x - 1"},
	{5, "x^4 + 38*x^2 - x - 8"},
	{1, "x^4 - 2*x^3 + 4*x^2 + 15"},
	{1, "x^4 + x^3 + 4*x^2 + 1"},
	{1, "x^4 + x^3 - x^2 - x - 1"},
	{3, "x^4 + 2*x^3 + x^2 - 4*x + 8"},
	{1, "x^4 - x^3 - x^2 - 2*x + 1"},
	{1, "x^4 - 4*x^2 + 1"},
	{1, "x^4 - 3*x^2 - 3*x + 1"},
	{1, "x^4 - x^3 - x + 18"},
	{12, "x^4 + 2*x^3 + 11*x^2 - x + 1"},
	{1, "x^4 + x^3 - x^2 + 4*x - 1"},
	{1, "x^4 + 5*x^3 - 1"},
	{1, "x^4 + x^3 + x^2 + x + 1"},
	{3, "x^4 + 19*x^3 + 3*x^2 + x + 2"},
	{2, "x^4 + 16*x^2 + x + 7"},
	{1, "x^4 + 2*x^3 - x^2 - 2*x + 9"},
	{1, "x^4 + 4*x^3 - x^2 - x - 48"},
	{1, "x^4 - x^3 + 3*x^2 + x - 3"},
	{2, "x^4 - x^3 + 129*x^2 - x + 4"},
	{1, "x^4 + 2*x^3 + 2*x^2 - 3*x - 1"},
	{1, "x^4 + x^3 + x^2 + 3*x + 1"},
	{1, "x^4 + x^2 + 3*x - 2"},
	{1, "x^4 - 5*x^3 + 3*x^2 - 3*x - 1"},
	{1, "x^4 - 2*x^2 - x - 5"},
	{1, "x^4 + 2*x^3 - 10*x^2 + x + 1"},
	{1, "x^4 - x^3 + 3*x^2 + x - 2"},
	{4, "x^4 - 476*x^3 - 3*x^2 - x - 3"},
	{1, "x^4 - x^3 - x^2 + x - 1"},
	{12, "x^4 + x^3 - 53*x^2 - 1"},
	{1, "x^4 + 12*x^3 + 1"},
	{1, "x^4 + x^3 - 3*x^2 - x - 1"},
	{1, "x^4 + 12*x^2 - 2*x - 5"},
	{1, "x^4 - x^3 + 2*x^2 + 2*x + 12"},
	{1, "x^4 + 5*x^2 - 5*x + 1"},
	{1, "x^4 + 3*x^3 + x^2 - x + 3"},
	{1, "x^4 + 5*x^3 - x^2 + 5*x + 3"},
	{2, "x^4 - x^3 + 4*x^2 + x + 4"},
	{1, "x^4 + 3*x^3 - 14*x^2 - 2*x - 4"},
	{1, "x^4 - 11*x^2 + x + 1"},
	{1, "x^4 - x^3 + x^2 - x - 2"},
	{1, "x^4 + 3*x^3 - 5*x^2 - 3*x + 2"},
	{1, "x^4 + 3*x^3 - 3*x^2 - x - 4"},
	{1, "x^4 - 3*x^3 - 1"},
	{2, "x^4 + x^3 + 28*x^2 + x + 2"},
	{1, "x^4 + 11*x^3 + 6*x^2 - 65*x + 1"},
	{1, "x^4 + x^3 - 29*x^2 + x - 2"},
	{3, "x^4 - x^3 + 2*x^2 + 7*x - 8"},
	{1, "x^4 - 2*x^3 - 21*x^2 + 4*x - 1"},
	{1, "x^4 + 2*x^2 + x - 1"},
	{1, "x^4 - 20*x^3 - 4*x^2 - 8*x + 2"},
	{1, "x^4 + x^3 - 3*x^2 - x + 5"},
	{1, "x^4 + x^2 - 1"},
	{1, "x^4 + 5*x^3 - x^2 - x - 1"},
	{1, "x^4 - x^3 - x - 2"},
	{1, "x^4 + x^2 + x + 1"},
	{1, "x^4 - x^3 - x^2 - 5*x - 1"},
	{1, "x^4 + 7*x^3 + x^2 - 8*x - 15"},
	{2, "x^4 + 7*x^3 + x^2 + 11*x + 1"},
	{1, "x^4 + x^3 - 10*x^2 - x + 2"},
	{1, "x^4 - x^2 + x + 2"},
	{1, "x^4 + 2*x^3 + 12*x^2 - 2*x - 1"},
	{1, "x^4 - x^3 + 4*x^2 + 2"},
	{1, "x^4 - x^3 - 9*x^2 + 3*x - 2"},
	{1, "x^4 + x^3 + 8*x^2 + x - 2"},
	{1, "x^4 + x^3 - 5"},
	{1, "x^4 - 2*x^2 + 17*x + 2"},
	{1, "x^4 - x^3 + 3*x - 1"},
	{1, "x^4 + x^3 - 2*x^2 + 1"},
	{1, "x^4 - 2*x^3 - x^2 - 2*x - 1"},
	{2, "x^4 + 9*x^3 - x^2 + x + 1"},
	{1, "x^4 - 2*x^3 + 2*x + 1"},
	{1, "x^4 - 3*x^3 + x^2 + 7*x - 2"},
	{1, "x^4 - 2*x^3 + x^2 - 24*x - 1"},
	{1, "x^4 - 6*x^3 + 3*x^2 + x - 3"},
	{1, "x^4 + x^3 + 7*x^2 - 69*x + 4"},
	{1, "x^4 + 5*x^3 - x + 1"},
	{3, "x^4 + 6*x^3 - x^2 + x - 1"},
	{1, "x^4 + x^3 - 4*x - 1"},
	{1, "x^4 + 12*x^2 - 2*x + 7"},
	{1, "x^4 + 2*x^2 + 5*x - 1"}}