package cmd

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"gopkg.in/resty.v1"

	"github.com/sacOO7/gowebsocket"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"

	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/keyboard"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/terminal/termbox"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgets/donut"
	"github.com/mum4k/termdash/widgets/text"
)

const (
	// donut widget constants
	playTypePercent playType = iota
	playTypeAbsolute
)

// optional port variable. example: `gex -p 30057`
var givenPort = flag.Int("p", 26657, "port to connect")
var givenHost = flag.String("h", "localhost", "host to connect")
var ssl = flag.Bool("s", false, "use SSL for connection")

// Info describes a list of types with data that are used in the explorer
type Info struct {
	blocks       *Blocks
	transactions *Transactions
}

// Blocks describe content that gets parsed for block
type Blocks struct {
	amount               int
	secondsPassed        int
	totalGasWanted       int64
	gasWantedLatestBlock int64
	maxGasWanted         int64
	lastTx               int64
}

// Transactions describe content that gets parsed for transactions
type Transactions struct {
	amount uint64
}

// playType indicates the type of the donut widget.
type playType int

// CreateGexCmd creates the gex command for cobra
func CreateGexCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "gex",
		Short: "GEX is a terminal explorer for the Cosmos SDK",
		Long: `GEX is a terminal explorer for the Cosmos SDK.
It allows you to explore blocks and transactions in real time.`,
		Run: func(_ *cobra.Command, _ []string) {
			Start()
		},
	}
}

