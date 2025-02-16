package mcir

import (
	"context"

	gc "github.com/cloud-barista/cb-tumblebug/src/api/grpc/common"
	"github.com/cloud-barista/cb-tumblebug/src/api/grpc/logger"
	pb "github.com/cloud-barista/cb-tumblebug/src/api/grpc/protobuf/cbtumblebug"

	"github.com/cloud-barista/cb-tumblebug/src/core/mcir"
)

// ===== [ Constants and Variables ] =====

// ===== [ Types ] =====

// ===== [ Implementations ] =====

// CheckResource - Resouce 체크
func (s *MCIRService) CheckResource(ctx context.Context, req *pb.ResourceQryRequest) (*pb.ExistsResponse, error) {
	logger := logger.NewLogger()

	logger.Debug("calling MCIRService.CheckResource()")

	exists, _, err := mcir.LowerizeAndCheckResource(req.NsId, req.ResourceType, req.ResourceId)
	if err != nil {
		logger.Debug(err)
	}

	resp := &pb.ExistsResponse{Exists: exists}
	return resp, nil
}

// ListLookupSpec - Spec 목록
func (s *MCIRService) ListLookupSpec(ctx context.Context, req *pb.LookupSpecListQryRequest) (*pb.ListSpiderSpecInfoResponse, error) {
	logger := logger.NewLogger()

	logger.Debug("calling MCIRService.ListLookupSpec()")

	content, err := mcir.LookupSpecList(req.ConnectionName)
	if err != nil {
		return nil, gc.ConvGrpcStatusErr(err, "", "MCIRService.ListLookupSpec()")
	}

	// MCIR 객체에서 GRPC 메시지로 복사
	var grpcObj pb.ListSpiderSpecInfoResponse
	err = gc.CopySrcToDest(&content, &grpcObj)
	if err != nil {
		return nil, gc.ConvGrpcStatusErr(err, "", "MCIRService.ListLookupSpec()")
	}

	return &grpcObj, nil
}

// GetLookupSpec - Spec 조회
func (s *MCIRService) GetLookupSpec(ctx context.Context, req *pb.LookupSpecQryRequest) (*pb.SpiderSpecInfoResponse, error) {
	logger := logger.NewLogger()

	logger.Debug("calling MCIRService.GetLookupSpec()")

	content, err := mcir.LookupSpec(req.ConnectionName, req.SpecName)
	if err != nil {
		return nil, gc.ConvGrpcStatusErr(err, "", "MCIRService.GetLookupSpec()")
	}

	// MCIR 객체에서 GRPC 메시지로 복사
	var grpcObj pb.SpiderSpecInfo
	err = gc.CopySrcToDest(&content, &grpcObj)
	if err != nil {
		return nil, gc.ConvGrpcStatusErr(err, "", "MCIRService.GetLookupSpec()")
	}

	resp := &pb.SpiderSpecInfoResponse{Item: &grpcObj}
	return resp, nil
}

// ListLookupImage - Image 목록
func (s *MCIRService) ListLookupImage(ctx context.Context, req *pb.LookupImageListQryRequest) (*pb.ListSpiderImageInfoResponse, error) {
	logger := logger.NewLogger()

	logger.Debug("calling MCIRService.ListLookupImage()")

	content, err := mcir.LookupImageList(req.ConnectionName)
	if err != nil {
		return nil, gc.ConvGrpcStatusErr(err, "", "MCIRService.ListLookupImage()")
	}

	// MCIR 객체에서 GRPC 메시지로 복사
	var grpcObj pb.ListSpiderImageInfoResponse
	err = gc.CopySrcToDest(&content, &grpcObj)
	if err != nil {
		return nil, gc.ConvGrpcStatusErr(err, "", "MCIRService.ListLookupImage()")
	}

	return &grpcObj, nil
}

// GetLookupImage - Image 조회
func (s *MCIRService) GetLookupImage(ctx context.Context, req *pb.LookupImageQryRequest) (*pb.SpiderImageInfoResponse, error) {
	logger := logger.NewLogger()

	logger.Debug("calling MCIRService.GetLookupImage()")

	content, err := mcir.LookupImage(req.ConnectionName, req.ImageName)
	if err != nil {
		return nil, gc.ConvGrpcStatusErr(err, "", "MCIRService.GetLookupImage()")
	}

	// MCIR 객체에서 GRPC 메시지로 복사
	var grpcObj pb.SpiderImageInfo
	err = gc.CopySrcToDest(&content, &grpcObj)
	if err != nil {
		return nil, gc.ConvGrpcStatusErr(err, "", "MCIRService.GetLookupImage()")
	}

	resp := &pb.SpiderImageInfoResponse{Item: &grpcObj}
	return resp, nil
}

// ===== [ Private Functions ] =====

// ===== [ Public Functions ] =====
