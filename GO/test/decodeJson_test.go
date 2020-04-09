package main

import (
	"reflect"
	"testing"
)

func TestDecode(t *testing.T) {
	branch := Branch{
		ID:         1,
		BranchName: "officeA",
		Prefecture: "Tokyo",
		Address:    "Shinjuku 1-2-3, Tokyo",
	}

	actual := Person{
		ID:        1,
		Branch:    branch,
		FirstName: "Keisuke",
		LastName:  "AAA",
		Age:       29,
		Sex:       0,
	}

	// wrongExpected := Person{
	// 	ID:        1,
	// 	Branch:    branch,
	// 	FirstName: "Keisuke",
	// 	LastName:  "BBB",
	// 	Age:       29,
	// 	Sex:       0,
	// }

	err := (&actual).encode("test.json")
	if err != nil {
		t.Error(err)
	}

	expected, err := unmarshal("test.json")
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nexpected : %+v \nbut was  : %+v", expected, actual)
	}
	// if !reflect.DeepEqual(wrongExpected, actual) {
	// 	t.Errorf("\nexpected : %+v \nbut was  : %+v", wrongExpected, actual)
	// }
}

func BenchmarkUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unmarshal("test.json")
	}
}

func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		decode("test.json")
	}
}

func TestEncode(t *testing.T) {
	t.Skip("Skip encoding for now")
}
