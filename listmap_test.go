package listmap

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_MyMap_Contains(t *testing.T) {
	Convey("Given a MyMap", t, func(c C) {

		var m *MyMap

		m = &MyMap{
			keys: []string{"a"},
			vals: []string{"a"},
		}
		So(m.Contains("a"), ShouldBeTrue)
		So(m.Contains("b"), ShouldBeFalse)

		m = &MyMap{
			keys: []string{"a", "d"},
			vals: []string{"1", "4"},
		}
		So(m.Contains("c"), ShouldBeFalse)
	})
}

func Test_MyMap_ToMap(t *testing.T) {
	Convey("Given a MyMap", t, func(c C) {

		var m *MyMap

		m = &MyMap{
			keys: []string{"a", "b", "d", "c"},
			vals: []string{"1", "2", "4", "3"},
		}

		So(m.ToMap(), ShouldResemble, map[string]string{
			"a": "1",
			"b": "2",
			"c": "3",
			"d": "4",
		})
	})
}

func Test_MyMap_Put(t *testing.T) {
	Convey("Given an empty MyMap", t, func(c C) {
		m := &MyMap{}

		m.Put("a", "2")
		So(m.keys, ShouldResemble, []string{"a"})
		So(m.vals, ShouldResemble, []string{"2"})

		m.Put("a", "1")
		So(m.keys, ShouldResemble, []string{"a"})
		So(m.vals, ShouldResemble, []string{"1"})

		m.Put("d", "4")
		So(m.keys, ShouldResemble, []string{"a", "d"})
		So(m.vals, ShouldResemble, []string{"1", "4"})

		m.Put("c", "3")
		So(m.keys, ShouldResemble, []string{"a", "c", "d"})
		So(m.vals, ShouldResemble, []string{"1", "3", "4"})
	})
}
