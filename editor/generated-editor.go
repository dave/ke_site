package main

import (
	_ "github.com/dave/ke_common/images"
	_ "github.com/dave/ke_common/images/types"
	_ "github.com/dave/ke_common/units"
	_ "github.com/dave/ke_common/units/types"
	_ "github.com/dave/ke_common/words"
	_ "github.com/dave/ke_common/words/types"
	_ "github.com/dave/ke_site"
	_ "github.com/dave/ke_site/types"
	"kego.io/process/editor"
	_ "kego.io/system"
	_ "kego.io/system/types"
)

func main() {
	editor.Edit("github.com/dave/ke_site")
}
