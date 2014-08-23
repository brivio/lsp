package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode/utf8"

	c "github.com/mitchellh/colorstring"
)

const (
	briefcaseRune     = '💼'
	gitRune           = '😻'
	musicRune         = '🎼'
	pythonRune        = '🐍'
	javaRune          = '🍵'
	documentRune      = '📄'
	commonPrefix      = "[blue]./"
	descriptionIndent = "                "
	columnSize        = 30 // characters in the filename column
	maxFileNameSize   = columnSize - 7
)

func render() {
	sort.Sort(byType(FileList))
	fmt.Printf("\n") // i like empty line before the list

	// summary
	fmt.Printf(c.Color("\n[cyan]" + strings.Repeat("-", 3*columnSize) + "\n"))
	fmt.Printf(c.Color("    lsp \"[red]%s[white]\"\n"), mode.targetPath)
	fmt.Printf(c.Color("     [red]%v[white] files, [red]%v[white] directories \n\n"), len(FileList), len(Trie.Ch["dirs"].Fls))
}

func renderFiles(fls []*FileInfo) {
	for _, fl := range fls {
		displayFileName := fl.f.Name()
		if utf8.RuneCount([]byte(displayFileName)) > maxFileNameSize {
			displayFileName = string([]rune(displayFileName)[0:maxFileNameSize]) + "[magenta][...]"
		}

		//indent
		if indentSize := columnSize - utf8.RuneCount([]byte(displayFileName)); indentSize > 0 {
			fmt.Printf(strings.Repeat(" ", indentSize) + "") // indent
		}

		fmt.Printf(c.Color(commonPrefix + fmt.Sprintf("[white]%s[blue]", displayFileName))) // column 1

		// central dividing space
		fmt.Printf("  ")

		fmt.Printf(c.Color(fmt.Sprintf("[red]%s[white]\n", fl.description))) // column 2
		// if fl.description != "" {
		// 	fmt.Printf(c.Color(fmt.Sprintf("[blue]%s[white]\n", fl.description))) // description line
		// }
	}
}

func printCentered(o string) {
	length := utf8.RuneCount([]byte(o))
	sideburns := (6 + 2*columnSize - length) / 2
	fmt.Printf(strings.Repeat(" ", sideburns))
	fmt.Printf(c.Color("[red]"+o+"[white]") + "\n")

}
