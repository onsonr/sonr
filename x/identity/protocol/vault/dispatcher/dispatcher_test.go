package dispatcher_test

import (
	"fmt"
	"testing"

	"github.com/cosmos/gogoproto/proto"
	"github.com/sonrhq/core/app"
	"github.com/sonrhq/core/pkg/common"
	"github.com/sonrhq/core/x/identity/protocol/vault/dispatcher"
)

func TestDispatcher(t *testing.T) {
	d := dispatcher.New()
	w, err := d.BuildNewDIDController()
	checkErr(t, err)
	t.Log(w.Address())
	err = w.CreateAccount("Ethereum", common.CoinType_CoinType_ETHEREUM)
	checkErr(t, err)
	err = w.CreateAccount("Bitcoin", common.CoinType_CoinType_BITCOIN)
	checkErr(t, err)
	bz, err := app.MakeEncodingConfig().Marshaler.MarshalJSON(w.Document())
	checkErr(t, err)
	t.Log(string(bz))
	fmt.Println(proto.MarshalTextString(w.Document()))
}

func checkErr(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

// snr16nzrp4x3sachmraq34uzr9tpzpp5tegcjam80z
// snr1qd3q2qfrax99264gcwts8jhentkttv7cgnl23k44u0w2j5n74cyqyxukmmh
