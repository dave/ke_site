package ke_site

import (
	images "github.com/davelondon/ke_common/images"
	units "github.com/davelondon/ke_common/units"
	words "github.com/davelondon/ke_common/words"
	system "kego.io/system"
)

var ptr8729020288 = &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "Rule to enforce this object always has a .uk server", Rules: []system.Rule(nil)}
var ptr8729429472 = &system.RuleBase{Selector: "{images:photo}>.server"}
var ptr8728723456 = &system.String_rule{Base: ptr8729020288, RuleBase: ptr8729429472, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{Value: "\\.uk$", Exists: true}, Format: system.String{}}
var ptr8728739840 = &system.Base{Type: system.Reference{Package: "github.com/davelondon/ke_site", Name: "gallery", Exists: true}, Id: system.Reference{Package: "github.com/davelondon/ke_site", Name: "site", Exists: true}, Description: "This is an instance of the gallery type containing two nested images", Rules: []system.Rule{ptr8728723456}}
var ptr8729020160 = &system.Base{Type: system.Reference{Package: "github.com/davelondon/ke_common/images", Name: "photo", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
var ptr8729019648 = &system.Base{Type: system.Reference{Package: "github.com/davelondon/ke_common/units", Name: "rectangle", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
var ptr8729191312 = &units.Rectangle{Base: ptr8729019648, Height: system.Number{Value: 600, Exists: true}, Width: system.Number{Value: 800, Exists: true}}
var ptr8728887296 = &images.Photo{Base: ptr8729020160, Protocol: system.String{Value: "https", Exists: true}, Server: system.String{Value: "www.brophy.uk", Exists: true}, Size: ptr8729191312, Path: system.String{Value: "/me.jpg", Exists: true}}
var ptr8729020416 = &system.Base{Type: system.Reference{Package: "github.com/davelondon/ke_common/words", Name: "translation", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
var ptr8729191984 = &words.Translation{Base: ptr8729020416, Translations: map[string]system.String{"es": system.String{Value: "Hola mundo.", Exists: true}, "fr": system.String{Value: "Bonjour le monde.", Exists: true}}, English: system.String{Value: "Hello, world.", Exists: true}}
var ptr8729428224 = &Gallery{Base: ptr8728739840, Images: map[string]images.Image{"dave": ptr8728887296}, Title: ptr8729191984}
var Site = ptr8729428224
