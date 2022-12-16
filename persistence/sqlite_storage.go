package persistence

import (
	"database/sql"
	"log"

	"github.com/jeromelesaux/fizzbuzz/model"
	_ "github.com/mattn/go-sqlite3"
)

var (
	sqliteDBName = "fizzbuzz.db"
)

func initialiseDB() {
	db, err := sql.Open("sqlite3", sqliteDBName)
	if err != nil {
		log.Printf("Error while opening sqlite database %s with error :%v\n", sqliteDBName, err)
		return
	}
	defer db.Close()
	stmt, err := db.Prepare("create table if not exists stats(int1 integer,int2 integer, lim integer, str1  varchar(512) null , str2 varchar(512) null, hits integer);")
	if err != nil {
		log.Printf("Error while preparing statement for sqlite database %s with error :%v\n", sqliteDBName, err)
		return
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Printf("Error while executing statement sqlite database %s with error :%v\n", sqliteDBName, err)
		return
	}
	stmt.Close()
}

func init() {
	initialiseDB()
}

func AddInDB(p model.Parameters) error {
	db, err := sql.Open("sqlite3", sqliteDBName)
	if err != nil {
		log.Printf("Error while opening sqlite database %s with error :%v\n", sqliteDBName, err)
		return err
	}
	defer db.Close()
	rows, err := db.Query("select hits from stats where int1=? and int2=? and str1=? and str2 =? and lim=?;",
		p.Int1,
		p.Int2,
		p.Str1,
		p.Str2,
		p.Limit,
	)
	if err != nil {
		log.Printf("Error while executing query  sqlite database %s with error :%v\n", sqliteDBName, err)
		return err
	}
	var count sql.NullInt64
	if rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			return err
		}
	}
	rows.Close()
	var q string
	hits := count.Int64

	if hits == 0 {
		q = "insert into stats(hits,int1,int2,str1,str2,lim) values(?,?,?,?,?,?);"
	} else {
		q = "update stats set hits = ? where int1=? and int2=? and str1=? and str2 =? and lim=?;"
	}
	hits++
	tx, err := db.Begin()
	if err != nil {
		log.Printf("Error while preparing query  sqlite database %s with error :%v\n", sqliteDBName, err)
		return err
	}
	stmt, err := db.Prepare(q)
	if err != nil {
		log.Printf("Error while preparing query  sqlite database %s with error :%v\n", sqliteDBName, err)
		return err
	}

	affected, err := tx.Stmt(stmt).Exec(hits, p.Int1, p.Int2, p.Str1, p.Str2, p.Limit)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
	} else {
		err := tx.Commit()
		if err != nil {
			return err
		}
	}
	rs, err := affected.RowsAffected()
	log.Printf("affected %d and error %v\n", rs, err)
	return err
}

func DeleteStatsDB() error {
	db, err := sql.Open("sqlite3", sqliteDBName)
	if err != nil {
		log.Printf("Error while opening sqlite database %s with error :%v\n", sqliteDBName, err)
		return err
	}
	defer db.Close()
	stmt, err := db.Prepare("delete from stats;")
	if err != nil {
		log.Printf("Error while preparing statement sqlite database %s with error :%v\n", sqliteDBName, err)
		return err
	}
	tx, err := db.Begin()
	if err != nil {
		log.Printf("Error while getting transaction sqlite database %s with error :%v\n", sqliteDBName, err)
		return err
	}

	_, err = tx.Stmt(stmt).Exec()
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
	} else {
		err := tx.Commit()
		if err != nil {
			return err
		}
	}
	return err
}

func GetMostFrequentDB() (model.Parameters, error) {
	var err error
	p := model.Parameters{}
	db, err := sql.Open("sqlite3", sqliteDBName)
	if err != nil {
		log.Printf("Error while opening sqlite database %s with error :%v\n", sqliteDBName, err)
		return p, err
	}
	defer db.Close()
	rows, err := db.Query("select int1, int2, lim, str1, str2, hits from stats order by hits desc limit 1;")
	if err != nil {
		log.Printf("Error while opening sqlite database %s with error :%v\n", sqliteDBName, err)
		return p, err
	}
	var int1, int2, limit, hits sql.NullInt64
	var str1, str2 sql.NullString

	if rows.Next() {
		err = rows.Scan(&int1, &int2, &limit, &str1, &str2, &hits)
	}
	rows.Close()
	if err != nil {
		log.Printf("Error while opening sqlite database %s with error :%v\n", sqliteDBName, err)
		return p, err
	}
	p.Hits = hits.Int64
	p.Int1 = int1.Int64
	p.Int2 = int2.Int64
	p.Limit = limit.Int64
	p.Str1 = str1.String
	p.Str2 = str2.String

	return p, nil
}
