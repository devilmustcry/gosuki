package fizzbuzz_test

import (
	"testing"

	"github.com/devilmustcry/gosuki/fizzbuzz"
)

func TestFizzBuzz(t *testing.T) {
	if fizzbuzz.FizzBuzz(1) != 1 {
		t.Errorf("FizzBuzz(1) should be equal 1")
	}
}

func TestFizzBuzz2(t *testing.T) {
	if fizzbuzz.FizzBuzz(2) != 2 {
		t.Errorf("FizzBuzz(2) should be equal 2")
	}
}
func TestFizzBuzz3(t *testing.T) {
	if fizzbuzz.FizzBuzz(3) != "fizz" {
		t.Errorf("FizzBuzz(3) should be equal 'fizz'")
	}
}
func TestFizzBuzz4(t *testing.T) {
	if fizzbuzz.FizzBuzz(4) != 4 {
		t.Errorf("FizzBuzz(4) should be equal 4")
	}
}
func TestFizzBuzz5(t *testing.T) {
	if fizzbuzz.FizzBuzz(5) != "buzz" {
		t.Errorf("FizzBuzz(5) should be equal 'buzz'")
	}
}
func TestFizzBuzz15(t *testing.T) {
	if fizzbuzz.FizzBuzz(15) != "fizzbuzz" {
		t.Errorf("FizzBuzz(15) should be equal 'fizzbuzz'")
	}
}
