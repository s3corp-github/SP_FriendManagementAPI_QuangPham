# **Project: friend-management**
### **Description**:
Using Go, go-chi, migration, sqlboiler, mockery, viper

### Database: 
postgesql

### Load Config file:
viper

### Routing: 
go-chi

### Models: 
-Users

-Relations

### API ENDPOINTS(localhost:8080)
    1. Get List User:
        Path: http:/localhost:8080/users
        Method: GET
        Success: 200
            {
              "Email": [
                  "andy@example.com",
                  "john@example.com",
                  "common@example.com",
                  "lisa@example.com"
              ]
            }
    2. Create User:
        Path: http:/localhost:8080/users
        Method: Post
        Payload:
            {
                "email":"test@example.com",
                "phone":"0123456789",
                "is_active": true
            }
        Success: 201
            {
                "ID": 1,
                "Email": "test@example.com",
                "Phone": "0123456789",
                "IsActive": true
            }
    3. Create friend connection between 2 email
        Path: http:/localhost:8080/relations/friend
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
    4. Retreive a list  for an email address
        Path: http:/localhost:8080/relations/friends
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
        Path: http:/localhost:8080/relations/commonfriends
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
        Path: http:/localhost:8080/relations/subscription
        Method: Post
        Payload:
            {
                "requestor":"andy@example.com",
                "target":"lisa@example.com"
            }
        Success: 201
            {
                "success":true
            }
    7. Block updates from an email
        Path: http:/localhost:8080/relations/block
        Method: Post
        Payload:
            {
                "requestor":"andy@example.com",
                "target":"lisa@example.com"
            }
        Success: 201
            {
                "success":true
            }
    8. Retrieve all email address can receive update from an email
        Path: http:/localhost:8080/relations/emailreceive
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
    

### Quick Run Project:

    1.Run command line: git clone https://github.com/s3corp-github/SP_FriendManagementAPI_QuangPham.git
    
    2.Run command line: docker compose up

    3.Using postman to get list user 

### Project architecture
- Workflow: Request => Internal/API/Rest => Internal/Controller => Internal/Repository => Database

- Three layers model:
    + Internal/API/Rest: Get request from httpRequest, decode, validate, call controllers, write httpResponse
    + Internal/Controller: Handle business logic, call repositories
    + Internal/Repository: Data access layer 

    

    

