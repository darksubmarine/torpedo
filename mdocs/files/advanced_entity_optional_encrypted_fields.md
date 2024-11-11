

The extended fields can be configured using the Go field tag: `torpedo.field`.  

!!! warning "Entity model"
    The tag must be added only at the Entity model, the file named `entity.go`

## Setting it as Optional 

The new field can be set as optional adding the tag: `torpedo.field:"optional"`. For instance:

```go
// SensorEntity Measurement sensor
type SensorEntity struct {
    *entityBase // DO NOT REMOVE IT

    myOptionalField           string  `torpedo.field:"optional"`
}
```


## Setting it as Encrypted

Also the field can be set as encrypted, so each time that a storage adapter saves the entity, the equivalent field in the DMO 
will be encrypted. Note that the field must be of type `string`

```go
// SensorEntity Measurement sensor
type SensorEntity struct {
    *entityBase // DO NOT REMOVE IT
	
    mySecureField             string  `torpedo.field:"encrypted"`
}
```

## Setting it as Optional and Encrypted, is possible ?

The answer is yes! ... the tag supports both values (comma separated). For instance:

```go
// SensorEntity Measurement sensor
type SensorEntity struct {
    *entityBase // DO NOT REMOVE IT
	
    myOptionalAndSecureField  string  `torpedo.field:"optional,encrypted"`
}
```
