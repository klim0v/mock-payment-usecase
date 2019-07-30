package payment_repository

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/introphin/proto-payment-repository"
	"google.golang.org/grpc"
)

type MockBatchAcceptPaymentRepository struct {
	*MockPaymentRepository
}

func (*MockBatchAcceptPaymentRepository) ListInvoices(ctx context.Context, in *payment_repository.ListInvoicesRequest, opts ...grpc.CallOption) (*payment_repository.ListInvoicesResponse, error) {
	return &payment_repository.ListInvoicesResponse{
		Invoices: []*payment_repository.Invoice{
			{
				Id:                    987,
				CreatedAt:             &timestamp.Timestamp{Seconds: 1564059069},
				UpdatedAt:             ptypes.TimestampNow(),
				WalletId:              100,
				Sum:                   512,
				Status:                payment_repository.InvoiceStatus_Rejected,
				Purpose:               "Payout",
				WalletUserId:          210,
				WalletPaymentSystemId: 5,
				WalletAccountNumber:   "R969235928630",
				WalletCurrency:        "RUB",
				WalletChecked:         true,
				EditedUserId:          1,
			},
			{
				Id:                    988,
				CreatedAt:             &timestamp.Timestamp{Seconds: 1564060931},
				UpdatedAt:             &timestamp.Timestamp{Seconds: 1564060931},
				WalletId:              101,
				Sum:                   434,
				Status:                payment_repository.InvoiceStatus_Processing,
				Purpose:               "Payout",
				WalletUserId:          214,
				WalletPaymentSystemId: 5,
				WalletAccountNumber:   "R969235928111",
				WalletCurrency:        "RUB",
				WalletChecked:         false,
				EditedUserId:          214,
			},
			{
				Id:                    989,
				CreatedAt:             &timestamp.Timestamp{Seconds: 1564060000},
				UpdatedAt:             &timestamp.Timestamp{Seconds: 1564060999},
				WalletId:              104,
				Sum:                   123,
				Status:                payment_repository.InvoiceStatus_Paid,
				Purpose:               "Payout",
				WalletUserId:          214,
				WalletPaymentSystemId: 5,
				WalletAccountNumber:   "Z291143207289",
				WalletCurrency:        "USD",
				WalletChecked:         true,
				EditedUserId:          1,
			},
		},
		Total: 3,
	}, nil
}

func (*MockBatchAcceptPaymentRepository) UpdateInvoice(ctx context.Context, in *payment_repository.UpdateInvoiceRequest, opts ...grpc.CallOption) (*payment_repository.Invoice, error) {
	return &payment_repository.Invoice{
		Id:                    in.Id,
		CreatedAt:             &timestamp.Timestamp{Seconds: 1564060000},
		UpdatedAt:             ptypes.TimestampNow(),
		WalletId:              in.WalletId,
		Sum:                   in.Sum,
		Status:                in.Status,
		Purpose:               in.Purpose,
		WalletUserId:          in.WalletUserId,
		WalletPaymentSystemId: in.WalletPaymentSystemId,
		WalletAccountNumber:   in.WalletAccountNumber,
		WalletCurrency:        in.WalletCurrency,
		WalletChecked:         in.WalletChecked,
		EditedUserId:          in.EditedUserId,
	}, nil
}