// Start starts the gex command
func Start() {
	view()

	// Init internal variables
	info := Info{}
	info.blocks = new(Blocks)
	info.transactions = new(Transactions)

	connectionSignal := make(chan string)
	t, err := termbox.New()
	if err != nil {
		panic(err)
	}
	defer t.Close()

	flag.Parse()

	networkInfo, err := getFromRPC("status")
	if err != nil {
		fmt.Println("Application not running on " + fmt.Sprintf("%s:%d", *givenHost, *givenPort))
		fmt.Println(err)
		os.Exit(1)
	}

	networkStatus := gjson.Parse(networkInfo)

	genesisRPC, _ := getFromRPC("genesis")

	genesisInfo := gjson.Parse(genesisRPC)

	ctx, cancel := context.WithCancel(context.Background())

	// START INITIALISING WIDGETS

	// Creates Network Widget
	currentNetworkWidget, err := text.New(text.RollContent(), text.WrapAtWords())
	if err != nil {
		panic(err)
	}
	if err := currentNetworkWidget.Write(networkStatus.Get("result.node_info.network").String()); err != nil {
		panic(err)
	}

	// Creates Health Widget
	healthWidget, err := text.New()
	if err != nil {
		panic(err)
	}
	if err := healthWidget.Write("⌛ loading"); err != nil {
		panic(err)
	}

	// Creates System Time Widget
	timeWidget, err := text.New()
	if err != nil {
		panic(err)
	}
	currentTime := time.Now()
	if err := timeWidget.Write(fmt.Sprintf("%s\n", currentTime.Format("2006-01-02\n03:04:05 PM"))); err != nil {
		panic(err)
	}

	// Creates Connected Peers Widget
	peerWidget, err := text.New()
	if err != nil {
		panic(err)
	}
	if err := peerWidget.Write("0"); err != nil {
		panic(err)
	}

	// Creates Seconds Between Blocks Widget
	secondsPerBlockWidget, err := text.New(text.RollContent(), text.WrapAtWords())
	if err != nil {
		panic(err)
	}
	if err := secondsPerBlockWidget.Write("0"); err != nil {
		panic(err)
	}

	// Creates Max Block Size Widget
	maxBlocksizeWidget, err := text.New()

	consensusParamsRPC, _ := getFromRPC("consensus_params")

	maxBlockSize := gjson.Get(consensusParamsRPC, "result.consensus_params.block.max_bytes").Int()
	if err != nil {
		panic(err)
	}
	if err := maxBlocksizeWidget.Write(fmt.Sprintf("%s", byteCountDecimal(maxBlockSize))); err != nil {
		panic(err)
	}

	// Creates Validators widget
	validatorWidget, err := text.New(text.RollContent(), text.WrapAtWords())
	if err != nil {
		panic(err)
	}
	if err := validatorWidget.Write("List available validators.\n\n"); err != nil {
		panic(err)
	}

	// Creates Validators widget
	gasMaxWidget, err := text.New(text.RollContent(), text.WrapAtWords())
	if err != nil {
		panic(err)
	}
	if err := gasMaxWidget.Write("How much gas.\n\n"); err != nil {
		panic(err)
	}

	// Creates Gas per Average Block Widget
	gasAvgBlockWidget, err := text.New(text.RollContent(), text.WrapAtWords())
	if err != nil {
		panic(err)
	}
	if err := gasAvgBlockWidget.Write("How much gas.\n\n"); err != nil {
		panic(err)
	}

	// Creates Gas per Average Transaction Widget
	gasAvgTransactionWidget, err := text.New(text.RollContent(), text.WrapAtWords())
	if err != nil {
		panic(err)
	}
	if err := gasAvgTransactionWidget.Write("How much gas.\n\n"); err != nil {
		panic(err)
	}

	// Creates Gas per Latest Transaction Widget
	latestGasWidget, err := text.New(text.RollContent(), text.WrapAtWords())
	if err != nil {
		panic(err)
	}
	if err := latestGasWidget.Write("How much gas.\n\n"); err != nil {
		panic(err)
	}

	// BIG WIDGETS

	// Block Status Donut widget
	green, err := donut.New(
		donut.CellOpts(cell.FgColor(cell.ColorGreen)),
		donut.Label("New Block Status", cell.FgColor(cell.ColorGreen)),
	)
	if err != nil {
		panic(err)
	}

	// Transaction parsing widget
	transactionWidget, err := text.New(text.RollContent(), text.WrapAtWords())
	if err != nil {
		panic(err)
	}
	if err := transactionWidget.Write("Transactions will appear as soon as they are confirmed in a block.\n\n"); err != nil {
		panic(err)
	}

	// Create Blocks parsing widget
	blocksWidget, err := text.New(text.RollContent(), text.WrapAtWords())
	if err != nil {
		panic(err)
	}
	if err := blocksWidget.Write(networkStatus.Get("result.sync_info.latest_block_height").String() + "\n"); err != nil {
		panic(err)
	}

	// END INITIALISING WIDGETS

	// The functions that execute the updating widgets.

	// system powered widgets
	go writeTime(ctx, info, timeWidget, 1*time.Second)

	// rpc widgets
	go writePeers(ctx, peerWidget, 1*time.Second)
	go writeHealth(ctx, healthWidget, 500*time.Millisecond, connectionSignal)
	go writeSecondsPerBlock(ctx, info, secondsPerBlockWidget, 1*time.Second)
	go writeAmountValidators(ctx, validatorWidget, 3000*time.Millisecond, connectionSignal)
	go writeGasWidget(ctx, info, gasMaxWidget, gasAvgBlockWidget, gasAvgTransactionWidget, latestGasWidget, 1000*time.Millisecond, connectionSignal, genesisInfo)

	// websocket powered widgets
	go writeBlocks(ctx, info, blocksWidget, connectionSignal)
	go writeTransactions(ctx, info, transactionWidget, connectionSignal)
	go writeBlockDonut(ctx, green, 0, 20, 700*time.Millisecond, playTypePercent, connectionSignal)

	// Draw Dashboard
	c, err := container.New(
		t,
		container.Border(linestyle.Light),
		container.BorderTitle("GEX: PRESS Q or ESC TO QUIT"),
		container.BorderColor(cell.ColorNumber(2)),
		container.SplitHorizontal(
			container.Top(
				container.SplitVertical(
					container.Left(
						container.SplitHorizontal(
							container.Top(
								container.SplitVertical(
									container.Left(
										container.SplitVertical(
											container.Left(
												container.Border(linestyle.Light),
												container.BorderTitle("Network"),
												container.PlaceWidget(currentNetworkWidget),
											),
											container.Right(
												container.Border(linestyle.Light),
												container.BorderTitle("Health"),
												container.PlaceWidget(healthWidget),
											),
										),
									),
									container.Right(
										container.SplitVertical(
											container.Left(
												container.Border(linestyle.Light),
												container.BorderTitle("System Time"),
												container.PlaceWidget(timeWidget),
											),
											container.Right(
												container.Border(linestyle.Light),
												container.BorderTitle("Connected Peers"),
												container.PlaceWidget(peerWidget),
											),
										),
									),
								),
							),
							container.Bottom(
								// INSERT NEW BOTTOM ROWS
								container.SplitVertical(
									container.Left(
										container.SplitVertical(
											container.Left(
												container.Border(linestyle.Light),
												container.BorderTitle("Latest Block"),
												container.PlaceWidget(blocksWidget),
											),
											container.Right(
												container.Border(linestyle.Light),
												container.BorderTitle("Max Block Size"),
												container.PlaceWidget(maxBlocksizeWidget),
											),
										),
									),
									container.Right(
										container.SplitVertical(
											container.Left(
												container.Border(linestyle.Light),
												container.BorderTitle("s Between Blocks"),
												container.PlaceWidget(secondsPerBlockWidget),
											),
											container.Right(
												container.Border(linestyle.Light),
												container.BorderTitle("Validators"),
												container.PlaceWidget(validatorWidget),
											),
										),
									),
								),
							),
						),
					),
					container.Right(
						container.Border(linestyle.Light),
						container.BorderTitle("Current Block Round"),
						container.PlaceWidget(green),
					),
				),
			),
			container.Bottom(
				container.SplitVertical(

					container.Left(
						container.SplitHorizontal(
							container.Top(
								container.SplitVertical(
									container.Left(
										container.SplitVertical(
											container.Left(
												container.Border(linestyle.Light),
												container.BorderTitle("Gas Max"),
												container.PlaceWidget(gasMaxWidget),
											),
											container.Right(
												container.Border(linestyle.Light),
												container.BorderTitle("Gas Ø Block"),
												container.PlaceWidget(gasAvgBlockWidget),
											),
										),
									),
									container.Right(
										container.SplitVertical(
											container.Left(
												container.Border(linestyle.Light),
												container.BorderTitle("Gas Ø Tx"),
												container.PlaceWidget(gasAvgTransactionWidget),
											),
											container.Right(
												container.Border(linestyle.Light),
												container.BorderTitle("Gas Latest Tx"),
												container.PlaceWidget(latestGasWidget),
											),
										),
									),
								),
							),
							container.Bottom(
							//empty
							),
						),
					), container.Right(
						container.Border(linestyle.Light),
						container.BorderTitle("Latest Confirmed Transactions"),
						container.PlaceWidget(transactionWidget),
					),
				),
			),
		),
	)
	if err != nil {
		panic(err)
	}

	quitter := func(k *terminalapi.Keyboard) {
		if k.Key == 'q' || k.Key == 'Q' || k.Key == keyboard.KeyEsc {
			cancel()
		}
	}

	if err := termdash.Run(ctx, t, c, termdash.KeyboardSubscriber(quitter)); err != nil {
		panic(err)
	}
}

