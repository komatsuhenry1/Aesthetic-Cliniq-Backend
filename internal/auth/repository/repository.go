package repository

type UserRepository interface {
	LoginUser(requestDto string) (string, error)
}

type userRepository struct {
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) LoginUser(requestDto string) (string, error) {
	return "token", nil
}
