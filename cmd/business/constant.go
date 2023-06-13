package business

const (
	YY           = "2006"
	YYMM         = "2006-01"
	YYMMDD       = "2006-01-02"
	YYMMDDHH     = "2006-01-02 15"
	YYMMDDHHMM   = "2006-01-02 15:04"
	YYMMDDHHMMSS = "2006-01-02 15:04:05"
)

const (
	/* 订单状态 begin */
	CANCELLED_STATE    = -10 //已取消
	WAIT_PAYMENT_STATE = 10  //新创建待付款
	PAID_STATE         = 20  //已支付待完成
	FINISHED_STATE     = 40  //已完成待结算
	SETTLED_STATE      = 50  //已结算
	IS_CLOSED          = -20 //已关闭
	/* 订单状态 end */
)
