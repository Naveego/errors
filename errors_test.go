package errors

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNew(t *testing.T) {

	Convey("Calling errors.New('My Error Message')", t, func() {

		err := New("My Error Message")

		Convey("Should return an error of type errors.Error", func() {
			_, ok := err.(Error)
			So(ok, ShouldBeTrue)
		})

		Convey("Should set the message to 'My Error Message'", func() {
			So(err.Error(), ShouldEqual, "My Error Message")
		})

		Convey("Shuld set the default error code to 5000000", func() {
			e, _ := err.(Error)
			So(e.Code, ShouldEqual, 5000000)
		})

	})
}

func TestNewWithCode(t *testing.T) {

	Convey("Calling errors.NewWithCode(400001, 'My Error Message')", t, func() {

		err := NewWithCode(200001, "My Error Message")

		Convey("Should return an error of type errors.Error", func() {
			_, ok := err.(Error)
			So(ok, ShouldBeTrue)
		})

		Convey("Should set the message to 'My Error Message'", func() {
			So(err.Error(), ShouldEqual, "My Error Message")
		})

		Convey("Shuld set the default error code to 200001", func() {
			e, _ := err.(Error)
			So(e.Code, ShouldEqual, 200001)
		})

	})

}

func TestHTTPStatus(t *testing.T) {

	Convey("An error with code less than 400000", t, func() {
		Convey("Should return HTTP StStatusauts 500", func() {
			err := &Error{Code: 300, Message: "Test"}
			So(err.HTTPStatus(), ShouldEqual, 500)
			err = &Error{Code: 100, Message: "Test"}
			So(err.HTTPStatus(), ShouldEqual, 500)
			err = &Error{Code: 300001, Message: "Test"}
			So(err.HTTPStatus(), ShouldEqual, 500)
		})
	})

	Convey("An error with code greater than 4000000", t, func() {
		Convey("Should return the HTTP Status identified by the first 3 digits", func() {

			Convey("So an error code of 4000005 should return 400", func() {
				err := &Error{Code: 4000005, Message: "Test"}
				So(err.HTTPStatus(), ShouldEqual, 400)
			})

			Convey("So an error code of 4220005 should return 422", func() {
				err := &Error{Code: 4220005, Message: "Test"}
				So(err.HTTPStatus(), ShouldEqual, 422)
			})

			Convey("So an error code of 4150005 should return 415", func() {
				err := &Error{Code: 4150005, Message: "Test"}
				So(err.HTTPStatus(), ShouldEqual, 415)
			})

		})
	})

}
