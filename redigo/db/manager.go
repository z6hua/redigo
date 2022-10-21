package db

type DB struct {
	Ht map[string]string
}

func (db DB) GetAll() map[string]string {
	return db.Ht
}

func (db DB) Get(key string) string {
	if val, ok := db.Ht[key]; ok {
		return val
	}
	return "nil"
}

func (db DB) Set(key string, val string) string {
	db.Ht[key] = val
	return "ok"
}

func (db DB) Del(key string) string {
	delete(db.Ht, key)
	return "ok"
}
