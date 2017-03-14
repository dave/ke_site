package main

import (
	"fmt"
	"os"

	_ "github.com/dave/ke_common/images"
	_ "github.com/dave/ke_common/images/types"
	_ "github.com/dave/ke_common/units"
	_ "github.com/dave/ke_common/units/types"
	_ "github.com/dave/ke_common/words"
	_ "github.com/dave/ke_common/words/types"
	_ "github.com/dave/ke_site"
	_ "github.com/dave/ke_site/types"
	"kego.io/editor/server"
	"kego.io/process"
	_ "kego.io/system"
	_ "kego.io/system/types"
)

// This starts the local editor server and opens a browser. This
// simulates the ke command without running it.
func main() {
	if err := server.Start("github.com/dave/ke_site", true); err != nil {
		fmt.Println(process.FormatError(err))
		os.Exit(1)
	}
}
