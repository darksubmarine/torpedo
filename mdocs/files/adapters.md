
The adapters are grouped in two categories: **input** and **output**.

The **input adapters** are all those that are needed to handle the incoming data to our entity. Like:

 - REST API
 - GraphQL
 - Command Line
 - Inbound queue
 - etc.

The **output adapters** are all those that are needed to handle the outcoming data from our entity.
Into the Torpedo context these adapters are known as `Repository`, just like:

 - MongoDB
 - SQL
 - Redis
 - Memory
 - Kafka
 - etc.

!!! info "Provided adapters"

    So far Torpedo provides you with the next out of the box adapters:

    **Input**:
    
    - HTTP REST API
    - GraphQL _(coming soon)_
    
    **Output (Repository)**:
    
    - Memory
    - Redis
    - MongoDB
    - SQL
    - Redis+MongoDB
    - Redis+SQL