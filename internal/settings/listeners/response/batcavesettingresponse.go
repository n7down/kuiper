package response

type BatCaveSettingResponse struct {
	DeepSleepDelay uint32 `json:"s"`
}

func GetBatCaveSettingDefault() BatCaveSettingResponse {
	return BatCaveSettingResponse{
		DeepSleepDelay: 15,
	}
}
