package response

type CreateBatCaveSettingResponse struct {
	DeviceID       string `json:"deviceID"`
	DeepSleepDelay uint32 `json:"deepSleepDelay"`
}
