package entity

// Product はテーブルのモデル
type Product struct {
	ID    int    `gorm:"primary_key;not null"       json:"id"`
	Name  string `gorm:"type:varchar(200);not null" json:"name"`
	Memo  string `gorm:"type:varchar(400)"          json:"memo"`
	State int    `gorm:"not null"                   json:"state"`
}

// 商品の購入状態を定義
const (
	NotPurchased = 0 // 未購入
	Purchased    = 1 // 購入済
)

func ChangeState(currentState int) int {
	changeState := NotPurchased

	if currentState == NotPurchased {
		changeState = Purchased
	}

	return changeState
}
