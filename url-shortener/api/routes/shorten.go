package routes

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/udaichauhan/url_shortener/api/database"
	"github.com/udaichauhan/url_shortener/api/models"
	"github.com/udaichauhan/url_shortener/api/utils"
)

func ShortenURL(c *gin.Context) {
	var body models.Request

	// Parse JSON request body
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot Parse JSON: " + err.Error()})
		return
	}

	// Creating the Redis client for rate limiting
	r2 := database.CreateClient(1)
	defer r2.Close()

	// Check if Redis connection is working
	_, err := r2.Ping(database.Ctx).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Redis connection failed"})
		return
	}

	// Rate limiting logic
	val, err := r2.Get(database.Ctx, c.ClientIP()).Result()
	if err == redis.Nil {
		// First time user, set quota
		err := r2.Set(database.Ctx, c.ClientIP(), os.Getenv("API_QUOTA"), 30*60*time.Second).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set rate limit"})
			return
		}
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Rate limit check failed"})
		return
	} else {
		// User exists, check remaining quota
		valInt, err := strconv.Atoi(val)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid rate limit value"})
			return
		}

		if valInt <= 0 {
			limit, _ := r2.TTL(database.Ctx, c.ClientIP()).Result()
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"error":             "Rate limit exceeded",
				"rate_limit_reset":  limit / time.Nanosecond / time.Minute,
			})
			return
		}
	}

	// URL validation
	if !govalidator.IsURL(body.URL) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
		return
	}

	// Domain validation
	if !utils.IsDifferentDomain(body.URL) {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"error": "You cant hack this one",
		})
		return
	}

	// Ensure proper HTTP format
	body.URL = utils.EnsureHttpPerfect(body.URL)

	// Generate ID
	var id string
	if body.CustomShort == "" {
		id = uuid.New().String()[:6]
	} else {
		id = body.CustomShort
	}

	// Create Redis client for URL storage
	r := database.CreateClient(0)
	defer r.Close()

	// Check if Redis connection is working
	_, err = r.Ping(database.Ctx).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "URL storage Redis connection failed"})
		return
	}

	// Check if custom short URL already exists
	val, err = r.Get(database.Ctx, id).Result()
	if err != nil && err != redis.Nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check existing URL"})
		return
	}
	
	if val != "" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "URL Custom Short Already Exists",
		})
		return
	}

	// Set default expiry
	if body.Expiry == 0 {
		body.Expiry = 24
	}

	// Store URL in Redis
	err = r.Set(database.Ctx, id, body.URL, time.Duration(body.Expiry)*time.Hour).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to store URL: " + err.Error(),
		})
		return
	}

	// Prepare response
	resp := models.Response{
		Expiry:          body.Expiry,
		XRateLimitReset: 30,
		XRateRemaining:  10,
		URL:             body.URL,
		CustomShort:     "",
	}

	// Decrement rate limit
	err = r2.Decr(database.Ctx, c.ClientIP()).Err()
	if err != nil {
		// Log error but don't fail the request
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update rate limit"})
		// return
	}

	// Get remaining rate limit
	val, err = r2.Get(database.Ctx, c.ClientIP()).Result()
	if err == nil {
		resp.XRateRemaining, _ = strconv.Atoi(val)
	}

	// Get TTL
	ttl, err := r2.TTL(database.Ctx, c.ClientIP()).Result()
	if err == nil {
		resp.XRateLimitReset = ttl / time.Nanosecond / time.Minute
	}

	// Set final short URL
	domain := os.Getenv("DOMAIN")
	if domain == "" {
		domain = "http://localhost:8000" // Default domain
	}
	resp.CustomShort = domain + "/" + id

	c.JSON(http.StatusOK, resp)
}