package bdd

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// $GOPATH/bin/goconvey web界面
func Test(t *testing.T) {
	Convey("Give 2 even numbers", t, func() {
		a := 2
		b := 4
		Convey("When add the two numbers", func() {
			c := a + b
			Convey("then the result is still even", func() {
				So(c%2, ShouldEqual, 0) // 断言
			})
		})
	})
}
