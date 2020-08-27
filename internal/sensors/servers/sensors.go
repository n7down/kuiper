package servers

import (
	"context"

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

func (s *SensorsServer) GetVoltageMeasurements(context.Context, *sensors_pb.GetVoltageMeasurementsRequest) (*sensors_pb.GetVoltageMeasurementsResponse, error) {
	return &sensors_pb.GetVoltageMeasurementsResponse{}, nil
}

func (s *SensorsServer) GetHumidityMeasurements(context.Context, *sensors_pb.GetHumidityMeasurementsRequest) (*sensors_pb.GetHumidityMeasurementsResponse, error) {
	return &sensors_pb.GetHumidityMeasurementsResponse{}, nil
}

func (s *SensorsServer) GetTemperatureMeasurements(context.Context, *sensors_pb.GetTemperatureMeasurementsRequest) (*sensors_pb.GetTemperatureMeasurementsResponse, error) {
	return &sensors_pb.GetTemperatureMeasurementsResponse{}, nil
}
