# Client golang templates for can  {add version number here}

### What this template should be (Goals)
- We should be able to generate a client that complies with the entire [OpenApi 3.0 spec](https://swagger.io/specification/).
- It should have thorough error handling that bubbles up descriptive errors to the caller.
- It should make an api available to the user that's descriptively named, easy to use and lines up almost identically to 
  the wording used in the spec file. 
- The api should generate unit tests in place that verify the behaviour of the generated code with a separate

### What this template is right now (Reality)
Given the age of this project it seems prudent to reduce the scope of what a template NEEDS to be. 
- We need to properly handle errors coming from the setup of and execution of an http call as well as the handling of the 
  returned responses. It seems prudent to keep this up to standard at every point along this projects lifetime and t
  therefore should be implemented in it's entirety from the initial release, even if only as a first iteration.

  At this point we have made a conscious point to omit any special case handling and only emit errors that are
  obvious in nature or likely to be hit. 

- The entirety of the OpenApi 3.0 spec has not yet been achieved in can and so that cannot be expected for this first 
  iteration of the client template. At this point we support the following OpenAPI sub-specifications in these client 
  templates. 
  - A client struct that composes of the net/http client, an anonymous auth function and a config that allows us to set 
    port, host, protocol, and content-type for all requests. The config is validated upon instantiation and will throw
    an error if invalid.
  - Full use of net/url for url composition supporting both query parameters and path parameters with both being 
    available in a parameters struct. 
  - An `AuthFunc` composed into and set on instantiation of the client that allows direct access to the request for auth 
    related functionality to be implemented. This is the only part of this generated client that is not guaranteed to be 
    thread safe.
  - JSON encoding of byte streams
  - ~~JSON decoding of byte streams~~ To be reattempted: we currently return a `[]byte` and expect the caller to make 
    type assertions on raw data in order to find out which one of the available structs was returned by the generated 
    code
  - Basic context propagation. Next step here is to look into embedding otel trace contexts
  - Models for all objects required for use of this client. Same as the go-gin server templates
  - Go style error handling in the case of an unsuccessful request occuring
