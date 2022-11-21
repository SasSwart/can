
# Design
```mermaid
graph LR
    Yaml --> Spec --> Output[File]
```

```go 
type File struct {
    content []byte
    path    string
}
```

## OpenAPI spec - _Can_ tree representation
```mermaid
graph TD
    Spec --> Path
    Spec --> Components
    Components --> S1[Schemas] --> S1
    Path --> PathItem --> Operation
    Operation --> Parameters --> S3[Schemas] --> S3
    Operation --> Responses
    Operation --> RequestBody
    RequestBody --> MediaType
    Responses --> MediaType --> S2[Schemas] --> S2
```

## Process Representation
```mermaid
graph LR 
    1[1. Unmarshal]
    2[2. ReadRefs] --- TT[Traverse Tree]
    3[3. Render] --- TT[Traverse Tree]
```

Spec is a Directed Acyclic Graph of composed nodes.

## Can Package Structure
```mermaid 
graph TD
    Can[main]
    Can --> Render
    Can --> Sanitize
    Can --> Server
    Can --> Model
    Can --> Route
    Can --> OpenAPI
```
