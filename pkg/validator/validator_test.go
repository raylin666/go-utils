package validator

import "testing"

func TestValidate(t *testing.T) {
	valid := New("zh", "label")
	s := valid.Validate(struct {
		Name string `label:"name" validate:"required,min=6"`
	}{
		Name: "raylin666",
	})

	if len(s) > 0 {
		t.Fatal(s)
	}

	t.Log("SUCCESS")
}