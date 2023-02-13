package controller

import (
	// 文字列と基本データ型の変換パッケージ
	strconv "strconv"

	// Gin
	"github.com/gin-gonic/gin"

	// エンティティ(データベースのテーブルの行に対応)
	entity "shoppingapp/models/entity"

	// DBアクセス用モジュール
	db "shoppingapp/models/db"
)

// FetchAllProducts は 全ての商品情報を取得する
func FetchAllProducts(c *gin.Context) {
	resultProducts := db.FindAllProducts()

	// URLへのアクセスに対してJSONを返す
	c.JSON(200, resultProducts)
}

// FindProduct は 指定したIDの商品情報を取得する
func FindProduct(c *gin.Context) {
	productIDStr := c.Query("productID")

	productID, _ := strconv.Atoi(productIDStr)

	resultProduct := db.FindProduct(productID)

	// URLへのアクセスに対してJSONを返す
	c.JSON(200, resultProduct)
}

// AddProduct は 商品をDBへ登録する
func AddProduct(c *gin.Context) {
	productName := c.PostForm("productName")
	productMemo := c.PostForm("productMemo")

	var product = entity.Product{
		Name:  productName,
		Memo:  productMemo,
		State: entity.NotPurchased,
	}

	db.InsertProduct(&product)
}

// ChangeStateProduct は 商品情報の状態を変更する
func ChangeStateProduct(c *gin.Context) {
	reqProductID := c.PostForm("productID")
	reqProductState := c.PostForm("productState")

	productID, _ := strconv.Atoi(reqProductID)
	productState, _ := strconv.Atoi(reqProductState)
	changeState := entity.ChangeState(productState)

	db.UpdateStateProduct(productID, changeState)
}

// DeleteProduct は 商品情報をDBから削除する
func DeleteProduct(c *gin.Context) {
	productIDStr := c.PostForm("productID")

	productID, _ := strconv.Atoi(productIDStr)

	db.DeleteProduct(productID)
}
