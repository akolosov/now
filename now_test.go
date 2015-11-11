package now

import (
  "time"
  "testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var format = "2006-01-02 15:04:05.999999999"

func TestNow(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Now Suite")
}

var _ = Describe("now", func() {

  Describe ("#TestBeginningOff", func() {
    n := time.Date(2013, 11, 18, 17, 51, 49, 123456789, time.UTC)

    It ("Should return right values",  func(){

      Expect(NewNow(n).BeginningOfMinute().Format(format)).Should(Equal("2013-11-18 17:51:00"))
      Expect(BeginningOfMinute()).ShouldNot(BeNil())

      Expect(NewNow(n).BeginningOfHour().Format(format)).Should(Equal("2013-11-18 17:00:00"))
      Expect(BeginningOfHour()).ShouldNot(BeNil())

      FirstDayMonday = true
      Expect(NewNow(n).BeginningOfDay().Format(format)).Should(Equal("2013-11-18 00:00:00"))
      Expect(BeginningOfDay()).ShouldNot(BeNil())

      FirstDayMonday = false
      Expect(NewNow(n).BeginningOfWeek().Format(format)).Should(Equal("2013-11-17 00:00:00"))

      FirstDayMonday = true
      Expect(NewNow(n).BeginningOfWeek().Format(format)).Should(Equal("2013-11-18 00:00:00"))
      Expect(BeginningOfWeek()).ShouldNot(BeNil())
      FirstDayMonday = false

      Expect(NewNow(n).BeginningOfMonth().Format(format)).Should(Equal("2013-11-01 00:00:00"))
      Expect(BeginningOfMonth()).ShouldNot(BeNil())

      Expect(NewNow(n).BeginningOfYear().Format(format)).Should(Equal("2013-01-01 00:00:00"))
      Expect(BeginningOfYear()).ShouldNot(BeNil())

      n = time.Date(2013, 11, 17, 17, 51, 49, 123456789, time.UTC)
      FirstDayMonday = true
      Expect(NewNow(n).BeginningOfWeek().Format(format)).Should(Equal("2013-11-11 00:00:00"))
    })
  })

  Describe ("#TestEndOff", func() {
    n := time.Date(2013, 11, 18, 17, 51, 49, 123456789, time.UTC)
    n1 := time.Date(2013, 02, 18, 17, 51, 49, 123456789, time.UTC)
    n2 := time.Date(1900, 02, 18, 17, 51, 49, 123456789, time.UTC)

    It ("Should return right values", func(){
      Expect(NewNow(n).EndOfMinute().Format(format)).Should(Equal("2013-11-18 17:51:59.999999999"))
      Expect(EndOfMinute()).ShouldNot(BeNil())

      Expect(NewNow(n).EndOfHour().Format(format)).Should(Equal("2013-11-18 17:59:59.999999999"))
      Expect(EndOfHour()).ShouldNot(BeNil())

      Expect(NewNow(n).EndOfDay().Format(format)).Should(Equal("2013-11-18 23:59:59.999999999"))
      Expect(EndOfDay()).ShouldNot(BeNil())

      FirstDayMonday = true
      Expect(NewNow(n).EndOfWeek().Format(format)).Should(Equal("2013-11-24 23:59:59.999999999"))
      Expect(EndOfWeek()).ShouldNot(BeNil())

      FirstDayMonday = false
      Expect(NewNow(n).EndOfWeek().Format(format)).Should(Equal("2013-11-23 23:59:59.999999999"))

      Expect(NewNow(n).EndOfMonth().Format(format)).Should(Equal("2013-11-30 23:59:59.999999999"))
      Expect(EndOfMonth()).ShouldNot(BeNil())

      Expect(NewNow(n).EndOfYear().Format(format)).Should(Equal("2013-12-31 23:59:59.999999999"))
      Expect(EndOfYear()).ShouldNot(BeNil())

      Expect(NewNow(n1).EndOfMonth().Format(format)).Should(Equal("2013-02-28 23:59:59.999999999"))

      Expect(NewNow(n2).EndOfMonth().Format(format)).Should(Equal("1900-02-28 23:59:59.999999999"))
    })
  })

  Describe ("#TestEndOff", func() {
    n := time.Date(2013, 11, 19, 17, 51, 49, 123456789, time.UTC)

    It ("Should return right values", func() {
      Expect(NewNow(n).NextDay().Format(format)).Should(Equal("2013-11-20 17:51:49.123456789"))
      Expect(NextDay()).ShouldNot(BeNil())

      Expect(NewNow(n).PrevDay().Format(format)).Should(Equal("2013-11-18 17:51:49.123456789"))
      Expect(PrevDay()).ShouldNot(BeNil())
    })
  })

  Describe("#TestMonthLength", func() {
    It ("Should return right values", func() {
      Expect(NewNow(time.Date(2012, 2, 1, 0, 0, 0, 0, time.UTC)).MonthLength()).Should(Equal(29))
      Expect(NewNow(time.Date(2011, 2, 1, 0, 0, 0, 0, time.UTC)).MonthLength()).Should(Equal(28))
      Expect(MonthLength()).ShouldNot(BeNil())
    })
  })

  Describe("#TestParse", func() {
    n := time.Date(2013, 11, 18, 17, 51, 49, 123456789, time.UTC)

    It ("Should return right values", func() {
      Expect(NewNow(n).MustParse("2002-10-12 22:14").Format(format)).Should(Equal("2002-10-12 22:14:00"))
      Expect(MustParse("2002-10-12 22:14")).ShouldNot(BeNil())

      Expect(NewNow(n).MustParse("20102-101-122 222:134").Format(format)).Should(Equal("0001-01-01 00:00:00"))
      Expect(MustParse("20102-101-122 222:134")).Should(Equal(time.Time{}))

      val, err := parseWithFormat("2002-10-122 22:14")
      Expect(err).ShouldNot(BeNil())

      val, err = Parse("2002-10-12 22:14")
      Expect(err).Should(BeNil())
      Expect(val).ShouldNot(BeNil())

      Expect(NewNow(n).MustParse("2002-10-12 02:04").Format(format)).Should(Equal("2002-10-12 02:04:00"))

      Expect(NewNow(n).MustParse("2002-10-12 22:14:56").Format(format)).Should(Equal("2002-10-12 22:14:56"))

      Expect(NewNow(n).MustParse("2002-10-12").Format(format)).Should(Equal("2002-10-12 00:00:00"))

      Expect(NewNow(n).MustParse("18:20").Format(format)).Should(Equal("2013-11-18 18:20:00"))

      Expect(NewNow(n).MustParse("18:20:39").Format(format)).Should(Equal("2013-11-18 18:20:39"))
    })
  })
})
