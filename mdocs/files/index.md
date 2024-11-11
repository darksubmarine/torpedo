# Torpedo project

## What is Torpedo?
<span class="dsMainColor">_Torpedo_</span> is a clean, decoupled, testable, and flexible framework designed to streamline development. It features a built-in code generator that adheres to hexagonal architecture principles, promoting separation of concerns and enhancing the maintainability of your code. By fostering a clear structure and supporting best practices, Torpedo enables developers to create scalable and robust applications with ease.

## What is NOT Torpedo?
<span class="dsMainColor">_Torpedo_</span> **isn't** a:
 
 - Database management tool.
 - Data migration tool
 - Mocking API service

## Why should I use Torpedo?
Torpedo offers a streamlined approach to coding by eliminating redundant tasks and reducing boilerplate code. By using Torpedo, you can significantly boost your productivity, allowing you to dedicate more time and energy to developing the core functionality of your application.

## What is Hexagonal Architecture?

Hexagonal architecture is a software design principle that encourages decoupling of software components in order to increase flexibility and maintainability.

There are many benefits to using hexagonal architecture in software development. This architecture makes it easier to add new features or make changes to existing ones, since components can be changed or added without affecting the rest of the system. Hexagonal architecture also makes it easier to test individual components in isolation, which can save time and resources during development.


#### Benefits of Hexagonal Architecture

There are many benefits to using hexagonal architecture in your software development projects. Hexagonal architecture is a great way to decouple your code, making it more testable and flexible. It also allows you to easily plug in different components or services without affecting the rest of your code. Hexagonal architecture is also a great way to design for scalability. By keeping your components loosely coupled, you can easily add new ones as your project grows. This makes it much easier to scale up your project without having to re-architect the entire thing. Overall, hexagonal architecture is a great way to structurally improve your code. It makes it more testable, scalable, and flexible. If you’re looking for ways to improve your software development process, hexagonal architecture is a great place to start.

#### Advantages

 - **Decoupling:** The core application logic is decoupled from the external systems, making the system easier to maintain and evolve.
Changes in external systems have minimal impact on the core logic.

 - **Testability:** The architecture facilitates easier unit and integration testing.
Core logic can be tested without needing the actual external systems, using mocks or stubs.

 - **Flexibility:** Allows easy replacement or modification of external systems (adapters) without changing the core logic.
Supports multiple types of user interfaces or input/output mechanisms.

 - **Scalability:** The architecture can scale better as different parts of the system can be developed, maintained, and scaled independently.
Supports the addition of new features with minimal disruption to existing functionality.

 - **Maintainability:** Clear separation of concerns makes the codebase easier to understand and maintain.
Simplifies troubleshooting and debugging by isolating issues to specific layers.

 - **Adaptability:** Facilitates integration with new technologies or systems.
Helps in porting the application to different platforms by merely adding new adapters.

## Why should I use Torpedo?

With
<span class="dsMainColor">_Torpedo_</span>, you will be able to reduce the boilerplate code winning time to focus on your application use cases.
Some features that you have covered by Torpedo are:

- Entity definition in `yaml` based schema
- Decouple, testable and flexible code generated
- Input adapters:
    - `HTTP`
    - `gRPC` _(coming soon)_
    - `GraphQL` _(coming soon)_
- Output adapters
    - `SQL`
    - `MongoDB`
    - `Memory`
    - `Redis`
    - `Redis+MongoDB` - _(Redis used as first cache level)_
    - `Redis+SQL` - _(Redis used as first cache level)_
- [CRUD operations](https://en.wikipedia.org/wiki/Create,_read,_update_and_delete) out of the box
- [TQL _(Torpedo Query Language)_](tql.html) data store agnostic query language
- Entity Documentation
- Easy stack migration
