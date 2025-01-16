package database

import (
	"database/sql"
	"errors"
	"time"

	"github.com/okppop/url-shortener/models"
	"github.com/okppop/url-shortener/utils"
)

var (
	ErrGenerateShortPathOver5Times = errors.New("generate short path over 5 times")
)

type Postgresql struct {
	db *sql.DB
}

func (p *Postgresql) CreateUrl(requestData models.UrlApiPOSTRequest) (*models.UrlApiPOSTResponse, error) {
	const sqlIsShortPathDuplicate string = "SELECT true FROM url WHERE short_path = $1;"
	const sqlInsert string = "INSERT INTO url (original_url, short_path, expired_at) VALUES ($1, $2, $3);"
	// const sqlInsert string = "insert into url (original_url, short_path, expired_at) values ($1, $2, current_timestamp + interval '1 hour' * $3);"

	var shortPath string
	var isDuplicate bool = true // assume duplicate at first

	for i := 0; i < 5; i++ {
		shortPath = utils.GenerateShortPath()
		err := p.db.QueryRow(sqlIsShortPathDuplicate, shortPath).Scan(&isDuplicate)
		if err != nil {
			if err == sql.ErrNoRows {
				isDuplicate = false
			}

			return nil, err
		}

		if !isDuplicate {
			break
		}
	}

	if isDuplicate {
		return nil, ErrGenerateShortPathOver5Times
	}

	expiredAt := time.Now().Add(time.Duration(requestData.ExpireHours) * time.Hour)

	_, err := p.db.Exec(sqlInsert, requestData.OriginalUrl, shortPath, expiredAt)
	if err != nil {
		return nil, err
	}

	return &models.UrlApiPOSTResponse{
		ShortPath: shortPath,
		ExpiredAt: expiredAt,
	}, nil
}
