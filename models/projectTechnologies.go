package models


type ProjectTechnologies struct {
    ProjectID    uint `gorm:"primaryKey"`
    TechnologyID uint `gorm:"primaryKey"`
    Project      Project `gorm:"foreignKey:ProjectID"`
    Technology   Technology `gorm:"foreignKey:TechnologyID"`
}
