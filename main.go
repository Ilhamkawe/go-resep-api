package main

import (
	"github.com/gin-gonic/gin"
	"go-resep-api/entity"
	"go-resep-api/handler"
	"go-resep-api/src/bahan"
	"go-resep-api/src/detailresep"
	"go-resep-api/src/kategori"
	"go-resep-api/src/resep"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

func main() {

	dsn := "host=localhost user=postgres password=yourpassword dbname=go-resep port=5432 sslmode=disable TimeZone=Asia/Jakarta"
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
	//==================================================
	bahanRepository := bahan.NewRepository(db)
	bahanService := bahan.NewService(bahanRepository)
	bahanHandler := handler.NewBahanHandler(bahanService)
	//==================================================
	//Kategori Dependencies
	//==================================================
	kategoriRepository := kategori.NewRepository(db)
	kategoriService := kategori.NewService(kategoriRepository)
	kategoriHandler := handler.NewKategoriHandler(kategoriService)
	//==================================================
	resepRepository := resep.NewRepository(db)
	detailRepository := detailresep.NewRepository(db)
	resepService := resep.NewService(resepRepository, detailRepository, kategoriRepository)
	detailService := detailresep.NewService(detailRepository)
	resepHandler := handler.NewResepHandler(resepService, detailService)

	router := gin.Default()

	api := router.Group("/api/v1")

	api.GET("/bahan", bahanHandler.Index)
	api.POST("/bahan/store", bahanHandler.Store)
	api.DELETE("/bahan/:id/destroy", bahanHandler.Destroy)
	api.GET("/bahan/:id/restore", bahanHandler.Restore)
	api.DELETE("/bahan/:id/delete", bahanHandler.PermDestroy)
	api.PUT("/bahan/:id/update", bahanHandler.Update)

	api.GET("/kategori", kategoriHandler.Index)
	api.POST("/kategori/store", kategoriHandler.Store)
	api.DELETE("/kategori/:id/destroy", kategoriHandler.Destroy)
	api.GET("/kategori/:id/restore", kategoriHandler.Restore)
	api.DELETE("/kategori/:id/delete", kategoriHandler.PermDestroy)
	api.PUT("/kategori/:id/update", kategoriHandler.Update)

	api.POST("/resep/store", resepHandler.Store)
	api.POST("/resep/store/detail", resepHandler.StoreDetail)
	api.GET("/resep/:id/detail", resepHandler.FindByID)
	api.GET("/resep", resepHandler.Index)
	api.DELETE("/resep/:id/detail", resepHandler.DeleteDetail)
	api.DELETE("/resep/:id/delete", resepHandler.Delete)

	router.Run()
}
