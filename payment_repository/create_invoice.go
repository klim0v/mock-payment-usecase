package payment_repository

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/introphin/proto-payment-repository"
	"google.golang.org/grpc"
)

type MockCreateInvoicePaymentRepository struct {
	*MockPaymentRepository
}

func (*MockCreateInvoicePaymentRepository) GetWallet(ctx context.Context, in *payment_repository.GetWalletRequest, opts ...grpc.CallOption) (*payment_repository.Wallet, error) {
	return &payment_repository.Wallet{
		Id:              in.Id,
		CreatedAt:       &timestamp.Timestamp{Seconds: 1564059069},
		UpdatedAt:       &timestamp.Timestamp{Seconds: 1564059999},
		UserId:          2,
		PaymentSystemId: 5,
		Name:            "",
		Currency:        "RUB",
		AutoPayment:     false,
		AccountNumber:   "R969235928111",
		Checked:         true,
		EditedUserId:    1,
		Type:            payment_repository.WalletType_WalletPersonal,
	}, nil
}

func (*MockCreateInvoicePaymentRepository) ListInvoices(ctx context.Context, in *payment_repository.ListInvoicesRequest, opts ...grpc.CallOption) (*payment_repository.ListInvoicesResponse, error) {
	return &payment_repository.ListInvoicesResponse{
		Invoices: []*payment_repository.Invoice{},
		Total:    0,
	}, nil
}

func (*MockCreateInvoicePaymentRepository) CreateInvoice(ctx context.Context, in *payment_repository.CreateInvoiceRequest, opts ...grpc.CallOption) (*payment_repository.Invoice, error) {
	return &payment_repository.Invoice{
		Id:                    2,
		CreatedAt:             ptypes.TimestampNow(),
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

func (*MockCreateInvoicePaymentRepository) ListSumAffiliatePayout(ctx context.Context, in *payment_repository.ListSumAffiliatePayoutRequest, opts ...grpc.CallOption) (*payment_repository.ListSumAffiliatePayoutResponse, error) {
	userId := in.UserIds[0]
	return &payment_repository.ListSumAffiliatePayoutResponse{
		SumAffiliatePayouts: []*payment_repository.SumAffiliatePayout{
			{
				UserId:   userId,
				Sum:      528.43,
				Currency: "RUB",
			},
			{
				UserId:   userId,
				Sum:      1300.65,
				Currency: "USD",
			},
			{
				UserId:   userId,
				Sum:      30.7,
				Currency: "EUR",
			},
		},
		Total: 3,
	}, nil
}
