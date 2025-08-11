package identifier

type UserID uint64

func (id UserID) Int() uint64 { return uint64(id) }
