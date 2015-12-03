package stringx_test

import (
	"testing"
	"github.com/drborges/rivers-contrib-stringx"
	"github.com/smartystreets/goconvey/convey"
	"github.com/smartystreets/assertions/should"
)

func TestStartsWith(t *testing.T) {

	convey.Convey("Given I have a slice of strings", t, func() {
		items := []string{"Borges", "Diego", "Diogo"}

		convey.Convey("When I filter items by prefix", func() {
			matched := stringx.From(items).StartingWith("Di").Collect()

			convey.Convey("Then the resulting slice is properly filtered", func() {
				convey.So(matched, should.Resemble, items[1:])
			})
		})

		convey.Convey("When I filter items by suffix", func() {
			matched := stringx.From(items).EndingWith("ges").Collect()

			convey.Convey("Then the resulting slice is properly filtered", func() {
				convey.So(matched, should.Resemble, items[:1])
			})
		})

		convey.Convey("When I filter items by regex", func() {
			matched := stringx.From(items).Matching(`Di.go`).Collect()

			convey.Convey("Then the resulting slice is properly filtered", func() {
				convey.So(matched, should.Resemble, items[1:])
			})
		})

		convey.Convey("When I filter items by their length", func() {
			matched := stringx.From(items).WithLength(5).Collect()

			convey.Convey("Then the resulting slice is properly filtered", func() {
				convey.So(matched, should.Resemble, items[1:])
			})
		})

		convey.Convey("When I apply to lower operation", func() {
			matched := stringx.From(items).ToLower().Collect()

			convey.Convey("Then the resulting slice is properly mapped", func() {
				convey.So(matched, should.Resemble, []string{"borges", "diego", "diogo"})
			})
		})

		convey.Convey("When I apply to upper operation", func() {
			matched := stringx.From(items).ToUpper().Collect()

			convey.Convey("Then the resulting slice is properly mapped", func() {
				convey.So(matched, should.Resemble, []string{"BORGES", "DIEGO", "DIOGO"})
			})
		})

		convey.Convey("When I apply title operation", func() {
			matched := stringx.From(items).ToLower().Title().Collect()

			convey.Convey("Then the resulting slice is properly mapped", func() {
				convey.So(matched, should.Resemble, items)
			})
		})

		convey.Convey("When I apply replace operation", func() {
			matched := stringx.From(items).Replace("Di", "iD").Collect()

			convey.Convey("Then the resulting slice is properly mapped", func() {
				convey.So(matched, should.Resemble, []string{"Borges", "iDego", "iDogo"})
			})
		})

		convey.Convey("When I apply prepend operation", func() {
			matched := stringx.From(items).Prepend("_").Collect()

			convey.Convey("Then the resulting slice is properly mapped", func() {
				convey.So(matched, should.Resemble, []string{"_Borges", "_Diego", "_Diogo"})
			})
		})

		convey.Convey("When I apply append operation", func() {
			matched := stringx.From(items).Append("_").Collect()

			convey.Convey("Then the resulting slice is properly mapped", func() {
				convey.So(matched, should.Resemble, []string{"Borges_", "Diego_", "Diogo_"})
			})
		})

		convey.Convey("When I apply trim operation", func() {
			matched := stringx.From(items).Append(" ").Prepend(" ").Trim().Collect()

			convey.Convey("Then the resulting slice is properly mapped", func() {
				convey.So(matched, should.Resemble, items)
			})
		})

		convey.Convey("When I apply split operation", func() {
			matched := stringx.From([]string{"Diego "}).Split().Collect()

			convey.Convey("Then the resulting slice is properly mapped", func() {
				convey.So(matched, should.Resemble, []string{"D", "i", "e", "g", "o", " "})
			})
		})

		convey.Convey("When I apply split by operation", func() {
			matched := stringx.From([]string{"Diego Borges"}).SplitBy(" ").Collect()

			convey.Convey("Then the resulting slice is properly mapped", func() {
				convey.So(matched, should.Resemble, []string{"Diego", "Borges"})
			})
		})
	})
}
