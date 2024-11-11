## MongoDB hooks

Each method to insert, update, fetch or delete the entity is provided with before and after hooks, pretty similar to service hooks, 
but in this case works in the storage layer. 

This storage hooks are useful for instance database monitoring or intercept the DMO object before save it or after save it, or what you need.

The MongoDB Repository contains a reference to a `HookBuilder` struct which implements the interfaces: 

 - `ISaveHooks` defines the before and after methods to be call at db save operation
 - `IFetchByIdHooks` defines the before and after methods to be call at db fetch by id operation
 - `IUpdateHooks` defines the before and after methods to be call at db update operation
 - `IDeleteByIdHooks` defines the before and after methods to be call at db delete by id operation
 - `IDeleteByHooks` defines the before and after methods to be call at db delete by field operation
 - `IQueryHooks` defines the before and after methods to be call at db query operation

!!! info "Builder pattern"
    Hooks are following the same builder pattern as the [Service's hooks](advanced_service_hooks.html#why-a-builder-is-needed) 
    to ensure the hook `atomicity` and `non-shared` features. 

### Repository constructor

The hook builder instance should be set up at repository creation. In order to do it a constructor function is provided 
with a hook builder parameter. By default, a `NoOpHook` is provided.

```go

// NewMongoDBRepositoryWithHooks repository constructor
func NewMongoDBRepositoryWithHooks(collection *mongo.Collection, cryptoKey []byte, hooks *HookBuilder) *mongoDBRepository {
	...
}
```

## Hook struct

The `HookBuilder` requires an instance of the `Hooks` struct which can be implemented with the needed hook functions.

```go
type Hooks struct {
	Save       func() ISaveHooks
	FetchById  func() IFetchByIdHooks
	Update     func() IUpdateHooks
	DeleteById func() IDeleteByIdHooks
	DeleteBy   func() IDeleteByHooks
	Query      func() IQueryHooks
}
```

Each hook function will return the specific built instance to run. 
For instance, imagine that we would like to track the entity saving time and we have an object that provide us with the logic to start and track the consumed time:

```go
type LatencyTrackerHooks struct {
	track.TimeTracker
}

func (l *LatencyTrackerHooks) BeforeSave(dmo *EntityDMOMongoDB) error {
	return l.StartTime()
}

func (l *LatencyTrackerHooks) AfterSave(dmo *EntityDMOMongoDB, err error) error {
	if err != nil {
	    return l.EndTimeAndTrackItWithMetricName("")	
    }
}

func CreateRepository() {
    // creating save hook time tracker.
	hook := Hooks{
        Save: func() ISaveHooks {
            return new(LatencyTrackerHooks)
        }
    }

	// creating repo with hooks.
    repo := NewMongoDBRepositoryWithHooks("myCollectionName", "some-crypto-key", NewHookBuilder(hooks))	
}
```

### Save hooks

The **save** hook instances must implements the interface below:
```go
type ISaveHooks interface {
	BeforeSave(dmo *EntityDMOMongoDB) error
	AfterSave(dmo *EntityDMOMongoDB, err error) error
}
```

### Fetch by ID hooks

The **FetchById** hook instances must implements the interface below:

```go
type IFetchByIdHooks interface {
	BeforeFetchById(id string) error
	AfterFetchById(dmo *EntityDMOMongoDB, err error) error
}
```

### Update hooks

The **Update** hook instances must implements the interface below:
```go
type IUpdateHooks interface {
	BeforeUpdate(dmo *EntityDMOMongoDB) error
	AfterUpdate(dmo *EntityDMOMongoDB, err error) error
}
```

### Delete by ID hooks

The **DeleteById** hook instances must implements the interface below:
```go
type IDeleteByIdHooks interface {
	BeforeDeleteById(id string) error
	AfterDeleteById(err error) error
}
```

### Delete hooks

The **Delete** hook instances must implements the interface below:
```go
type IDeleteByHooks interface {
	BeforeDeleteBy(field, val string) error
	AfterDeleteBy(count int64, err error) error
}
```

### Query hooks

The **Query** hook instances must implements the interface below:
```go
type IQueryHooks interface {
	BeforeQuery(q *tql.Query) error
	AfterQuery(err error) error
}
```
