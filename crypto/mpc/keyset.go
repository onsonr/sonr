package mpc

type Keyset interface {
	Address() string
	Val() *ValKeyshare
	ValJSON() string
	User() *UserKeyshare
	UserJSON() string
}

type keyset struct {
	val  *ValKeyshare
	user *UserKeyshare
	addr string
}

func (k keyset) Address() string {
	return k.addr
}

func (k keyset) Val() *ValKeyshare {
	return k.val
}

func (k keyset) User() *UserKeyshare {
	return k.user
}

func (k keyset) ValJSON() string {
	return k.val.String()
}

func (k keyset) UserJSON() string {
	return k.user.String()
}
