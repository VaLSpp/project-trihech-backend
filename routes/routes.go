package routes

import (
    "project-trihech-backend/controllers"
    "project-trihech-backend/middleware"

    "github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    // Public routes
    r.POST("/register", controllers.Register)
    r.POST("/login", controllers.Login)


    // Protected routes
    protected := r.Group("/")
    protected.Use(middleware.AuthMiddleware())
    {
        // User routes
        protected.GET("/useprotecteds", controllers.GetUsers)
        protected.GET("/users/:id", controllers.GetUser)
        protected.POST("/users", controllers.CreateUser)
        protected.PUT("/users/:id", controllers.UpdateUser)
        protected.DELETE("/users/:id", controllers.DeleteUser)

        // UserDetails routes
        protected.GET("/user_details/:user_id", controllers.GetUserDetails)
        protected.POST("/user_details", controllers.CreateUserDetails)
        protected.PUT("/user_details/:user_id", controllers.UpdateUserDetails)
        protected.DELETE("/user_details/:user_id", controllers.DeleteUserDetails)

        // Role routes
        protected.GET("/roles", controllers.GetRoles)
        protected.GET("/roles/:id", controllers.GetRole)
        protected.POST("/roles", controllers.CreateRole)
        protected.PUT("/roles/:id", controllers.UpdateRole)
        protected.DELETE("/roles/:id", controllers.DeleteRole)

        // Permission routes
        protected.GET("/permissions", controllers.GetPermissions)
        protected.GET("/permissions/:id", controllers.GetPermission)
        protected.POST("/permissions", controllers.CreatePermission)
        protected.PUT("/permissions/:id", controllers.UpdatePermission)
        protected.DELETE("/permissions/:id", controllers.DeletePermission)

        // RolesPermissions routes
        protected.GET("/roles_permissions", controllers.GetRolesPermissions)
        protected.GET("/roles_permissions/:role_id/:permission_id", controllers.GetRolePermission)
        protected.POST("/roles_permissions", controllers.CreateRolePermission)
        protected.DELETE("/roles_permissions/:role_id/:permission_id", controllers.DeleteRolePermission)

        // Project routes
        protected.GET("/projects", controllers.GetProjects)
        protected.GET("/projects/:id", controllers.GetProject)
        protected.POST("/projects", controllers.CreateProject)
        protected.PUT("/projects/:id", controllers.UpdateProject)
        protected.DELETE("/projects/:id", controllers.DeleteProject)

        // Technology routes
        protected.GET("/technologies", controllers.GetTechnologies)
        protected.GET("/technologies/:id", controllers.GetTechnology)
        protected.POST("/technologies", controllers.CreateTechnology)
        protected.PUT("/technologies/:id", controllers.UpdateTechnology)
        protected.DELETE("/technologies/:id", controllers.DeleteTechnology)

        // ProjectTechnologies routes
        protected.GET("/project_technologies", controllers.GetProjectTechnologies)
        protected.GET("/project_technologies/:project_id/:technology_id", controllers.GetProjectTechnology)
        protected.POST("/project_technologies", controllers.CreateProjectTechnology)
        protected.DELETE("/project_technologies/:project_id/:technology_id", controllers.DeleteProjectTechnology)
    }

    return r
}
