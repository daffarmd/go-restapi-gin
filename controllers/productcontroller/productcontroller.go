package productcontroller

import (
	"encoding/json"
	"net/http"

	"root/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var products []models.Product

	result := models.DB.Raw("SELECT * FROM products").Scan(&products)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}

func Show(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	result := models.DB.Raw("SELECT * FROM products WHERE id = ?", id).Scan(&product)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product})

}

func Create(c *gin.Context) {

	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result := models.DB.Exec("INSERT INTO products (nama_product, deskripsi) VALUES (?, ?)",
		product.NamaProduct, product.Deskripsi)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product})

}

func Update(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result := models.DB.Exec("UPDATE products SET nama_product = ?, deskripsi = ? WHERE id = ?",
		product.NamaProduct, product.Deskripsi, id)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product})

}

func Delete(c *gin.Context) {

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()

	query := "DELETE FROM products WHERE id = ?"
	result := models.DB.Exec(query, id)

	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
