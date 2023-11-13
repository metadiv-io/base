# Base Package Documentation

## Installation

To incorporate the base package into your Go project, execute the following go get command:

```bash
go get -u github.com/metadiv-io/base
```

## Key Features

### Repository

The Repository in the base package provides a set of fundamental database operations for entities. It includes the following methods:

- Save(tx *gorm.DB, entity *T) (*T, error): Saves a single entity.
- SaveAll(tx *gorm.DB, entities []T) ([]T, error): Saves multiple entities.
- Delete(tx *gorm.DB, entity *T) error: Deletes a single entity.
- DeleteAll(tx *gorm.DB, entities *[]T) error: Deletes multiple entities.
- DeleteBy(tx *gorm.DB, clause *sql.Clause) error: Deletes entities based on a specified clause.
- FindOne(tx *gorm.DB, clause *sql.Clause) (*T, error): Retrieves a single entity based on a clause.
- FindAll(tx *gorm.DB, clause *sql.Clause) ([]T, error): Retrieves multiple entities based on a clause.
- FindAllComplex(tx *gorm.DB, clause *sql.Clause, sort *sql.Sort, page *sql.Pagination) ([]T, *sql.Pagination, error): Retrieves complex entities with sorting and pagination.
- Count(tx *gorm.DB, clause *sql.Clause) (int64, error): Counts entities based on a clause.
- FindByID(tx *gorm.DB, id uint) (*T, error): Retrieves an entity by its unique identifier.

### Mapper

The Mapper in the base package facilitates the mapping of data between structs. It includes the following methods:

- Map2Model(from interface{}) *T: Maps data from an interface to a single model.
- Map2Models(from interface{}) []T: Maps data from an interface to a slice of models.

## Setup AfterMap2Model

The AfterMap2Model feature enables additional customization during the mapping process. It can be utilized by creating a mapper struct and implementing the desired mapping logic. As an example:

```go
type UserMapper[T any] struct {
    BaseMapper: base.Mapper[T] {
        AfterMap2Model: func(from interface{}, to *T) *T {
            e := from.(entity.User) // must not be a pointer
            to.Name = e.Name
            return to
        }
    }
}
```

This allows users to define specific logic to be executed after the mapping process, enhancing the flexibility of the base package.
