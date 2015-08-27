package main

import (
	_ "github.com/davelondon/ke_common/images"
	_ "github.com/davelondon/ke_common/images/types"
	_ "github.com/davelondon/ke_common/units"
	_ "github.com/davelondon/ke_common/units/types"
	_ "github.com/davelondon/ke_common/words"
	_ "github.com/davelondon/ke_common/words/types"
	_ "github.com/davelondon/ke_site"
	_ "github.com/davelondon/ke_site/types"
	"honnef.co/go/js/console"
	"kego.io/editor"
	_ "kego.io/system"
	_ "kego.io/system/types"
)

func main() {
	if err := editor.Edit("github.com/davelondon/ke_site"); err != nil {
		console.Error(err)
	}
}
