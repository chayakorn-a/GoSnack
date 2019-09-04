package tdd

import "testing"

// This is unit test

// 1. file with postfix "_test.go" is test file
// 2. Test func has to start with preFix "Test"
// 3. After "Test" have to start with upper case. it's camel case
// 4. Signature of function test is parameter always be *testing.T
func TestInput1ShouldBe1(t *testing.T) {
	actual := FizzBuzz(1)
	expected := "1"
	if actual != expected {
		t.Error("Expected :", expected, " but got", actual)
	}
}

func TestInput2ShouldBe2(t *testing.T) {
	actual := FizzBuzz(2)
	expected := "2"
	if actual != expected {
		t.Error("Expected :", expected, " but got", actual)
	}
}

func TestInput3ShouldBeFizz(t *testing.T) {
	actual := FizzBuzz(3)
	expected := "Fizz"
	if actual != expected {
		t.Error("Expected :", expected, " but got", actual)
	}
}

func TestInput4ShouldBe4(t *testing.T) {
	actual := FizzBuzz(4)
	expected := "4"
	if actual != expected {
		t.Error("Expected :", expected, " but got", actual)
	}
}

func TestInput5ShouldBeBuzz(t *testing.T) {
	actual := FizzBuzz(5)
	expected := "Buzz"
	if actual != expected {
		t.Error("Expected :", expected, " but got", actual)
	}
}

func TestInput6ShouldBeFizz(t *testing.T) {
	actual := FizzBuzz(6)
	expected := "Fizz"
	if actual != expected {
		t.Error("Expected :", expected, " but got", actual)
	}
}
func TestInput7ShouldBe7(t *testing.T) {
	actual := FizzBuzz(7)
	expected := "7"
	if actual != expected {
		t.Error("Expected :", expected, " but got", actual)
	}
}
func TestInput8ShouldBe8(t *testing.T) {
	actual := FizzBuzz(8)
	expected := "8"
	if actual != expected {
		t.Error("Expected :", expected, " but got", actual)
	}
}

// triple A rule for test
// Arrange เตรียมของ, Action สิ่งที่ทำ, Assertion สิ่งที่คาดหวัง
func TestInput9ShouldBeFizz(t *testing.T) {
	// Arrange
	var actual string

	// Action
	actual = FizzBuzz(9)

	// Assertion
	expected := "Fizz"
	if actual != expected {
		t.Error("Expected :", expected, " but got", actual)
	}
}

/* Test case from Teacher
func TestFizzBuzz(t *testing.T) {
 testcases := []struct {
  input    int
  expected string
 }{
  {1, "1"},
  {2, "2"},
  {3, "Fizz"},
  {4, "4"},
  {5, "Buzz"},
  {6, "Fizz"},
  {7, "7"},
  {8, "8"},
  {9, "Fizz"},
  {10, "Buzz"},
  {11, "11"},
  {12, "Fizz"},
  {13, "13"},
  {14, "14"},
  {15, "FizzBuzz"},
 }

 for _, v := range testcases {
  actual := FizzBuzz(v.input)
  if actual != v.expected {
   t.Error("expected :", v.expected, " but got ", actual)
  }
 }
}*/
