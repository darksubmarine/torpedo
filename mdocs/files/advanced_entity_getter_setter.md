The entity object is the core data model and all other objects like DTOs, DMOs and QROs turns around it.

!!! info "Entity data model object"
    This Entity objects ara defined as simple objects with fields, getters and setters.

Basically the field names follows a simple convention...

Each **entity** field:

- Is not exportable and must be lower case.
- The entity must provide a getter and setter following:
    - **Getter**: must be the same as field name but exportable, so the capitalized field name.
    - **Setter**: must start with the keyword `Set` followed by the capitalized field name (camel case).

For instance:

```go
// UserEntity a system user
type UserEntity struct {
    *entityBase

	// the user system role 
    role string
}

// Role get the user role
func(u *UserEntity) Role() string {
	return u.role
}

// SetRole sets the user role
func(u *UserEntity) SetRole(r string) {
    u.role = r
}
```

## Modifying Getter and Setter method names

If for some reason the getter and setter methods cannot follow the convention, it is possible modify its names. 
Torpedo introduces 2 tags to bind the setter and getter method names to the specific field:

 - `torpedo.setter`
 - `torpedo.getter`


Following the previous sample:

```go
// UserEntity a system user
type UserEntity struct {
    *entityBase

	// the user system role 
    role string `torpedo.getter="GetSystemRole" torpedo.setter="SetSystemRole"`
}

// GetSystemRole get the user role
func(u *UserEntity) GetSystemRole() string {
	return u.role
}

// SetSystemRole sets the user role
func(u *UserEntity) SetSystemRole(r string) {
    u.role = r
}
```