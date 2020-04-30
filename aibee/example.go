package main

import "fmt"

func get(url string) (string, error) {
	return "", nil
}

func download() error {
	url := []string{""}
	c := make(chan string)
	defer close(c)
	data := []string{}

	errChan := make(chan error)
	dataChan := make(chan string)
	defer close(errChan)
	defer close(dataChan)

	go func() {
		for _, v := range url {
			c <- v
		}
	}()

	for i := 1; i<10; i++ {
		go func() {
			for u := range c {
				d, err := get(u)
				if err != nil {
					d, err = get(u)
					if err != nil {
						errChan <- err
					}
				} else {
					dataChan <- d
				}
			}
			//for {
			//	select {
			//	case u := <- c:
			//		d, err := get(u)
			//		if err != nil {
			//			d, err = get(u)
			//			if err != nil {
			//				errChan <- err
			//			}
			//		} else {
			//			dataChan <- d
			//		}
			//		count ++
			//	}
			//}


			//u := <- c
			//d, err := get(u)
			//if err != nil {
			//	d, err = get(u)
			//	if err != nil {
			//		errChan <- err
			//	}
			//} else {
			//	dataChan <- d
			//}
			//count ++
		}()
	}

	for  {
		select {
		case err := <- errChan:
			return err
		case d := <- dataChan:
			data = append(data, d)
			if len(data) == len(url) {
				// 写文件
				return nil
			}
		}
	}


}

/*

book:
id count amount(单价)

user:
id money（余额）

userbook（流水）:
userid bookid count（数量）



根据bookid，查询书单价 bookamount
根据userid，查询余额 useramount
判断余额是否充足 useramount 》= bookamount * count




 */

type node struct {
	a int
	pre *node
	next *node
}

type list struct {
	head *node
	tail *node
	length int
}

func (l *list) delLeast () {
	if l.tail == nil {
		return
	}

	tmp := l.tail.pre
	if tmp != nil {
		tmp.next = nil
	}
	l.tail = tmp
}

//func (l *list) addHead () {
//	if l.head == nil {
//		return
//	}
//
//	tmp := l.tail.pre
//	if tmp != nil {
//		tmp.next = nil
//	}
//	l.tail = tmp
//}


func (l *list) add (data int) {
	if l.head == nil {
		newNode := &node{
			a:    data,
			pre:  nil,
			next: nil,
		}
		l.head = newNode
		l.tail = newNode
		l.length = 1
		return
	}

	tmp := l.head
	for tmp != nil {
		if data >= tmp.a {
			newNode := &node{
				a:    data,
				pre:  tmp.pre,
				next: tmp,
			}
			tmp.pre.next = newNode
			tmp.pre = newNode
			l.length ++
			return
		}
		tmp = tmp.next
	}

	newNode := &node{
		a:    data,
		pre:  l.tail,
		next: nil,
	}
	l.tail.next = newNode
	l.tail = newNode
}

func f() {
	data := []int{3,2,3,1,2,4,5,5,6}
	k := 4

	l := list{
		head:   nil,
		tail:   nil,
		length: 0,
	}

	for _, v := range data {
		if l.length < k {
			l.add(v)
		} else {
			if v > l.tail.a {
				l.add(v)
				l.delLeast()
			}
		}
	}

	fmt.Println(l.tail.a)
}
