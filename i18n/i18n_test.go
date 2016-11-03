package i18n

import "testing"

func TestLoad(t *testing.T) {
	dir := `F:\workspace\go\web\resources\lang`
	err := Load(dir)
	if err != nil {
		t.Fatal(err)
	}
}
