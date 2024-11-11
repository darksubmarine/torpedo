As we have mentioned previously at [Entity definition](basic_entity_definition.html#type) the below data type are supported by Torpedo
as part of the code generation:

- `string`: represents a string data type.
- `integer`: represents the integer numbers and it is mapped to `int64`
- `float`: represents the float numbers and it is mapped to `float64`
- `date`: represents a timestamp number and it is mapped to `int64`
- `boolean`: represents a boolean value and it is mapped to `bool`

However sometimes you will need to add, for instance, a list of string (`[]string`) as custom data type. This one is not supported by auto-generation code,
but can be added manually intercepting the "entity points of contact" during the application life cycle. Thanks that we have based our App
on the Hexagonal Architectural pattern we have well identified those entity point of contact:

- **DTO:** Data Transfer Object are the data entry point, so, the `[]string` should be added as member of the entity DTO.
- **Entity model:** The entity schema is generated from its yaml spec, however there is an Entity class to add the custom `[]string`.
- **DMO:** Data Mapper Object are the data output point and defines how the entity field model should be stored. So, `[]string` must be added as member of the DMO object.
- **QRO:** Query Result Object are the data result object for query ([TQL](tql.html)) operations. So, `[]string` must be added as member of the QRO object.

## Lets following this with an example!

Imagine that we are writing a software to measure and track values from measurement sensors. Our sensor data schema should be defined like:

??? abstract ".torpedo/entities/sensor.yaml"
    ```yaml
    version: torpedo.darksub.io/v1.0
    kind: entity
    spec:
      name: "sensor"
      plural: "sensors"
      description: "Measurement sensor"
    
      schema:
        reserved:
          id:
            type: ulid
    
        fields:
          - name: name
            type: string
            description: "The sensor full name"
    
          - name: serial
            type: string
            description: "The sensor serial number"
    
      adapters:
        input:
          - type: http
    
        output:
          - type: memory
    
          - type: mongodb
    
          - type: redis
            metadata:
              ttl: 30000
    
          - type: redis+mongodb
    
    ```

So, guessing that one of the use cases is: **Save the last 5 sensor measurements** we can tackle this creating the [use case definition](quickstart_use_cases.html)
and coding it. But there is a faster way to do it, writing less code. 

And the answer is yes!, extending the defined sensor entity with a new custom field, like: `[]float64`. 

As we have discussed previously, an entity definition creates automatically its own CRUD actions, and we can extend those entity point of contact:

### Entity data model object - Entity

The first step is adding the custom field as part of the entity data model. 

!!! abstract "./domain/entities/sensor/entity.go"
    ```go hl_lines="8 17 20-23"
    // Package sensor Measurement sensor
    package sensor
    
    // SensorEntity Measurement sensor
    type SensorEntity struct {
        *entityBase // DO NOT REMOVE IT
    
        measures []float64 //(1)!
    }
    
    // New is a SensorEntity constructor function
    func New() *SensorEntity {
        return &SensorEntity{entityBase: newEntityBase()}
    }
    
    // Measures The sensor measurements
    func (e *SensorEntity) Measures() []float64 { return e.measures } //(2)!
    
    // SetMeasures The sensor measurements
    func (e *SensorEntity) SetMeasures(measures []float64) error { //(3)!
        e.measures = measures
        return nil
    }
    ```
    
    1. Measures slice
       2. Measures getter method
       3. Measures setter method


### Http input object - DTO

Once that we have the entity data model updated with the new field, is time to add it into the application input data flow. In order to
achieve this we need to extend the entity DTO object as like:

!!! abstract "./domain/entities/sensor/inputs/http/gin/dto.go"
    ```go hl_lines="5 9"
    // Package gin input
    package gin
    
    type CustomDTO struct {
        Measures_ []float64 `json:"measures,omitempty"` //(1)!
    } //@name sensor.CustomDTO
    
    // Measures getter method
    func (dto *CustomDTO) Measures() []float64 { return dto.Measures_ } //(2)!
    ```
    
    1. Adding measures as part of custom fields. Note the naming convention, starts with capital letter and ends with underscore `Measures_`.
    2. Getter method needed to populate entity data model from DTO.


### Storage mapper (output) object - DMO

So far we can send measures from the input data flow to the entity model, but now we need to save this info in our defined storage.
At the beginning we have defined the storage output as a composition of **MongoDB** and **Redis**, also for testing purpose we have set **Memory** adapter as well. 
So, what we need to do is updating the DMO object on each storage adapter:

!!! warning "Remember"
    Remember update the DMO object of each storage adapter.

#### Memory
!!! abstract "./domain/entities/sensor/outputs/memory/dmo.go"
    ```go hl_lines="11 15"
    // Package memory is an output adapter to store entities in memory
    package memory
    
    import (
        "bitbucket.org/darksubmarine/machine/domain/entities/sensor"
    )
    
    type DMO struct {
        *sensor.EntityDMO // Do not remove it. This will let you add custom encrypted fields and more.
    
        Measures_ []float64 `json:"measures,omitempty"` //(1)!
    }
    
    // Measures getter method
    func (dmo *DMO) Measures() []float64 { return dmo.Measures_ } //(2)!
    
    ```

    1. Adding measures as part of custom fields. Note the naming convention, starts with capital letter and ends with underscore `Measures_`.
    2. Getter method needed to populate entity data model from DMO when fetch from DB.


#### MongoDB
!!! abstract "./domain/entities/sensor/outputs/mongodb/dmo.go"
    ```go hl_lines="11 15"
    // Package mongodb is an output adapter to store entities in MongoDB
    package mongodb
    
    import (
        "bitbucket.org/darksubmarine/machine/domain/entities/sensor"
    )
    
    type DMO struct {
        *sensor.EntityDMO `bson:"-"` // Do not remove it. This will let you add custom encrypted fields and more.
    
        Measures_ []float64 `bson:"measures"` //(1)!
    }
    
    // Measures getter method
    func (dmo *DMO) Measures() []float64 { return dmo.Measures_ } //(2)!
    
    ```
    
    1. Adding measures as part of custom fields. Note the naming convention, starts with capital letter and ends with underscore `Measures_`.
    2. Getter method needed to populate entity data model from DMO when fetch from DB.


##### What happens with SQL adapter?
By design MongoDB supports arrays as data type, but with SQL engines this is a little different, some ones add it as JSON other ones 
implements vectors. So, in order to keep this feature aligned alongside all supported SQL engines [Torpedo introduces `Array` data type as JSON strings](https://github.com/darksubmarine/torpedo-lib-go/blob/v0.5.2/storage/sql_utils/data_type/types.go) 
and following the `sensor` example, the sql DMO should look like:

!!! abstract  "**./domain/entities/sensor/outputs/sql/dmo.go**"
    ```go
    package sql
    
    import (
        "bitbucket.org/darksubmarine/machine/domain/entities/sensor"
        "github.com/darksubmarine/torpedo-lib-go/storage/sql_utils/data_type"
    )
    
    type DMO struct {
        *sensor.EntityDMO // Do not remove it. This will let you add custom encrypted fields and more.
    
        Measures_ data_type.JsonArrayFloat `db:"measures"` //(1)!
    }
    ```

    1. Supported arrays: 
        - `JsonArrayFloat`
        - `JsonArrayInteger`
        - `JsonArrayString`
        - `JsonArrayDate`
        - `JsonArrayBoolean`
    

#### Redis
!!! abstract "./domain/entities/sensor/outputs/redis/dmo.go"
    ```go hl_lines="9 13"
    // Package redis implements Redis output
    package redis
    
    import "bitbucket.org/darksubmarine/machine/domain/entities/sensor"
    
    type DMO struct {
        *sensor.EntityDMO // Do not remove it. This will let you add custom encrypted fields and more.
    
        Measures_ []float64 `json:"measures,omitempty"` //(1)!
    }
    
    // Measures getter method
    func (dmo *DMO) Measures() []float64 { return dmo.Measures_ } //(2)!
    
    ```

    1. Adding measures as part of custom fields. Note the naming convention, starts with capital letter and ends with underscore `Measures_`.
    2. Getter method needed to populate entity data model from DMO when fetch from DB.



### Query result object - QRO

And last but no least, each time that we call the entity endpoint query (`[POST] /api/v1/sensors/query`) the result object must know 
how to map the measures `[]float64` slice from the extended entity data model.

!!! abstract "./domain/entities/sensor/qro.go"
    ```go hl_lines="5"
    // Package sensor Measurement sensor
    package sensor
    
    type QRO struct {
        Measures_ []float64 `json:"measures,omitempty"` //(1)!
    }
    ```
    
    1. Adding measures as part of custom fields. Note the naming convention, starts with capital letter and ends with underscore `Measures_`.

### Try it!

Running the application example, we can create a sensor with the following curl command:
```text
curl --location 'http://localhost:8081/api/v1/sensors' \
     --header 'Content-Type: application/json' \
     --data '{
         "name":"sensor-01",
         "serial": "ABC-001",
         "measures": [0.123, 0.45, 5.21]
     }'
```

The response should be something like:
```json
{
    "id": "01J5BAP6EKJ1S9049RESK9A1H8",
    "created": 1723735568339,
    "updated": 1723735568339,
    "name": "sensor-01",
    "serial": "ABC-001",
    "measures": [
        0.123,
        0.45,
        5.21
    ]
}
```

And in order to verify that data has been saved, we can call a fetch:
```text
curl --location 'http://localhost:8081/api/v1/sensors/01J5BAP6EKJ1S9049RESK9A1H8'
```

The response would look like:

```json
{
    "id": "01J5BAP6EKJ1S9049RESK9A1H8",
    "created": 1723735939539,
    "updated": 1723735939539,
    "name": "sensor-01",
    "serial": "ABC-001",
    "measures": [
        0.123,
        0.45,
        5.21
    ]
}
```