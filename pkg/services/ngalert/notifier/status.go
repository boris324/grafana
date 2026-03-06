package notifier

import (
	"context"

	apimodels "github.com/grafana/grafana/pkg/services/ngalert/api/tooling/definitions"
)

// TODO: We no longer do apimodels at this layer, move it to the API.
func (am *alertmanager) GetStatus(_ context.Context) (apimodels.GettableStatus, error) {
	status := am.Base.AppliedConfig()
	if status == nil {
		return *apimodels.NewGettableStatus(&apimodels.PostableApiAlertingConfig{}), nil
	}

	amConfig := NotificationsConfigurationToPostableAPIConfig(*status)
	return *apimodels.NewGettableStatus(&amConfig), nil
}
