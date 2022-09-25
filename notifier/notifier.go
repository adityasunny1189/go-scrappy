package notifier

import (
	"encoding/json"
	"log"
	"os"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func Notify(msg string) {
	accountSid := os.Getenv("ACCNTID")
	authToken := os.Getenv("AUTHTKN")
	targetNum := os.Getenv("TARGETNUM")
	myNum := os.Getenv("MYNUM")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(targetNum)
	params.SetFrom(myNum)
	params.SetBody(msg)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		log.Println(err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		log.Println("Response: " + string(response))
	}
}
