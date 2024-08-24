package storage

import "context"

func (s *sqlModel) GetAmountFollow(ctx context.Context, column string, userId int) (int, error) {
	var amount int
	if err := s.db.Table("follow").Where(column+"=?", userId).Select("COUNT(" + column + ") AS amount").Group(column).Scan(&amount).Error; err != nil {
		return 0, err
	}
	return amount, nil
}
