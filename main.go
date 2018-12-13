package main

import (
	"fmt"

	assertion "internal/assertion"
	types "types"
)

type TgoAssertion interface {
	Should(matcher types.TgaMatcher, optionalDescription ...interface{}) bool
	ShouldNot(matcher types.TgaMatcher, optionalDescription ...interface{}) bool

	To(matcher types.TgaMatcher, optionalDescription ...interface{}) bool
	ToNot(matcher types.TgaMatcher, optionalDescription ...interface{}) bool
	NotTo(matcher types.TgaMatcher, optionalDescription ...interface{}) bool
}

func Assert(expr interface{}) {
	fmt.Println(expr.(bool))
}

func Describe(text string, body func()) {
	fmt.Println("Describe: ", text)
	body()
}

func It(text string, body func()) {
	fmt.Println("It: ", text)
	body()
}

func ExpectWithOffset(offset int, actual interface{}, extra ...interface{}) TgoAssertion {
	return assertion.New(actual, offset, extra...)
}

func Expect(actual interface{}, extra ...interface{}) TgoAssertion {
	return ExpectWithOffset(0, actual, extra...)
}

func Before(body interface{})     {}
func BeforeEach(body interface{}) {}
func After(body interface{})      {}
func AfterEach(body interface{})  {}

func main() {
	fmt.Println("vim-go")

	_ = assertion.New("", 0)

	Describe("test 1", func() {
		It("case 1", func() {
			fmt.Println("exec")
			Assert(1 == 2)
		})
	})
}
