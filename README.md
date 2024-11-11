# Torpedo

> Before moving on, please consider giving us a GitHub star ⭐️. Thank you!


## What is Torpedo?
<span class="dsMainColor">_Torpedo_</span> is a clean, decoupled, testable, and flexible framework designed to streamline development. It features a built-in code generator that adheres to hexagonal architecture principles, promoting separation of concerns and enhancing the maintainability of your code. By fostering a clear structure and supporting best practices, Torpedo enables developers to create scalable and robust applications with ease.

## What is NOT Torpedo?
<span class="dsMainColor">_Torpedo_</span> **isn't** a:

- Database management tool.
- Data migration tool
- Mocking API service

## Why should I use Torpedo?
Torpedo offers a streamlined approach to coding by eliminating redundant tasks and reducing boilerplate code. By using Torpedo, you can significantly boost your productivity, allowing you to dedicate more time and energy to developing the core functionality of your application.

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


## What's next?

Please read the [Torpedo documentation](https://darksubmarine.com/docs/torpedo/index.html)

- Architecture
  - [Overview](https://darksubmarine.com/docs/torpedo/architecture.html)
  - [Application container](https://darksubmarine.com/docs/torpedo/arch_application_container.html)
- Quick start
  - [Install](https://darksubmarine.com/docs/torpedo/quickstart_install.html) 
  - [Create a project](https://darksubmarine.com/docs/torpedo/quickstart_create_project.html)
  - [Booking Fly (Example repo)](https://github.com/darksubmarine/booking-fly)

 