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
	var body models.Request;

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : "Cannot Parse JSON"})
		return;
	}

	//creating the client
	r2 := database.CreateClient(1);
	defer r2.Close();

	val, err := r2.Get(database.Ctx, c.ClientIP()).Result()
	if err == redis.Nil{
		_ = r2.Set(database.Ctx, c.ClientIP(), os.Getenv("API_QUOTA"), 30*60*time.Second)
	}else {
		val , _ = r2.Get(database.Ctx, c.ClientIP()).Result();
		valInt, _ := strconv.Atoi(val);

		if valInt <= 0 {
			limit, _ := r2.TTL(database.Ctx, c.ClientIP()).Result();
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"error" : "rate limit exceeded",
				"rate_limit_reset" : limit/time.Nanosecond/time.Minute,
			})
			return;
		}
	}

	if !govalidator.IsURL(body.URL){
		c.JSON(http.StatusBadRequest, gin.H{"error" : "Invalid URL"});
		return;
	}

	if !utils.IsDifferentDomain(body.URL){
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"error" : "You cant hack this one",
		})
		return;
	}

	body.URL = utils.EnsureHttpPerfect(body.URL);

	var id string;

	if body.CustomShort == "" {
		id = uuid.New().String()[:6];
	}else{
		id = body.CustomShort;
	}

	r := database.CreateClient(0);
	defer r.Close();

	val, _ = r.Get(database.Ctx, id).Result();
	if val != "" {
		c.JSON(http.StatusForbidden, gin.H{
			"error" : "URL Custom Short Already Exists",
		});
		return;
	}

	if body.Expiry == 0 {
		body.Expiry = 24;
	}

	r.Set(database.Ctx, id, body.URL, body.Expiry*3600*time.Second).Err();
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : "Unable to connect to the redis server",
		});
		return;
	}

	resp := models.Response{
		Expiry : body.Expiry,
		XRateLimitReset : 30,
		XRateRemaining : 10,
		URL : body.URL,
		CustomShort : "",
	};

	r2.Decr(database.Ctx, c.ClientIP());

	val, _ = r2.Get(database.Ctx, c.ClientIP()).Result();
	resp.XRateRemaining, _ = strconv.Atoi(val);

	ttl, _ := r2.TTL(database.Ctx, c.ClientIP()).Result();
	resp.XRateLimitReset = ttl/time.Nanosecond/time.Minute;

	resp.CustomShort = os.Getenv("DOMAIN") + "/" + id;

	c.JSON(http.StatusOK, resp);
}