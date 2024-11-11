# TODO

- [x] Defines data types (string, int, float, boolean, date) 
  - [x] String (string)
  - [x] Integer (int64)
  - [x] Float (float64)
  - [x] Boolean (boolean)
  - [x] List (slice) done as custom field, not supported in yaml spec
- [ ] BATCH operations
- [x] Review and finish the UPDATE logic
 - [x] HTTP Error response struct
 - [x] HTTP Partial DTO check method to send bad request
 - [x] Error codes
 - [x] Gin middlewares support
   - [x] Entity Create (POST)
   - [x] Entity Read (GET)
   - [x] Entity Update (PUT)
   - [x] Entity Delete (DEL)
 - [x] TQL at `core entity` and at endpoint `/<entity_plural_name>/query`
   - [x] TQL support
     - [x] MongoDB
       - [x] Filter (find)
       - [x] Sort
       - [x] Projection
       - [x] Pagination
     - [x] SQL
   - [x] Error handler with error codes 
     - [x] 400x if query has errors
     - [x] 500x if search errored
 - [ ] Unique fields support
 - [ ] Input support
   - [x] HTTP Gin
   - [ ] GraphQL
   - [ ] gRPC
 - [x] Output support
   - [x] Memory _(for unit tests)_
   - [x] MongoDB
   - [x] Redis
   - [x] Composite Redis _(as cache)_ + MongoDB
   - [x] SQL output support
   - [x] Composite Redis _(as cache)_ + SQL
 - [ ] Entity relationship
   - [x] HasMany
   - [x] HasOne
   - [x] Nested loading
   - [ ] Eager loading
   - [ ] Lazy loading
 - [x] Encryption at storage
 - [ ] Entity Unit Tests
 - [x] Logs
 - [ ] Metrics

## Error codes:

**400 bad request**

 - 4001 - Error binding JSON
 - 4002 - Partial entity incomplete some field is missing
 - 4003 - Entity building from DTO

**500 internal error**

 - 5001 - Entity creation error
 - 5002 - Entity update error
 - 5003 - Entity read error
 - 5004 - Entity remove error
 - 5005 - Entity Query (TQL) error

## Concepts

 - **Data Mapper Object (DMO)**: Object to map Entity data to output adapters like storages.
 - **Data Transfer Object (DTO)**: Object to transfer Entity data from input adapters like API REST.
 - **Query Result Object (QRO)**: Object to return query results mapping the entity into this object. It is similar to a DTO.