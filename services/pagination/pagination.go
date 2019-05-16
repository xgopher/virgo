package pagination

import (
	"math"

	"github.com/jinzhu/gorm"
)

// Param 分页参数
type Param struct {
	DB      *gorm.DB
	Page    int
	PerPage int
	OrderBy []string
}

// Paginator 分页返回
type Paginator struct {
	Total       int `json:"total"`
	From        int `json:"from"`
	To          int `json:"to"`
	PerPage     int `json:"per_page"`
	CurrentPage int `json:"current_page"`
	LastPage    int `json:"last_page"`
	PrevPage    int `json:"prev_page"`
	NextPage    int `json:"next_page"`
}

// Meta 分页头信息
type Meta struct {
	Paginator `json:"pagination"`
}

// Pagging 分页
func Pagging(p *Param, result interface{}) (*Meta, error) {
	db := p.DB

	if p.Page < 1 {
		p.Page = 1
	}
	if p.PerPage == 0 {
		p.PerPage = 10
	}
	if len(p.OrderBy) > 0 {
		for _, o := range p.OrderBy {
			db = db.Order(o)
		}
	}

	done := make(chan bool, 1)
	var paginator Paginator
	var count int
	var offset int

	go countRecords(db, result, done, &count)

	if p.Page == 1 {
		offset = 0
	} else {
		offset = (p.Page - 1) * p.PerPage
	}

	err := db.Limit(p.PerPage).Offset(offset).Find(result).Error
	
	<-done

	paginator.Total = count
	paginator.CurrentPage = p.Page
	paginator.PerPage = p.PerPage
	paginator.LastPage = int(math.Ceil(float64(count) / float64(p.PerPage)))
	paginator.From = offset + 1
	paginator.To = getTo(paginator, p)
	
	if p.Page > 1 {
		paginator.PrevPage = p.Page - 1
	} else {
		paginator.PrevPage = p.Page
	}

	if p.Page == paginator.LastPage {
		paginator.NextPage = p.Page
	} else {
		paginator.NextPage = p.Page + 1
	}

	r := Meta{
		Paginator: paginator,
	}
	
	return &r, err
}

func countRecords(db *gorm.DB, anyType interface{}, done chan bool, count *int) {
	db.Model(anyType).Count(count)
	done <- true
}

func getTo(paginator Paginator, p *Param) int{
	offset := paginator.From - 1
	if paginator.LastPage == paginator.CurrentPage {
		count := paginator.Total % p.PerPage 
		return offset + count
	}

	return offset + p.PerPage 
}
