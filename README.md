# Golangの原則(インターフェースは使用する側で定義する)

### **インターフェイスを受け入れ、構造体を返す** Accept interfaces, return structs

---

## 間違いな例
postgres.go
構造体をReturnするのではなく、インターフェースをReturnしている
これでは構造体のフィールドを活用できない。
```go
package db

type Store interface {
	Insert(item interface{}) error
	Get(id int) error
}

type MyStore struct {
	db *sql.DB
}

func NewStore() Store { ... } //func to initialise DB
func (s *Store) Insert(item interface{}) error { ... } //insert item
func (s *Store) Get(id int) error { ... } //get item by id
```
user.go
```go
package user

type UserService struct {
	store db.Store
}

func NewUserService(s db.Store) *UserService {
	return &UserService{
		store: s,
	}
}
func (u *UserService) CreateUser() { ... }
func (u *UserService) RetrieveUser(id int) User { ... }
```

## 正しい例
戻り値は構造体で引数でInterfaceを受け取っている
db.go
```go
package db
type Store struct {
	db *sql.DB
}
func NewStore() *Store { ... } //func to initialise DB
func (s *Store) Insert(item interface{}) error { ... } //insert item
func (s *Store) Get(id int) error { ... } //get item by id
```
使用する側でインターフェースを使用する良い例
user.go
```go
package user
type UserStore interface {
	Insert(item interface{}) error
	Get(id int) error
}
type UserService struct {
	store UserStore
}
// Accepting interface here!
func NewUserService(s UserStore) *UserService {
	return &UserService{
		store: s,
	}
}
func (u *UserService) CreateUser() { ... }
func (u *UserService) RetrieveUser(id int) User { ... }
```

---

## 実際に標準ライブラリーでも使用箇所でインターフェースを定義している
### ioパッケージ
```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
func Copy(dst Writer, src Reader) (written int64, err error)
```

### httpパッケージ
```go
type Handler interface { 
    ServeHTTP(ResponseWriter, *Request) 
} 
func ListenAndServe(addr string, handler Handler) error
```

## 参考文献
[https://medium.com/@mbinjamil/using-interfaces-in-go-the-right-way-99384bc69d39](https://medium.com/@mbinjamil/using-interfaces-in-go-the-right-way-99384bc69d39)
[https://zenn.dev/pranc1ngpegasus/articles/a8c92235bec641](https://zenn.dev/pranc1ngpegasus/articles/a8c92235bec641)

[非常に良い: https://bryanftan.medium.com/accept-interfaces-return-structs-in-go-d4cab29a301b](https://bryanftan.medium.com/accept-interfaces-return-structs-in-go-d4cab29a301b)



