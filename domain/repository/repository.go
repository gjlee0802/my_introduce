// repository.go
// model과 관련한 interface 선언

package repository
import "github.com/gjlee0802/my_introduce/domain/model"

type UserRepo interface {
	GetByID(id int) (*model.User, error)
	GetByName(name string) (*model.User, error)
	Create(user *model.User) (*model.User, error)
}