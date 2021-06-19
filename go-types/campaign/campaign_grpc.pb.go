// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package campaign

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CampaignServiceClient is the client API for CampaignService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CampaignServiceClient interface {
	CreateCampaign(ctx context.Context, in *CreateCampaignRequest, opts ...grpc.CallOption) (*CreateCampaignResponse, error)
	ListCampaigns(ctx context.Context, in *ListCampaignsRequest, opts ...grpc.CallOption) (*CampaignsResponse, error)
	FilterCampaigns(ctx context.Context, in *FilterCampaignsRequest, opts ...grpc.CallOption) (*CampaignsResponse, error)
	GetCampaignDetails(ctx context.Context, in *GetCampaignDetailsRequest, opts ...grpc.CallOption) (*CampaignDetails, error)
	CreateCampaignType(ctx context.Context, in *CreateCampaignTypeRequest, opts ...grpc.CallOption) (*CreateCampaignTypeResponse, error)
	ListCampaignTypes(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CampaignTypesResponse, error)
	// beneficiary
	CreateBeneficiary(ctx context.Context, in *CreateBeneficiaryRequest, opts ...grpc.CallOption) (*CreateBeneficiaryResponse, error)
}

type campaignServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCampaignServiceClient(cc grpc.ClientConnInterface) CampaignServiceClient {
	return &campaignServiceClient{cc}
}

