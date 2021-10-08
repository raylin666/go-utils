package validator

import "testing"

func TestValidate(t *testing.T) {
	valid := New("zh")
	s := valid.Validate(struct {
		Name string `label:"name" validate:"required,min=6"`
	}{
		Name: "cc",
	})

	if len(s) > 0 {
		t.Fatal(s)
	}

	t.Log("SUCCESS")
}