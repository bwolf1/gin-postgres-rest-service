package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/bwolf1/gin-postgres-rest-service/pkg/config"
	"github.com/bwolf1/gin-postgres-rest-service/pkg/service"
	"github.com/gin-gonic/gin"
)

type Rest struct {
	products *service.Product
}

// Create an interface for the server so that it can be mocked in tests
type Server interface {
	listProducts(c *gin.Context)
	getProduct(c *gin.Context)
	getProductVersions(c *gin.Context)
}

// New creates a new server
func New(cfg *config.Config) (*Rest, error) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(AuthMiddleware())
	products, err := service.New(cfg)
	if err != nil {
		return nil, err
	}

	r := &Rest{products}
	router.GET("/products", r.listProducts)
	router.GET("/products/:id", r.getProduct)
	router.GET("/products/:id/versions", r.getProductVersions)
	// TODO: Get the host and port from the config so that it can be changed
	// easily for different environments (e.g. dev, staging, prod) without having
	// to recompile the binary
	router.Run("localhost:8080")

	return r, nil
}

// Add ad authentication middleware for user authentication
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Check for the Authorization header and validate the JWT token
	}
}

func (r *Rest) listProducts(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	start, _ := strconv.Atoi(c.DefaultQuery("start", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	serviceName := c.Query("name")

	sort := c.DefaultQuery("sort", "name")
	products, err := r.products.ListProducts(c, serviceName, start, pageSize, sort)
	if err != nil {
		log.Printf("failed to list products: %v", err)
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": fmt.Sprintf("failed to list products: %v", err),
			},
		)
		return
	}

	c.JSON(http.StatusOK, products)
}

func (r *Rest) getProduct(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	serviceID := c.Param("id")
	product, err := r.products.GetProduct(c, serviceID)
	if err != nil {
		log.Printf("failed to get product by ID %q: %v", serviceID, err)
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": fmt.Sprintf("failed to get product by ID %q: %v", serviceID, err),
			},
		)
		return
	}

	c.JSON(http.StatusOK, product)
}

func (r *Rest) getProductVersions(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	versions, err := r.products.GetProductVersions(c, c.Param("id"))
	if err != nil {
		log.Printf("failed to get product versions: %v", err)
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": fmt.Sprintf("failed to get product versions: %v", err),
			},
		)
		return
	}

	c.JSON(http.StatusOK, versions)
}
