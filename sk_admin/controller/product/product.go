/**
    @Author:     ZonzeeLi
    @Project:    sk_admin
    @CreateDate: 1/7/2022
    @UpdateDate: xxx
    @Note:       商品api
**/

package product

import (
	"fmt"
	"log"
	"sk_admin/model"
	"sk_admin/service"

	"github.com/gin-gonic/gin"
)

func CreateProduct(ctx *gin.Context) {
	var product = &model.Product{}
	if err := ctx.ShouldBindJSON(product); err != nil {
		log.Printf("ProductServer.CreateProduct, err : %v", err)
		ctx.JSON(400, map[string]interface{}{
			"code": 400,
			"msg":  "failed",
		})
		return
	}

	//product.ProductName = ctx.PostForm("product_name")
	//product.Total, _ = com.StrTo(ctx.PostForm("product_total")).Int()
	//product.Status, _ = com.StrTo(ctx.PostForm("status")).Int()
	fmt.Println(product.ProductName, product.Total, product.Status)
	productServer := service.NewProductServer()
	err := productServer.CreateProduct(product)
	if err != nil {
		log.Printf("ProductServer.CreateProduct, err : %v", err)
		ctx.JSON(400, map[string]interface{}{
			"code": 400,
			"msg":  "failed",
		})
		return
	}

	ctx.JSON(200, map[string]interface{}{
		"code": 200,
		"msg":  "success",
	})
	return
}

func GetPorductList(ctx *gin.Context) {
	productService := service.NewProductServer()
	productList, err := productService.GetProductList()
	if err != nil {
		log.Printf("ProductService.productList, err : %v", err)
		ctx.JSON(400, map[string]interface{}{
			"code": 400,
			"msg":  "failed",
		})
		return
	}

	ctx.JSON(200, map[string]interface{}{
		"code": 200,
		"msg":  "success",
		"data": productList,
	})
	return
}
