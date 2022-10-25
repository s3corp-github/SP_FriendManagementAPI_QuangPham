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
    1. Create User:
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
    2. Create friend connection between 2 email
        Path: http:/localhost:8080/relations/createfriendrelation
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
    3. Retreive a list  for an email address
        Path: 

### Quick Run Project:

    git clone https://github.com/congthang12312/golang-training.git
    
    cd golang-training
    
    go build main.go
    
    ./main.exe
