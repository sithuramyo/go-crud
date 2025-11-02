package services

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PaginationQuery struct {
	Page   int
	Limit  int
	Search string
}

func ParsePagination(c *gin.Context) PaginationQuery {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}

	search := c.DefaultQuery("q", "")

	return PaginationQuery{
		Page:   page,
		Limit:  limit,
		Search: search,
	}
}

func List[T any](c *gin.Context, db *gorm.DB, searchFields []string, preloads ...string) {
	var items []T
	pq := ParsePagination(c)

	offset := (pq.Page - 1) * pq.Limit

	var model T
	query := db.Model(&model)

	for _, p := range preloads {
		query = query.Preload(p)
	}

	if pq.Search != "" && len(searchFields) > 0 {
		cond := ""
		args := []interface{}{}
		for i, field := range searchFields {
			if i > 0 {
				cond += " OR "
			}
			cond += field + " ILIKE ?"
			args = append(args, "%"+pq.Search+"%")
		}
		query = query.Where(cond, args...)
	}

	var total int64
	query.Count(&total)

	if err := query.Limit(pq.Limit).Offset(offset).Find(&items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       items,
		"page":       pq.Page,
		"limit":      pq.Limit,
		"total":      total,
		"totalPages": (total + int64(pq.Limit) - 1) / int64(pq.Limit),
	})
}

func Get[T any](c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var item T
	if err := db.First(&item, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	c.JSON(http.StatusOK, item)
}

func Create[T any](c *gin.Context, db *gorm.DB) {
	var item T
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Create(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, item)
}

func Update[T any](c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var item T
	if err := db.First(&item, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	var input T
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&item).Updates(input)
	c.JSON(http.StatusOK, item)
}

func Delete[T any](c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var item T
	if err := db.Delete(&item, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully"})
}
