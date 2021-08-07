// usecase.go
// usecase 패키지가 맡는 역할
// usecase는 사용자에게 보여줄 출력을 위해 해당 출력을 생성하기 위한 처리 단계를 기술하는 곳

package usecase
import "github.com/gjlee0802/my_introduce/domain/model"

type RegistrationUsecase interface {
	RegisterUser(name, pass string) (*model.User, error)
	MatchUser(name, pass string) (*model.User, error)
}