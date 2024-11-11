Sometimes when you are working on a multi-domain service or your domain has been thought as a library instead of a service,
the domain business logic can be exported.

This is possible thanks to the class called `domain.Service` as part of the domain package. 

> Remember, Domain Service implements a [Facade pattern](https://en.wikipedia.org/wiki/Facade_pattern),
so we need to add the method definition and the use case bind to execute.

### Binding entity (service) with the Domain Service

Torpedo adds out of the box a bind between the domain service and the entity service. So, each time that you need to call
an entity service method, like `Create` or `Read` or even your own entity method, an instance of it will be available via 
the `domain.Service` class.

For instance, following the Blog App exampel, if your entity is named `post` the entry point method will be `domain.Service.Post()`  

### Binding use case with the Domain Service

Registering a use case as part of the domain service should be easier, based on the Blog App example:

!!! quote "Remember that our case is"
    **Name:** Onboarding

    **Description** As a system a new user must be added if the email has not been registered previously.


lets adding a method named `RegisterNewUser` to our `domain.Service` in order to call the use case implementation.

At the domain service file we should add the code like below:

???+ abstract "domain/service.go"
    ```go hl_lines="13 20 23 26 31-33"
    // Package domain domain entry point
    package domain

    import (
    "github.com/darksubmarine/blog-app/domain/entities/user"
    "github.com/darksubmarine/blog-app/domain/use_cases/onboarding"
    )
    
    type IDomainService interface {
        iDomainServiceBase
    
        // RegisterNewUser defines the use case as part of the domain service
        RegisterNewUser(userModel *user.UserEntity) (string, error) //(1)!
    }
    
    type Service struct {
        *serviceBase
    
        // onboardingUC wiring an onboarding use case instance
        onboardingUC *onboarding.UseCase //(2)!
    }
    
    func NewService(ctx *Context, useCaseOnboarding *onboarding.UseCase) *Service { //(3)!
        return &Service{
	    	serviceBase:  &serviceBase{ctx: ctx},
		    ucOnboarding: useCaseOnboarding, //(4)!
    	}
    }
    
    // RegisterNewUser add new user if the email has not been registered previously
    func (s *Service) RegisterNewUser(userModel *user.UserEntity) (string, error) { //(5)!
        return s.onboardingUC.RegisterNewUser(userModel)
    }
    ```

    1. Adding the use case method at domain service definition
    2. This is a reference to the Use Case instance
    3. Adding the use case reference as parameter in the Domain Service constructor
    4. Binding the use case reference
    5. Method implementation that calls the referenced use case logic.
       <br>This is the [Facade pattern](https://en.wikipedia.org/wiki/Facade_pattern) implementation.  
