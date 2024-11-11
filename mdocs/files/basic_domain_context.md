In hexagonal architecture, a domain context encapsulates all the components related to a specific aspect of the domain, providing a clear and cohesive structure for organizing business logic. By defining clear boundaries and encapsulating domain-specific logic, the domain context helps to ensure consistency, modularity, and maintainability, aligning with the principles of hexagonal architecture and promoting a clean and adaptable application design.

## Benefits of a Domain Context

 1. **Encapsulation**: Clearly defines the boundaries within which business rules are applied.
 2. **Consistency**: Ensures that business rules and invariants are consistently enforced.
 3. **Modularity**: Facilitates modularity, allowing different parts of the system to evolve independently.
 4. **Maintainability**: Improves maintainability by isolating domain-specific logic.
 5. **Scalability**: Supports scalability by enabling the independent scaling of different domain contexts.

## Domain Context implementation

Torpedo implements a `Context` class with access to each domain entity service. This adds access to all the core business logic and entities data. 

!!! tip "Context uses"
    The Context class is really useful when some entity meta-data is needed to perform some actions, for instance, in other entity use case.