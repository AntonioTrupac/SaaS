package service

import (
	logging "github.com/AntonioTrupac/hannaWebshop/loggers"
	"github.com/AntonioTrupac/hannaWebshop/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductService interface {
	GetProducts() ([]*model.Product, error)
	GetProductById(id int) (*model.Product, error)
	CreateAProduct(input *model.Product) error
}

type Products struct {
	DB  *gorm.DB
	Log *logging.ZapLogger
}

func NewProducts(db *gorm.DB) ProductService {
	return &Products{
		DB: db,
	}
}

func (p Products) GetProducts() ([]*model.Product, error) {
	var products []*model.Product

	if err := p.DB.Preload("Category").Preload("Image").Find(&products).Error; err != nil {

		p.Log.Errorf("Error during GetProducts execution! %v", err)
		return nil, err
	}

	return products, nil
}

func (p Products) GetProductById(id int) (*model.Product, error) {
	var product *model.Product

	if err := p.DB.Preload("Category").Preload("Image").Where("id = ?", id).Find(&product).Error; err != nil {
		p.Log.Errorf("Error during GetProductById execution! %v ", err)

		return nil, err
	}

	return product, nil
}

func (p Products) CreateAProduct(input *model.Product) error {
	return p.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Omit(clause.Associations).Create(input).Error; err != nil {
			p.Log.Errorf("ERROR WHILE EXECUTING TRANSACTION! %v", err)

			return err
		}

		for _, value := range input.Image {
			value.ProductId = int(input.ID)
		}

		if err := tx.CreateInBatches(input.Image, 100).Error; err != nil {
			p.Log.Errorf("Error while inserting in the Image table! %v", err)

			return err
		}

		for _, categoryValue := range input.Category {
			categoryValue.ProductId = int(input.ID)
		}

		if err := tx.CreateInBatches(input.Category, 100).Error; err != nil {
			p.Log.Errorf("Error while inserting in the Category table! %v", err)

			return err
		}

		return nil
	})
}
