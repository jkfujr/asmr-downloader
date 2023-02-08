package patch

import (
	"bufio"
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)
import _ "github.com/mattn/go-sqlite3"

var haveDoneTxt = "have-download.txt"

func PatchHavenDownload2DB() {
	db, err := sql.Open("sqlite3", "../asmr.db")
	_ = db
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.OpenFile(haveDoneTxt, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatalf("read file line error: %v", err)
			return
		}
		println(line)
		line = strings.Trim(line, "\n")
		id := strings.Replace(line, "RJ", "", 1)
		realId, _ := strconv.Atoi(id)
		id = "RJ" + strconv.Itoa(realId)
		println(id)
		tx, err := db.Begin()

		if err != nil {
			log.Fatal("开启事务失败: ", err)
		}
		_, err = tx.Exec("update asmr_download set download_flag = 1 where rjid = ?", id)
		if err != nil {
			tx.Rollback()
			fmt.Println("数据下载完成状态更新失败: ", err)
			fmt.Println("正在进行数据回滚...")
		}
		err = tx.Commit()
		if err != nil {
			fmt.Println("数据提交失败：", err)
		}

	}
}
