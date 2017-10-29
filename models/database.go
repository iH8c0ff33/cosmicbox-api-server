package models

var db Database

type Database struct {
	connection string
	events     []Event
}

func Initialize() error {
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("connection: ")
	// connection, err := reader.ReadString('\n')
	// if err != nil {
	// 	return err
	// }
	db.connection = "conn"
	db.events = []Event{}

	return nil
}
