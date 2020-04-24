package request

import (
	"net/url"
	"regexp"
)

type GetBatCaveSettingsRequest struct {
	DeviceID string `json:"-"`
}

func (r *GetBatCaveSettingsRequest) Validate() url.Values {
	errs := url.Values{}

	if r.DeviceID == "" {
		errs.Add("deviceID", "The deviceID field is required!")
	}

	if len(r.DeviceID) != 12 {
		errs.Add("deviceID", "The deviceID field needs to be a valid mac!")
	}

	regex, _ := regexp.Compile("a-fA-F0-9]{12}")
	isMacAddress := regex.MatchString(r.DeviceID)
	if !isMacAddress {
		errs.Add("deviceID", "The deviceID field needs to be a valid mac!")
	}

	return errs
}
