package captcha

import (
	"testing"
)

func TestCaptcha(t *testing.T) {
	r := Random{10, 20, 0}
	expected := GenerateCaptcha(r)
	if expected != 30 {
		t.Errorf("Test Error should equal %d", expected)
	}
}
