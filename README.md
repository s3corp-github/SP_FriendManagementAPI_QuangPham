# **Project: go-training**
### **Description**:
Sample crud operation using Go, go-chi, migration, sqlboiler

### Database: 
postgesql

### Routing: 
go-chi

### Models: 
-Users

-Products

### API ENDPOINTS(localhost:5000)
    1. Create User:
        Path : http:/localhost:5000/users
        Method: Post
        Payload:
            {
                "name": "Test",
                "email":"test@gmail.com",
                "phone":"0123456789",
                "role":"ADMIN",
                "is_active": false
            }
        Success: 201
            {
                "ID": 41,
                "Name": "Test",
                "Email": "test@gmail.com",
                "Phone": "0123456789",
                "Role": "ADMIN",
                "IsActive": false
            }



### Quick Run Project:

    git clone https://github.com/congthang12312/golang-training.git
    
    cd golang-training
    
    go build main.go
    
    ./main.exe

