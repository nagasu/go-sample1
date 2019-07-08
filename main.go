package main

import (
	"sync"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

type EMPLOYEE struct {
	ID     string
	NUMBER string
}

func insert(start int, st *sql.Stmt, dt1 time.Time, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := start; i < start + 10000; i++ {
		_, err := st.Exec(i, "AAA", "BBB")

		if err != nil {
			fmt.Println(err)
		}

		if i % 10000 == 0 {
			dt2 := time.Now()
			fmt.Println(dt2.Sub(dt1).Seconds())
			fmt.Println(i)
		}
	}
}

func main() {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=mina-lib sslmode=disable")
	defer db.Close()

	if err != nil {
		fmt.Println(err)
	}

	// DELETE
	_, err = db.Exec("DELETE FROM test")

	dt1 := time.Now()

	// INSERT
	wg := &sync.WaitGroup{}
	st, err := db.Prepare("INSERT INTO test VALUES($1, $2, $3)")

	row := 0
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go insert(row, st, dt1, wg)
		row += 10000
	}

	wg.Wait()
	dt3 := time.Now()
	fmt.Println("complete ", dt3.Sub(dt1).Seconds())
}
