In hexagonal architecture, input adapters are the components that allow external systems to interact with the core application. They handle the inbound communication and translate external requests (`DTOs`) into a format that the core application logic can process (`Entity model`). Input adapters act as the primary interface between users or external systems and the application's core use cases.

### Role of Input Adapters in Hexagonal Architecture

1. **Handling External Requests**:
    - Input adapters receive and handle requests from external sources, such as HTTP requests from web clients, messages from a message queue, or commands from a CLI.

2. **Translating Requests**:
    - They translate these external requests into calls to the application's use cases, often encapsulated in service methods.

3. **Interfacing with Ports**:
    - Input adapters interact with the input ports defined in the core application. These ports expose the use cases of the application in a technology-agnostic manner.

4. **Response Handling**:
    - After processing the request, input adapters also handle the response, translating the result from the core application back into a format suitable for the external system (e.g., HTTP response, CLI output).

### Common Types of Input Adapters

1. **Web Controllers**:
    - Handle HTTP requests and responses. Commonly used in web applications.

2. **CLI Handlers**:
    - Handle command-line interface inputs and outputs.

3. **Message Listeners**:
    - Handle messages from a message queue or a pub/sub system.


### Benefits of Input Adapters in Hexagonal Architecture

- **Decoupling**: Input adapters decouple the external interfaces from the core application logic, promoting flexibility and maintainability.
- **Separation of Concerns**: They handle the specifics of request handling and response formatting, allowing the core application to focus on business logic.
- **Testability**: The core application logic can be tested independently of the input adapters, enhancing test coverage and reliability.
- **Flexibility**: Different input adapters can be implemented for various interfaces (e.g., web, CLI, messaging) without changing the core application logic.

Input adapters in hexagonal architecture ensure that the application remains adaptable and maintainable, supporting various external interfaces while keeping the core business logic isolated and testable.
