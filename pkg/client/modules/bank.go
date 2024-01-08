package modules

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

type BankClient struct {}

func (bk BankClient) GetBalance(ctx context.Context, address sdk.AccAddress, denom string) (*banktypes.QueryBalanceResponse, error) {
    cc, err := getGrpcConn()
    if err != nil {
        return nil, err
    }
    bankClient := banktypes.NewQueryClient(cc)
    bankRes, err := bankClient.Balance(ctx, banktypes.NewQueryBalanceRequest(address, denom))
    if err != nil {
        return nil, err
    }
    return bankRes, nil
}
