# TODOS API WITH GO
### 시작하게 된 배경
기존 프로젝트의 3계층 구조가 api(presentation) - service(application) - repository(data)로 나뉘어져 있었는데, 
개인적으로 presentation 계층에서 다음과 같은 것들이 불편하게 느껴졌다.
- query, path 파라미터 등을 파싱하고 `Atoi` 등을 사용해서 타입을 변경해주는 것(여기까지는 문제가 없음)
  - 특히 `strconv.Atoi()`를 하는 과정에서 매번 에러 핸들링이 들어가는데 이것들이 섞여있는게 불편했음
- 리턴 타입이 항상 `eho.HandlerFunc`라 presentation 계층을 보고 어떤 DTO로 응답이 나가는지 알기 어려웠다는 점
- 서버 패키지 레벨에서 "서비스"를 초기화 하고 Injection 한다는 점
  - 서비스가 많아진다면 불편해질 것 같다고 막연히 생각하였음

### 따라서 다음과 같이 바꾸면 어떨까 생각이 들었다.
- 서버 패키지 레벨에서 "컨트롤러"를 초기화 한다.
  - 특히 `*Echo.GET` 등의 함수들을 보면 `path`, `handlerFunc`, `...MiddlewareFunc`를 받아 `*Route`를 반환하도록 되어있다.
```go
// Initialize Controllers
todoController := controllers.NewTodoController(db, cfg.Setting())
// ...생략

v1 := TodoServer.Group("/api/v1/todos)
{
    v1.GET("/", api.HandleGetTodos(todoController))
    v1.GET("/todoId", api.HandleGetTodoById(todoController))
// ...생략
}
``` 
- 따라서 API 레벨의 `Handle~`의 핸들 함수들은 컨트롤러를 인자로 받아 `echo.HandlerFunc` 를 반환하도록 한다.
  - 이 핸들러 레벨에서는 파라미터를 파싱 및 예외응답처리를 하고, 성공했을 때 `echo.HandlerFunc`에 응답데이터를 담아 보내기만 한다.
- 컨트롤러 레벨에서는 파싱된 인자를 받아 presentation layer에 맞는 DTO들로 포장해서 응답을 한다.
- 서비스 레이어부터는 기존 3계층과 같다. 

### 하고 나서 느낀점
- 쓸데없이 복잡했다. 아직 잘 몰라서 그런지 모르지만 GO 에서 하는 예외처리들이 안그래도 불편한데 더 많아지기만 한 것 같다.
  - 하면 할수록 GO 스타일과 멀어지는 것 같은 느낌이 들었다.
    - 사실 GO 스타일이 무엇인지 잘 모르긴 함
  - 각 레이어별로 예외처리가 다르게 들어가야 할 것 같은데 그것도 부담스럽게 느껴진다.

### 다른 느낀점
- sqlboiler는 아예 DB 테이블을 세팅하고 시작해야 한다.
  - ddl 저장소를 별도로 두고 사용해야 할 듯
  - MSA를 하다보면 ddl 관리를 따로 해야할 필요성이 바로 생기는데 아마 그래서 그런것이 아닐까.
- 굉장히 빠르게 서버 개발을 할 수 있을 것 같다고 느꼈다.
- 처음에 "서비스가 많아진다면 불편하겠지" 생각해서 이걸 시작했다고 써놨는데 애초에 서버를 작게 쪼갠다면 별 문제 없을 것 같다.
  - 그래서 MSA 얘기를 하면 GO가 나오나보다.
- sqlboiler로 models를 만든 후에 DTO에 매핑하려고 하는데 
  - optional이 없다는 것이 충격, 근데 깊게 찾아보지 않고 그냥 `null.String` 으로 대충 해놨는데 맞는건지 모르겠다.
  - datetime도 무슨 `.Format("2006-01-02 15:04:05)` 이런게 생기던데 뭔지 잘 모르겠다.
- GO 얘기하면 라이브러리 생태계 이 얘기가 반드시 나오는데 아주 아주 편하고 좋았다. 굿!
