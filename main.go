package main

import (
	"github.com/gin-gonic/gin"
	"go-resep-api/entity"
	"go-resep-api/handler"
	"go-resep-api/src/bahan"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

func main() {

	dsn := "host=localhost user=postgres password=p@ssw0rdnd5 dbname=go-resep port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	err = db.AutoMigrate(
		&entity.Bahan{},
		&entity.Kategori{},
		&entity.Resep{},
		&entity.DetailResep{},
	)

	if err != nil {
		log.Fatal(err.Error())
	}

	//Bahan Dependencies
	//=================================================
	bahanRepository := bahan.NewRepository(db)
	bahanService := bahan.NewService(bahanRepository)
	bahanHandler := handler.NewBahanHandler(bahanService)
	//==================================================

	router := gin.Default()

	api := router.Group("/api/v1")

	api.GET("/bahan", bahanHandler.Index)
	api.POST("/bahan/store", bahanHandler.Store)
	api.DELETE("/bahan/:id/destroy", bahanHandler.Destroy)
	api.GET("/bahan/:id/restore", bahanHandler.Restore)
	api.DELETE("/bahan/:id/delete", bahanHandler.PermDestroy)
	api.PUT("/bahan/:id/update", bahanHandler.Update)
	router.Run()
}
