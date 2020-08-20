package handler

import (
	"fmt"
	"net/http"
	"strconv"

	//_ "github.com/go-sql-driver/mysql"

	"github.com/ArtyDev57/lxd-warehouse/api/model"
	"github.com/ArtyDev57/lxd-warehouse/api/util"
	"github.com/labstack/echo"
)

// CreateProductBrand function
func CreateProductBrand(c echo.Context) error {
	name := c.FormValue("name")

	db, err := util.ConnectDB()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "DB Connection Error")
	}
	defer db.Close()

	stmt := "INSERT INTO tbl_product_brand VALUES(DEFAULT, ?)"

	result, err := db.Exec(stmt, name)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "DB Execution Error")
	}

	id, err := result.LastInsertId()
	fmt.Println("Inserted ID: ", id)

	return c.String(http.StatusCreated, "Successful")
}

// UpdateProductBrand function
func UpdateProductBrand(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	name := c.FormValue("name")
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Input Error")
	}

	db, err := util.ConnectDB()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "DB Connection Error")
	}
	defer db.Close()

	stmt := "UPDATE tbl_product_brand SET brand_name = ? WHERE id = ?"
	result, err := db.Exec(stmt, name, id)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "DB Execution Error")
	}

	rows, err := result.RowsAffected()
	fmt.Println("Rows Updated: ", rows)

	if rows == 0 {
		return c.String(http.StatusOK, "Nothing Updated")
	}

	return c.String(http.StatusOK, "Successful")
}

// GetProductBrands fuction
func GetProductBrands(c echo.Context) error {
	db, err := util.ConnectDB()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "DB Connection Error")
	}
	defer db.Close()

	var (
		id   int
		name string
	)

	query := "SELECT id, brand_name from tbl_product_brand"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "DB Execution Error")
	}
	defer rows.Close()

	result := make([]model.ProductBrand, 0)

	for rows.Next() {
		if err := rows.Scan(&id, &name); err != nil {
			fmt.Println(err)
			break
		}
		result = append(result, model.ProductBrand{ID: id, Name: name})
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Rows Error")
	}

	return c.JSON(http.StatusOK, result)
}

// CreateProductType function
func CreateProductType(c echo.Context) error {
	name := c.FormValue("name")

	db, err := util.ConnectDB()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "DB Connection Error")
	}
	defer db.Close()

	stmt := "INSERT INTO tbl_product_type VALUES(DEFAULT, ?)"
	result, err := db.Exec(stmt, name)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "DB Execution Error")
	}

	id, err := result.LastInsertId()
	fmt.Println("Inserted ID: ", id)

	return c.String(http.StatusCreated, "Successful")
}

// UpdateProductType function
func UpdateProductType(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	name := c.FormValue("name")
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Invalid Input")
	}

	db, err := util.ConnectDB()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "DB Connection Error")
	}
	defer db.Close()

	stmt := "UPDATE tbl_product_type SET type_name = ? WHERE id = ?"
	result, err := db.Exec(stmt, name, id)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "DB Execution Error")
	}

	rows, err := result.RowsAffected()
	fmt.Println("Rows Affected: ", rows)

	if rows == 0 {
		return c.String(http.StatusOK, "Nothing Updated")
	}

	return c.String(http.StatusOK, "Successful")
}

// GetProductTypes function
func GetProductTypes(c echo.Context) error {
	db, err := util.ConnectDB()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "DB Connection Error")
	}
	defer db.Close()

	var (
		id   int
		name string
	)
	query := "SELECT id, type_name from tbl_product_type"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "DB Execution Error")
	}
	defer rows.Close()

	result := make([]model.ProductType, 0)

	for rows.Next() {
		if err := rows.Scan(&id, &name); err != nil {
			fmt.Println(err)
			break
		}
		result = append(result, model.ProductType{ID: id, Name: name})
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Rows Error")
	}

	return c.JSON(http.StatusOK, result)
}

