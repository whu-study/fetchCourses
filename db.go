package main

import (
	"encoding/json"
	"fetchCourses/generator"
	"os"
)

func UpdateDB() {

	var allInfo AllJson
	var infos []OneJson

	file1, _ := os.ReadFile("jsons/general_details.json")
	json.Unmarshal(file1, &infos)

	file2, _ := os.ReadFile("jsons/general.json")
	json.Unmarshal(file2, &allInfo)

	if len(infos) != len(allInfo.TmpList) {
		panic("数据长度不一致")
	}

	resCourse := make([]CourseInfo, 0)
	resTime := make([]TimeInfo, 0)

	table1IdCount := 0
	table2IdCount := 0

	for i, info := range infos {
		if info.Jxdd == "--" || info.Jxdd == "" {
			continue
		}
		table1IdCount++

		teacherInfos := extractTeacher(info.Jsxx)
		resCourse = append(resCourse, CourseInfo{
			ID:               uint32(table1IdCount),
			Years:            info.Year,
			Semester:         "2",
			CourseNum:        allInfo.TmpList[i].KchId,
			CourseName:       allInfo.TmpList[i].Kcmc,
			Faculty:          info.Kkxymc,
			Credit:           allInfo.TmpList[i].Xf,
			CourseComplexion: info.Kcxzmc,
			CourseType:       allInfo.TmpList[i].Kclxmc,
			Grade:            "",
			Major:            "",
			Teacher:          teacherInfos[0].TeacherName,
			TeacherTitle:     teacherInfos[0].TeacherTitle,
		})

		weekInfo := extractTime(info.Sksj)
		addressInfo := extractAddress(info.Jxdd)
		for _, w := range weekInfo {
			table2IdCount++

			resTime = append(resTime, TimeInfo{
				ID:           uint32(table2IdCount),
				CourseInfoId: uint32(table1IdCount),
				WeekAndTime:  generator.WeekLesson2Bin(w.Weeks, w.Lessons),
				DayOfWeek:    uint8(w.Week),
				Area:         uint8(addressInfo.Area),
				Building:     addressInfo.Building,
				Classroom:    addressInfo.Room,
			})
		}

	}
	//logger.Info("start")
	//println(database.Client.Error)
	//if err := router.Routers().Run(":" + config.EnvCfg.ServerPort); err != nil {
	//	logger.Error("run server error: ", err)
	//	return
	//}

	println(len(resCourse), len(resTime))

	batchInsert[CourseInfo](resCourse)
	batchInsert[TimeInfo](resTime)

}

func batchInsert[T CourseInfo | TimeInfo](data []T) {
	batchSize := 500
	for i := 0; i < len(data); i += batchSize {
		end := i + batchSize
		if end > len(data) {
			end = len(data)
		}

		if err := Client.Save(data[i:end]).Error; err != nil {
			panic(err)
			return
		}
	}
}
