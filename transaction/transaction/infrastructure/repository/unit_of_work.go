package repository

type UnitOfWorkImpl struct {
	newEntities []interface{}
	dirtyEntities []interface{}
	cleanEntities []interface{}
	deletedEntities []interface{}
}

func (u *UnitOfWorkImpl) RegisterNew(entity interface{}) {
	u.newEntities = append(u.newEntities, entity)
}

func (u *UnitOfWorkImpl) RegisterDirty(entity interface{}) {
	u.dirtyEntities = append(u.dirtyEntities, entity)
}

func (u *UnitOfWorkImpl) RegisterClean(entity interface{}) {
	u.cleanEntities = append(u.cleanEntities, entity)
}

func (u *UnitOfWorkImpl) RegisterDeleted(entity interface{}) {
	u.deletedEntities = append(u.deletedEntities, entity)
}

func (u *UnitOfWorkImpl) Commit() error {
	// トランザクションを貼り、DB保存実行
	for _, e := range u.newEntities {
		u.save(e)
	}
	// ...

	return nil
}

func (u *UnitOfWorkImpl) save(entity interface{}) {
	// 略
}
