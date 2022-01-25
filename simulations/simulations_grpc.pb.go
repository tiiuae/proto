// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package simulations

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SimulationCoordinatorClient is the client API for SimulationCoordinator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SimulationCoordinatorClient interface {
	ListSimulations(ctx context.Context, in *ListSimulationsRequest, opts ...grpc.CallOption) (*ListSimulationsResponse, error)
	CreateSimulation(ctx context.Context, in *CreateSimulationRequest, opts ...grpc.CallOption) (*CreateSimulationResponse, error)
	DestroySimulation(ctx context.Context, in *DestroySimulationRequest, opts ...grpc.CallOption) (*DestroySimulationResponse, error)
	ListModels(ctx context.Context, in *ListModelsRequest, opts ...grpc.CallOption) (*ListModelsResponse, error)
	AddModel(ctx context.Context, in *AddModelRequest, opts ...grpc.CallOption) (*AddModelResponse, error)
	RemoveModel(ctx context.Context, in *RemoveModelRequest, opts ...grpc.CallOption) (*RemoveModelResponse, error)
	Logs(ctx context.Context, in *LogsRequest, opts ...grpc.CallOption) (SimulationCoordinator_LogsClient, error)
}

type simulationCoordinatorClient struct {
	cc grpc.ClientConnInterface
}

func NewSimulationCoordinatorClient(cc grpc.ClientConnInterface) SimulationCoordinatorClient {
	return &simulationCoordinatorClient{cc}
}

