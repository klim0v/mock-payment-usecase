package order

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/introphin/proto-order"
	"google.golang.org/grpc"
)

type Order struct {
}

func (o *Order) ListUserPayments(ctx context.Context, in *order.ListUserPaymentsRequest, opts ...grpc.CallOption) (*order.ListUserPaymentsResponse, error) {
	panic("implement me")
}

func (o *Order) ResetHold(ctx context.Context, in *order.ResetHoldRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	panic("implement me")
}

func (o *Order) CreateOrder(ctx context.Context, in *order.CreateOrderRequest, opts ...grpc.CallOption) (*order.Order, error) {
	panic("implement me")
}

func (o *Order) ListOrders(ctx context.Context, in *order.ListOrdersRequest, opts ...grpc.CallOption) (*order.ListOrdersResponse, error) {
	panic("implement me")
}
