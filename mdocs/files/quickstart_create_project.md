
Lets begging with our first application built with Torpedo!. We will be building a fly reservation app called `Booking Fly` following these simple steps:

 1. [Initialize the project](quickstart_create_project.html#torpedo-init)
     - [Entity definition](quickstart_create_project.html#entity-definition)
     - [Use cases definition](quickstart_create_project.html#use-case-definition)
     - [App definition](quickstart_create_project.html#application-definition)
 2. [Code generation](quickstart_create_project.html#torpedo-fire)
 3. [Dependency Injection](quickstart_dependency.html)
 4. [Writing use case: Booking a fly](quickstart_use_cases.html#lets-start-writing-the-use-case)
 5. [Run it!](quickstart_run.html)


We encourage you to follow all the steps to create it, however if you need to check something here is a link to the application repo: [Booking Fly](https://github.com/darksubmarine/booking-fly). 

## Torpedo Init

Creating your first application based on an entity definition requires a torpedo initialization. Following the next command
your project will be initialized.

First you need to create a project dir and be located on it:

```shell
mkdir booking-fly && cd booking-fly
```

After that you would initialize the torpedo project:

```shell
~/projects/booking-fly> torpedo init
```

### `.torpedo` dir struct

The `.torpedo` dir struct is so simple. You need a directory to put your entities definitions and a directory to put the
use cases definitions. The next is an example of it:

```text
booking-fly/
  |_ .torpedo
    |
    |_ entities           Entities yaml files goes here
    | |_ docs             Entities MD documentation
    | | |_ user.md
    | |_ user.yaml
    |
    |_ use_cases           Use cases yaml files goes here
    | |_ docs              Use cases MD documentation
    | | |_ booking_fly.md
    | |_ booking_fly.yaml
    |
    |_ app.yaml        The application yaml definition
```

### Entity definition
The following snippets illustrates how to create the entities definition for the **Booking Fly** example:
For further information refers to the section [entity schema](schema_entity.html)

??? abstract ".torpedo/entities/user.yaml"
    ```yaml
    version: torpedo.darksub.io/v1.0
    kind: entity
    spec:
        name: "user"
        plural: "users" #(1)!
        description: "The frequent flyer user"
        doc: |
            The user entity represents a system user but also a frequent flyer. 
            This entity is only for the example purpose.
        schema:
            reserved:
                id:
                    type: ulid #(2)!
            
            fields:
              - name: name
                type: string
                description: "The user full name"
        
              - name: email
                type: string
                description: "The user contact email"
    
              - name: password #(3) it is not recommended to save passwords, this is an example only
                type: string
                encrypted: true
                description: "The user system password"
        
              - name: plan
                type: string
                description: "The user membership plan"
                validate:
                  list:
                    values:
                      - GOLD
                      - SILVER
                      - BRONZE
        
              - name: miles
                type: integer
                description: "The accumulated flyer miles"
        
        relationships:
            - name: trips
              type: $rel
              ref: ".torpedo/entities/trip.yaml"
              cardinality: hasMany
              load:
                type: nested
                metadata:
                    maxItems: 100
    
        adapters:
            input:
                - type: http
        
            output:
              - type: memory #(4)!
    ```

    1. The __plural__ attribute is used for instance to create the entity RESTful API path
    2. The __id__ field is reserved and can be only one of: **ulid** or **uuid**.
    3.  The encrypted field attribute is used to store the field value into the repository as encrypted instead of plain value.
        This example illustrates how it works, but a password shouldn't be stored in the repo.
    4. The **memory** output adapter is usually for testing purpose.


??? abstract ".torpedo/entities/trip.yaml"
    ```yaml
    version: torpedo.darksub.io/v1.0
    kind: entity
    spec:
        name: trip
        plural: trips
        description: "The user fly trip reservations"
        doc: |
            The trip entity handles all data related with the frequent flyer trip
        schema:
            reserved:
                id:
                    type: ulid
    
            fields:
              - name: departure
                type: string
                description: "The trip departure airport"
        
              - name: arrival
                type: string
                description: "The trip arrival airport"
        
              - name: miles
                type: integer
                description: "The trip miles"
        
              - name: from
                type: date
                description: "The trip from date"
        
              - name: to
                type: date
                description: "The trip to date"
        
        adapters:
            input:
                - type: http
        
            output:
                - type: memory

    ```

### Use case definition

The use cases are the corner stone of any application. Your custom business logic should be defined as use cases. Torpedo introduces
a simple but powerful use case definition. The following example is the fly reservation use case:

??? abstract ".torpedo/use_cases/booking_fly.yaml"
    ```yaml
    version: torpedo.darksub.io/v1.0
    kind: useCase
    spec:
        name: "BookingFly"
        description: "Fly reservation use case"
        doc: | #(1)!
            Given a frequent flyer user should be able to do a booking fly from our well known fly routes, selecting the
            departure airport and the arrival airport, also setting up the from-to fly dates. If the booking is successful, so the
            system should calculate the user awards and upgrade it
        domain:
            entities:
                - user.yaml
                - trip.yaml

    ```

    1. The documentation can be placed in a markdown file into the folder `.torpedo/use_cases/docs/booking_fly.md` and you can refer it from the
        yaml file as `doc: booking_fly.md`

!!! tip "Documentation as Markdown files"
    The documentation can be placed in a markdown file into the folder `.torpedo/use_cases/docs/booking_fly.md` and you can link it from the
    yaml file as `doc: booking_fly.md`

### Application definition

The application definition is the most important file because here is described your app. 
The following example shows how to define the **Booking Fly** app:

??? abstract ".torpedo/app.yaml"
    ```yaml hl_lines="8"
    version: torpedo.darksub.io/v1.0
    kind: app
    spec:
      name: "Booking Fly System"
      description: "Application example"
      stack:
        lang: go
        package: "github.com/darksubmarine/booking-fly" #(1)!
      domain:
        entities:
          - user.yaml
          - trip.yaml
        useCases:
          - booking_fly.yaml
    
    ```

    1. Replace it with your package (git) name

## Torpedo Fire

Once that you have the project dir, the entity definitions, the use cases definitions and your app definition, running 
the Torpedo command line tool the code will be generated:

```shell
~/projects/booking-fly> torpedo fire
```
The ouput should be like:
```shell
mkdir: ~/projects/booking-fly/dependency
mkdir: ~/projects/booking-fly/domain
mkdir: ~/projects/booking-fly/domain/entities
mkdir: ~/projects/booking-fly/domain/use_cases
mkdir: ~/projects/booking-fly/domain/testing
mkdir: ~/projects/booking-fly/domain/inputs
mkdir: ~/projects/booking-fly/domain/outputs
Entities parsed:
	- ~/projects/booking-fly/.torpedo/entities/user.yaml
	- ~/projects/booking-fly/.torpedo/entities/trip.yaml
Use cases parsed:
	- ~/projects/booking-fly/.torpedo/use_cases/booking_fly.yaml

  + Build has been successfully!
```


This command will generate the project struct and the autogenerated code. So something like the following project should be created:

```text
booking-fly/
  |_ .torpedo
  |_ dependency
  |_ domain
     |_ entities
        |_ ...
     |_ inputs
     |_ outputs
     |_ testing
     |_ use_cases
         |_ ...
     |_ ...   
  |_ config-dev.yaml
  |_ go.mod
  |_ go.sum
  |_ main.go
    
```

Now your project is ready to run with the following command: 

```shell
~/projects/booking-fly> ENVIRONMENT=dev go run main.go
```

!!! warning "Missing API endpoints"
    You should notice about that **Use Case API endpoints are not available so far**. To get it working as expected we need to proceed 
    with the Dependency Injection section.

