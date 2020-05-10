package response

type UpdateBatCaveSettingResponse struct {
	DeviceID       string `json:"deviceID"`
	DeepSleepDelay uint32 `json:"deepSleepDelay"`
}
