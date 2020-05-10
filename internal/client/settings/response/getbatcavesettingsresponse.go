package response

type GetBatCaveSettingResponse struct {
	DeviceID       string `json:"deviceID"`
	DeepSleepDelay uint32 `json:"deepSleepDelay"`
}
