package dbops

import (
	"VideoServer/api/defs"
	"VideoServer/api/utils"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func AddUserCredential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO users (login_name, pwd) VALUES (?, ?)")

	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(loginName, pwd)
	if err != nil {
		return nil
	}

	defer stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string)  (string, error) {
	stmtOut,err := dbConn.Prepare("SELECT pwd FROM users where login_name = ?")

	if err != nil {
		log.Printf("%s", err)
		return "", err
	}

	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}

	defer stmtOut.Close()

	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name = ? and pwd = ?")

	if err != nil {
		log.Printf("delete user error : %s", err)
		return err
	}

	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtDel.Close()

	return nil
}

func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	// create uuid
	vid ,err := utils.NewUUID()

	if err != nil {
		return nil, err
	}

	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05") // M d Y, HH:MM:SS
	stmtIns, err := dbConn.Prepare(`INSERT INTO video_info " +
		"(id,author_id, name,display_ctime) VALUES(?,?,?,?)`)

	if err != nil {
		return nil, err
	}

	_,err = stmtIns.Exec(vid,aid,name,ctime)

	if err != nil {
		return nil, err
	}

	res := &defs.VideoInfo{
		Id: vid,
		AuthorId: aid,
		Name: name,
		DisplayCtime: ctime,
	}

	defer stmtIns.Close()

	return res, nil
}

//删除
func DeleteVideo(vid string) error{
	stmtDel, err := dbConn.Prepare(`DELETE FROM video_info where id = ?`)

	if err != nil {
		return err
	}

	_,err = stmtDel.Exec(vid)

	if err != nil {
		log.Printf("delete from video failed : %s", err)
	}

	defer stmtDel.Close()
	return nil
}

//改
func UpdateVideoInfo(vid string, name string) (*defs.VideoInfo, error) {
	stmtUp, err := dbConn.Prepare("UPDATE video_info SET name = ? WHERE id = ?")

	if err != nil {
		return nil, err
	}

	_,err = stmtUp.Exec(name, vid)

	if err != nil {
		log.Printf("update video info failed ! id : %s, name : %s, msg : %s", vid, name, err)
		return nil, err
	}

	defer stmtUp.Close()
	return nil, err
}

func GetVideoInfo(vid string) (string, error)  {
	var name string
	stmtGet, err := dbConn.Prepare("SELECT name FROM video_info where id = ?")

	if err != nil {
		return "", err
	}

	err = stmtGet.QueryRow(vid).Scan(&name)

	if err != nil {
		log.Printf("get video info failed ; id : %s , msg : %s", vid, err)
		return "", err
	}

	defer stmtGet.Close()
	return name, nil
}

func AddNewComments(vid string , aid int, content string) error {
	id , err := utils.NewUUID()
	if err != nil {
		return err
	}

	stmtIns, err := dbConn.Prepare("INSERT INTO comments (id,video_id, author_id, content) VALUES (?,?,?,?)")

	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(id, vid, aid, content)

	if err != nil {
		return err
	}
	defer stmtIns.Close()

	return nil
}

func ListComments(vid string, from , to int) ([]*defs.Comment, error) {
	stmtOut, err := dbConn.Prepare(`SELECT comments.id, users.login_name,comments.content FROM comments 
										   Inner join users on comments.author_id = users.id where comments.video_id = ? 
										   and comments.time > FROM_UNIXTIME(?) AND comments.time <= FROM_UNIXTIME (?)`)
	var res []*defs.Comment

	rows, err := stmtOut.Query(vid, from , to)

	if err != nil {
		return res, err
	}

	for rows.Next() {
		var id, name, content string

		if err := rows.Scan(&id, &name, &content); err != nil {
			return res, err
		}

		c := &defs.Comment{
			Id:id,
			VideoId:vid,
			Author: name,
			Content: content,
		}

		res = append(res, c)
	}

	defer stmtOut.Close()
	return res, nil
}
