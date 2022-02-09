package main

type Status int

const (
	InvalidLogin Status = iota + 1
	NotFound
)

type StatusErr struct {
	Status  Status
	Message string
}

func (se StatusErr) Error() string {
	return se.Message
}

func LoginAndGetData(uid, pwd, file string) ([]byte, error) {
	// err := login(uid, pwd)
	// if err != nil {
	// 	return nil, StatusErr{
	// 		Status:  InvalidLogin,
	// 		Message: fmt.Sprintf("invalid credentials for user %s", uid),
	// 	}
	// }
	// data, err := getData(file)
	// if err != nil {
	// 	return nil, StatusErr{
	// 		Status:  NotFound,
	// 		Message: fmt.Sprintf("file %s not found", file),
	// 	}
	// }
	// return data, nil
	return nil, nil
}
