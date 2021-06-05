package user

const UserTableName string = "user"

type UserModel struct {
	ID       int `gorm:"type:int;primaryKey;autoIncrement"`
	Name     string
	Nik      string
	Phone    string
	Saldo    int `gorm:"type:int;not null;default:0"`
	Password string
	Username string
	Version  int
}

func (M *UserModel) TableName() string {
	return UserTableName
}
