package global

import (
	"backend/pkg/cache"
	"database/sql"
	"fmt"
)

type Storage struct {
	ID         string `json:"id"`
	Key        string `json:"key"`
	Value      string `json:"value"`
	Readonly   bool   `json:"readonly"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

type storageDatabaseCache struct {
	storage   cache.ILocalDataCache
	db        *sql.DB
	tablename string
}

func NewStorageDatabaseCache(db *sql.DB) *storageDatabaseCache {
	return &storageDatabaseCache{
		storage:   cache.NewLocalDataCache(),
		db:        db,
		tablename: "storage",
	}
}

// init local memeory data from database
func (p *storageDatabaseCache) InitCacheFromDB() error {
	query := fmt.Sprintf(`SELECT id, key, value, readonly, create_time, update_time FROM %s`, p.tablename)
	rows, err := p.db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		tmp := new(Storage)
		err := rows.Scan(&tmp.ID, &tmp.Key, &tmp.Value, &tmp.Readonly, &tmp.CreateTime, &tmp.UpdateTime)
		if err != nil {
			return err
		}
		p.storage.Add(tmp.Key, tmp)
	}
	return nil
}

func (p *storageDatabaseCache) DataLength() int {
	return p.storage.Count()
}

// insert data to local memory and database
func (p *storageDatabaseCache) Insert(key string, value string, readonly bool) (*Storage, error) {
	tmp := &Storage{
		Key:      key,
		Value:    value,
		Readonly: readonly,
	}

	query := fmt.Sprintf(`INSERT INTO %s(key, value, readonly) 
		VALUES($1, $2, $3)
		RETURNING id, create_time, update_time;`, p.tablename)
	err := p.db.QueryRow(query, key, value, readonly).Scan(&tmp.ID, &tmp.CreateTime, &tmp.UpdateTime)
	if err != nil {
		return nil, err
	}

	p.storage.Add(tmp.Key, tmp)
	return tmp, nil
}

// get all data from local memory
func (p *storageDatabaseCache) Select() ([]*Storage, error) {
	tmps := make([]*Storage, 0)
	mapTmp := p.storage.GetAll()
	for _, val := range mapTmp {
		if tmp, ok := val.(*Storage); ok {
			tmps = append(tmps, tmp)
		}
	}
	return tmps, nil
}

// get data from local memory
func (p *storageDatabaseCache) SelectFromDB(key string) (*Storage, bool) {

	tmp := new(Storage)
	query := fmt.Sprintf(`SELECT id, key, value, readonly, create_time, update_time FROM %s WHERE key=$1;`, p.tablename)
	err := p.db.QueryRow(query, key).Scan(&tmp.ID, &tmp.Key, &tmp.Value, &tmp.Readonly, &tmp.CreateTime, &tmp.UpdateTime)
	if err != nil {
		return nil, false
	}

	return tmp, true
}

// get data from local memory
func (p *storageDatabaseCache) SelectOne(key string) (*Storage, bool) {
	if val, ok := p.storage.Get(key); ok {
		if tmp, ok := val.(*Storage); ok {
			return tmp, (tmp != nil)
		}
	}
	return nil, false
}

// get data from local memory
// return (*Storage, bool), [select or insert object, is exist objct]
func (p *storageDatabaseCache) SelectOrInsertOne(key string, value string, readonly bool) (*Storage, bool) {
	if val, ok := p.storage.Get(key); ok {
		if tmp, ok := val.(*Storage); ok {
			return tmp, ok
		}
	} else {
		s, err := p.Insert(key, value, readonly)
		if err == nil {
			return s, false
		}
	}
	return nil, false
}

// update data to local memory and database
func (p *storageDatabaseCache) Update(key string, value string) error {
	tmp, ok := p.SelectOne(key)
	if !ok {
		return fmt.Errorf("can't update date, key: %s, readonly is %t", tmp.Key, tmp.Readonly)
	}
	if !tmp.Readonly {
		query := fmt.Sprintf(`UPDATE %s 
		SET value=$1, update_time=now() 
		WHERE key=$2`, p.tablename)
		_, err := p.db.Exec(query, value, key)
		if err != nil {
			return err
		}

		tmp.Value = value
		p.storage.Add(tmp.Key, tmp)
		return nil
	} else {
		return fmt.Errorf("can't update date, key: %s, readonly is %t", tmp.Key, tmp.Readonly)
	}

}

// delete data to local memory and database
func (p *storageDatabaseCache) Delete(key string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE key=$1`, p.tablename)
	_, err := p.db.Exec(query, key)
	if err != nil {
		return err
	}

	p.storage.Remove(key)
	return nil
}
