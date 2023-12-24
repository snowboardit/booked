package reserved_test

import (
	"slices"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/snowboardit/reserved/pkg/reserved"
)

var (
	languages = []string{
		"java (jvm)",
		"python (py)",
		"javascript (js, jsx, node)",
		"c-sharp (cs, csharp)",
		"c-plus-plus (cpp, c++)",
		"c",
		"typescript (ts, tsx)",
		"ruby (rb)",
		"swift",
		"kotlin",
		"go",
		"rust",
		"php",
		"dart",
		"objective-c (objc, obj-c)",
		"scala",
		"perl",
		"lua",
		"haskell",
		"mysql",
		"postgres (pg, psql)",
		"sql-server",
		"sql-lite (sqlite)",
		"mongodb (mongo)",
		"redis (rd)",
		"oracle",
		"cassandra",
		"mariadb (maria)",
	}
	testWord  = "const"
	testWords = []string{
		"const",
		"let",
		"var",
	}
)
var _ = Describe("Reserved", func() {
	It("should list all languages", func() {
		r := reserved.New()
		langs := slices.Clip(r.Languages())
		Expect(langs).To(Equal(languages))
	})

	It("should return all languages that a test word is reserved in", func() {
		r := reserved.New()
		checked := r.Check(testWord)

		Expect(checked).To(HaveKeyWithValue("typescript", ContainElement(testWord)))
		Expect(checked).To(HaveKeyWithValue("java", ContainElement(testWord)))
		Expect(checked).To(HaveKeyWithValue("c-sharp", ContainElement(testWord)))
		Expect(checked).To(HaveKeyWithValue("c-plus-plus", ContainElement(testWord)))
		Expect(checked).To(HaveKeyWithValue("go", ContainElement(testWord)))
		Expect(checked).To(HaveKeyWithValue("rust", ContainElement(testWord)))
		Expect(checked).To(HaveKeyWithValue("php", ContainElement(testWord)))
		Expect(checked).To(HaveKeyWithValue("dart", ContainElement(testWord)))
		Expect(checked).To(HaveKeyWithValue("objective-c", ContainElement(testWord)))
		Expect(checked).To(HaveKeyWithValue("javascript", ContainElement(testWord)))
	})

	It("should return all languages that test words are reserved in", func() {
		r := reserved.New()
		checked := r.Check(testWords...)

		Expect(checked).To(HaveKeyWithValue("c", []string{testWords[0]}))
		Expect(checked).To(HaveKeyWithValue("go", []string{testWords[0], testWords[2]}))
		Expect(checked).To(HaveKeyWithValue("javascript", testWords))
		Expect(checked).To(HaveKeyWithValue("rust", testWords[:2]))
		Expect(checked).To(HaveKeyWithValue("haskell", []string{testWords[1]}))
		Expect(checked).To(HaveKeyWithValue("c-sharp", []string{testWords[0]}))
		Expect(checked).To(HaveKeyWithValue("typescript", testWords))
		Expect(checked).To(HaveKeyWithValue("kotlin", []string{testWords[2]}))
		Expect(checked).To(HaveKeyWithValue("c-plus-plus", []string{testWords[0]}))
		Expect(checked).To(HaveKeyWithValue("swift", testWords[1:]))
		Expect(checked).To(HaveKeyWithValue("php", []string{testWords[0], testWords[2]}))
		Expect(checked).To(HaveKeyWithValue("scala", []string{testWords[2]}))
		Expect(checked).To(HaveKeyWithValue("dart", []string{testWords[0], testWords[2]}))
		Expect(checked).To(HaveKeyWithValue("objective-c", []string{testWords[0]}))
	})
})

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Reserved Suite")
}
