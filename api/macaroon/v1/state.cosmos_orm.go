// Code generated by protoc-gen-go-cosmos-orm. DO NOT EDIT.

package macaroonv1

import (
	context "context"
	ormlist "cosmossdk.io/orm/model/ormlist"
	ormtable "cosmossdk.io/orm/model/ormtable"
	ormerrors "cosmossdk.io/orm/types/ormerrors"
)

type GrantTable interface {
	Insert(ctx context.Context, grant *Grant) error
	InsertReturningId(ctx context.Context, grant *Grant) (uint64, error)
	LastInsertedSequence(ctx context.Context) (uint64, error)
	Update(ctx context.Context, grant *Grant) error
	Save(ctx context.Context, grant *Grant) error
	Delete(ctx context.Context, grant *Grant) error
	Has(ctx context.Context, id uint64) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, id uint64) (*Grant, error)
	HasBySubjectOrigin(ctx context.Context, subject string, origin string) (found bool, err error)
	// GetBySubjectOrigin returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetBySubjectOrigin(ctx context.Context, subject string, origin string) (*Grant, error)
	List(ctx context.Context, prefixKey GrantIndexKey, opts ...ormlist.Option) (GrantIterator, error)
	ListRange(ctx context.Context, from, to GrantIndexKey, opts ...ormlist.Option) (GrantIterator, error)
	DeleteBy(ctx context.Context, prefixKey GrantIndexKey) error
	DeleteRange(ctx context.Context, from, to GrantIndexKey) error

	doNotImplement()
}

type GrantIterator struct {
	ormtable.Iterator
}

func (i GrantIterator) Value() (*Grant, error) {
	var grant Grant
	err := i.UnmarshalMessage(&grant)
	return &grant, err
}

type GrantIndexKey interface {
	id() uint32
	values() []interface{}
	grantIndexKey()
}

// primary key starting index..
type GrantPrimaryKey = GrantIdIndexKey

type GrantIdIndexKey struct {
	vs []interface{}
}

func (x GrantIdIndexKey) id() uint32            { return 0 }
func (x GrantIdIndexKey) values() []interface{} { return x.vs }
func (x GrantIdIndexKey) grantIndexKey()        {}

func (this GrantIdIndexKey) WithId(id uint64) GrantIdIndexKey {
	this.vs = []interface{}{id}
	return this
}

type GrantSubjectOriginIndexKey struct {
	vs []interface{}
}

func (x GrantSubjectOriginIndexKey) id() uint32            { return 1 }
func (x GrantSubjectOriginIndexKey) values() []interface{} { return x.vs }
func (x GrantSubjectOriginIndexKey) grantIndexKey()        {}

func (this GrantSubjectOriginIndexKey) WithSubject(subject string) GrantSubjectOriginIndexKey {
	this.vs = []interface{}{subject}
	return this
}

func (this GrantSubjectOriginIndexKey) WithSubjectOrigin(subject string, origin string) GrantSubjectOriginIndexKey {
	this.vs = []interface{}{subject, origin}
	return this
}

type grantTable struct {
	table ormtable.AutoIncrementTable
}

func (this grantTable) Insert(ctx context.Context, grant *Grant) error {
	return this.table.Insert(ctx, grant)
}

func (this grantTable) Update(ctx context.Context, grant *Grant) error {
	return this.table.Update(ctx, grant)
}

func (this grantTable) Save(ctx context.Context, grant *Grant) error {
	return this.table.Save(ctx, grant)
}

func (this grantTable) Delete(ctx context.Context, grant *Grant) error {
	return this.table.Delete(ctx, grant)
}

func (this grantTable) InsertReturningId(ctx context.Context, grant *Grant) (uint64, error) {
	return this.table.InsertReturningPKey(ctx, grant)
}

func (this grantTable) LastInsertedSequence(ctx context.Context) (uint64, error) {
	return this.table.LastInsertedSequence(ctx)
}

func (this grantTable) Has(ctx context.Context, id uint64) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, id)
}

