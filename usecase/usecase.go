// usecase.go
// usecase 패키지가 맡는 역할
// usecase는 사용자에게 보여줄 출력을 위해 해당 출력을 생성하기 위한 처리 단계를 기술하는 곳
// Entities는 이 Use cases에 대해 전혀 알지 못합니다. 하지만 Use cases는 Entities를 알고있죠.
// 하지만 이 Use cases 계층의 변경이 Entities에 영향을 미쳐서는 안됩니다.
// 또한 Use cases가 DB, UI 같은 외부 환경의 변경으로 인해 영향을 받지 않아야 합니다.
// 왜냐면 그냥 정말 단순히 Use cases이기 때문에
// 데이터가 어디에 저장되어있든 외부 프레임워크가 바뀌든 상관없는거에요.
//  하지만! 앱 전체의 작동 방식 변경(?)은 Use cases에 영향을 미칠 수 있습니다

package usecase
import "github.com/gjlee0802/my_introduce/domain/model"

type RegistrationUsecase interface {
	RegisterUser(name, pass string) (*model.User, error)
	MatchUser(name, pass string) (*model.User, error)
}