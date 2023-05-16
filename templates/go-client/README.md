# Client templates for can 

##### What a template should be 
We should be able to generate a client that complies with the entire openApi 3.0 spec.
It should have thorough error handling that bubbles up descriptive errors to the caller.

##### What a template is right now
Given the age of this project it seems prudent to reduce the scope of what a template NEEDS to be. 
We need to properly handle errors coming from the setup of and execution of an http call as well as the handling of the 
returned responses. The initial goal is one where we omit any special case handling and only emit errors that are
obvious in nature or likely to be hit. 
The entirety of the OpenApi 3.0 spec is not going to be achieved in this first iteration. We aim to have the following
features available in can {{ INSERT NEXT VERSION NUMBER HERE }}:
- A config struct that allows integration through setting the following fields
  - BasePath
  - *Port
  - Auth helper methods
  - JSON encoding of byte streams
  - JSON decoding of byte streams
- A request object created through the use of {{ PackageName }}.{{ EndpointName }}{{ MethodName }}() with the following
  features:
  - An optional ctx parameter
  - Url constructed
  - Auth headers preapplied
  - A {{ EndpointName }}{{ MethodName }}.Do() method that 
    - takes a parameters struct and constructs the correct body OR
    - takes a parameters struct and constructs the correct query parameters
    - tests for a success response and throws an error if failed
    - reads the body and parses the json returned in the response
    - returns a response struct
    - 
