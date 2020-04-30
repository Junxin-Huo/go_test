package main

import (
	"fmt"
)

// 租金扣除规则（不算逾期等）
func pay(price, day int) int {
	if price >= 100 {
		if day <= 15 {
			return 5 * day
		}
		return 5*15 + 3*(day-15)
	} else if price >= 50 {
		if day <= 15 {
			return 3 * day
		}
		return 3*15 + 2*(day-15)
	} else {
		return 1 * day
	}
}

// 借阅书后的余额
func borrow(total, price, reserveDay, actualDay int) int {
	// 余额小于当前所借书价格，不能借该书
	if total < price {
		return total
	}

	// 预借需要花费
	prePay := pay(price, reserveDay)
	if prePay > price {
		prePay = price
	}

	// 预借时，余额不够
	if total < prePay {
		return total
	}

	// 实际花费
	aftPay := pay(price, actualDay)
	// 先判断超期额外扣除余额，最后再判断扣除金额不超过价格
	if actualDay > reserveDay {
		aftPay = aftPay + (actualDay - reserveDay)
	}
	if aftPay > price {
		aftPay = price
	}

	return total - aftPay
}

func main() {
	total := 300

	price := 0
	reserve := 0
	actual := 0
	for {
		n, _ := fmt.Scan(&price)
		if n == 0 {
			break
		}
		n, _ = fmt.Scan(&reserve)
		if n == 0 {
			break
		}
		n, _ = fmt.Scan(&actual)
		if n == 0 {
			break
		}

		total = borrow(total, price, reserve, actual)

		//fmt.Printf("%d, %d, %d, %d\n", price, reserve, actual, total)
	}
	fmt.Printf("%d\n", total)
}
