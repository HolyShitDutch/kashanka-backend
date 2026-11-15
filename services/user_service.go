package services

type UserService struct {
	db DB
}

func NewUserService(db DB) *UserService {
	return &UserService{db}
}

func (s *UserService) AuthenticateByEmailPassword()
