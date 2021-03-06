package main

import (
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
}

// func DoSomeInserts(ctx context.Context, db *sql.DB, value1, value2 string)
//                   (err error) {
//     tx, err := db.BeginTx(ctx, nil)
//     if err != nil {
//         return err
//     }
//     defer func() {
//         if err == nil {
//             err = tx.Commit()
//         }
//         if err != nil {
//             tx.Rollback()
//         }
//     }()
//     _, err = tx.ExecContext(ctx, "INSERT INTO FOO (val) values $1", value1)
//     if err != nil {
//         return err
//     }
//     // use tx to do more database inserts here
//     return nil
// }
