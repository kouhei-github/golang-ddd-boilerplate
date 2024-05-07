# 1. Golangの原則(インターフェースは使用する側で定義する)

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
func (s *MyStore) Insert(item interface{}) error { ... } //insert item
func (s *MyStore) Get(id int) error { ... } //get item by id
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



---

# 2. DBのマイグレーションファイルについて
Golangでmigrationを管理は、[golang-migrate](https://github.com/golang-migrate/migrate)で行う



## マイグレーションファイルの作成
コンテナ上の **/app**で下記コマンドを打つ
```shell
# make migrate_create name=create_table_user
migrate create -ext sql -dir ./migrations -seq 作成するファイル名(create_users)
```
なぜ **/app** なのかというと、migrationsファイルをロジックを含むpublicフォルダに作成したくないから。

## マイグレートの適用
```shell
make migrate_up
```

## マイグレートのロールバック
```shell
make migrate_down
```

## Golangの構造体からSQLを作成してもらう
```text
golangでGORMを使用しています。 下記構造体でテーブルを作成したいです。 SQL文を書いてください


type User struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	UserName string `json:"user_name"`
	Email    string `gorm:"not null" gorm:"unique" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Image    string `json:"image"`
	UserAuth UserAuth
}

type UserAuth struct {
	ID           int    `gorm:"primaryKey" json:"id"`
	PasswordHash string `json:"-"`
	PasswordSalt string `json:"password_salt"`
	UserID       int    `json:"user_id"`
}


databaseのマイグレーションにはgolang-migrateを使用しています。
up.sqlとdown.sqlを作成してください
```
---

## 参考文献
[hQiita: golang-migrateの導入](https://qiita.com/shuyaeer/items/3f8a93cac6dcc4323f5f)

