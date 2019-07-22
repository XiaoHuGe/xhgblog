package models

import (
	"strconv"
	"time"
)

type Archive struct {
	ArchiveDate time.Time //month
	Year        int       `json:"year"`
	Month       int       `json:"month"`
	Total       int       `json:"total"`
}

func GetArchive() ([]*Archive, error) {
	var archives []*Archive
	//db.Table("xhgblog_article").Select("DATE_FORMAT(created_at,'%Y-%m') as date, count(*) as total").Group("date").Order("date desc").Scan(&archives)
	rows, err := db.Table("xhgblog_article").Select("DATE_FORMAT(created_at,'%Y-%m') as month, count(*) as num").Group("month").Order("month desc").Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var archive Archive
		var month string
		var num string
		rows.Scan(&month, &num)
		archive.ArchiveDate, _ = time.Parse("2006-01", month)
		archive.Year = archive.ArchiveDate.Year()
		archive.Month = int(archive.ArchiveDate.Month())

		archive.Total, _ = strconv.Atoi(num)
		archives = append(archives, &archive)
	}
	return archives, nil
}
