package main

func main() {

}

// func fileHandler() {
// 	file, err := os.Open(fileName)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()
// 	data := make([]byte, 100)
// 	for {
// 		count, err := file.Read(data)
// 		if err != nil {
// 			return err
// 		}
// 		if count == 0 {
// 			return nil
// 		}
// 		process(data[:count])
// 	}
// }
