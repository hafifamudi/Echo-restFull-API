package models

import (
	"net/http"

	"github.com/hafif/echoFramework/db"
)

type Buku struct {
	ID        int    `json:"id"`
	NamaBuku  string `json:"nama_buku"`
	HargaBuku string `json:"harga_buku"`
	JenisBuku string `json:"jenis_buku"`
}

func FetchAllData() (Response, error) {
	var obj Buku
	var arrObj []Buku
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM bukuFantasy"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.ID, &obj.NamaBuku, &obj.HargaBuku, &obj.JenisBuku)

		if err != nil {
			return res, err
		}

		arrObj = append(arrObj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrObj

	return res, nil
}

func StoreData(nama string, harga string, jenis string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT bukufantasy (nama_buku,harga_buku,jenis_buku) VALUES (?,?,?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nama, harga, jenis)
	if err != nil {
		return res, err
	}

	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"lastInsertedID": lastInsertedID,
	}

	return res, nil

}

func UpdateData(ID int, nama string, harga string, jenis string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE bukufantasy SET nama_buku=?,harga_buku=?,jenis_buku=? WHERE id=?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nama, harga, jenis, ID)
	if err != nil {
		return res, err
	}

	rowAffect, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"row_affected": rowAffect,
	}

	return res, nil
}

func DeleteData(ID int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM bukufantasy WHERE id=?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(ID)
	if err != nil {
		return res, err
	}

	rowAffect, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"row_affected": rowAffect,
	}

	return res, nil
}
