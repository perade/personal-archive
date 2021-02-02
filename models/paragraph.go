package models

import (
	"gorm.io/gorm"
	"time"
)

type Paragraph struct {
	ID                int64             `gorm:"column:id;primarykey" json:"id"`
	NoteID            int64             `gorm:"column:note_id;type:integer;not null" json:"noteID"`
	Seq               int               `gorm:"column:seq;type:integer;not null" json:"seq"`
	Content           string            `gorm:"column:content;type:text" json:"content"`
	ReferenceArticles ReferenceArticles `gorm:"foreignKey:ParagraphID" json:"referenceArticles"`
	ReferenceWebs     ReferenceWebs     `gorm:"foreignKey:ParagraphID" json:"referenceWebs"`
	Created           time.Time         `gorm:"column:created;type:datetime;not null" json:"created"`
	LastModified      time.Time         `gorm:"column:last_modified;type:datetime;not null" json:"lastModified"`
}

func (p *Paragraph) TableName() string {
	return "paragraph"
}

func (p *Paragraph) BeforeSave(db *gorm.DB) error {
	if p.Created.IsZero() {
		p.Created = time.Now()
	}
	p.LastModified = time.Now()
	return nil
}

type Paragraphs []*Paragraph

func (p Paragraphs) ExtractIDs() []int64 {
	ids := []int64{}
	for _, paragraph := range p {
		ids = append(ids, paragraph.ID)
	}
	return ids
}

func (p Paragraphs) ExtractReferenceArticleIDs() []int64 {
	ids := []int64{}
	for _, paragraph := range p {
		ids = append(ids, paragraph.ReferenceArticles.ExtractIDs()...)
	}
	return ids
}

func (p Paragraphs) ExtractReferenceArticleArticleIDs() []int64 {
	ids := []int64{}
	for _, paragraph := range p {
		ids = append(ids, paragraph.ReferenceArticles.ExtractArticleIDs()...)
	}
	return ids
}

func (p Paragraphs) ExtractReferenceWebIDs() []int64 {
	ids := []int64{}
	for _, paragraph := range p {
		ids = append(ids, paragraph.ReferenceWebs.ExtractIDs()...)
	}
	return ids
}

func (p Paragraphs) MaxSeq() int {
	max := -1
	for _, paragraph := range p {
		if paragraph.Seq > max {
			max = paragraph.Seq
		}
	}
	return max
}
