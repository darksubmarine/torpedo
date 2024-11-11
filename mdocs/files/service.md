
The service code instance is where your business logic should be placed. A service has as main dependency
a `Repository` object and knows how to handle its `Entity`.

Basically each time that Torpedo generates the code two classes are written. A `Service` class and a `ServiceBase` class.
The last one will contain the CRUD operations and the Query operation.


!!! danger "As developer"

    As a developer your own business logic **MUST BE** written into the Service class in order to avoid
    that Torpedo code generation tool overwrite your code!
    
    We strongly recommend write your uses cases into the Service class and not as part of the ServiceBase.

The diagram below illustrates how the classes are generated:

``` mermaid
classDiagram
  ServiceBase <|-- Service
  ServiceBase : IRepository repo
  
  ServiceBase: +Create(Entity entity) Entity
  ServiceBase: +Read(String id) Entity
  ServiceBase: +Update(Entity entity) Entity
  ServiceBase: +Delete(String id)
  ServiceBase: +Query(q tql.Query) Result
  
```


