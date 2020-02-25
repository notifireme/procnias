package procnias

import "net/url"

func (a App) sendSMS() {
	a.twilio.Messages.SendMessage("", "", "Sent via go  ✓", nil)
}

func (a App) makeCall() {
	content, _ := url.Parse("http://magory.net/assets/MagoryNET-HappyFlutes.mp3")
	a.twilio.Calls.MakeCall("", "", content)
}
