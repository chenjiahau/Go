package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"ivanfun.com/mis/db/driver"
	"ivanfun.com/mis/model"
	"ivanfun.com/mis/util"
)

func main() {
	target := os.Args[1]

	pgConn := ConnectDB()
	defer pgConn.SQL.Close()
	
	if target == "color" {
		HandleColor(pgConn)
	}

	if target == "memberrole" {
		HandleMemberRole(pgConn)
	}
}

func ConnectDB() *driver.DBConn {
	pgConn, err := driver.ConnectSQL(driver.PostgreSQLDataSourceName)
	if err != nil {
		log.Fatal("cannot connect to database")
	}

	model.NewDbConfig(pgConn)
	return pgConn
}

func HandleColor(PgConn *driver.DBConn, ) {
	var c model.ColorInterface = &model.Color{}
	var cc model.ColorCategoryInterface = &model.ColorCategory{}

	// Delete all colors
	err := c.DeleteAll()
	if err != nil {
		log.Fatal(err)
	}

	// Delete all color categories
	err = cc.DeleteAll()
	if err != nil {
		log.Fatal(err)
	}

	// Read and insert color categories and colors
	filePath := "./init/color"
	files, err := os.ReadDir(filePath)
   if err != nil {
			util.WriteErrorLog(err.Error())
      return
   }

   for _, file := range files {
      fileName := strings.Split(file.Name(), ".")[0]
			id, err := cc.Create(fileName)

			if err != nil {
				util.WriteErrorLog(err.Error())
			}

			util.WriteInfoLog(fmt.Sprintf("Color category %s created with id %d", fileName, id))

			colors, err := os.ReadFile(fmt.Sprintf("%s/%s.json", filePath, fileName))
			if err != nil {
				util.WriteErrorLog(err.Error())
				return
			}

			var colorList []model.Color
			err = json.Unmarshal(colors, &colorList)
			if err != nil {
				util.WriteErrorLog(err.Error())
				return
			}

			for _, color := range colorList {
				id, err := c.Create(id, color.Name, color.HexCode, color.RGBCode)
				if err != nil {
					util.WriteErrorLog(err.Error())
				}

				util.WriteInfoLog(fmt.Sprintf("Color %s created with id %d", color.Name, id))
			}	

   }
}

func HandleMemberRole(PgConn *driver.DBConn) {
	var mr model.MemberRoleInterface = &model.MemberRole{}

	// Delete all member roles
	err := mr.DeleteAll()
	if err != nil {
		log.Fatal(err)
	}

	// Read and insert member roles
	filePath := "./init/member/Role.json"
	memberRoles, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	var memberRoleList []model.MemberRole
	err = json.Unmarshal(memberRoles, &memberRoleList)
	if err != nil {
		log.Fatal(err)
	}

	for _, memberRole := range memberRoleList {
		id, err := mr.Create(memberRole.Title, memberRole.Abbr)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Member role %s created with id %d\n", memberRole.Title, id)
	}
}