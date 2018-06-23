package ch9

var deposits = make(chan int)
var balances = make(chan int)
var withdraws = make(chan *withdraw)

type withdraw struct {
	amount int
	ok chan bool
}

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <- balances
}

func Withdraw(amount int) bool {
	draw := new(withdraw)
	draw.amount = amount
	draw.ok = make(chan bool)		// 把同道初始化 如果不初始化,同道的值为nil, 那么不管是发送还是接收都会当前goroutine造成永久阻塞!

	withdraws <- draw

	ok := <- draw.ok

	return ok
}

func teller() {
	var balance int
	for {
		select {
		case amount := <- deposits:
			balance += amount
		case balances <- balance:
		case draw := <- withdraws:
			if draw.amount <= balance {
				balance -= draw.amount

				draw.ok <- true
			}else {

				draw.ok <- false

			}
		}
	}
}

func init() {
	go teller()
}
