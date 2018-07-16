package subject

import "testing"

func TestGetSubject(t *testing.T) {
	expected := "Subject"
	if GetSubject() != expected {
		t.Errorf("GetSubject() should be equal %s\n", expected)
	}
}
