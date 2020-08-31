package servers

import (
	"github.com/n7down/kuiper/internal/sensors/persistence"

	sensors_pb "github.com/n7down/kuiper/internal/pb/sensors"
)

type SensorsServer struct {
	persistence persistence.Persistence
	sensors_pb.UnimplementedSensorsServiceServer
}

func NewSensorsServer(persistence persistence.Persistence) *SensorsServer {
	return &SensorsServer{
		persistence: persistence,
	}
}

func (s *SensorsServer) GetVoltageMeasurements(*sensors_pb.GetVoltageMeasurementsRequest, sensors_pb.SensorsService_GetVoltageMeasurementsServer) error {
	return nil
}

func (s *SensorsServer) GetHumidityMeasurements(*sensors_pb.GetHumidityMeasurementsRequest, sensors_pb.SensorsService_GetHumidityMeasurementsServer) error {
	return nil
}

func (s *SensorsServer) GetTemperatureMeasurements(*sensors_pb.GetTemperatureMeasurementsRequest, sensors_pb.SensorsService_GetTemperatureMeasurementsServer) error {
	return nil
}
