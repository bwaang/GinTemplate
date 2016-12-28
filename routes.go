// routes.go

package main

func initializeRoutes() {
  fileRoutes := router.Group("/file/v1")
  {
    fileRoutes.GET("/health", healthCheck)
    fileRoutes.GET("/version", versionCheck)
  }
}