// writeTime writes the current system time to the timeWidget.
// Exits when the context expires.
func writeTime(ctx context.Context, info Info, t *text.Text, delay time.Duration) {
	ticker := time.NewTicker(delay)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			currentTime := time.Now()
			t.Reset()
			if err := t.Write(fmt.Sprintf("%s\n", currentTime.Format("2006-01-02\n03:04:05 PM"))); err != nil {
				panic(err)
			}
			info.blocks.secondsPassed++
		case <-ctx.Done():
			return
		}
	}
}

// writeHealth writes the status to the healthWidget.
// Exits when the context expires.
func writeHealth(ctx context.Context, t *text.Text, delay time.Duration, connectionSignal chan string) {
	reconnect := false
	healthRPC, _ := getFromRPC("health")
	health := gjson.Get(healthRPC, "result")
	t.Reset()
	if health.Exists() {
		t.Write("✔️ good")
	} else {
		t.Write("✖️ not connected")
	}

	ticker := time.NewTicker(delay)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			healthRPC, _ := getFromRPC("health")
			health := gjson.Get(healthRPC, "result")
			if health.Exists() {
				t.Reset()
				t.Write("✔️ good")
				if reconnect == true {
					connectionSignal <- "reconnect"
					connectionSignal <- "reconnect"
					connectionSignal <- "reconnect"
					reconnect = false
				}
			} else {
				t.Reset()
				t.Write("✖️ not connected")
				if reconnect == false {
					connectionSignal <- "no_connection"
					connectionSignal <- "no_connection"
					connectionSignal <- "no_connection"
					reconnect = true
				}
			}
		case <-ctx.Done():
			return
		}
	}
}

// writePeers writes the connected Peers to the peerWidget.
// Exits when the context expires.
func writePeers(ctx context.Context, t *text.Text, delay time.Duration) {
	netInfoRPC, _ := getFromRPC("net_info")

	peers := gjson.Get(netInfoRPC, "result.n_peers").String()
	t.Reset()
	if peers != "" {
		t.Write(peers)
	}
	if err := t.Write(peers); err != nil {
		panic(err)
	}

	ticker := time.NewTicker(delay)
	t.Reset()
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			t.Reset()
			netInfoRPC, _ := getFromRPC("net_info")
			peers := gjson.Get(netInfoRPC, "result.n_peers").String()
			if peers != "" {
				t.Reset()
				t.Write(peers)
			}

		case <-ctx.Done():
			return
		}
	}
}

