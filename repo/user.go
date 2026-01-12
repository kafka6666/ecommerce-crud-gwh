package repo

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

type UserRepo interface {
	Create(u *User) (*User, error)
	ListAll() ([]*User, error)
	Find(email, password string) (*User, error)
}

type userRepo struct {
	userList []*User
}

func NewUserRepo() UserRepo {
	return &userRepo{}
}

func (r *userRepo) Create(u *User) (*User, error) {
	u.ID = len(r.userList) + 1
	r.userList = append(r.userList, u)
	return u, nil
}

func (r *userRepo) ListAll() ([]*User, error) {
	if len(r.userList) == 0 {
		return nil, nil
	}
	return r.userList, nil
}

func (r *userRepo) Find(email, password string) (*User, error) {
	for _, u := range r.userList {
		if u.Email == email && u.Password == password {
			return u, nil
		}
	}

	return nil, nil
}
