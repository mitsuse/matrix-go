package rewriters

import (
	"bytes"
	"testing"
)

type testPair struct {
	X int
	Y int
}

func TestReflect(t *testing.T) {
	test := testPair{
		X: 1,
		Y: 2,
	}

	if x, y := Reflect().Rewrite(test.X, test.Y); x == test.X && y == test.Y {
		return
	}

	t.Fatal("The result pair should equal to the input pair.")
}

func TestReverse(t *testing.T) {
	test := testPair{
		X: 1,
		Y: 2,
	}

	if x, y := Reverse().Rewrite(test.X, test.Y); x == test.Y && y == test.X {
		return
	}

	t.Fatal("The result pair should be reversed.")
}

func TestReflectSerialize(t *testing.T) {
	writer := bytes.NewBuffer([]byte{})

	if err := Reflect().Serialize(writer); err != nil {
		t.Fatalf("An expected error occured on serialization: %s", err)
	}

	reader := bytes.NewReader(writer.Bytes())

	rewriter, err := Deserialize(reader)

	if err != nil {
		t.Fatalf("An expected error occured on deserialization: %s", err)
	}

	if rewriter != Reflect() {
		t.Fatal("Deserialization failed for a serialized rewriter.")
	}
}

func TestReverseSerialize(t *testing.T) {
	writer := bytes.NewBuffer([]byte{})

	if err := Reverse().Serialize(writer); err != nil {
		t.Fatalf("An expected error occured on serialization: %s", err)
	}

	reader := bytes.NewReader(writer.Bytes())

	rewriter, err := Deserialize(reader)

	if err != nil {
		t.Fatalf("An expected error occured on deserialization: %s", err)
	}

	if rewriter != Reverse() {
		t.Fatal("Deserialization failed for a serialized rewriter.")
	}
}

func TestDeserializeFailsForIncompatibleId(t *testing.T) {
	test := []byte("...")
	test = append(test, version)
	test = append(test, typeReflect)

	reader := bytes.NewReader(test)

	if _, err := Deserialize(reader); err == nil {
		t.Fatal("The incompatible id should cause an error on deserialization.")
	}
}

func TestDeserializeFailsForIncompatibleVersion(t *testing.T) {
	test := []byte(id)
	test = append(test, 255)
	test = append(test, typeReflect)

	reader := bytes.NewReader(test)

	if _, err := Deserialize(reader); err == nil {
		t.Fatal("The incompatible version should cause an error on deserialization.")
	}
}

func TestDeserializeFailsForUnknownRewriter(t *testing.T) {
	test := []byte(id)
	test = append(test, version)
	test = append(test, 255)

	reader := bytes.NewReader(test)

	if _, err := Deserialize(reader); err == nil {
		t.Fatal("The unknown rewriter should cause an error on deserialization.")
	}
}

func TestReflectTranspose(t *testing.T) {
	if Reflect().Transpose() == Reverse() {
		return
	}

	t.Fatal("The transpose of Reflect should be Reverse.")
}

func TestReverseTranspose(t *testing.T) {
	if Reverse().Transpose() == Reflect() {
		return
	}

	t.Fatal("The transpose of Reverse should be Reflect.")
}
