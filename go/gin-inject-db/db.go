package main

// A mock DB for demonstration purposes

type DB struct {
	db map[string]string
}

func NewDB() *DB {
	return &DB{db: make(map[string]string)}
}

func PreloadDB(m map[string]string) *DB {
	db := &DB{db: make(map[string]string)}
	for k, v := range m {
		db.db[k] = v
	}
	return db
}

func (db *DB) List() []string {
	resp := make([]string, 0, len(db.db))
	for k := range db.db {
		resp = append(resp, k)
	}
	return resp
}

func (db *DB) Put(k, v string) {
	db.db[k] = v
}

func (db *DB) Get(k string) (v string, ok bool) {
	v, ok = db.db[k]
	return
}

func (db *DB) Del(k string) (ok bool) {
	if _, ok = db.db[k]; ok {
		delete(db.db, k)
	}
	return
}
