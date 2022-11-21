# **Project: friend-management**
### Quick Run Project:
- Setup: make setup
- Run app: make run

### **Technology**:
- Using Go 1.19
- Postgesql 11
- go-chi 
- DB migration
- sqlboiler
- mockery
- viper

### Project architecture
- Workflow: Request => Internal/API/Rest => Internal/Controller => Internal/Repository => Database

- Three layers model:
  + Internal/API/Rest: Get request from httpRequest, decode, validate, call controllers, write httpResponse
  + Internal/Controller: Handle business logic, call repositories
  + Internal/Repository: Data access layer

### API ENDPOINTS(localhost:8080)
    1. Get List User:
        Path: http:/localhost:8080/users
        Method: GET
        Success: 200
          [
            {
            "Name": "andy",
            "Email": "andy@example.com"
            }
          ]

    2. Create User:
        Path: http:/localhost:8080/users
        Method: Post
        Payload:
            {
                "email":"test@example.com",
                "name":"test"
            }
        Success: 201
            {
                "ID": 1,
                "Email": "test@example.com"
                
            }
    3. Create friend connection between 2 email
        Path: http:/localhost:8080/friends/friend
        Method: Post
        Payload:
            {
                "friends": [
                    "andy@example.com",
                    "common@example.com"
                ]
            }
        Success: 201
            {
                "success":"true"
            }
    4. Retreive a list friend for an email address
        Path: http:/localhost:8080/friends/friends
        Method: Post
        Payload:
            {
                "email":"andy@example.com"
            }
        Success: 200
            {
                "success":true,
                "friends":[
                    "john@example.com"
                ],
                "count":1
            }
    5. Retreive common friend list between 2 email address
        Path: http:/localhost:8080/friends/commonfriends
        Method: Post
        Payload:
            {
                "friends":[
                    "andy@example.com",
                    "john@example.com"
                ]
            }
        Success: 200
            {
                "success":true,
                "friends":[
                    "common@example.com"
                ],
                "count":1
            }
    6. Subcribe to updates from an email
        Path: http:/localhost:8080/friends/subscription
        Method: Post
        Payload:
            {
                "requester":"andy@example.com",
                "target":"lisa@example.com"
            }
        Success: 201
            {
                "success":true
            }
    7. Block updates from an email
        Path: http:/localhost:8080/friends/block
        Method: Post
        Payload:
            {
                "requester":"andy@example.com",
                "target":"lisa@example.com"
            }
        Success: 201
            {
                "success":true
            }
    8. Retrieve all email address can receive update from an email
        Path: http:/localhost:8080/friends/emailreceive
        Method: Post
        Payload: 
            {
                "sender": "andy@example.com",
                "text": "Hello World! lisa@example.com"
            }
        Success: 200
            {
                "success": true,
                "recipients": [
                    "common@example.com",
                    "lisa@example.com"
                ]
            }
