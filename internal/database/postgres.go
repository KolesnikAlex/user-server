package database

import (
	"github.com/jmoiron/sqlx"
	"user-server/app/service"
)

func NewSQLRepo(db *sqlx.DB) *SQLRepository {
	return &SQLRepository{
		//logger: l,
		db: db,
	}
}

type SQLRepository struct {
	//logger *zap.SugaredLogger
	db *sqlx.DB
}

func (rep SQLRepository) GetUser(id int64) (result service.User, err error) {
	err = rep.db.Get(&result, getUserQuery(), id)
	if err != nil {
		//log.Error().Err(err).Msg("err get coinsID from postgres")
		return service.User{}, err
	}
	return result, err
}

func (rep SQLRepository) AddUser(user service.User) error {
	_, err := rep.db.Exec(getAddUserQuery(), user.ID, user.Name, user.Login, user.Password)
	if err != nil {
		//log.Error().Err(err).Msg("err add user to postgres")
		return err
	}
	return err
}

func (rep SQLRepository) RemoveUser(id int64) error {
	_, err := rep.db.Exec(getRemoveUserQuery(), id)
	if err != nil {
		//log.Error().Err(err).Msg("err del user from postgres")
		return err
	}
	return err
}

func (rep SQLRepository) UpdateUser(user service.User) (err error) {
	_, err = rep.db.Exec(getUpdateUserQuery(), user.Name, user.Login, user.Password, user.ID)
	if err != nil {
		//log.Error().Err(err).Msg("err UpdateUser in postgres")
		return  err
	}
	return err
}

func (rep SQLRepository) Close() error {
	err := rep.db.Close()
	if err != nil {
		//log.Error().Err(err).Msg("err close postgres")
		return err
	}
	return err
}
