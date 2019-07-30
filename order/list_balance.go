package order

import (
	"context"
	"github.com/introphin/proto-order"
	"google.golang.org/grpc"
)

type MockListBalanceOrder struct {
	*Order
}

func (o *MockListBalanceOrder) ListUserPayments(ctx context.Context, in *order.ListUserPaymentsRequest, opts ...grpc.CallOption) (*order.ListUserPaymentsResponse, error) {
	userId := in.UserId[0]
	return &order.ListUserPaymentsResponse{
		UserPayments: []*order.UserPayment{
			{
				UserId:   userId,
				Currency: "RUB",
				Sum:      99923,
				Type:     order.PaymentType_PaymentDefault,
			},
			{
				UserId:   userId,
				Currency: "USD",
				Sum:      12432,
				Type:     order.PaymentType_PaymentDefault,
			},
			{
				UserId:   userId,
				Currency: "EUR",
				Sum:      5680,
				Type:     order.PaymentType_PaymentDefault,
			},
		},
		Total: 0,
	}, nil
}
