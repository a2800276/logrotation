package logrotation

import (
	"testing"
	"time"
)

func TestFn(t *testing.T) {
	lr := Logrotation{
		BaseFilename: "test",
		Suffix:       "log",
		BaseDir:      "something",
	}
	now, err := time.Parse("20060102150405", "20060102150405")
	if err != nil {
		t.Fatalf("%v", err)
	}
	is := lr.makeFN(now)
	should := "something/test.20060102-150405.log"
	if is != should {
		t.Fatalf("should: %s is: %s", should, is)
	}

}

func TestFnDefaultDir(t *testing.T) {

	lr := Logrotation{
		BaseFilename: "test",
		Suffix:       "log",
	}
	now, err := time.Parse("20060102150405", "20060102150405")
	if err != nil {
		t.Fatalf("%v", err)
	}
	is := lr.makeFN(now)
	should := "./test.20060102-150405.log"
	if is != should {
		t.Fatalf("should: %s is: %s", should, is)
	}
}
func TestFnDateTree(t *testing.T) {

	lr := Logrotation{
		BaseFilename: "test",
		Suffix:       "log",
		UseDateTree:  true,
	}
	now, _ := time.Parse("20060102150405", "20060102150405")
	is := lr.makePathString(now)
	should := "./2006/01/02"
	if is != should {
		t.Fatalf("should: %s is: %s", should, is)

	}
}
