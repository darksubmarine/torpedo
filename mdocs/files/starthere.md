
## Installation
Download the installable for your distribution

### OSX
#### Intel 
```shell
curl --silent https://darksubmarine.com/downloads/torpedo-darwin-amd64-latest-installer.bin > torpedo-darwin-amd64-latest-installer.bin && \
sh torpedo-darwin-amd64-latest-installer.bin && \
rm torpedo-darwin-amd64-latest-installer.bin
```

#### ARM
```shell
curl --silent https://darksubmarine.com/downloads/torpedo-darwin-arm64-latest-installer.bin > torpedo-darwin-arm64-latest-installer.bin && \
sh torpedo-darwin-arm64-latest-installer.bin && \
rm torpedo-darwin-arm64-latest-installer.bin
```

### Linux
#### Intel
```shell
curl --silent https://darksubmarine.com/downloads/torpedo-linux-amd64-latest-installer.bin > torpedo-linux-amd64-latest-installer.bin && \
sh torpedo-linux-amd64-latest-installer.bin && \
rm torpedo-linux-amd64-latest-installer.bin
```

#### ARM
```shell
curl --silent https://darksubmarine.com/downloads/torpedo-linux-arm64-latest-installer.bin > torpedo-linux-arm64-latest-installer.bin && \
sh torpedo-linux-arm64-latest-installer.bin && \
rm torpedo-linux-arm64-latest-installer.bin
```

## Quick start

Creating your first CRUD application based on an entity definition.

### Creating the project dir

The required dir struct is so simple. You need a directory to put your entities definitions and a directory to put the
entity code. The next is an example of it:

```text
my-app/
  |_ .torpedo
  | |_ entities        All the entities yaml files
  |   |_ user.yaml
  |
  |_ domain            Main place to store the generated code
```

### Entity definition
The following snippets illustrates how to create an entity definition:

```yaml
#
# Entity: User
# Description: System user entity
# Date: 2022-12-16
# Author: Sebastian <sebastian@darksubmarine.com>
#

version: "1.0"
kind: entity
entity:
  name: "user"
  plural: "users"
  description: "System user entity"

  schema:
    fields:
      - name: name
        type: string
        description: "The user full name"
        doc: "The user full name"

      - name: socialNumber
        type: string
        description: "The user social number"
        doc: "The user social number"
        encrypted: true
        validate:
          regex:
            pattern: '^(?!(000|666|9))\d{3}-(?!00)\d{2}-(?!0000)\d{4}$'
            go: '(?m)^\d{3}-\d{2}-\d{4}$'

  adapters:
    input:
      - type: http
        metadata:
          map:
            socialNumber: social_number

    output:
      - type: memory
```

### Entity code generation

Once that you have the project dir and the entities definitions running the Torpedo command line tool the entity code
can be generated:

```shell
torpedo fire entity \
  --file my-app/.torpedo/entities/user.yaml \
  --output my-app/domain/ \
  --stack go \
  --package "github.com/myuser/my-app/domain"
```