func (c *campaignServiceClient) CreateCampaign(ctx context.Context, in *CreateCampaignRequest, opts ...grpc.CallOption) (*CreateCampaignResponse, error) {
	out := new(CreateCampaignResponse)
	err := c.cc.Invoke(ctx, "/podkrepibg.campaign.CampaignService/CreateCampaign", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campaignServiceClient) ListCampaigns(ctx context.Context, in *ListCampaignsRequest, opts ...grpc.CallOption) (*CampaignsResponse, error) {
	out := new(CampaignsResponse)
	err := c.cc.Invoke(ctx, "/podkrepibg.campaign.CampaignService/ListCampaigns", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campaignServiceClient) FilterCampaigns(ctx context.Context, in *FilterCampaignsRequest, opts ...grpc.CallOption) (*CampaignsResponse, error) {
	out := new(CampaignsResponse)
	err := c.cc.Invoke(ctx, "/podkrepibg.campaign.CampaignService/FilterCampaigns", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campaignServiceClient) GetCampaignDetails(ctx context.Context, in *GetCampaignDetailsRequest, opts ...grpc.CallOption) (*CampaignDetails, error) {
	out := new(CampaignDetails)
	err := c.cc.Invoke(ctx, "/podkrepibg.campaign.CampaignService/GetCampaignDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campaignServiceClient) CreateCampaignType(ctx context.Context, in *CreateCampaignTypeRequest, opts ...grpc.CallOption) (*CreateCampaignTypeResponse, error) {
	out := new(CreateCampaignTypeResponse)
	err := c.cc.Invoke(ctx, "/podkrepibg.campaign.CampaignService/CreateCampaignType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campaignServiceClient) ListCampaignTypes(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CampaignTypesResponse, error) {
	out := new(CampaignTypesResponse)
	err := c.cc.Invoke(ctx, "/podkrepibg.campaign.CampaignService/ListCampaignTypes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campaignServiceClient) CreateBeneficiary(ctx context.Context, in *CreateBeneficiaryRequest, opts ...grpc.CallOption) (*CreateBeneficiaryResponse, error) {
	out := new(CreateBeneficiaryResponse)
	err := c.cc.Invoke(ctx, "/podkrepibg.campaign.CampaignService/CreateBeneficiary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampaignServiceServer is the server API for CampaignService service.
// All implementations must embed UnimplementedCampaignServiceServer
// for forward compatibility
type CampaignServiceServer interface {
	CreateCampaign(context.Context, *CreateCampaignRequest) (*CreateCampaignResponse, error)
	ListCampaigns(context.Context, *ListCampaignsRequest) (*CampaignsResponse, error)
	FilterCampaigns(context.Context, *FilterCampaignsRequest) (*CampaignsResponse, error)
	GetCampaignDetails(context.Context, *GetCampaignDetailsRequest) (*CampaignDetails, error)
	CreateCampaignType(context.Context, *CreateCampaignTypeRequest) (*CreateCampaignTypeResponse, error)
	ListCampaignTypes(context.Context, *empty.Empty) (*CampaignTypesResponse, error)
	// beneficiary
	CreateBeneficiary(context.Context, *CreateBeneficiaryRequest) (*CreateBeneficiaryResponse, error)
	mustEmbedUnimplementedCampaignServiceServer()
}

// UnimplementedCampaignServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCampaignServiceServer struct {
}

func (UnimplementedCampaignServiceServer) CreateCampaign(context.Context, *CreateCampaignRequest) (*CreateCampaignResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCampaign not implemented")
}
func (UnimplementedCampaignServiceServer) ListCampaigns(context.Context, *ListCampaignsRequest) (*CampaignsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCampaigns not implemented")
}
func (UnimplementedCampaignServiceServer) FilterCampaigns(context.Context, *FilterCampaignsRequest) (*CampaignsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FilterCampaigns not implemented")
}
func (UnimplementedCampaignServiceServer) GetCampaignDetails(context.Context, *GetCampaignDetailsRequest) (*CampaignDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCampaignDetails not implemented")
}
func (UnimplementedCampaignServiceServer) CreateCampaignType(context.Context, *CreateCampaignTypeRequest) (*CreateCampaignTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCampaignType not implemented")
}
func (UnimplementedCampaignServiceServer) ListCampaignTypes(context.Context, *empty.Empty) (*CampaignTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCampaignTypes not implemented")
}
func (UnimplementedCampaignServiceServer) CreateBeneficiary(context.Context, *CreateBeneficiaryRequest) (*CreateBeneficiaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBeneficiary not implemented")
}
func (UnimplementedCampaignServiceServer) mustEmbedUnimplementedCampaignServiceServer() {}

// UnsafeCampaignServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CampaignServiceServer will
// result in compilation errors.
type UnsafeCampaignServiceServer interface {
	mustEmbedUnimplementedCampaignServiceServer()
}

func RegisterCampaignServiceServer(s grpc.ServiceRegistrar, srv CampaignServiceServer) {
	s.RegisterService(&CampaignService_ServiceDesc, srv)
}

func _CampaignService_CreateCampaign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCampaignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignServiceServer).CreateCampaign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/podkrepibg.campaign.CampaignService/CreateCampaign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignServiceServer).CreateCampaign(ctx, req.(*CreateCampaignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampaignService_ListCampaigns_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCampaignsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignServiceServer).ListCampaigns(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/podkrepibg.campaign.CampaignService/ListCampaigns",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignServiceServer).ListCampaigns(ctx, req.(*ListCampaignsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampaignService_FilterCampaigns_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilterCampaignsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignServiceServer).FilterCampaigns(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/podkrepibg.campaign.CampaignService/FilterCampaigns",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignServiceServer).FilterCampaigns(ctx, req.(*FilterCampaignsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampaignService_GetCampaignDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCampaignDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignServiceServer).GetCampaignDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/podkrepibg.campaign.CampaignService/GetCampaignDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignServiceServer).GetCampaignDetails(ctx, req.(*GetCampaignDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampaignService_CreateCampaignType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCampaignTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignServiceServer).CreateCampaignType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/podkrepibg.campaign.CampaignService/CreateCampaignType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignServiceServer).CreateCampaignType(ctx, req.(*CreateCampaignTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampaignService_ListCampaignTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignServiceServer).ListCampaignTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/podkrepibg.campaign.CampaignService/ListCampaignTypes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignServiceServer).ListCampaignTypes(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampaignService_CreateBeneficiary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBeneficiaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignServiceServer).CreateBeneficiary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/podkrepibg.campaign.CampaignService/CreateBeneficiary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignServiceServer).CreateBeneficiary(ctx, req.(*CreateBeneficiaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CampaignService_ServiceDesc is the grpc.ServiceDesc for CampaignService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CampaignService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "podkrepibg.campaign.CampaignService",
	HandlerType: (*CampaignServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCampaign",
			Handler:    _CampaignService_CreateCampaign_Handler,
		},
		{
			MethodName: "ListCampaigns",
			Handler:    _CampaignService_ListCampaigns_Handler,
		},
		{
			MethodName: "FilterCampaigns",
			Handler:    _CampaignService_FilterCampaigns_Handler,
		},
		{
			MethodName: "GetCampaignDetails",
			Handler:    _CampaignService_GetCampaignDetails_Handler,
		},
		{
			MethodName: "CreateCampaignType",
			Handler:    _CampaignService_CreateCampaignType_Handler,
		},
		{
			MethodName: "ListCampaignTypes",
			Handler:    _CampaignService_ListCampaignTypes_Handler,
		},
		{
			MethodName: "CreateBeneficiary",
			Handler:    _CampaignService_CreateBeneficiary_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "campaign.proto",
}
