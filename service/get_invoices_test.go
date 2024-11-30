package service

import (
	"testing"
)

func Test_GetInvoices(t *testing.T) {
	//testTime := time.Date(2024, 11, 30, 0, 0, 0, 0, time.UTC)
	//testDueDate := testTime.AddDate(0, 0, 10) // 10 days past

	//testCases := []struct {
	//	name      string
	//	dbFactory func(mockCtrl *gomock.Controller) db.DB
	//	request   api.GetInvoicesRequestObject
	//	expected  api.GetInvoicesResponseObject
	//	err       bool
	//}{
	//	{
	//		name: "retrieve_invoices_successfully",
	//dbFactory: func(mockCtrl *gomock.Controller) db.DB {
	//	dbService := db.NewMockDB(mockCtrl)
	//
	//	dbService.EXPECT().GetInvoices(gomock.Any(), &testTime, nil).
	//		Return(
	//			[]model.Invoice{
	//				{
	//					ID:             "in_example",
	//					PublishedBy:    "0be62c928261",
	//					PublishedAt:    testTime,
	//					ReceivedBy:     "bdb2476fd525",
	//					Commission:     400,
	//					CommissionRate: 4,
	//					Tax:            40,
	//					TaxRate:        10,
	//					PaymentAmount:  10000,
	//					BilledAmount:   10440,
	//					DueDate:        testTime.AddDate(0, 1, 0),
	//					Status:         model.InvoiceStatusUnPaid,
	//				},
	//			})
	//	return dbService
	//},
	//request: api.GetInvoicesRequestObject{
	//	Params: api.GetInvoicesParams{
	//		PeriodStartAt: &testTime,
	//	},
	//},
	//expected: api.GetInvoices200JSONResponse{
	//	{
	//		BilledAmount:   10440,
	//		Commission:     400,
	//		CommissionRate: 4,
	//		DueDate:        testTime.AddDate(0, 1, 0),
	//		Id:             "in_example",
	//		PaymentAmount:  10000,
	//		PublishedAt:    testTime,
	//		Status:         "UNPAID",
	//		Tax:            40,
	//		TaxRate:        10,
	//	},
	//},
	//err: false,
	//	},
	//}

	//for _, testCase := range testCases {
	// if someone still uses an old Go version,
	//   the same pointer is referenced in the loop, hence re-define local variable, just in case.
	//tc := testCase
	//t.Run(tc.name, func(t *testing.T) {
	//	mockCtrl := gomock.NewController(t)
	//	defer mockCtrl.Finish()
	//
	//	t.Parallel()
	//
	//	service := New(
	//		tc.dbFactory(mockCtrl),
	//	)
	//
	//	actual, err := service.GetInvoices(context.Background(), tc.request)
	//	if !tc.err {
	//		if err != nil {
	//			t.Error("want nil but error is received")
	//		}
	//
	//		if diff := cmp.Diff(
	//			tc.expected,
	//			actual,
	//		); diff != "" {
	//			t.Error("want (-want), but got (-actual)", diff)
	//		}
	//	} else {
	//		if err == nil {
	//			t.Error("want error but error is nil")
	//		}
	//	}
	//})
	//}
}
