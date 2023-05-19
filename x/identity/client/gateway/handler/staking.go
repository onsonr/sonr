package handler

import (
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/gofiber/fiber/v2"
	"github.com/sonrhq/core/internal/local"
	"github.com/sonrhq/core/x/identity/internal/tx/cosmos"
	"github.com/sonrhq/core/x/identity/client/gateway/middleware"

	"google.golang.org/grpc"
)

// This function lists all validators using gRPC in a Fiber web framework.
func ListValidators(c *fiber.Ctx) error {
	conn, err := grpc.Dial(local.Context().GrpcEndpoint(), grpc.WithInsecure())
	if err != nil {
		return err
	}
	resp, err := stakingtypes.NewQueryClient(conn).Validators(c.Context(), &stakingtypes.QueryValidatorsRequest{})
	if err != nil {
		return err
	}
	return c.JSON(resp.Validators)
}

// This function lists the delegators of a validator in a blockchain network using gRPC in Go.
func ListDelegators(c *fiber.Ctx) error {
	conn, err := grpc.Dial(local.Context().GrpcEndpoint(), grpc.WithInsecure())
	if err != nil {
		return err
	}
	resp, err := stakingtypes.NewQueryClient(conn).ValidatorDelegations(c.Context(), &stakingtypes.QueryValidatorDelegationsRequest{
		ValidatorAddr: c.Params("address"),
	})
	if err != nil {
		return err
	}
	return c.JSON(resp)
}

// The function submits a transaction to create a validator on a blockchain network using the first
// account of a user and provided parameters.
func SubmitCreateValidator(c *fiber.Ctx) error {
	usr, err := middleware.FetchUser(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	accs, err := usr.Controller.ListAccounts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if len(accs) == 0 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "no accounts found",
		})
	}
	msg := cosmos.BuildMsgCreateValidator(accs[0], cosmos.WithDescriptionMoniker(c.Query("moniker")), cosmos.WithDescriptionIdentity(c.Query("identity")), cosmos.WithDescriptionWebsite(c.Query("website")), cosmos.WithSecurityContact(c.Query("securityContact")), cosmos.WithStakeAmount(500000))
	// resp, err := accs[0].SendSonrTx(msg)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"error": err.Error(),
	// 	})
	// }
	return c.JSON(msg)
}

// The function submits a transaction to edit a validator on the blockchain network using the first
// account of a user and provided parameters.
func SubmitEditValidator(c *fiber.Ctx) error {
	usr, err := middleware.FetchUser(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	accs, err := usr.Controller.ListAccounts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if len(accs) == 0 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "no accounts found",
		})
	}
	msg := cosmos.BuildMsgEditValidator(accs[0], cosmos.WithDescriptionMoniker(c.Query("moniker")), cosmos.WithDescriptionIdentity(c.Query("identity")), cosmos.WithDescriptionWebsite(c.Query("website")), cosmos.WithSecurityContact(c.Query("securityContact")))
	// resp, err := accs[0].SendSonrTx(msg)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"error": err.Error(),
	// 	})
	// }
	return c.JSON(msg)
}

// The function submits a transaction to delegate to a validator on the blockchain network using the first
// account of a user and provided parameters.
func SubmitDelegate(c *fiber.Ctx) error {
	usr, err := middleware.FetchUser(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	accs, err := usr.Controller.ListAccounts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if len(accs) == 0 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "no accounts found",
		})
	}
	msg := cosmos.BuildMsgDelegate(accs[0], c.Query("validator"), cosmos.WithDelegateAmount(c.QueryInt("amount")))
	// resp, err := accs[0].SendSonrTx(msg)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"error": err.Error(),
	// 	})
	// }
	return c.JSON(msg)
}

// The function submits a transaction to undelegate from a validator on the blockchain network using the first
// account of a user and provided parameters.
func SubmitUndelegate(c *fiber.Ctx) error {
	user, err := middleware.FetchUser(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	accounts, err := user.Controller.ListAccounts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if len(accounts) == 0 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "no accounts found",
		})
	}

	msg := cosmos.BuildMsgUndelegate(accounts[0], c.Query("validator"), cosmos.WithDelegateAmount(c.QueryInt("amount")))
	// resp, err := accounts[0].SendSonrTx(msg)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"error": err.Error(),
	// 	})
	// }
	return c.JSON(msg)
}

// The function submits a transaction to cancel a undelgation from a validator on the blockchain network using the first
// account of a user and provided parameters.
func SubmitCancelUnbondingDelegation(c *fiber.Ctx) error {
	user, err := middleware.FetchUser(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	accounts, err := user.Controller.ListAccounts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if len(accounts) == 0 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "no accounts found",
		})
	}

	msg := cosmos.BuildMsgCancelUndelegate(accounts[0], c.Query("validator"), cosmos.WithDelegateAmount(c.QueryInt("amount")), cosmos.WithCreationHeight(int64(c.QueryInt("creationHeight"))))
	// resp, err := accounts[0].SendSonrTx(msg)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"error": err.Error(),
	// 	})
	// }
	return c.JSON(msg)
}
