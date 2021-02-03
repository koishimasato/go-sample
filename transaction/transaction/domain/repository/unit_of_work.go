package repository

type UnitOfWork interface {
	RegisterNew(interface{})
	RegisterDirty(interface{})
	RegisterClean(interface{})
	RegisterDeleted(interface{})
	Commit() error
}

var Current UnitOfWork

type Entity struct {
}

func (e *Entity) MarkNew() {
	Current.RegisterNew(e)
}

func (e *Entity) MarkClean() {
	Current.RegisterClean(e)
}

func (e *Entity) MarkDirty() {
	Current.RegisterDirty(e)
}

func (e *Entity) MarkDeleted() {
	Current.RegisterDeleted(e)
}
