package util

import (
	"google.golang.org/grpc/grpclog"
	"os"
)

var GrpcLoggerV2 grpclog.LoggerV2

func GetGrpcLoggerV2() grpclog.LoggerV2 {
	if GrpcLoggerV2 != nil {
		return GrpcLoggerV2
	}
	GrpcLoggerV2 = grpclog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout)
	grpclog.SetLoggerV2(GrpcLoggerV2)
	return GrpcLoggerV2
}
