package static

import (
	_ "embed"
)

//go:embed index.html
var IndexHTML []byte

//go:embed sw.js
var WorkerJS []byte
