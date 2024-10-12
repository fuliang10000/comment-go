package main

import (
	"flag"
	"fmt"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gen/field"

	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/rawsql"
)

var (
	sqlFile = flag.String("sql", "", "sql path")
	sqlDns  = flag.String("dns", "debian-sys-maint:l3BkmCtvLyjPCMlw@tcp(127.0.0.1:3306)/comment?charset=utf8mb4&parseTime=True&loc=Local", "sql dns")
)

func main() {
	flag.Parse()
	var db *gorm.DB
	var err error
	if len(*sqlFile) > 0 {
		fmt.Println(strings.Split(*sqlFile, ","))
		db, err = gorm.Open(rawsql.New(rawsql.Config{
			FilePath: strings.Split(*sqlFile, ","),
		}))

	} else {
		db, err = gorm.Open(mysql.Open(*sqlDns), &gorm.Config{})
	}
	if err != nil {
		panic(err)
	}
	c := gen.Config{
		OutPath:           "dao/query",
		WithUnitTest:      true,
		FieldNullable:     true,
		FieldCoverable:    true,
		FieldSignable:     false,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
		Mode:              gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	}
	// gorm 生成工具
	g := gen.NewGenerator(c)
	g.WithOpts(gen.FieldModify(func(f gen.Field) gen.Field {
		if f.ColumnName == "deletedAt" {
			f.Type = "soft_delete.DeletedAt"
			f.GORMTag.Set(field.TagKeyGormDefault, "0")
		}
		return f
	}))

	g.UseDB(db) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(g.GenerateAllTable()...)

	// Generate the code
	g.Execute()
}
