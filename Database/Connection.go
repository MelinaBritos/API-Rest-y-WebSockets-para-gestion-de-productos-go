package Database

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Conexiondb() *sql.DB {
	//conexion a PostgreSQL
	DSN, err := ObtenerDSN()
	if err != nil {
		log.Fatal(err.Error())
	}

	db, err := sql.Open("postgres", DSN)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Exec("SET TIME ZONE 'America/Argentina/Buenos_Aires';")
	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}

func CrearTablas(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS Products(id SERIAL PRIMARY KEY, name varchar(100) NOT NULL, description varchar(500),
	 	price NUMERIC(10, 2) NOT NULL, stock int NOT NULL, created_at TIMESTAMPTZ, updated_at TIMESTAMPTZ);

		CREATE TABLE IF NOT EXISTS Categories(id SERIAL PRIMARY KEY, name varchar(100) NOT NULL, description varchar(500),
	 	created_at TIMESTAMPTZ, updated_at TIMESTAMPTZ);

		CREATE TABLE IF NOT EXISTS Product_Category(product_id INTEGER NOT NULL, category_id INTEGER NOT NULL, PRIMARY KEY (product_id, category_id), CONSTRAINT fk_product
        FOREIGN KEY (product_id)  REFERENCES products (id)  ON DELETE CASCADE, CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES categories (id)  ON DELETE CASCADE);
		
		CREATE TABLE IF NOT EXISTS Product_History(id SERIAL PRIMARY KEY, product_id INTEGER NOT NULL, price NUMERIC(10, 2) NOT NULL, stock int NOT NULL, changed_at TIMESTAMPTZ, CONSTRAINT fk_product_history
        FOREIGN KEY (product_id)  REFERENCES products (id) ON DELETE CASCADE);`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func ObtenerDSN() (string, error) {

	err := godotenv.Load(".env.example")
	if err != nil {
		return os.Getenv("DSN"), err
	}
	return os.Getenv("DSN"), nil

}
