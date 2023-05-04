# base

## Installation

```bash
go get -u github.com/metadiv-io/base
```

## Highlights

* base.Repository[T any]

    - base.Repository.Save(tx *gorm.DB, entity *T) (*T, error)

    - base.Repository.SaveAll(tx *gorm.DB, entities *[]T) (*[]T, error)

    - base.Repository.Delete(tx *gorm.DB, entity *T) error

    - base.Repository.DeleteAll(tx *gorm.DB, entities *[]T) error

    - base.Repository.DeleteBy(tx *gorm.DB, clause *sql.Clause) error

    - base.Repository.FindOne(tx *gorm.DB, clause *sql.Clause) (*T, error)

    - base.Repository.FindAll(tx *gorm.DB, clause *sql.Clause) ([]T, error)

    - base.FindAllComplex(tx *gorm.DB, clause *sql.Clause, sort *sql.Sort, page *sql.Pagination) ([]T, *sql.Pagination, error)

    - base.Count(tx *gorm.DB, clause *sql.Clause) (int64, error)

    - base.FindByID(tx *gorm.DB, id uint) (*T, error)

* base.Mapper[T any]

    - base.Mapper.Map2Model(from interface{}) *T

    - base.Mapper.Map2Models(from interface{}) []T

## Setup AfterMap2Model

```go
type UserMapper[T any] struct {
    BaseMapper: base.Mapper[T] {
        AfterMap2Model: func(from interface{}, to *T) *T {
            e := from.(entity.User) // must not pointer
            to.Name = e.Name
            return to
        }
    }
}
```
