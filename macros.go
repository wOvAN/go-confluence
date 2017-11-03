// macros
package confluence

import (
	"fmt"
)

const (
	macroHTML = `
	    <ac:structured-macro ac:name="html">
            <ac:plain-text-body><![CDATA[%s]]>
			</ac:plain-text-body>
        </ac:structured-macro>`
)

// Inserts RAW HTML, JS etc.
// https://confluence.atlassian.com/conf55/confluence-user-s-guide/creating-content/using-the-editor/working-with-macros/html-macro
func MacroHTML(aHTML string) string {
	return fmt.Sprintf(macroHTML, aHTML)
}
