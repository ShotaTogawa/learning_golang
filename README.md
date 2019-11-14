# What is Rest(Representational State Transfer)?
- An architect style
- client-server communication via HTTP(Request / Response)

## Request

- Basic HTTP verbs
    - GET:    Requesting for resource
    - POST:   Creating a new resource
    - PUT:    Update a existing resource
    - DELETE: Delete a existing resource
- Resource paths (end points)
    - Paths are used for identifying resource location
- HTTP Header
    - Content-Type
    - Auth
- HTTP Request Payload

## Response
- The status of the request
    - status code
        - 200 - Success
        - 201 - Success
        - 400 - Bad Request(ex: invalid json object...)
        - 401 - Not Authorized(ex: you need to login...)
        - 403 - forbidden(ex: admins only...)
        - 500 - sever error
    - Response Payload(could be JSON object)