func (c *simulationCoordinatorClient) ListSimulations(ctx context.Context, in *ListSimulationsRequest, opts ...grpc.CallOption) (*ListSimulationsResponse, error) {
	out := new(ListSimulationsResponse)
	err := c.cc.Invoke(ctx, "/simulations.SimulationCoordinator/ListSimulations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simulationCoordinatorClient) CreateSimulation(ctx context.Context, in *CreateSimulationRequest, opts ...grpc.CallOption) (*CreateSimulationResponse, error) {
	out := new(CreateSimulationResponse)
	err := c.cc.Invoke(ctx, "/simulations.SimulationCoordinator/CreateSimulation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simulationCoordinatorClient) DestroySimulation(ctx context.Context, in *DestroySimulationRequest, opts ...grpc.CallOption) (*DestroySimulationResponse, error) {
	out := new(DestroySimulationResponse)
	err := c.cc.Invoke(ctx, "/simulations.SimulationCoordinator/DestroySimulation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simulationCoordinatorClient) ListModels(ctx context.Context, in *ListModelsRequest, opts ...grpc.CallOption) (*ListModelsResponse, error) {
	out := new(ListModelsResponse)
	err := c.cc.Invoke(ctx, "/simulations.SimulationCoordinator/ListModels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simulationCoordinatorClient) AddModel(ctx context.Context, in *AddModelRequest, opts ...grpc.CallOption) (*AddModelResponse, error) {
	out := new(AddModelResponse)
	err := c.cc.Invoke(ctx, "/simulations.SimulationCoordinator/AddModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simulationCoordinatorClient) RemoveModel(ctx context.Context, in *RemoveModelRequest, opts ...grpc.CallOption) (*RemoveModelResponse, error) {
	out := new(RemoveModelResponse)
	err := c.cc.Invoke(ctx, "/simulations.SimulationCoordinator/RemoveModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simulationCoordinatorClient) Logs(ctx context.Context, in *LogsRequest, opts ...grpc.CallOption) (SimulationCoordinator_LogsClient, error) {
	stream, err := c.cc.NewStream(ctx, &SimulationCoordinator_ServiceDesc.Streams[0], "/simulations.SimulationCoordinator/Logs", opts...)
	if err != nil {
		return nil, err
	}
	x := &simulationCoordinatorLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SimulationCoordinator_LogsClient interface {
	Recv() (*LogsResponse, error)
	grpc.ClientStream
}

type simulationCoordinatorLogsClient struct {
	grpc.ClientStream
}

func (x *simulationCoordinatorLogsClient) Recv() (*LogsResponse, error) {
	m := new(LogsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SimulationCoordinatorServer is the server API for SimulationCoordinator service.
// All implementations must embed UnimplementedSimulationCoordinatorServer
// for forward compatibility
type SimulationCoordinatorServer interface {
	ListSimulations(context.Context, *ListSimulationsRequest) (*ListSimulationsResponse, error)
	CreateSimulation(context.Context, *CreateSimulationRequest) (*CreateSimulationResponse, error)
	DestroySimulation(context.Context, *DestroySimulationRequest) (*DestroySimulationResponse, error)
	ListModels(context.Context, *ListModelsRequest) (*ListModelsResponse, error)
	AddModel(context.Context, *AddModelRequest) (*AddModelResponse, error)
	RemoveModel(context.Context, *RemoveModelRequest) (*RemoveModelResponse, error)
	Logs(*LogsRequest, SimulationCoordinator_LogsServer) error
	mustEmbedUnimplementedSimulationCoordinatorServer()
}

// UnimplementedSimulationCoordinatorServer must be embedded to have forward compatible implementations.
type UnimplementedSimulationCoordinatorServer struct {
}

func (UnimplementedSimulationCoordinatorServer) ListSimulations(context.Context, *ListSimulationsRequest) (*ListSimulationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSimulations not implemented")
}
func (UnimplementedSimulationCoordinatorServer) CreateSimulation(context.Context, *CreateSimulationRequest) (*CreateSimulationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSimulation not implemented")
}
func (UnimplementedSimulationCoordinatorServer) DestroySimulation(context.Context, *DestroySimulationRequest) (*DestroySimulationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DestroySimulation not implemented")
}
func (UnimplementedSimulationCoordinatorServer) ListModels(context.Context, *ListModelsRequest) (*ListModelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListModels not implemented")
}
func (UnimplementedSimulationCoordinatorServer) AddModel(context.Context, *AddModelRequest) (*AddModelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddModel not implemented")
}
func (UnimplementedSimulationCoordinatorServer) RemoveModel(context.Context, *RemoveModelRequest) (*RemoveModelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveModel not implemented")
}
func (UnimplementedSimulationCoordinatorServer) Logs(*LogsRequest, SimulationCoordinator_LogsServer) error {
	return status.Errorf(codes.Unimplemented, "method Logs not implemented")
}
func (UnimplementedSimulationCoordinatorServer) mustEmbedUnimplementedSimulationCoordinatorServer() {}

// UnsafeSimulationCoordinatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SimulationCoordinatorServer will
// result in compilation errors.
type UnsafeSimulationCoordinatorServer interface {
	mustEmbedUnimplementedSimulationCoordinatorServer()
}

func RegisterSimulationCoordinatorServer(s grpc.ServiceRegistrar, srv SimulationCoordinatorServer) {
	s.RegisterService(&SimulationCoordinator_ServiceDesc, srv)
}

func _SimulationCoordinator_ListSimulations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSimulationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimulationCoordinatorServer).ListSimulations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simulations.SimulationCoordinator/ListSimulations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimulationCoordinatorServer).ListSimulations(ctx, req.(*ListSimulationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimulationCoordinator_CreateSimulation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSimulationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimulationCoordinatorServer).CreateSimulation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simulations.SimulationCoordinator/CreateSimulation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimulationCoordinatorServer).CreateSimulation(ctx, req.(*CreateSimulationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimulationCoordinator_DestroySimulation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DestroySimulationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimulationCoordinatorServer).DestroySimulation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simulations.SimulationCoordinator/DestroySimulation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimulationCoordinatorServer).DestroySimulation(ctx, req.(*DestroySimulationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimulationCoordinator_ListModels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListModelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimulationCoordinatorServer).ListModels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simulations.SimulationCoordinator/ListModels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimulationCoordinatorServer).ListModels(ctx, req.(*ListModelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimulationCoordinator_AddModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimulationCoordinatorServer).AddModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simulations.SimulationCoordinator/AddModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimulationCoordinatorServer).AddModel(ctx, req.(*AddModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimulationCoordinator_RemoveModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimulationCoordinatorServer).RemoveModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simulations.SimulationCoordinator/RemoveModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimulationCoordinatorServer).RemoveModel(ctx, req.(*RemoveModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimulationCoordinator_Logs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LogsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SimulationCoordinatorServer).Logs(m, &simulationCoordinatorLogsServer{stream})
}

type SimulationCoordinator_LogsServer interface {
	Send(*LogsResponse) error
	grpc.ServerStream
}

type simulationCoordinatorLogsServer struct {
	grpc.ServerStream
}

func (x *simulationCoordinatorLogsServer) Send(m *LogsResponse) error {
	return x.ServerStream.SendMsg(m)
}

// SimulationCoordinator_ServiceDesc is the grpc.ServiceDesc for SimulationCoordinator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SimulationCoordinator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "simulations.SimulationCoordinator",
	HandlerType: (*SimulationCoordinatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListSimulations",
			Handler:    _SimulationCoordinator_ListSimulations_Handler,
		},
		{
			MethodName: "CreateSimulation",
			Handler:    _SimulationCoordinator_CreateSimulation_Handler,
		},
		{
			MethodName: "DestroySimulation",
			Handler:    _SimulationCoordinator_DestroySimulation_Handler,
		},
		{
			MethodName: "ListModels",
			Handler:    _SimulationCoordinator_ListModels_Handler,
		},
		{
			MethodName: "AddModel",
			Handler:    _SimulationCoordinator_AddModel_Handler,
		},
		{
			MethodName: "RemoveModel",
			Handler:    _SimulationCoordinator_RemoveModel_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Logs",
			Handler:       _SimulationCoordinator_Logs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "simulations/simulations.proto",
}
