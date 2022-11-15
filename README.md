# gin-in-a-can

[The existing openapi generator for Gin](https://openapi-generator.tech/docs/generators/go-gin-server)  is decent, but it lacks certain features that I need.
Ideally, I would contribute to it directly, but I need this sooner than I am able to learn how their build system, etc. work.

The features I need are:
* To be able to register an existing gin Engine with predefined middleware.
* To be able to use the generated code without modifying it.

## TODO
* use components to avoid duplicate ref structs
* use nullable types in parameters and request bodies
* Add parameter validation based on the `required` flag, regex patterns and other OpenAPI format specifications
* Add API Fuzz testing

# HackDay Plans
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

Spec is a Directed Acyclic Graph of composed nodes:
```go 
type Node struct {
    outputPath  string
    fileName    string 
    content     []byte
    parent      *Node
    children    []Node
}
```

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
