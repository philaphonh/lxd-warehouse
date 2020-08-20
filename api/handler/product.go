package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/ArtyDev57/lxd-warehouse/api/model"
	"github.com/ArtyDev57/lxd-warehouse/api/util"
	"github.com/labstack/echo"
)

// CreateProduct function
func CreateProduct(c echo.Context) error {
	name := c.FormValue("name")

	brandID, err := strconv.Atoi(c.FormValue("brandId"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Invalid Input")
	}

	typeID, err := strconv.Atoi(c.FormValue("typeId"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Invalid Input")
	}

	description := c.FormValue("description")

	price, err := strconv.ParseFloat(c.FormValue("price"), 64)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Invalid Input")
	}

	file, err := c.FormFile("image")
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Invalid File")
	}

	src, err := file.Open()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error Opening File")
	}
	defer src.Close()

	timestamp := time.Now().Format("2006-01-02T15:04:05")

	filename := "api/upload/" + timestamp + "_" + file.Filename

	dst, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error Creating File")
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error Copying File")
	}

	hostname := c.Request().Host
	if protocol := c.Request().Proto; protocol != "HTTP/2" {
		filename = "http://" + hostname + "/upload/" + timestamp + "_" + file.Filename
	} else {
		filename = "https://" + hostname + "/upload/" + timestamp + "_" + file.Filename
	}

	db, err := util.ConnectDB()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error DB Conection")
	}
	defer db.Close()

	stmt := "INSERT INTO tbl_products VALUES(DEFAULT, ?, ?, ?, ?, ?, DEFAULT, ?)"
	result, err := db.Exec(stmt, name, brandID, typeID, description, price, filename)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error DB Execution")
	}

	id, err := result.LastInsertId()
	fmt.Println("Inserted ID: ", id)

	return c.String(http.StatusCreated, "Successful")
}

// GetAllProducts function
func GetAllProducts(c echo.Context) error {
	db, err := util.ConnectDB()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error DB Connection")
	}
	defer db.Close()

	var (
		id          int
		name        string
		pbrand      string
		ptype       string
		description string
		price       float64
		qty         int
		image       string
	)

	query := `SELECT 
	tbl_products.id,
	tbl_products.product_name,
	tpb.brand_name, tpt.type_name,
	tbl_products.product_description,
	tbl_products.product_price,
	tbl_products.product_qty,
	tbl_products.product_image_url
	from tbl_products 
	INNER JOIN tbl_product_brand tpb on tbl_products.product_brand = tpb.id 
	INNER JOIN tbl_product_type tpt on tbl_products.product_type = tpt.id`
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error DB Execution")
	}
	defer rows.Close()

	result := make([]model.Product, 0)

	for rows.Next() {
		if err := rows.Scan(&id, &name, &pbrand, &ptype, &description, &price, &qty, &image); err != nil {
			fmt.Println(err)
			break
		}
		result = append(result, model.Product{ID: id, Name: name, Brand: pbrand, Type: ptype, Description: description, Price: price, Qty: qty, Image: image})
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Rows Error")
	}

	return c.JSON(http.StatusOK, result)
}

// GetProduct function
func GetProduct(c echo.Context) error {
	pid := c.Param("id")

	db, err := util.ConnectDB()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error DB Connection")
	}
	defer db.Close()

	var (
		id          int
		name        string
		pbrand      int
		ptype       int
		description string
		price       float64
		qty         int
		image       string
	)

	query := `SELECT 
	tbl_products.id,
	tbl_products.product_name,
	tbl_products.product_brand, 
	tbl_products.product_type,
	tbl_products.product_description,
	tbl_products.product_price,
	tbl_products.product_qty,
	tbl_products.product_image_url
	from tbl_products 
	WHERE tbl_products.id = ?`
	row := db.QueryRow(query, pid)
	if err = row.Err(); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error DB Execution")
	}
	if err := row.Scan(&id, &name, &pbrand, &ptype, &description, &price, &qty, &image); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error Scanning Row")
	}

	result := &model.ProductRaw{ID: id, Name: name, BrandID: pbrand, TypeID: ptype, Description: description, Price: price, Qty: qty, Image: image}
	// if Roerr != nil {
	// 	fmt.Println(err)
	// 	return c.String(http.StatusInternalServerError, "Error DB Execution")
	// }
	// defer rows.Close()

	// result := make([]model.ProductRaw, 0)

	// for rows.Next() {
	// 	if err := rows.Scan(&id, &name, &pbrand, &ptype, &description, &price, &qty, &image); err != nil {
	// 		fmt.Println(err)
	// 		break
	// 	}
	// 	result = append(result, model.ProductRaw{ID: id, Name: name, BrandID: pbrand, TypeID: ptype, Description: description, Price: price, Qty: qty, Image: image})
	// }

	// if err := rows.Err(); err != nil {
	// 	fmt.Println(err)
	// 	return c.String(http.StatusInternalServerError, "Rows Error")
	// }

	return c.JSON(http.StatusOK, result)
}

