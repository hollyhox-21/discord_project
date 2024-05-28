package pkg

import "embed"

// SwaggerFile holds our static swagger file.
//
//go:embed */*.swagger.json
var SwaggerFile embed.FS
