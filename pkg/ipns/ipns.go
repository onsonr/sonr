package ipns

import (
	"fmt"
	"time"

	shell "github.com/ipfs/go-ipfs-api"
)

func Publish(shell *shell.Shell, rec *IPNSRecord) (string, error) {
	id, err := shell.ID()
	if err != nil {
		return "", err
	}
	resp, err := shell.PublishWithDetails(rec.Builder.cid, id.ID, time.Hour, rec.Ttl, true)
	if err != nil {
		return "", err
	}

	fmt.Print(resp.Name)

	return resp.Name, nil
}

func Resolve(shell *shell.Shell, id string) (string, error) {
	resp, err := shell.Resolve(id)
	if err != nil {
		return "", err
	}

	fmt.Print(resp)
	return resp, err
}
