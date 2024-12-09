package internal

import (
	_ "embed"
)

//go:embed index.html
var IndexHTML []byte

//go:embed main.js
var MainJS []byte

//go:embed sw.js
var WorkerJS []byte
