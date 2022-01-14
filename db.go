package lorm

type DB struct {
	db   DBer
	isTx bool

	dbConfig DbConfig
	ctx      OrmContext

	dialect Dialect

	//where tokens
	whereTokens []string

	extraWhereSql []byte

	orderByTokens []string

	limit  int64
	offset int64

	//where values
	args      []interface{}
	batchArgs [][]interface{}
}

func (db DB) OrmConf(c *OrmConf) DB {
	if c == nil {
		return db
	}
	db.ctx.conf = *c
	return db
}

type OrmSelect struct {
	base EngineBase
}

type OrmFrom struct {
	base EngineBase
}

type OrmWhere struct {
	base EngineBase
}
