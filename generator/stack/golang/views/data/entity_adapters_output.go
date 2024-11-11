package data

type OutputAdapters struct {
	Memory       *MemoryAdapter
	MongoDB      *MongoDBAdapter
	Redis        *RedisAdapter
	Sql          *SqlAdapter
	RedisSql     bool
	RedisMongoDB bool
}
