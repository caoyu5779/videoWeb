package dbops

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M)  {
	clearTables()
	m.Run()
	//clearTables()
}

func TestUserWorkFlow(t *testing.T)  {
	t.Run("Add", testAddUser)
	//t.Run("Get", testGetUser)
	//t.Run("Delete", testDeleteUser)
	//t.Run("reGet", testRegetUser)
}

func testAddUser(t *testing.T)  {
	err := AddUserCredential("chaosLee", "123")

	if err != nil {
		t.Errorf("Error of AddUser: %v", err)
	}
}

func testGetUser(t *testing.T)  {
	pwd, err := GetUserCredential("chaosLee")

	if pwd != "123" || err != nil {
		t.Errorf("Error of AddUser: %v", err)
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("chaosLee", "123")

	if err != nil {
		t.Errorf("Error of DeleteUser : %v" ,err)
	}
}

func testRegetUser(t *testing.T)  {
	pwd, err := GetUserCredential("chaosLee")

	if err != nil {
		t.Errorf("Error of RegetUser : %v", err)
	}

	if pwd != "" {
		t.Errorf("Deleting user test failed")
	}
}

func TestComments(t *testing.T)  {
	clearTables()
	t.Run("AddUser", testAddUser)
	t.Run("Add Comments", testAddComments)
	t.Run("List Comments", testListComments)
}

func testAddComments(t *testing.T)  {
	vid := "12345"
	aid := 1
	content := "Like this video"

	err := AddNewComments(vid, aid, content)

	if err != nil {
		t.Errorf("Error of AddComments: %v", err)
	}
}

func testListComments(t *testing.T) {
	vid := "12345"
	//to := time.Now().Unix()
	//
	//from,_ := time.ParseDuration("-24*7h")
	from := 1514764800
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/ 1000000000, 10))

	res, err := ListComments(vid, from, to)

	if err != nil {
		t.Errorf("Error of ListComments: %v", err)
	}

	for i, ele := range res {
		fmt.Printf("comment : %d, %v \n", i, ele)
	}


}