package doc_test

import (
	"bytes"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func ExampleGenManTree() {
	cmd := &cobra.Command{
		Use:   "test",
		Short: "my test program",
	}
	header := &doc.GenManHeader{
		Title:   "MINE",
		Section: "3",
	}
	doc.GenManTree(cmd.TemplateData(), header, "/tmp")
}

func ExampleGenMan() {
	cmd := &cobra.Command{
		Use:   "test",
		Short: "my test program",
	}
	header := &doc.GenManHeader{
		Title:   "MINE",
		Section: "3",
	}
	out := new(bytes.Buffer)
	doc.GenMan(cmd.TemplateData(), header, out)
	fmt.Print(out.String())
}
