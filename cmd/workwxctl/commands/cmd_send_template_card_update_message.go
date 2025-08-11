package commands

import (
	"log"

	"github.com/bobbywjb/go-workwx/v2"
	"github.com/urfave/cli/v2"
)

func cmdSendTemplateCardUpdateMessage(c *cli.Context) error {
	cfg := mustGetConfig(c)

	buttonText := c.String(flagButtonText)

	responseCode := c.String(flagResponseCode)

	app := cfg.MakeWorkwxApp()

	err := app.SendTemplateCardUpdateMessage(&workwx.TemplateCardUpdateMessage{
		AtAll:        1,
		ResponseCode: responseCode,
		Button: struct {
			ReplaceName string `json:"replace_name"`
		}{buttonText},
	}, false)
	log.Printf("err:%v", err)
	return err
}
