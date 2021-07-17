package connect

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/account"
	"github.com/stripe/stripe-go/accountlink"

	"../../../structs"
)

func Get(c *fiber.Ctx) error {
	stripe.Key = "sk_test_51IKbe1HgfFQID7blLncgaPQrfmygWq1SlDqkFSsfYWp42vHTSaC4A3g7TLQ5yws6iZAcy2Geieperz2HWjmKuDX000GFd365M3"

	var token structs.Token
	var user structs.User

	var connect structs.Connect
	var connectID int
	var connectedAccountID string

	err, stat := token.Get("Value", c.Params("value"))
	if err != nil {
		c.Status(stat)
		return err
	}

	err, stat = user.Get("ID", strconv.Itoa(token.UserID))
	if err != nil {
		c.Status(stat)
		return err
	}

	err, stat = connect.Get("UserID", strconv.Itoa(user.ID))
	if err != nil {
		aParams := &stripe.AccountParams{
			Email: stripe.String(user.Email),
			Type:  stripe.String(string(stripe.AccountTypeExpress)),
		}

		aAcct, err := account.New(aParams)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return errors.New("Failed to generate a stripe account for linking")
		}

		connect.UserID = user.ID
		connect.ConnectID = aAcct.ID
		connect.Post(&connectID)

		connectedAccountID = aAcct.ID
	} else {
		connectedAccountID = connect.ConnectID
	}

	aLParams := &stripe.AccountLinkParams{
		Account:    stripe.String(connectedAccountID),
		RefreshURL: stripe.String("https://syncedstudios.com/stripe/connect/reauth"),
		ReturnURL:  stripe.String("https://syncedstudios.com/stripe/connect/return"),
		Type:       stripe.String("account_onboarding"),
	}

	aLAcct, err := accountlink.New(aLParams)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return c.SendString(aLAcct.URL)
}
