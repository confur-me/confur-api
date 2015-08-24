package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"reflect"
)

func FindBy(field string, value interface{}) func(db *gorm.DB) *gorm.DB {
	condition := "= ?"

	return func(db *gorm.DB) *gorm.DB {
		fmt.Println(reflect.TypeOf(value))
		switch reflect.TypeOf(value).Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.String, reflect.Float32, reflect.Float64:
			v := fmt.Sprintf("%v", reflect.ValueOf(value))
			return db.Where(fmt.Sprintf("%v %v", field, condition), v)
		case reflect.Slice:
			condition = "IN (?)"
			return db.Where(fmt.Sprintf("%v %v", field, condition), reflect.ValueOf(value))
		case reflect.Struct:
			// TODO: work time values
			panic("not implemented yet")
		}
		panic(fmt.Sprintf("Unsupported type (%v) for %v", reflect.TypeOf(value), reflect.ValueOf(value)))
	}
}

func Paginate(limit int, page int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Limit(limit).Offset((page - 1) * limit)
	}
}

func GetRange(limit int, offset int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Limit(limit).Offset(offset)
	}
}

func Unpaginate(db *gorm.DB) *gorm.DB {
	return db.Offset(-1).Limit(-1)
}

func Active(db *gorm.DB) *gorm.DB {
	return db.Where("is_active = true")
}