// writeAmountValidators writes the status to the healthWidget.
// Exits when the context expires.
func writeAmountValidators(ctx context.Context, t *text.Text, delay time.Duration, connectionSignal chan string) {
	reconnect := false
	validatorsRPC, _ := getFromRPC("validators")
	validators := gjson.Get(validatorsRPC, "result")
	t.Reset()
	if validators.Exists() {
		t.Write("0")
	} else {
		t.Write("0")
	}

	ticker := time.NewTicker(delay)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			validatorsRPC, _ := getFromRPC("validators")
			validators := gjson.Get(validatorsRPC, "result")
			if validators.Exists() {
				t.Reset()
				t.Write(validators.Get("total").String())
				if reconnect == true {
					connectionSignal <- "reconnect"
					connectionSignal <- "reconnect"
					connectionSignal <- "reconnect"
					reconnect = false
				}
			} else {
				t.Reset()
				t.Write("0")
				if reconnect == false {
					connectionSignal <- "no_connection"
					connectionSignal <- "no_connection"
					connectionSignal <- "no_connection"
					reconnect = true
				}
			}
		case <-ctx.Done():
			return
		}
	}
}

// writeGasWidget writes the status to the healthWidget.
// Exits when the context expires.
func writeGasWidget(ctx context.Context, info Info, tMax *text.Text, tAvgBlock *text.Text, tAvgTx *text.Text, tLatest *text.Text, delay time.Duration, _ chan string, _ gjson.Result) {
	tMax.Write("0")
	tAvgBlock.Write("0")
	tLatest.Write("0")
	tAvgTx.Write("0")

	ticker := time.NewTicker(delay)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			tMax.Reset()
			tAvgBlock.Reset()
			tAvgTx.Reset()
			tLatest.Reset()

			totalGasWanted := uint64(info.blocks.totalGasWanted)
			totalBlocks := uint64(info.blocks.amount)
			totalGasPerBlock := uint64(0)

			// don't divide by 0
			if totalBlocks > 0 {
				totalGasPerBlock = uint64(totalGasWanted / totalBlocks)
			}

			totalTransactions := uint64(info.transactions.amount)

			// don't divide by 0
			averageGasPerTx := uint64(0)
			if totalTransactions > 0 {
				averageGasPerTx = uint64(totalGasWanted / info.transactions.amount)
			}

			tMax.Write(fmt.Sprintf("%v", numberWithComma(info.blocks.maxGasWanted)))
			tAvgBlock.Write(fmt.Sprintf("%v", numberWithComma(int64(totalGasPerBlock))))
			tLatest.Write(fmt.Sprintf("%v", numberWithComma(info.blocks.lastTx)))
			tAvgTx.Write(fmt.Sprintf("%v", numberWithComma(int64(averageGasPerTx))))
		case <-ctx.Done():
			return
		}
	}
}

// writeSecondsPerBlock writes the status to the Time per block.
// Exits when the context expires.
func writeSecondsPerBlock(ctx context.Context, info Info, t *text.Text, delay time.Duration) {

	t.Reset()

	ticker := time.NewTicker(delay)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			t.Reset()
			blocksPerSecond := 0.00
			if info.blocks.secondsPassed != 0 {
				blocksPerSecond = float64(info.blocks.secondsPassed) / float64(info.blocks.amount)
			}

			t.Write(fmt.Sprintf("%.2f seconds", blocksPerSecond))
		case <-ctx.Done():
			return
		}
	}
}

// WEBSOCKET WIDGETS

// writeBlocks writes the latest Block to the blocksWidget.
// Exits when the context expires.
func writeBlocks(ctx context.Context, info Info, t *text.Text, connectionSignal <-chan string) {
	socket := gowebsocket.New(getWSURL() + "/websocket")

	socket.OnTextMessage = func(message string, _ gowebsocket.Socket) {
		currentBlock := gjson.Get(message, "result.data.value.block.header.height")
		if currentBlock.String() != "" {
			t.Reset()
			err := t.Write(fmt.Sprintf("%v", numberWithComma(int64(currentBlock.Int()))))
			if err != nil {
				panic(err)
			}
			info.blocks.amount++
			info.blocks.maxGasWanted = gjson.Get(message, "result.data.value.result_end_block.consensus_param_updates.block.max_gas").Int()
		}

	}

	socket.Connect()

	socket.SendText("{ \"jsonrpc\": \"2.0\", \"method\": \"subscribe\", \"params\": [\"tm.event='NewBlock'\"], \"id\": 1 }")

	for {
		select {
		case s := <-connectionSignal:
			if s == "no_connection" {
				socket.Close()
			}
			if s == "reconnect" {
				writeBlocks(ctx, info, t, connectionSignal)
			}
		case <-ctx.Done():
			log.Println("interrupt")
			socket.Close()
			return
		}
	}
}

