package window

import (
	"crypto/rand"
	"fmt"
	"log"
	"net"
	"net/http"
	"runtime"
	"sync"

	// "github.com/skip2/go-qrcode"
	"github.com/gobuffalo/packr"
	"github.com/zserge/lorca"
)

type counter struct {
	sync.Mutex
	count int
}

func (c *counter) Add(n int) {
	c.Lock()
	defer c.Unlock()
	c.count = c.count + n
}

func (c *counter) Value() int {
	c.Lock()
	defer c.Unlock()
	return c.count
}

// ^ Opens a Window For QR Code ^ //
func OpenQRWindow(json string) {
	// Encode to QR
	print(json)
	// //qrData, err := qrcode.Encode(json, qrcode.Medium, 256)
	// err := qrcode.WriteFile(json, qrcode.Medium, 256, "qrcode.png")
	// if err != nil {
	// 	log.Panicln(err)
	// }

	args := []string{}
	if runtime.GOOS == "linux" {
		args = append(args, "--class=Lorca")
	}
	ui, err := lorca.New("", "", 480, 320, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	// A simple way to know when UI is ready (uses body.onload event in JS)
	err = ui.Bind("start", func() {
		log.Println("UI is ready")
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create and bind Go object to the UI
	c := &counter{}
	ui.Bind("counterAdd", c.Add)
	ui.Bind("counterValue", c.Value)

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	box := packr.NewBox("../assets/www")
	go http.Serve(ln, http.FileServer(box))
	ui.Load(fmt.Sprintf("http://%s", ln.Addr()))

	// You may use console.log to debug your JS code, it will be printed via
	// log.Println(). Also exceptions are printed in a similar manner.
	ui.Eval(`
		console.log("Hello, world!");
		console.log('Multiple values:', [1, false, {"x":5}]);
	`)
	defer ui.Close()
	// Wait until UI window is closed
	<-ui.Done()
}

// ^ Generates UUID ^ //
func UUID() string {
	// Generate UUID
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Panicln(err)
	}

	// Format String
	return fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
