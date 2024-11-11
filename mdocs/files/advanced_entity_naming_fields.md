Basically the field names follows a simple convention:

## Entity

At **entity** level the field...: 

 - Is not exportable and must be lower case.
 - The entity must provide a getter and setter following:
    - Getter: must be the same as field name but exportable, so the capitalized field name.
    - Setter: must start with the keyword `Set` followed by the capitalized field name (camel case). 

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

!!! tip "Advanced getter and Setter"
    Getter and Setter methods can follow your own naming rule. This is not recommended but you can do it. 
    Please read the next section [Entity Getter and Setter](advanced_entity_getter_setter.html) 

## Data Transfer Object - DTO

The DTOs are different to the Entity because the field sometimes must be exported, for instance, to JSON serialize.

In this case the DTO field names must follow this conventions:

- The field is exported, so starts with capital letter.
- In order to avoid overlapping with getter method, the field name must ends with a underscore `_`
- If the DTO is mapped as JSON, the json tag must be added.

For instance:

```go
package gin

type CustomDTO struct {
	Role_ *string `json:"role"`
} //@name user.CustomDTO

// Role getter method
func (dto *CustomDTO) Role() *string { 
	return dto.Role_ 
}
```


## Data Mapper Object - DMO

This kind of object is pretty similar to DTO and **follows the same convention**.

For instance:

!!! abstract "MongoDB DMO"
      ```go
      type DMO struct {
          *user.EntityDMO `bson:"-"`
      
          Role_ *string `bson:"role"`
      }

      // Role getter method
      func (dmo *DMO) Role() *string {
         return dmo.Role_
      }
      ```

!!! abstract "SQL DMO"
      ```go
      type DMO struct {
         *user.EntityDMO

          Role_ *string `db:"role"`
      }

      // Role getter method
      func (dmo *DMO) Role() *string {
         return dmo.Role_
      }
      ```

## Query Result Object - QRO

QRO is similar to DTO and **follows the same convention**. This one only exports to JSON the query results, so no getters and setters are needed

For instance:

```go
type QRO struct {
    Role_ *string `json:"role,omitempty"`
}
```

## Setting custom names at DTO, DMO and QRO

Torpedo supports custom field names in DTOs, DMOs and QROs objects. Each custom field name MUST be mapped to the respective entity field.

### Mapping custom DTO field

The custom DTO's field can be mapped with its respective entity field via Go tag. The `entity` is the core data model and all other objects turns around this one.

Torpedo support the tag `torpedo.dto` in order to map this field and follow the pattern:

```text
torpedo.dto:"<adapter>=<fieldName>"
```

where:

 - `<adapter>` is the input adapter. For instance `http`
 - `<fieldName>` is the custom DTO field name to map. 


For instance: 

**defining a DTO:**
```go hl_lines="4"
package gin

type CustomDTO struct {
	UserRole *string `json:"role"`
} //@name user.CustomDTO
```

The entity looks like:

```go hl_lines="6"
// UserEntity a system user
type UserEntity struct {
    *entityBase

	// the user system role 
    role string `torpedo.dto:"http=UserRole"`
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


### Mapping custom DMO field

DMO field mapping is pretty similar to DTOs mapping, but the tag name is different. 

Torpedo supports the tag `torpedo.dmo` in order to map this field and follow the pattern:

```text
torpedo.dmo:"<adapter>=<fieldName>"
```

where:

- `<adapter>` is the output adapter. For instance `memory`
- `<fieldName>` is the custom DMO field name to map.


For instance:

**defining the sql and mongoDB DMO:**
```go hl_lines="4"
type DMO struct {
    *user.EntityDMO

    SqlUserRole *string `db:"role"`
}
```

```go hl_lines="4"
type DMO struct {
    *user.EntityDMO

    UserRole *string `bson:"role"`
}
```

The entity looks like:

```go hl_lines="6"
// UserEntity a system user
type UserEntity struct {
    *entityBase

	// the user system role 
    role string `torpedo.dmo:"sql=SqlUserRole,mongodb=UserRole"`
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




### Mapping custom QRO field

Finally, QRO objects follows the same behaviour that previous ones.

Torpedo supports the tag `torpedo.qro` in order to map the field following the pattern:

```text
torpedo.qro:"<fieldName>"
```


where:

- `<adapter>` is the output adapter. For instance `memory`
- `<fieldName>` is the custom DMO field name to map.


For instance:

**defining QRO:**
```go hl_lines="2"
type QRO struct {
    UserRole *string `json:"role"`
}
```

The entity looks like:

```go hl_lines="6"
// UserEntity a system user
type UserEntity struct {
    *entityBase

	// the user system role 
    role string `torpedo.qro:"UserRole"`
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