// CreateSupplier function
func CreateSupplier(c echo.Context) error {
	name := c.FormValue("name")
	address := c.FormValue("address")

	db, err := util.ConnectDB()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "DB Coonection Error")
	}
	defer db.Close()

	stmt := "INSERT INTO tbl_suppliers VALUES(DEFAULT, ?, ?)"
	result, err := db.Exec(stmt, name, address)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "DB Execution Error")
	}

	id, err := result.LastInsertId()
	fmt.Println("Inserted ID: ", id)

	return c.String(http.StatusCreated, "Successful")
}

// UpdateSupplier function
func UpdateSupplier(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	name := c.FormValue("name")
	address := c.FormValue("address")

	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Invalid Input")
	}

	db, err := util.ConnectDB()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "DB Connection Error")
	}
	defer db.Close()

	stmt := "UPDATE tbl_suppliers SET supplier_name = ?, supplier_address = ? WHERE id = ?"
	result, err := db.Exec(stmt, name, address, id)
	if err != nil {
		fmt.Println(err)
		return c.HTML(http.StatusInternalServerError, "DB Execution Error")
	}

	rows, err := result.RowsAffected()
	fmt.Println("Rows affected: ", rows)

	if rows == 0 {
		return c.String(http.StatusOK, "Nothing Updated")
	}

	return c.String(http.StatusOK, "Successful")
}

// GetSuppliers function
func GetSuppliers(c echo.Context) error {
	db, err := util.ConnectDB()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "DB Connection Error")
	}
	defer db.Close()

	var (
		id      int
		name    string
		address string
	)
	query := "SELECT id, supplier_name, supplier_address FROM tbl_suppliers"

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "DB Execution Error")
	}
	defer rows.Close()

	result := make([]model.Supplier, 0)

	for rows.Next() {
		if err := rows.Scan(&id, &name, &address); err != nil {
			fmt.Println(err)
			break
		}
		result = append(result, model.Supplier{
			ID:      id,
			Name:    name,
			Address: address,
		})
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Rows Error")
	}

	return c.JSON(http.StatusOK, result)
}

// CreateDistributor function
func CreateDistributor(c echo.Context) error {
	name := c.FormValue("name")
	address := c.FormValue("address")

	db, err := util.ConnectDB()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "DB Connection Error")
	}
	defer db.Close()

	stmt := "INSERT INTO tbl_distributors VALUES (DEFAULT, ?, ?)"
	result, err := db.Exec(stmt, name, address)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "DB Execution Error")
	}

	id, err := result.LastInsertId()
	fmt.Println("Inserted ID: ", id)

	return c.String(http.StatusCreated, "Successful")
}

// UpdateDistributor function
func UpdateDistributor(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	name := c.FormValue("name")
	address := c.FormValue("address")
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Invalid Input")
	}

	db, err := util.ConnectDB()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "DB Connection Error")
	}
	defer db.Close()

	stmt := "UPDATE tbl_distributors SET distributor_name = ?, distributor_address = ? WHERE id = ?"
	result, err := db.Exec(stmt, name, address, id)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "DB Execution Error")
	}

	rows, err := result.RowsAffected()
	fmt.Println(rows)

	if rows == 0 {
		return c.String(http.StatusOK, "Nothing Update")
	}

	return c.String(http.StatusOK, "Successful")
}

// GetDistributors function
func GetDistributors(c echo.Context) error {
	db, err := util.ConnectDB()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "DB Connection Error")
	}
	defer db.Close()

	var (
		id      int
		name    string
		address string
	)

	query := "SELECT id, distributor_name, distributor_address FROM tbl_distributors"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "DB Execution Error")
	}
	defer rows.Close()

	result := make([]model.Distributor, 0)

	for rows.Next() {
		if err := rows.Scan(&id, &name, &address); err != nil {
			fmt.Println(err)
		}
		result = append(result, model.Distributor{ID: id, Name: name, Address: address})
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Rows Error")
	}

	return c.JSON(http.StatusOK, result)
}
