package dist

import "embed"

// Dist holds our static swagger content.
//
//go:embed *
var Dist embed.FS
