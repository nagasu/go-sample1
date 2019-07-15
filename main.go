package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

type EMPLOYEE struct {
	ID     string
	NUMBER string
}

func helloworld() {
	defer func() {
		fmt.Println("End")
		err := recover()
		if err != nil {
			fmt.Println("Recover!:", err)
		}
	}()

	fmt.Println("Start")
	panic("Panic!")
}

func main() {
	// test()
	// // helloworld()
	// val := "test"
	// fmt.Printf("%v, %T", val, val)

	// f, error := os.Open("filename.ext")
	// if error != nil {
	// 	return
	// }
	// defer f.Close()
	dbInsert2()
}

func insert(startIndex int, endIndex int, st *sql.Stmt, dt1 time.Time, wg *sync.WaitGroup) {
	defer wg.Done()
	s := ""
	for i := startIndex; i < endIndex; i++ {
		s += "a"
		// _, err := st.Exec(i, "AAA", "BBB")

		// if err != nil {
		// 	fmt.Println(err)
		// }

		// if i % 10000 == 0 {
		// 	dt2 := time.Now()
		// 	fmt.Println(dt2.Sub(dt1).Seconds())
		// 	fmt.Println(i)
		// }
	}
}

func dbInsert2() {

	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=mina-lib sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	// DELETE
	_, err = db.Exec("DELETE FROM test")
	if err != nil {
		fmt.Println(err)
	}

	dt1 := time.Now()

	var m2 = make([]byte, 0, 100000)

	// INSERT
	const totalNumber = 100000
	m2 = append(m2, "INSERT INTO test VALUES(0, 'AAA', 'BBB')"...)
	for i := 1; i < totalNumber; i++ {
		// m2 = append(m2, ",("+strconv.Itoa(i)+", 'AAA', 'BBB')"...)
		s := fmt.Sprintf(",(%d, 'AAA', 'BBB')", strconv.Itoa(i))
		m2 = append(m2, s...)
	}
	// fmt.Println(string(m2))
	// dt2 := time.Now()
	_, err = db.Exec(string(m2))
	if err != nil {
		fmt.Println(err)
	}

	dt3 := time.Now()
	fmt.Println("complete ", dt3.Sub(dt1).Seconds())
	// fmt.Println("complete ", dt3.Sub(dt2).Seconds())
}

func main2() {
	// db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=mina-lib sslmode=disable")
	// defer db.Close()

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// DELETE
	// _, err = db.Exec("DELETE FROM test")

	dt1 := time.Now()

	// INSERT
	// wg := &sync.WaitGroup{}
	// st, err := db.Prepare("INSERT INTO test VALUES($1, $2, $3)")

	// const totalNumber = 1000000
	// const perInsert = 1000000
	// startIndex := 0
	// for i := 0; i < totalNumber/perInsert; i++ {
	// 	wg.Add(1)
	// 	fmt.Println(startIndex, startIndex+perInsert)
	// 	go insert(startIndex, startIndex+perInsert, st, dt1, wg)
	// 	startIndex += perInsert
	// }
	// s := ""
	num := 0
	for i := 0; i < 1000000; i++ {
		// s += strconv.Itoa(i)
		num = i
	}
	fmt.Println(num)

	// wg.Wait()
	dt3 := time.Now()
	fmt.Println("complete ", dt3.Sub(dt1).Seconds())
}
