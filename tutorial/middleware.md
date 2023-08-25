# middleware
## What is middleware
Middleware is a function whose return type is HandlerFunc

    type HandlerFunc func(*Context)

    func DemoMiddleware() gin.HandlerFunc{
        return func(c *gin.Context){
            // operations before current api
            c.Next() 
            // operations after current api
        }
    }