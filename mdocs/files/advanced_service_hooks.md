## Service hooks

Each entity service implements a hooks mechanism that allows developers to execute actions at specific points in 
the CRUD operations workflow. 

The service is provided with a method named `initHooks` where the hook builder object should be set in the reserved variable named `hookBuilder` .

Initially the service is created with no operational hooks, but you can implement it or even better creates a custom hook object to handle your own logic. 
```go
// Service defines your use cases. Extends from ServiceBase to get the CRUD operations
type Service struct {
    *ServiceBase // DO NOT REMOVE this line. ServiceBase implements IServiceBase interface
}

func (s *Service) initHooks() {
    s.hookBuilder = newServiceHooksBuilder(
        // builder hooks function for CREATE operation
        func() iServiceCreateHooks {
            return newNoopServiceHooks()
        },
        
        // builder hooks function for READ operation
        func() iServiceReadHooks {
            return newNoopServiceHooks()
        },
        
        // builder hooks function for UPDATE operation
        func() iServiceUpdateHooks {
            return newNoopServiceHooks()
        },
        
        // builder hooks function for DELETE operation
        func() iServiceDeleteHooks {
            return newNoopServiceHooks()
        },
    )
}
```

### Why a builder is needed?

The hook scope should be within each CRUD action at service layer. Means that the builder object will be called per each CRUD operation, 
creating an object that implement one of the interfaces: `iServiceCreateHooks`, `iServiceReadHooks`, `iServiceUpdateHooks` and `iServiceDeleteHooks`.

The created object will live within the CRUD operation lives. So on this way a hook has the features:

 - **Atomic:** The hook object is created, executed and destroyed within the CRUD operation execution.
 - **Non-Shared:** The object instance is not shared between CRUD operations doing it secure on asynchronous requests. 

## Create hook functions
The create hook functions are executed before and after entity creation in storage layer that means before and after call the repo save method. 

This hook object must implement the interface.

```go
type iServiceCreateHooks interface {
	beforeCreate(ctx context.IDataMap, entity *Entity) error
	afterCreate(ctx context.IDataMap, entity *Entity) error
}
```

### Before Create

The before hook will be called before execute the `repo.Save(entity)` method. This hooks must implement the function like:

```go
func beforeCreate(ctx context.IDataMap, entity *Entity) error {
	// do your before create actions.
}
```

The function input parameters are a context.IDataMap interface useful to share context data from the HTTP request. 
[For further information see REST API context](advanced_rest_api_context.html) 

### After Create

The after create hook is called immediately after saving the entity into the repo if it was successfully. And implements the same
function that the before create hook.

```go
func afterCreate(ctx context.IDataMap, entity *Entity) error {
	// do your after create actions.
}
```

## Read hook functions

The read hook is executed before and after calls the repo mehtod `FetchByID(id string) (*Entity, error)` and must implement the interface:

```go
type iServiceReadHooks interface {
	beforeRead(ctx context.IDataMap, id string) error
	afterRead(ctx context.IDataMap, entity *Entity) error
}
```

### Before Read

Before read is executed before calls `repo.FetchById(id string)` method and the entiy is not available because this one has not been fetched so far.

The function to implment should be:
```go
func(ctx context.IDataMap, id string) error { return nil }
```

### After Read

Once that the entity has been fetched from the repo successfully, the `afterRead` hook will be executed. The hook function to implement is:

```go
func(ctx context.IDataMap, entity *Entity) error { return nil }
```

!!! warning "Entity scope"
    Note that the entity parameter is a pointer to; so, if some hcanges are applied to the entity, those will be bubble up till input adapter.

## Update hook functions

The before/after update hooks are pretty similar to the creation hooks and must implement the interface:

```go
type iServiceUpdateHooks interface {
	beforeUpdate(ctx context.IDataMap, entity *Entity) error
	afterUpdate(ctx context.IDataMap, entity *Entity) error
}
```

### Before Update

The "before hook" will be called before execute the `repo.Update(entity)` method. This hooks must implement the function like:

```go
func beforeUpdate(ctx context.IDataMap, entity *Entity) error {
	// do your before create actions.
}
```

### After Update

The after update hook is called immediately after updating the entity into the repo if it was successfully. And implements the same
function that the before update hook.

```go
func afterCreate(ctx context.IDataMap, entity *Entity) error {
	// do your after create actions.
}
```

## Delete hook functions

Delete hooks are called when the `service.Delete(entityId)` method is called and must implement the interface:

```go
type iServiceDeleteHooks interface {
	beforeDelete(ctx context.IDataMap, id string) error
	afterDelete(ctx context.IDataMap, id string) error
}
```

### Before Delete

The "before hook" will be called previous execute the `repo.DeleteById(entityId)` method. This hooks must implement the function like:

```go
func beforeDelete(ctx context.IDataMap, id string) error { return nil }
```

### After Delete

The "after delete hook" is called immediately after deleting the entity into the repo if it was successfully. And implements the same
function that the before delete hook.

```go
func afterDelete(ctx context.IDataMap, id string) error { return nil }
```
