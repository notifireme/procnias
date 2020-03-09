package procnias

import (
	"strings"

	alertmanager "github.com/prometheus/alertmanager/template"
)

func (a App) logic(data alertmanager.Data) error {
	a.logger.Printf("Alerts: GroupLabels=%v, CommonLabels=%v", data.GroupLabels, data.CommonLabels)
	for _, alert := range data.Alerts {
		a.logger.Printf("Alert: status=%s,Labels=%v,Annotations=%v", alert.Status, alert.Labels, alert.Annotations)
		severity := alert.Labels["severity"]
		switch strings.ToUpper(severity) {
		case "HIGH":
			if err := a.sendSMS(alert); err != nil {
				return err
			}
		case "WARNING":
			if err := a.sendSMS(alert); err != nil {
				return err
			}
		default:
			a.logger.Printf("no action on severity: %s", severity)
		}
	}
	return nil
}
