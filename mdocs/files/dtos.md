
The DTOs are the well known **Data Transfer Object** and into the Torpedo ecosystem DTOs are the chosen objects
to use as input. 

Each defined input adapter will have their own DTO instance. Sometimes could be the same one or sometimes not.
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
    input:
      - type: http
```

And the **input adapter** as `http` the generated DTO will be an object with the capability to be serialized and deserialize
as JSON implementing the generated interface.

=== "Go"

    ```go
    // IEntityDTO interface to defines the Data Transfer Object implementation
    type IEntityDTO interface {
        ToEntity() (*AuthorEntity, error)
        HydrateFromEntity(entity *AuthorEntity)
    
        Id() string
        Created() int64
        Updated() int64
    
        Name() string
        Email() string
    }
    ```


Which can be serialized as:

```json
{
  "id": "01HCYTHEVRWQMCMKFNBE7KFH9A",
  "created": 1697546575,
  "updated": 1697546575,
  "name": "Arthur",
  "email": "arthur@someemail.com"
}
```


