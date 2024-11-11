package v1

const (
	Version      = "version"
	Kind         = "kind"
	Entity       = "entity"
	Name         = "name"
	Plural       = "plural"
	Description  = "description"
	Docs         = "docs"
	Schema       = "schema"
	Reserved     = "reserved"
	Type         = "type"
	Fields       = "fields"
	Unique       = "unique"
	Encrypted    = "encrypted"
	Validate     = "validate"
	Regex        = "regex"
	Pattern      = "pattern"
	Go           = "go"
	Java         = "java"
	Value        = "value"
	Enum         = "enum"
	Range        = "range"
	Min          = "min"
	Max          = "max"
	Adapters     = "adapters"
	Map          = "map"
	Rel          = "$rel"
	Input        = "input"
	HTTP         = "http"
	ResourceName = "resourceName"
	Output       = "output"
	Memory       = "memory"
	MongoDB      = "mongodb"
	Redis        = "redis"
	SQL          = "sql"
	RedisMongoDB = "redis+mongodb"
	RedisSQL     = "redis+sql"
	TTL          = "ttl"
	Collection   = "collection"
	Table        = "table"
	HasMany      = "hasMany"
	HasOne       = "hasOne"
	BelongsTo    = "belongsTo"
	Nested       = "nested"
	MaxItems     = "maxItems"
)

// Entity field types
const (
	// Integer data type mapped to: Go=int64, Java=java.lang.Long
	Integer = "integer"

	// Float data type mapped to: Go=float64, Java=java.lang.Float
	Float = "float"

	// Boolean data type mapped to: Go=bool, Java=java.lang.Boolean
	Boolean = "boolean"

	// String data type mapped to: Go=string, Java=java.lang.String
	String = "string"

	// Date data type mapped to: Go=int64, Java=java.lang.Long
	// date representation as UTC milliseconds
	Date = "date"

	// UUID data type mapped to: Go=string, Java=java.lang.String
	UUID = "uuid"

	// ULID data type mapped to: Go=string, Java=java.lang.String
	ULID = "ulid"
)
