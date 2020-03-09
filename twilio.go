package procnias

import (
	"bytes"
	"net/url"
	"text/template"

	alertmanager "github.com/prometheus/alertmanager/template"
)

const msg = `Allert: {{.Status}}`

func (a App) sendSMS(alert alertmanager.Alert) error {
	var b bytes.Buffer
	tmpl, err := template.New("").Parse(msg)
	if err != nil {
		return err
	}
	if err := tmpl.Execute(&b, &alert); err != nil {
		return err
	}
	if _, err := a.twilio.Messages.SendMessage("", "", b.String(), nil); err != nil {
		return err
	}
	return nil
}

func (a App) makeCall() {
	content, _ := url.Parse("http://magory.net/assets/MagoryNET-HappyFlutes.mp3")
	a.twilio.Calls.MakeCall("", "", content)
}
