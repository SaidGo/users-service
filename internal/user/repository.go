package user

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository { return &Repository{db: db} }

func (r *Repository) Migrate() error { return r.db.AutoMigrate(&User{}) }

func (r *Repository) Create(u *User) error { return r.db.Create(u).Error }

func (r *Repository) GetByID(id uint64) (*User, error) {
	var u User
	if err := r.db.First(&u, id).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *Repository) Update(u *User) error {
	return r.db.Model(&User{}).Where("id = ?", u.ID).Updates(u).Error
}

func (r *Repository) Delete(id uint64) error {
	return r.db.Delete(&User{}, id).Error
}

func (r *Repository) List(offset, limit int) ([]User, int64, error) {
	var res []User
	var total int64
	if err := r.db.Model(&User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := r.db.Order("id asc").Offset(offset).Limit(limit).Find(&res).Error; err != nil {
		return nil, 0, err
	}
	return res, total, nil
}
