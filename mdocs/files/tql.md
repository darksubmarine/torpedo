Torpedo Query Language (TQL)

TQL is a data store agnostic query language which is available out of the box to query any entity. Each time that an entity is created a `/query` endpoint
is added to this one letting users fetch entities applying filters, pagination and more!. 

For further information about entity endpoints please refer to [Entity API endpoints](basic_entity_input_restapi.html#api-endpoints)

!!! danger "Disclaimer"
    String sensitive match or not will deppends on each storage engine.


TQL defines a JSON based query language with next capabilities:

- Querying
- Sorting
- Projection
- Pagination

## Query format

The query object contains 3 sections: `filter`, `projection` and `pagination`. 
The query skeleton is described as a JSON object like:

```json
{
  "filter": {},
  "projection": [],
  "pagination": {}
}
```

### Filter section

This section is used to apply a filter over the records that you would like to fetch. 
The filter is composed by two fields:

Type (`type`) can be set as: 

 - `all` (and) all field query in the filter must match with the records  
 - `any` (or) any field query in the filter must match with the records

And fields (`fields`) a list of fields query to be applied, each field object in the query contains 3 keys:
 
 - `field`: the field name to query in the condition
 - `operator`: the operator condition
 - `value`: the value to be queryed

An example of that could be:
```json
{
  "field": "created",
  "operator": ">=",
  "value": 1666875856369
}
```
or for a list operator could be:
```json
{
    "field": "plan",
    "operator": "[?]",
    "value": [
      "silver",
      "gold",
      "platinum"
    ]
}
```

#### Filter operators
The next table describes the available filter operators per field

| Operator | Description                  | Type                     |
|-------|------------------------------|--------------------------|
| ==    | Equal                        | int, float, string, bool |
| !=    | Not Equal                    | int, float, string, bool |
| >     | Greater than                 | int, float               |
| >=    | Greater than or Equal        | int, float               |
| <     | Less than                    | int, float               |
| <=    | Less than or Equal           | int, float               |
| \[n\] | Between includes limits      | int, float               |
| (n)   | Between excludes limits      | int, float               |
| [n)   | Between includes left limit  | int, float               |
| (n]   | Between includes right limit | int, float               |
| \[?\] | In list                      | int, float, string       |
| .s.   | Contains                     | string                   |
| s..   | Prefix                       | string                   |
| ..s   | Suffix                       | string                   |

!!! info "Boundary Operators behaviour"
    The boundary operators follows the next behaviour
    ```text
    (𝑎,𝑏)⇒{𝑥∈ℝ:𝑎<𝑥<𝑏}
     
    [𝑎,𝑏]⇒{𝑥∈ℝ:𝑎≤𝑥≤𝑏}
    ```

### Projection section

The projection is a list of field names to return as part of the expected result. This is usefull when you don't
need the full object fields. An example of this one could be:

```json
"projection": ["name", "id","updated"]
```

### Pagination section

This section helps developers to run paginated queries over your datastore. Two alternatives are available: `cursor` and `offset`.

#### Offset pagination

The offset pagination uses the well known storage `offset` and `limit` to run the query and let you sort the results by the specified fields. 
The query should be configured with the number of items to get and the page number to retrieve.   

```json
"pagination": {
  "items":30,
  "offset":{
    "page":10,
    "sort": [{"field": "name", "type": "desc"}]
  }
}
```


#### Cursor pagination

The cursor pagination is useful to query page by page calling with the given `nextToken`. Also a sorting by only one field is supported.

!!! warning "Entity ID"
    Due to the cursor pagination works querying records via the entity ID, this one must be `lexicographically (timed) sortable` because
    of this the **UUID** cannot be used for this kind of pagination, instead use the provided **ULID** 

```json
"pagination":{
  "items":30,
  "cursor":{
    "meta": {
      "type":"asc",
      "sort": {"field": "name", "type": "desc"}
    },
  "nextToken":"YjQ5ZjAzZTktZTBlMS00ZjRkLWI2NDAtNGRlMjcxMTI0ZjA2"
}
```



### Query example
```json
{
  "filter": {
        "type": "all",
        "fields": [
            {
                "field": "created",
                "operator": ">=",
                "value": 1666875856369
            },
            {
                "field": "age",
                "operator": ">=",
                "value": 18
            },
            {
                "field": "plan",
                "operator": "[?]",
                "value": [
                  "silver",
                  "gold",
                  "platinum"
                ]
            },
            {
                "field": "date",
                "operator": "[?]",
                "value": [
                  1665757374948,
                  1664752374987
                ]
            }
        ]
    },
  
  "projection": ["name", "id","updated"],
 
  "pagination":{
        "items":30,
        "offset":{
          "page":10,
          "sort": [{"field": "name", "type": "desc"}]
        }
      }
    
}
```