// UpdateProduct function
func UpdateProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Invalid Input")
	}

	name := c.FormValue("name")

	brandID, err := strconv.Atoi(c.FormValue("brandId"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Invalid Input")
	}

	typeID, err := strconv.Atoi(c.FormValue("typeId"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Invalid Input")
	}

	description := c.FormValue("description")

	price, err := strconv.ParseFloat(c.FormValue("price"), 64)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Invalid Input")
	}

	db, err := util.ConnectDB()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error DB Conection")
	}
	defer db.Close()

	file, err := c.FormFile("image")
	if err != nil {
		fmt.Println("No File for Uploading")
	} else {
		src, err := file.Open()
		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "Error Opening File")
		}
		defer src.Close()

		timestamp := time.Now().Format("2006-01-02T15:04:05")

		filename := "api/upload/" + timestamp + "_" + file.Filename

		dst, err := os.Create(filename)
		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "Error Creating File")
		}
		defer dst.Close()

		if _, err := io.Copy(dst, src); err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "Error Copying File")
		}

		hostname := c.Request().Host
		if protocol := c.Request().Proto; protocol != "HTTP/2" {
			filename = "http://" + hostname + "/upload/" + timestamp + "_" + file.Filename
		} else {
			filename = "https://" + hostname + "/upload/" + timestamp + "_" + file.Filename
		}

		stmt := "UPDATE tbl_products SET product_image_url = ? WHERE id = ?"
		result, err := db.Exec(stmt, filename, id)
		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "Error DB Execution")
		}

		rows, err := result.RowsAffected()
		fmt.Println("Rows affected: ", rows)

		if rows == 0 {
			fmt.Println("No File Updated")
		}
	}

	stmt := `UPDATE tbl_products SET 
	product_name = ?, 
	product_brand = ?, 
	product_type = ?, 
	product_description = ?,
	product_price = ?
	WHERE id = ?
	`
	result, err := db.Exec(stmt, name, brandID, typeID, description, price, id)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error DB Execution")
	}

	rows, err := result.RowsAffected()
	fmt.Println("Rows affected: ", rows)

	if rows == 0 {
		return c.String(http.StatusOK, "Nothing Updated")
	}

	return c.String(http.StatusOK, "Successful")
}

// DeleteProduct function
func DeleteProduct(c echo.Context) error {

	return c.String(http.StatusOK, "Successful")
}

// ImportProduct function
func ImportProduct(c echo.Context) error {
	supplier, err := strconv.Atoi(c.FormValue("supplier"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Invalid Input")
	}
	user, err := strconv.Atoi(c.FormValue("user"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Invalid Input")
	}
	timestamp := time.Now().Format("2006-01-02T15:04:05")

	db, err := util.ConnectDB()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error DB Connection")
	}
	defer db.Close()

	stmt := "INSERT INTO tbl_import VALUES(DEFAULT, ?, ?, ?, ?)"
	result, err := db.Exec(stmt, supplier, user, timestamp, "ORDERED")
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error DB Execution")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error Inserting Into DB")
	}
	fmt.Println("Inserted ID: ", id)

	return c.String(http.StatusCreated, strconv.FormatInt(id, 10))
}

// ImportProductDetail function
func ImportProductDetail(c echo.Context) error {
	importID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Invalid Input")
	}

	p := new(model.ImportedProduct)
	if err := c.Bind(p); err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Failed")
	}

	db, err := util.ConnectDB()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error DB Connection")
	}
	defer db.Close()

	stmt := "INSERT INTO tbl_import_detail VALUES(DEFAULT, ?, ?, ?)"
	result, err := db.Exec(stmt, importID, p.ID, p.Quantity)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error DB Execution")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error Inserting Into DB")
	}
	fmt.Println("Inserted ID: ", id)

	return c.String(http.StatusCreated, "Successful")
}

