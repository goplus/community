package core

import (
	"context"
	"database/sql"
	"log"
	"strconv"
	"time"
)

type User struct {
	ID       string
	Name     string
	Password string
	Avatar   string
	Ctime    time.Time
	Mtime    time.Time
}

// GetUser return author
func (p *Community) GetUser(ctx context.Context, id string) (user *User, err error) {
	user = &User{}
	sqlStr := "select id,name,avatar from users where id=?"
	err = p.db.QueryRow(sqlStr, id).Scan(&user.ID, &user.Name, &user.Avatar)
	if err == sql.ErrNoRows {
		log.Println("not found the author")
		return user, ErrNotExist
	} else if err != nil {
		return user, err
	}
	return
}

// PutUser adds new user (ID == "") or edits an existing user (ID != "").
func (p *Community) PutUser(ctx context.Context, user *User) (id string, err error) {
	// new user
	if user.ID == "" {
		sqlStr := "insert into users (name, password, avatar, ctime, mtime) values (?, ?, ?, ?, ?)"
		res, err := p.db.Exec(sqlStr, &user.Name, &user.Password, &user.Avatar, time.Now(), time.Now())
		if err != nil {
			return "", err
		}
		idInt, err := res.LastInsertId()
		return strconv.FormatInt(idInt, 10), nil
	}
	// edit user
	sqlStr := "update users set name=?, avatar=?, mtime=? where id=?"
	_, err = p.db.Exec(sqlStr, &user.Name, &user.Avatar, time.Now(), &user.ID)
	return user.ID, err
}

// DeleteUser delete the user.
func (p *Community) DeleteUser(ctx context.Context, id string) (err error) {
	// begin Transaction
	tx, err := p.db.Begin()
	if err != nil {
		return
	}

	// delete user
	sqlStr := "delete from users where id=?"
	_, err = p.db.Exec(sqlStr, id)
	if err != nil {
		tx.Rollback()
	}

	// delete articles
	err = p.DeleteArticles(ctx, id)
	if err != nil {
		tx.Rollback()
	}

	err = tx.Commit()
	return
}