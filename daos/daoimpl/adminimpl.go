package daoimpl

import (
	"OnlineTestGo/models"
	"OnlineTestGo/tos"
	"OnlineTestGo/utility"
	"log"
)

type AdminImpl struct{}

func (dao AdminImpl) FetchData() []tos.Admin {

	var datas []models.Admin
	var admindata []tos.Admin

	utility.GetLogger()
	log.Println("entering FetchData()")

	log.Println("executing query and Fetching data from db ")

	query := " select a.uid,r.fname,r.lname,a.testtype,a.score from answers a, registration r where r.id=a.uid"

	db, conn := connectaws()
	defer db.Close()
	defer conn.Close()

	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var data models.Admin
		err = rows.Scan(&data.Uid, &data.Fname, &data.Lname, &data.Testtype, &data.Score)

		if err != nil {
			log.Fatal(err)
		}

		datas = append(datas, data)

	}

	intitialUid := 0
	var testtype []string
	var score []int
	var admin tos.Admin
	for _, dataRow := range datas {

		if intitialUid != dataRow.Uid {
			admin.Uid = dataRow.Uid
			admin.Fname = dataRow.Fname
			admin.Lname = dataRow.Lname
			admindata = append(admindata, admin)
			intitialUid = dataRow.Uid
		}

	}

	n := 0
	intitialUid = datas[0].Uid
	for i := 0; i < len(datas); i++ {

		if intitialUid != datas[i].Uid {
			admindata[n].Type = testtype
			admindata[n].Score = score
			n++
			testtype = make([]string, 0)
			score = make([]int, 0)
			intitialUid = datas[i].Uid
		}

		testtype = append(testtype, datas[i].Testtype)
		admindata[n].Type = testtype
		score = append(score, datas[i].Score)
		admindata[n].Score = score

	}
	return admindata
}