// GetAllImport function
func GetAllImport(c echo.Context) error {
	db, err := util.ConnectDB()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error DB Connection")
	}
	defer db.Close()

	var (
		id       int
		supplier string
		user     string
		time     string
		status   string
	)

	query := `
	SELECT tbl_import.id, 
	ts.supplier_name, 
	tu.username, 
	tbl_import.import_time, 
	tbl_import.import_status 
	FROM tbl_import
    INNER JOIN tbl_suppliers ts ON tbl_import.supplier_id = ts.id
    INNER JOIN tbl_users tu ON tbl_import.user_id = tu.id
	`

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error DB Execution")
	}
	defer rows.Close()

	result := make([]model.Import, 0)

	for rows.Next() {
		if rows.Scan(&id, &supplier, &user, &time, &status); err != nil {
			fmt.Println(err)
			break
		}
		result = append(result,
			model.Import{
				ID:         id,
				Supplier:   supplier,
				User:       user,
				ImportTime: time,
				Status:     status,
			})
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Rows Error")
	}

	return c.JSON(http.StatusOK, result)
}

// GetImportDetail function
func GetImportDetail(c echo.Context) error {
	importID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadGateway, "Invalid Input")
	}

	var (
		id      int
		imID    int
		product string
		qty     int
	)

	db, err := util.ConnectDB()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error DB Connection")
	}
	defer db.Close()

	query := `
	SELECT tbl_import_detail.id,
       tbl_import_detail.import_id,
       tbl_products.product_name,
       tbl_import_detail.product_qty
	FROM tbl_import_detail INNER JOIN tbl_products on tbl_import_detail.product_id = tbl_products.id
	WHERE tbl_import_detail.import_id = ?
	`

	rows, err := db.Query(query, importID)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error DB Execution")
	}

	result := make([]model.ImportDetail, 0)

	for rows.Next() {
		if rows.Scan(&id, &imID, &product, &qty); err != nil {
			fmt.Println(err)
			break
		}
		result = append(result,
			model.ImportDetail{
				ID:       id,
				ImportID: imID,
				Product:  product,
				Qty:      qty,
			})
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Rows Error")
	}

	return c.JSON(http.StatusOK, result)
}

// ConfirmImport function
func ConfirmImport(c echo.Context) error {
	importID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Invalid Input")
	}

	db, err := util.ConnectDB()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error DB Connection")
	}
	defer db.Close()

	stmt := "UPDATE tbl_import SET import_status = 'IMPORTED' WHERE id = ?"
	result, err := db.Exec(stmt, importID)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error DB Execution")
	}

	row, err := result.RowsAffected()
	fmt.Println("Rows affected: ", row)

	if row == 0 {
		return c.String(http.StatusOK, "Nothing Updated")
	}

	var (
		pid  int
		pqty int
	)

	query := `
	SELECT tbl_import_detail.product_id, 
	tbl_import_detail.product_qty 
	FROM tbl_import_detail WHERE tbl_import_detail.import_id = ?`

	detailRows, err := db.Query(query, importID)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error DB Execution")
	}
	defer detailRows.Close()

	detailResult := make([]model.ImportedProductDetail, 0)

	for detailRows.Next() {
		if err := detailRows.Scan(&pid, &pqty); err != nil {
			fmt.Println(err)
			break
		}
		detailResult = append(detailResult, model.ImportedProductDetail{ProductID: pid, ProductQty: pqty})
	}

	if err := detailRows.Err(); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Rows Error")
	}

	stmt = "UPDATE tbl_products SET tbl_products.product_qty = tbl_products.product_qty + ? WHERE id = ?"

	for _, item := range detailResult {
		result, err := db.Exec(stmt, item.ProductQty, item.ProductID)
		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "Error DB Execution")
		}
		row, err := result.RowsAffected()
		fmt.Println("Rows affected: ", row)
	}

	return c.String(http.StatusOK, "Successful")
}

// ExportProduct function
func ExportProduct(c echo.Context) error {
	distributor, err := strconv.Atoi(c.FormValue("distributor"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Invalid Input")
	}
	user, err := strconv.Atoi(c.FormValue("user"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Invalid Input")
	}
	timestamp := time.Now().Format("2006-01-02T15:04:05")

	db, err := util.ConnectDB()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error DB Connection")
	}
	defer db.Close()

	stmt := "INSERT INTO tbl_export VALUES(DEFAULT, ?, ?, ?)"
	result, err := db.Exec(stmt, distributor, user, timestamp)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error DB Execution")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error Inserting Into DB")
	}
	fmt.Println("Inserted ID: ", id)

	return c.String(http.StatusCreated, strconv.FormatInt(id, 10))
}

