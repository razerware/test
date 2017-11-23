package service

import "errors"

var Ch=make(chan int,10)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("Divided by zero!")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}


func (t *Arith) Lzy(args Args, reply *int) error{
	*reply = args.A * args.B
	Ch<-*reply
	return nil
}