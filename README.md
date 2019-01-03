# protoc-gen-rest

a simple protobuf to golang and typescript generator. Because I'm lazy to write both client and server for every project.

```protobuf

message Req {
  string name = 1;
}

message Res {
  string message = 1;
}

service Greeting {
  rpc sayHello(Req) returns (Res);
  rpc sayGoodbye(Req) returns (Res);
}

```

```go

package rest

type Error struct {
  code int
  message string
  meta map[string]interface{}
}

func (e Error) Error() string {
  return fmt.Sprintf("")
}

func NewError(code int, message string, meta map[string]interface{}) error {
  return &Error{
    code: code,
    message: message,
    meta: meta,
  }
}

func jsonReader(r io.Reader, body interface{}) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(body)
}

```

```go

import (
  "net/http"

  restUtil "github.com/nulloop/rest-util"
)

type Req struct {
  Name string `json:"name"`
}

type Res struct {
  Message string `json:"message"`
}

type Greeting interface {
  SayHello(Req) (Res, error)
  SayGoodbye(Req) (Res, error)
}

type greeting struct {}

func (r *greeting) sayHello(w http.ResponseWriter, r *http.Request) {
  req := &Req{}
  err := restUtil.ReadJson(r.Body, req)
  if err != nil {

  }
}

func (r *greeting) sayGoodbye(w http.ResponseWriter, r *http.Request) {

}

func RegisterGretting() http.Router {
  r := chi.NewRouter()

  handler := greeting{}
  r.Post("/greeting/sayhello", handler.sayHello)
  r.Post("/greeting/saygoodbye", handler.sayGoodbye)

  return r
}

```

```ts
```