func (this grantTable) Get(ctx context.Context, id uint64) (*Grant, error) {
	var grant Grant
	found, err := this.table.PrimaryKey().Get(ctx, &grant, id)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &grant, nil
}

func (this grantTable) HasBySubjectOrigin(ctx context.Context, subject string, origin string) (found bool, err error) {
	return this.table.GetIndexByID(1).(ormtable.UniqueIndex).Has(ctx,
		subject,
		origin,
	)
}

func (this grantTable) GetBySubjectOrigin(ctx context.Context, subject string, origin string) (*Grant, error) {
	var grant Grant
	found, err := this.table.GetIndexByID(1).(ormtable.UniqueIndex).Get(ctx, &grant,
		subject,
		origin,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &grant, nil
}

func (this grantTable) List(ctx context.Context, prefixKey GrantIndexKey, opts ...ormlist.Option) (GrantIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return GrantIterator{it}, err
}

func (this grantTable) ListRange(ctx context.Context, from, to GrantIndexKey, opts ...ormlist.Option) (GrantIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return GrantIterator{it}, err
}

func (this grantTable) DeleteBy(ctx context.Context, prefixKey GrantIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this grantTable) DeleteRange(ctx context.Context, from, to GrantIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this grantTable) doNotImplement() {}

var _ GrantTable = grantTable{}

func NewGrantTable(db ormtable.Schema) (GrantTable, error) {
	table := db.GetTable(&Grant{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&Grant{}).ProtoReflect().Descriptor().FullName()))
	}
	return grantTable{table.(ormtable.AutoIncrementTable)}, nil
}

type MacaroonTable interface {
	Insert(ctx context.Context, macaroon *Macaroon) error
	InsertReturningId(ctx context.Context, macaroon *Macaroon) (uint64, error)
	LastInsertedSequence(ctx context.Context) (uint64, error)
	Update(ctx context.Context, macaroon *Macaroon) error
	Save(ctx context.Context, macaroon *Macaroon) error
	Delete(ctx context.Context, macaroon *Macaroon) error
	Has(ctx context.Context, id uint64) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, id uint64) (*Macaroon, error)
	HasBySubjectOrigin(ctx context.Context, subject string, origin string) (found bool, err error)
	// GetBySubjectOrigin returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetBySubjectOrigin(ctx context.Context, subject string, origin string) (*Macaroon, error)
	List(ctx context.Context, prefixKey MacaroonIndexKey, opts ...ormlist.Option) (MacaroonIterator, error)
	ListRange(ctx context.Context, from, to MacaroonIndexKey, opts ...ormlist.Option) (MacaroonIterator, error)
	DeleteBy(ctx context.Context, prefixKey MacaroonIndexKey) error
	DeleteRange(ctx context.Context, from, to MacaroonIndexKey) error

	doNotImplement()
}

type MacaroonIterator struct {
	ormtable.Iterator
}

func (i MacaroonIterator) Value() (*Macaroon, error) {
	var macaroon Macaroon
	err := i.UnmarshalMessage(&macaroon)
	return &macaroon, err
}

type MacaroonIndexKey interface {
	id() uint32
	values() []interface{}
	macaroonIndexKey()
}

// primary key starting index..
type MacaroonPrimaryKey = MacaroonIdIndexKey

type MacaroonIdIndexKey struct {
	vs []interface{}
}

func (x MacaroonIdIndexKey) id() uint32            { return 0 }
func (x MacaroonIdIndexKey) values() []interface{} { return x.vs }
func (x MacaroonIdIndexKey) macaroonIndexKey()     {}

func (this MacaroonIdIndexKey) WithId(id uint64) MacaroonIdIndexKey {
	this.vs = []interface{}{id}
	return this
}

type MacaroonSubjectOriginIndexKey struct {
	vs []interface{}
}

func (x MacaroonSubjectOriginIndexKey) id() uint32            { return 1 }
func (x MacaroonSubjectOriginIndexKey) values() []interface{} { return x.vs }
func (x MacaroonSubjectOriginIndexKey) macaroonIndexKey()     {}

