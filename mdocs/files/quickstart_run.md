The project is ready to be executed. The env var called `ENVIRONMENT` will be needed in order to read the local configuration.
Now your project is ready to run with the following command:

```shell
~/projects/booking-fly> ENVIRONMENT=dev go run main.go
```

!!! tip "Booking Fly App"
    The Booking Fly application example is available to clone from our GitHub repository: [Booking Fly](https://github.com/darksubmarine/booking-fly).


## How to call application endpoints?

Add a user to your app:
```shell
curl --location 'http://localhost:8081/api/v1/users' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Jhon Doe",
    "email": "jhon.doe@example.com",
    "password": "super-secret",
    "plan": "BRONZE",
    "miles": 1500
}'
```

The response should be like:
```json
{
    "id": "01J9MB1EQMHJ9XCEXJBMBDH47Y",
    "created": 1728333265652,
    "updated": 1728333265652,
    "name": "Jhon Doe",
    "email": "jhon.doe@example.com",
    "password": "super-secret",
    "plan": "BRONZE",
    "miles": 1500
}
```
Fetch the given `user id` and try to perform a fly reservation calling:
```shell
curl --location 'http://localhost:8081/api/v1/booking' \
--header 'Content-Type: application/json' \
--data '{
    "departure": "CFO",
    "arrival": "JFK",
    "miles": 2700,
    "from": 1726837954000,
    "to": 1727529154000,
    "userId": "01J9MB1EQMHJ9XCEXJBMBDH47Y"
}'
```
And the response looks like:
```json
{
    "id": "01J9MB1S7DHCNHMH0NGGNQNN77",
    "created": 1728333276397,
    "updated": 1728333276397,
    "departure": "CFO",
    "arrival": "JFK",
    "miles": 2700,
    "from": 1726837954000,
    "to": 1727529154000,
    "userId": "01J9MB1EQMHJ9XCEXJBMBDH47Y"
}
```

Once that a reservation is success the user state is upgraded following the use case rules. Fetch the user to see the new status:
```shell
curl --location 'http://localhost:8081/api/v1/users/01J9MB1EQMHJ9XCEXJBMBDH47Y'
```

So, the user status has been upgraded! ... the new plan is `SILVER` and the accumulated miles are 4500.
```json
{
    "id": "01J9MBJYFKTSNS6BNJKWZ6ZEKD",
    "created": 1728333838835,
    "updated": 1728333842563,
    "name": "Jhon Doe",
    "email": "jhon.doe@example.com",
    "password": "super-secret",
    "plan": "SILVER",
    "miles": 4500
}
```

## What's next?

After running the example app successfully we encourage you to continue reading the next topics:

 - [Adding a repository like MongoDB or SQL](basic_entity_output_adapters.html)
 - [Querying entity data using TQL](tql.html)
 - [Adding custom fields to the entity](advanced_entity_naming_fields.html)
 - [How to document the REST API](advanced_rest_api_oas.html)
 - [Adding context data to your HTTP requests](advanced_rest_api_context.html)
 - [How Dependency injection works?](advanced_di.html)
 