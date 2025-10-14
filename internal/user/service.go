package user

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service { return &Service{repo: repo} }

func (s *Service) Migrate() error { return s.repo.Migrate() }

func (s *Service) Create(email, name string) (*User, error) {
	u := &User{Email: email, Name: name}
	return u, s.repo.Create(u)
}

func (s *Service) Get(id uint64) (*User, error) { return s.repo.GetByID(id) }

func (s *Service) Update(id uint64, email, name string) (*User, error) {
	u := &User{ID: id, Email: email, Name: name}
	return u, s.repo.Update(u)
}

func (s *Service) Delete(id uint64) error { return s.repo.Delete(id) }

func (s *Service) List(page, pageSize int) ([]User, int64, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize
	return s.repo.List(offset, pageSize)
}