// ExportProductDetail function
func ExportProductDetail(c echo.Context) error {
	exportID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Invalid Input")
	}

	p := new(model.ExportedProduct)
	if err := c.Bind(p); err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Failed")
	}

	db, err := util.ConnectDB()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error DB Connection")
	}
	defer db.Close()

	stmt := "INSERT INTO tbl_export_detail VALUES(DEFAULT, ?, ?, ?)"
	result, err := db.Exec(stmt, exportID, p.ID, p.Quantity)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error DB Execution")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error Inserting Into DB")
	}
	fmt.Println("Inserted ID: ", id)

	return c.String(http.StatusCreated, "Successful")
}

// GetAllExport fucntion
func GetAllExport(c echo.Context) error {
	db, err := util.ConnectDB()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error DB Connection")
	}
	defer db.Close()

	var (
		id          int
		distributor string
		user        string
		time        string
	)

	query := `
	SELECT tbl_export.id, 
	td.distributor_name, 
	tu.username, 
	tbl_export.export_time
	FROM tbl_export
    INNER JOIN tbl_distributors td ON tbl_export.distributor_id = td.id
    INNER JOIN tbl_users tu ON tbl_export.user_id = tu.id
	`

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error DB Execution")
	}
	defer rows.Close()

	result := make([]model.Export, 0)

	for rows.Next() {
		if rows.Scan(&id, &distributor, &user, &time); err != nil {
			fmt.Println(err)
			break
		}
		result = append(result,
			model.Export{
				ID:          id,
				Distributor: distributor,
				User:        user,
				ExportTime:  time,
			})
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Rows Error")
	}

	return c.JSON(http.StatusOK, result)
}

// GetExportDetail function
func GetExportDetail(c echo.Context) error {
	exportID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadGateway, "Invalid Input")
	}

	var (
		id      int
		exID    int
		product string
		qty     int
	)

	db, err := util.ConnectDB()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error DB Connection")
	}
	defer db.Close()

	query := `
	SELECT tbl_export_detail.id,
       tbl_export_detail.export_id,
       tbl_products.product_name,
       tbl_export_detail.product_qty
	FROM tbl_export_detail INNER JOIN tbl_products on tbl_export_detail.product_id = tbl_products.id
	WHERE tbl_export_detail.export_id = ?
	`

	rows, err := db.Query(query, exportID)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error DB Execution")
	}

	result := make([]model.ExportDetail, 0)

	for rows.Next() {
		if rows.Scan(&id, &exID, &product, &qty); err != nil {
			fmt.Println(err)
			break
		}
		result = append(result,
			model.ExportDetail{
				ID:       id,
				ExportID: exID,
				Product:  product,
				Qty:      qty,
			})
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Rows Error")
	}

	return c.JSON(http.StatusOK, result)
}

// ConfirmExport function
func ConfirmExport(c echo.Context) error {
	exportID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Invalid Input")
	}

	db, err := util.ConnectDB()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error DB Connection")
	}
	defer db.Close()

	var (
		pid  int
		pqty int
	)

	query := `
	SELECT tbl_export_detail.product_id, 
	tbl_export_detail.product_qty 
	FROM tbl_export_detail WHERE tbl_export_detail.export_id = ?`

	detailRows, err := db.Query(query, exportID)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error DB Execution")
	}
	defer detailRows.Close()

	detailResult := make([]model.ExportedProductDetail, 0)

	for detailRows.Next() {
		if err := detailRows.Scan(&pid, &pqty); err != nil {
			fmt.Println(err)
			break
		}
		detailResult = append(detailResult, model.ExportedProductDetail{ProductID: pid, ProductQty: pqty})
	}

	if err := detailRows.Err(); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Rows Error")
	}

	stmt := "UPDATE tbl_products SET tbl_products.product_qty = tbl_products.product_qty - ? WHERE id = ?"

	for _, item := range detailResult {
		result, err := db.Exec(stmt, item.ProductQty, item.ProductID)
		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "Error DB Execution")
		}
		row, err := result.RowsAffected()
		fmt.Println("Rows affected: ", row)
	}

	return c.String(http.StatusOK, "Successful")
}
