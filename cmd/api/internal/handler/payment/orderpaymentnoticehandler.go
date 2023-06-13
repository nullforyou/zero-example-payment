package payment

import (
	"go-zero-base/utils/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"payment/cmd/api/internal/logic/payment"
	"payment/cmd/api/internal/svc"
	"payment/cmd/api/internal/types"
)

func OrderPaymentNoticeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PaymentNoticeReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParseParamErrResponse(r, w, err)
			return
		}

		if err := svcCtx.Validator.Validate.StructCtx(r.Context(), req); err != nil {
			response.ValidateErrResponse(r, w, err, svcCtx.Validator.Trans)
			return
		}

		l := payment.NewOrderPaymentNoticeLogic(r.Context(), svcCtx)
		resp, err := l.OrderPaymentNotice(&req)
		response.Response(r, w, resp, err)
	}
}
