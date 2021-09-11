package localTickit

type LocalTicket struct {
	LocalBalance int64
	LocalSales int64
}

func NewLocalTicket(balance int64, sales int64) *LocalTicket {
	return &LocalTicket{
		LocalBalance: balance,
		LocalSales: sales,
	}
}

// 本地扣庫存, 並返回boolean, 當"賣出量"大於"庫存量"時, return false 表示已售完
func (tickets *LocalTicket) LocalDeductionTicket() bool {
	tickets.LocalSales++
	return tickets.LocalBalance >= tickets.LocalSales
}
