package implementations

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CRUD[T any] struct {
	DB           *gorm.DB
	SearchFields []string
	Preloads     []string
}

func NewCRUD[T any](db *gorm.DB, searchFields []string, preloads ...string) *CRUD[T] {
	return &CRUD[T]{DB: db, SearchFields: searchFields, Preloads: preloads}
}

func (r *CRUD[T]) List(c *gin.Context)   { List[T](c, r.DB, r.SearchFields, r.Preloads...) }
func (r *CRUD[T]) Get(c *gin.Context)    { Get[T](c, r.DB) }
func (r *CRUD[T]) Create(c *gin.Context) { Create[T](c, r.DB) }
func (r *CRUD[T]) Update(c *gin.Context) { Update[T](c, r.DB) }
func (r *CRUD[T]) Delete(c *gin.Context) { Delete[T](c, r.DB) }
