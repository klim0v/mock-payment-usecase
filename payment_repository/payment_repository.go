package payment_repository

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/introphin/proto-payment-repository"
	"google.golang.org/grpc"
)

type MockPaymentRepository struct {
}

func (m *MockPaymentRepository) ListWallets(ctx context.Context, in *payment_repository.ListWalletsRequest, opts ...grpc.CallOption) (*payment_repository.ListWalletsResponse, error) {
	panic("implement me")
}

func (m *MockPaymentRepository) GetWallet(ctx context.Context, in *payment_repository.GetWalletRequest, opts ...grpc.CallOption) (*payment_repository.Wallet, error) {
	panic("implement me")
}

func (m *MockPaymentRepository) CreateWallet(ctx context.Context, in *payment_repository.CreateWalletRequest, opts ...grpc.CallOption) (*payment_repository.Wallet, error) {
	panic("implement me")
}

func (m *MockPaymentRepository) UpdateWallet(ctx context.Context, in *payment_repository.UpdateWalletRequest, opts ...grpc.CallOption) (*payment_repository.Wallet, error) {
	panic("implement me")
}

func (m *MockPaymentRepository) DeleteWallet(ctx context.Context, in *payment_repository.DeleteWalletRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	panic("implement me")
}

func (m *MockPaymentRepository) ListInvoices(ctx context.Context, in *payment_repository.ListInvoicesRequest, opts ...grpc.CallOption) (*payment_repository.ListInvoicesResponse, error) {
	panic("implement me")
}

func (m *MockPaymentRepository) GetInvoice(ctx context.Context, in *payment_repository.GetInvoiceRequest, opts ...grpc.CallOption) (*payment_repository.Invoice, error) {
	panic("implement me")
}

func (m *MockPaymentRepository) CreateInvoice(ctx context.Context, in *payment_repository.CreateInvoiceRequest, opts ...grpc.CallOption) (*payment_repository.Invoice, error) {
	panic("implement me")
}

func (m *MockPaymentRepository) UpdateInvoice(ctx context.Context, in *payment_repository.UpdateInvoiceRequest, opts ...grpc.CallOption) (*payment_repository.Invoice, error) {
	panic("implement me")
}

func (m *MockPaymentRepository) DeleteInvoice(ctx context.Context, in *payment_repository.DeleteInvoiceRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	panic("implement me")
}

func (m *MockPaymentRepository) ListPaymentSystems(ctx context.Context, in *payment_repository.ListPaymentSystemsRequest, opts ...grpc.CallOption) (*payment_repository.ListPaymentSystemsResponse, error) {
	panic("implement me")
}

func (m *MockPaymentRepository) GetPaymentSystem(ctx context.Context, in *payment_repository.GetPaymentSystemRequest, opts ...grpc.CallOption) (*payment_repository.PaymentSystem, error) {
	panic("implement me")
}

func (m *MockPaymentRepository) CreatePaymentSystem(ctx context.Context, in *payment_repository.CreatePaymentSystemRequest, opts ...grpc.CallOption) (*payment_repository.PaymentSystem, error) {
	panic("implement me")
}

func (m *MockPaymentRepository) UpdatePaymentSystem(ctx context.Context, in *payment_repository.UpdatePaymentSystemRequest, opts ...grpc.CallOption) (*payment_repository.PaymentSystem, error) {
	panic("implement me")
}

func (m *MockPaymentRepository) DeletePaymentSystem(ctx context.Context, in *payment_repository.DeletePaymentSystemRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	panic("implement me")
}

func (m *MockPaymentRepository) ListBatches(ctx context.Context, in *payment_repository.ListBatchesRequest, opts ...grpc.CallOption) (*payment_repository.ListBatchesResponse, error) {
	panic("implement me")
}

func (m *MockPaymentRepository) GetBatch(ctx context.Context, in *payment_repository.GetBatchRequest, opts ...grpc.CallOption) (*payment_repository.Batch, error) {
	panic("implement me")
}

func (m *MockPaymentRepository) CreateBatch(ctx context.Context, in *payment_repository.CreateBatchRequest, opts ...grpc.CallOption) (*payment_repository.Batch, error) {
	panic("implement me")
}

func (m *MockPaymentRepository) DeleteBatch(ctx context.Context, in *payment_repository.DeleteBatchRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	panic("implement me")
}

func (m *MockPaymentRepository) ListAdvertiserPayments(ctx context.Context, in *payment_repository.ListAdvertiserPaymentsRequest, opts ...grpc.CallOption) (*payment_repository.ListAdvertiserPaymentsResponse, error) {
	panic("implement me")
}

func (m *MockPaymentRepository) GetAdvertiserPayment(ctx context.Context, in *payment_repository.GetAdvertiserPaymentRequest, opts ...grpc.CallOption) (*payment_repository.AdvertiserPayment, error) {
	panic("implement me")
}

func (m *MockPaymentRepository) CreateAdvertiserPayment(ctx context.Context, in *payment_repository.CreateAdvertiserPaymentRequest, opts ...grpc.CallOption) (*payment_repository.AdvertiserPayment, error) {
	panic("implement me")
}

func (m *MockPaymentRepository) UpdateAdvertiserPayment(ctx context.Context, in *payment_repository.UpdateAdvertiserPaymentRequest, opts ...grpc.CallOption) (*payment_repository.AdvertiserPayment, error) {
	panic("implement me")
}

func (m *MockPaymentRepository) DeleteAdvertiserPayment(ctx context.Context, in *payment_repository.DeleteAdvertiserPaymentRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	panic("implement me")
}

func (m *MockPaymentRepository) ListBatchInvoice(ctx context.Context, in *payment_repository.ListBatchInvoiceRequest, opts ...grpc.CallOption) (*payment_repository.ListBatchInvoiceResponse, error) {
	panic("implement me")
}

func (m *MockPaymentRepository) CreateBatchInvoice(ctx context.Context, in *payment_repository.CreateBatchInvoiceRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	panic("implement me")
}

func (m *MockPaymentRepository) DeleteBatchInvoice(ctx context.Context, in *payment_repository.DeleteBatchInvoiceRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	panic("implement me")
}

func (m *MockPaymentRepository) ListSumAdvertiserPayment(ctx context.Context, in *payment_repository.ListSumAdvertiserPaymentRequest, opts ...grpc.CallOption) (*payment_repository.ListSumAdvertiserPaymentResponse, error) {
	panic("implement me")
}

func (m *MockPaymentRepository) ListSumAffiliatePayout(ctx context.Context, in *payment_repository.ListSumAffiliatePayoutRequest, opts ...grpc.CallOption) (*payment_repository.ListSumAffiliatePayoutResponse, error) {
	panic("implement me")
}
