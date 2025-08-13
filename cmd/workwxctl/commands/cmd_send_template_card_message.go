package commands

import (
	"log"

	"github.com/bobbywjb/go-workwx/v2"
	"github.com/google/uuid"
	"github.com/urfave/cli/v2"
)

func cmdSendTemplateCardMessage(c *cli.Context) error {
	cfg := mustGetConfig(c)

	templateMessageType := c.String(flagTemplateMessageType)

	buttonText := c.String(flagButtonText)

	_ = buttonText

	responseCode := c.String(flagResponseCode)

	app := cfg.MakeWorkwxApp()

	switch templateMessageType {
	case "normal":
		resp, err := app.SendTemplateCardMessageWithResponse(&workwx.Recipient{
			UserIDs: []string{"3t5er"},
		}, workwx.TemplateCard{

			CardType: workwx.CardTypeMultipleInteraction,
			MainTitle: &workwx.MainTitle{
				Title: "ttt",
				Desc:  "123",
			},
			TaskID: uuid.NewString(),
			SelectList: []workwx.SelectList{
				{

					QuestionKey: "leader",
					Title:       "所选领导",
					OptionList: []workwx.OptionList{
						{
							ID:   "3t5er",
							Text: "ttttt",
						},
					},
				},
			},
			SubmitButton: &workwx.SubmitButton{
				Text: "提交",
				Key:  "ListReply",
			},
		}, false)
		if err != nil {
			return err
		}
		log.Println(resp)
	case "update":

		err := app.SendTemplateCardUpdateMessage(&workwx.TemplateCardUpdateMessage{
			AtAll:        1,
			ResponseCode: responseCode,
			TemplateCard: workwx.TemplateCard{

				CardType: workwx.CardTypeMultipleInteraction,
				MainTitle: &workwx.MainTitle{
					Title: "tst",
					Desc:  "tst",
				},
				SelectList: []workwx.SelectList{
					{

						QuestionKey: "leader",
						Title:       "所选领导",
						SelectedID:  "3t5er",
						Disable:     true,
						OptionList: []workwx.OptionList{
							{
								ID:   "3t5er",
								Text: "3t5er",
							},
						},
					},
				},
				SubmitButton: &workwx.SubmitButton{
					Text: "已提交",
					Key:  "Submit",
				},
				ReplaceText: "已提交",
			},
		}, false)
		log.Printf("err:%v", err)
		return err
	}
	return nil

}
