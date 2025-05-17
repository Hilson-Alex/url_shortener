package shortener

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"math/rand/v2"
	"time"

	"github.com/Hilson-Alex/url_shortener/base62"
	"github.com/Hilson-Alex/url_shortener/connection"
	"github.com/roylee0704/gron"
)

type URLConnection struct {
	db *sql.DB
}

type rowScanner interface {
	Scan(dest ...any) error
}

func URLRepository() *URLConnection {
	return &URLConnection{db: connection.GetDB()}
}

type ShortURL struct {
	Key         string `json:"key"`
	OriginalUrl string `json:"originalUrl"`
	ExpireDate  int64  `json:"expireDate"`
	ShortUrl    string `json:"shortUrl"`
}

func init() {
	cron := gron.New()
	cron.AddFunc(gron.Every(24*time.Hour), func() {
		URLRepository().clearExpiredUrls()
	})
	cron.Start()
}

func (repo *URLConnection) listUrls() ([]*ShortURL, error) {
	rows, err := repo.db.Query(`SELECT key, original_url, expire_date FROM short_urls`)
	if err != nil {
		return nil, err
	}
	return parseAll(rows)
}

func (repo *URLConnection) findByKey(key string) (*ShortURL, error) {
	row := repo.db.QueryRow(`
		SELECT key, original_url, expire_date FROM short_urls
		WHERE key = ?
	`, key)
	return parseEntry(row)
}

func (repo *URLConnection) findByUrl(url string) (*ShortURL, error) {
	row := repo.db.QueryRow(`
		SELECT key, original_url, expire_date FROM short_urls
		WHERE original_url = ?
	`, url)
	return parseEntry(row)
}

func (repo *URLConnection) saveURL(entry *ShortURL) error {
	entry.Key = base62.Encode(time.Now().Unix()) + base62.Encode(rand.Int64N(124))
	entry.normalizeExpireDate()
	fmt.Println(entry.ExpireDate)
	var processedDate = entry.ExpireDate
	if saved, err := repo.findByUrl(entry.OriginalUrl); err == nil {
		entry.Key = saved.Key
		if saved.ExpireDate > processedDate {
			entry.ExpireDate = saved.ExpireDate
		}
	}
	_, err := repo.db.Exec(
		`INSERT OR REPLACE INTO short_urls 
			(key, original_url, expire_date) 
		VALUES (?, ?, ?)`,
		entry.Key, entry.OriginalUrl, entry.ExpireDate,
	)
	if err != nil {
		return err
	}
	fmt.Println(processedDate)

	entry.ExpireDate = processedDate
	return nil
}

func (repo *URLConnection) clearExpiredUrls() {
	log.Println("Cleaning Expired URLS")
	res, err := repo.db.Exec(`DELETE FROM short_urls WHERE expire_date <= ?`, time.Now().Unix())
	if err != nil {
		log.Println("ERROR, Cleaning failed:", err.Error())
		return
	}
	rows, _ := res.RowsAffected()
	log.Printf("Deleted %v rows \n", rows)
}

func (dto *ShortURL) GenShortUrl(host string) {
	dto.ShortUrl = host + dto.Key
}

func (dto *ShortURL) normalizeExpireDate() {
	var days = dto.ExpireDate
	dto.ExpireDate = time.Now().
		AddDate(0, 0, int(days)+1).
		Truncate(24 * time.Hour).
		Unix()
}

func parseEntry(entry rowScanner) (*ShortURL, error) {
	var result = &ShortURL{}
	var err = entry.Scan(&result.Key, &result.OriginalUrl, &result.ExpireDate)
	return result, err
}

func parseAll(rows *sql.Rows) ([]*ShortURL, error) {
	var urlList = []*ShortURL{}
	for rows.Next() {
		var shortUrl, err = parseEntry(rows)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return urlList, nil
			}
			return urlList, err
		}
		urlList = append(urlList, shortUrl)
	}
	return urlList, nil
}