func (this MacaroonSubjectOriginIndexKey) WithSubject(subject string) MacaroonSubjectOriginIndexKey {
	this.vs = []interface{}{subject}
	return this
}

func (this MacaroonSubjectOriginIndexKey) WithSubjectOrigin(subject string, origin string) MacaroonSubjectOriginIndexKey {
	this.vs = []interface{}{subject, origin}
	return this
}

type macaroonTable struct {
	table ormtable.AutoIncrementTable
}

func (this macaroonTable) Insert(ctx context.Context, macaroon *Macaroon) error {
	return this.table.Insert(ctx, macaroon)
}

func (this macaroonTable) Update(ctx context.Context, macaroon *Macaroon) error {
	return this.table.Update(ctx, macaroon)
}

func (this macaroonTable) Save(ctx context.Context, macaroon *Macaroon) error {
	return this.table.Save(ctx, macaroon)
}

func (this macaroonTable) Delete(ctx context.Context, macaroon *Macaroon) error {
	return this.table.Delete(ctx, macaroon)
}

func (this macaroonTable) InsertReturningId(ctx context.Context, macaroon *Macaroon) (uint64, error) {
	return this.table.InsertReturningPKey(ctx, macaroon)
}

func (this macaroonTable) LastInsertedSequence(ctx context.Context) (uint64, error) {
	return this.table.LastInsertedSequence(ctx)
}

func (this macaroonTable) Has(ctx context.Context, id uint64) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, id)
}

func (this macaroonTable) Get(ctx context.Context, id uint64) (*Macaroon, error) {
	var macaroon Macaroon
	found, err := this.table.PrimaryKey().Get(ctx, &macaroon, id)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &macaroon, nil
}

func (this macaroonTable) HasBySubjectOrigin(ctx context.Context, subject string, origin string) (found bool, err error) {
	return this.table.GetIndexByID(1).(ormtable.UniqueIndex).Has(ctx,
		subject,
		origin,
	)
}

func (this macaroonTable) GetBySubjectOrigin(ctx context.Context, subject string, origin string) (*Macaroon, error) {
	var macaroon Macaroon
	found, err := this.table.GetIndexByID(1).(ormtable.UniqueIndex).Get(ctx, &macaroon,
		subject,
		origin,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &macaroon, nil
}

func (this macaroonTable) List(ctx context.Context, prefixKey MacaroonIndexKey, opts ...ormlist.Option) (MacaroonIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return MacaroonIterator{it}, err
}

func (this macaroonTable) ListRange(ctx context.Context, from, to MacaroonIndexKey, opts ...ormlist.Option) (MacaroonIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return MacaroonIterator{it}, err
}

func (this macaroonTable) DeleteBy(ctx context.Context, prefixKey MacaroonIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this macaroonTable) DeleteRange(ctx context.Context, from, to MacaroonIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this macaroonTable) doNotImplement() {}

var _ MacaroonTable = macaroonTable{}

func NewMacaroonTable(db ormtable.Schema) (MacaroonTable, error) {
	table := db.GetTable(&Macaroon{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&Macaroon{}).ProtoReflect().Descriptor().FullName()))
	}
	return macaroonTable{table.(ormtable.AutoIncrementTable)}, nil
}

type StateStore interface {
	GrantTable() GrantTable
	MacaroonTable() MacaroonTable

	doNotImplement()
}

type stateStore struct {
	grant    GrantTable
	macaroon MacaroonTable
}

func (x stateStore) GrantTable() GrantTable {
	return x.grant
}

func (x stateStore) MacaroonTable() MacaroonTable {
	return x.macaroon
}

func (stateStore) doNotImplement() {}

var _ StateStore = stateStore{}

func NewStateStore(db ormtable.Schema) (StateStore, error) {
	grantTable, err := NewGrantTable(db)
	if err != nil {
		return nil, err
	}

	macaroonTable, err := NewMacaroonTable(db)
	if err != nil {
		return nil, err
	}

	return stateStore{
		grantTable,
		macaroonTable,
	}, nil
}
