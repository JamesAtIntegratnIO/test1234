package main

import (
    "log"
    "net/http"
    "os"
    
    "github.com/gin-gonic/gin"
    
    
    
    
)





func main() {
    
    
    
    
    // Initialize router
    r := gin.Default()
    
    
    
    // Routes
    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Welcome to test1234",
            "version": "1.0.0",
        })
    })
    
    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "status": "ok",
            "service": "test1234",
        })
    })
    
    // Start server
    port := os.Getenv("PORT")
    if port == "" {
        port = "5173"
    }
    
    log.Printf("Server starting on port %s", port)
    r.Run(":" + port)
}




