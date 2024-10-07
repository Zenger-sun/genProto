package model

import (
	"database/sql"
	"errors"
	"fmt"
	"genProto/server"
	"log/slog"
)

type DataPlayer struct {
	Id        uint32
	UserNanme string
	Avatar	  uint32
}

func (d *DataPlayer) Load(db *sql.DB, id uint32) (*DataPlayer, error) {
	rows, e := db.Query(fmt.Sprintf("SELECT id, user_name, avatar FROM player WHERE id in (%d)", id))
	if e != nil {
		return nil, e
	}
	defer rows.Close()

	for rows.Next() {
		var player DataPlayer
		if e := rows.Scan(&player.Id, &player.UserNanme, &player.Avatar); e != nil {
			slog.Warn("rows.Scan: %v\n", e)
			return nil, e
		}

		return &player, nil
	}

	return nil, errors.New("Not Exist!")
}

func (d *DataPlayer) Update(db *sql.DB, option server.DbOption) (err error) {
	var sql string

	switch option {
	case server.INSERT:
		sql = fmt.Sprintf("INSERT INTO player(`user_name`, `avatar`) VALUES('%s', %d)", d.UserNanme, d.Avatar)
		result, err := db.Exec(sql)
		if err != nil {
			return
		}

		id, err := result.LastInsertId()
		if err != nil {
			return
		}

		d.Id = uint32(id)

	case server.UPDATE:
		sql = fmt.Sprintf("UPDATA player SET `user_name` = '%s', `avatar` = %d WHERE id = %d", d.UserNanme, d.Avatar, d.Id)
		_, err = db.Exec(sql)
	case server.DELETE:
		sql = fmt.Sprintf("DELETE FROM player WHERE id = %d", d.Id)
		_, err = db.Exec(sql)
	}

	return
}