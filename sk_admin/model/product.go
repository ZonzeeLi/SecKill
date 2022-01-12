/**
    @Author:     ZonzeeLi
    @Project:    sk_admin
    @CreateDate: 1/7/2022
    @UpdateDate: xxx
    @Note:       商品数据库mysql处理
**/

package model

import (
	"log"
	"sk_admin/config"
)

type Product struct {
	ProductId   int    `json:"product_id"`   //商品Id
	ProductName string `json:"product_name"` //商品名称
	Total       int    `json:"total"`        //商品数量
	Status      int    `json:"status"`       //商品状态
}

type ProductModel struct {
}

func NewProductModel() *ProductModel {
	return &ProductModel{}
}

func (p *ProductModel) getTableName() string {
	return "product"
}

func (p *ProductModel) GetProductList() ([]map[string]interface{}, error) {
	conn := config.SecAdminConfCtx.DbConf.DbConn
	list, err := conn.Table(p.getTableName()).Get()
	if err != nil {
		log.Printf("Error : %v", err)
		return nil, err
	}
	return list, nil
}

func (p *ProductModel) CreateProduct(product *Product) error {
	conn := config.SecAdminConfCtx.DbConf.DbConn
	_, err := conn.Table(p.getTableName()).Data(map[string]interface{}{
		"product_name": product.ProductName,
		"total":        product.Total,
		"status":       product.Status,
	}).Insert()
	if err != nil {
		log.Printf("Error : %v", err)
		return err
	}
	return nil
}