// writeBlockDonut continuously changes the displayed percent value on the donut by the
// step once every delay. Exits when the context expires.
func writeBlockDonut(ctx context.Context, d *donut.Donut, start, step int, delay time.Duration, pt playType, connectionSignal <-chan string) {
	socket := gowebsocket.New(getWSURL() + "/websocket")

	socket.OnTextMessage = func(message string, _ gowebsocket.Socket) {
		step := gjson.Get(message, "result.data.value.step")
		progress := 0

		if step.String() == "RoundStepNewHeight" {
			progress = 100
		}

		if step.String() == "RoundStepCommit" {
			progress = 80
		}

		if step.String() == "RoundStepPrecommit" {
			progress = 60
		}

		if step.String() == "RoundStepPrevote" {
			progress = 40
		}

		if step.String() == "RoundStepPropose" {
			progress = 20
		}

		if err := d.Percent(progress); err != nil {
			panic(err)
		}

	}

	socket.Connect()

	socket.SendText("{ \"jsonrpc\": \"2.0\", \"method\": \"subscribe\", \"params\": [\"tm.event='NewRoundStep'\"], \"id\": 3 }")

	for {
		select {
		case s := <-connectionSignal:
			if s == "no_connection" {
				socket.Close()
			}
			if s == "reconnect" {
				writeBlockDonut(ctx, d, start, step, delay, pt, connectionSignal)
			}
		case <-ctx.Done():
			log.Println("interrupt")
			socket.Close()
			return
		}
	}
}

// writeTransactions writes the latest Transactions to the transactionsWidget.
// Exits when the context expires.
func writeTransactions(ctx context.Context, info Info, t *text.Text, connectionSignal <-chan string) {
	socket := gowebsocket.New(getWSURL() + "/websocket")

	socket.OnTextMessage = func(message string, _ gowebsocket.Socket) {
		currentTx := gjson.Get(message, "result.data.value.TxResult.result.log")
		currentTime := time.Now()
		if currentTx.String() != "" {
			if err := t.Write(fmt.Sprintf("%s\n", currentTime.Format("2006-01-02 03:04:05 PM")+"\n"+currentTx.String())); err != nil {
				panic(err)
			}

			info.blocks.totalGasWanted = info.blocks.totalGasWanted + gjson.Get(message, "result.data.value.TxResult.result.gas_wanted").Int()
			info.blocks.lastTx = gjson.Get(message, "result.data.value.TxResult.result.gas_wanted").Int()
			info.transactions.amount++
		}
	}

	socket.Connect()

	socket.SendText("{ \"jsonrpc\": \"2.0\", \"method\": \"subscribe\", \"params\": [\"tm.event='Tx'\"], \"id\": 2 }")

	for {
		select {
		case s := <-connectionSignal:
			if s == "no_connection" {
				socket.Close()
			}
			if s == "reconnect" {
				writeTransactions(ctx, info, t, connectionSignal)
			}
		case <-ctx.Done():
			log.Println("interrupt")
			socket.Close()
			return
		}
	}
}

// UTIL FUNCTIONS

// Get Data from RPC Endpoint
func getFromRPC(endpoint string) (string, error) {
	resp, err := resty.R().
		SetHeader("Cache-Control", "no-cache").
		SetHeader("Content-Type", "application/json").
		Get(getHTTPURL() + "/" + endpoint)

	return resp.String(), err
}

// byteCountDecimal calculates bytes integer to a human readable decimal number
func byteCountDecimal(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "kMGTPE"[exp])
}

func numberWithComma(n int64) string {
	in := strconv.FormatInt(n, 10)
	numOfDigits := len(in)
	if n < 0 {
		numOfDigits-- // First character is the - sign (not a digit)
	}
	numOfCommas := (numOfDigits - 1) / 3

	out := make([]byte, len(in)+numOfCommas)
	if n < 0 {
		in, out[0] = in[1:], '-'
	}

	for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
		out[j] = in[i]
		if i == 0 {
			return string(out)
		}
		if k++; k == 3 {
			j, k = j-1, 0
			out[j] = ','
		}
	}
}

func getHTTPURL() string {
	return getURL("http", *ssl)
}

func getWSURL() string {
	return getURL("ws", *ssl)
}

func getURL(protocol string, secure bool) string {
	if secure {
		protocol = protocol + "s"
	}

	return fmt.Sprintf("%s://%s:%d", protocol, *givenHost, *givenPort)
}

func view() {

}
