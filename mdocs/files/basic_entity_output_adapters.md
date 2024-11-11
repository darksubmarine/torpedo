
In hexagonal architecture, output adapters act as the interfaces through which the application interacts with external systems for persisting and retrieving data. These adapters implement the repository pattern, providing a way to store and access data in a technology-agnostic manner. They translate the core application’s domain objects into data models suitable for the external systems (like databases, caches, file systems) and vice versa.

## Role of Output Adapters as Repositories

1. **Interfacing with External Systems**:
    - Output adapters manage interactions with external systems like databases (SQL/NoSQL), caches, or other storage mechanisms.

2. **Data Translation**:
    - They translate domain objects into a format suitable for storage in the external system and translate stored data back into domain objects for the application.

3. **Decoupling**:
    - Output adapters decouple the core application logic from the specifics of data storage, promoting flexibility and maintainability.

4. **Implementation of Repository Interfaces**:
    - They implement repository interfaces defined in the core application, ensuring that the core logic remains technology-agnostic.

## Benefits of Output Adapters as Repositories

- **Decoupling**: They decouple the core application logic from data storage specifics, promoting flexibility and ease of maintenance.
- **Separation of Concerns**: The core application logic remains focused on business rules, while output adapters handle data persistence and retrieval.
- **Testability**: The core application can be tested independently of the storage mechanism, enhancing test coverage and reliability.
- **Flexibility**: Different storage mechanisms (SQL, NoSQL, Cache) can be used and switched without altering the core application logic.

Output adapters in hexagonal architecture ensure that the application remains adaptable, maintainable, and scalable. By implementing the repository pattern, they provide a clean interface for data storage and retrieval, enabling the core application to remain agnostic to the specifics of the underlying storage technology.

## Torpedo Repository Implementation

All repositories implemented by Torpedo follows the same pattern:

- Port interface:
    - A base interface to define autogenerated methods (CRUD)
    - A main interface that aggregates the base one to define user repository methods.
- Repository classes:
    - A base repository that implements the base port interface.
    - A main repository that aggregates the base repo and is a safe place for users to code their own methods.
- Data Mapper Objects:
    - A base DMO object that implements the mapping from the entity schema
    - A user DMO Object that can handle custom fields to be stored along side of entity fields.

For instance a redis repository implementation looks like:

``` mermaid
classDiagram
  class BaseDMO
  class DMO
  class RedisRepository
  class redisRepositoryBase
  
  interface IRepository
  interface IRepositoryBase
  
 
      IRepositoryBase <|-- IRepository
      DMO <|-- BaseDMO
      
      IRepositoryBase <|.. redisRepositoryBase 
      IRepository <|.. RedisRepository
     
      redisRepositoryBase <|-- RedisRepository
      redisRepositoryBase : redis.client data 
      redisRepositoryBase : []byte cryptoKey 
    
      redisRepositoryBase: +Save(entity) error
      redisRepositoryBase: +FetchByID(id) (Entity, error)
      redisRepositoryBase: +Update(entity) error
      redisRepositoryBase: +DeleteByID(id) error
 
```

### Data Mapper Object (DMO)

DMOs are the selected objects to map entity data into a secondary adapter or storage adapter.
This objects are responsible to encrypt or decrypt field values at save/update or fetch operations.
Each time that a repository is created the encryption key must be provided.

!!! warning "AES key"
    The key argument should be the AES key, either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.