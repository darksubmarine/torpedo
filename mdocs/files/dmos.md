
The DMOs are the well known **Data Mapper Object** and into the Torpedo ecosystem DMOs are the chosen objects
to use as **output**.

Each defined output adapter will have their own DMO instance. Sometimes could be the same one or sometimes not.

A DMO instance has the responsibility to map the entity data to the specific adapter. For instance, an adapter of type
`mongodb` will have a DMO with the capability to encode (map) the entity data to `bson` format.

Let's illustrate this with the Author's Blog Entity sample.

Having the previous defined author entity schema as:


```yaml
version: torpedo.darksub.io/v1.0
kind: entity
spec:
  name: "author"
  plural: "authors"
  description: "The blog post author"

  schema:
    reserved:
      # By default, an ID is assigned as ULID format (string).
      # the field name id is reserved, but you can configure it from this section
      id:
        type: ulid

    fields:
      - name: name
        type: string
        description: "The author full name"
        doc: "The author full name"

      - name: email
        type: string
        description: "The author contact email"
        doc: "The author contact email"

  adapters:
    output:
      - type: mongodb
        metadata:
          collection: "authors"
```

And the **output adapter** as `mongodb` the generated **DMO** will be an object with the capability to be stored in MongoDB
implementing the generated interface.

=== "Go"
    ```go
    // IEntityDMO interface to defines the Data Mapper Object implementation
    type IEntityDMO interface {
        ToEntity() (*AuthorEntity, error)
        HydrateFromEntity(entity *AuthorEntity) error
        Id() string
        Created() int64
        Updated() int64
        Name() string
        Email() string
    }
    ```

Following the golang implementation, the generated DMO will be:

=== "Go"
    ```go
    // EntityDMOMongoDB Data Mapper Object (DMO) to store entity into MongoDB
    type EntityDMOMongoDB struct {
        *author.EntityDMO `bson:"-"`
    
        Id_      string `bson:"_id"`
        Created_ int64  `bson:"created"`
        Updated_ int64  `bson:"updated"`
    
        Name_      string `bson:"name"`
        Email_     string `bson:"email"`
    }
    ```