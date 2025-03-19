package main

import (
	"encoding/json"
	"fetchCourses/generator"
	"os"
	"path"
)

type filePair struct {
	OneJsonArrayFilePath string
	AllJsonFilePath      string
}

func UpdateDB() {

	filePaths := []filePair{
		{
			OneJsonArrayFilePath: "general_details.json",
			AllJsonFilePath:      "general.json",
		},
		{
			OneJsonArrayFilePath: "major_details.json",
			AllJsonFilePath:      "major.json",
		},
		{
			OneJsonArrayFilePath: "physics_details.json",
			AllJsonFilePath:      "physics.json",
		},
	}
	updateDBByJson("jsons/", filePaths...)
}

func updateDBByJson(fileDir string, filenames ...filePair) {

	//println(path.Join(fileDir, "a.json"))
	var infos []OneJson
	var allInfo AllJson

	table1IdCount := 0
	table2IdCount := 0
	for _, filename := range filenames {
		file1, _ := os.ReadFile(path.Join(fileDir, filename.OneJsonArrayFilePath))
		json.Unmarshal(file1, &infos)

		file2, _ := os.ReadFile(path.Join(fileDir, filename.AllJsonFilePath))
		json.Unmarshal(file2, &allInfo)

		doUpdateDB(allInfo, infos, &table1IdCount, &table2IdCount)
	}

}

func doUpdateDB(allInfo AllJson, infos []OneJson, table1IdCount *int, table2IdCount *int) {

	if len(infos) != len(allInfo.TmpList) {
		println(len(infos), len(allInfo.TmpList))
		panic("数据长度不一致")
	}

	resCourse := make([]CourseInfo, 0)
	resTime := make([]TimeInfo, 0)

	for i, info := range infos {
		if info.Jxdd == "--" || info.Jxdd == "" {
			continue
		}
		*table1IdCount++

		teacherInfos := extractTeacher(info.Jsxx)
		resCourse = append(resCourse, CourseInfo{
			ID:               uint32(*table1IdCount),
			Years:            info.Year,
			Semester:         "2",
			CourseNum:        allInfo.TmpList[i].KchId,
			CourseName:       allInfo.TmpList[i].Kcmc,
			Faculty:          info.Kkxymc,
			Credit:           allInfo.TmpList[i].Xf,
			CourseComplexion: info.Kcxzmc,
			CourseType:       info.Kclbmc,
			Grade:            "",
			Major:            "",
			Teacher:          teacherInfos[0].TeacherName,
			TeacherTitle:     teacherInfos[0].TeacherTitle,
		})

		weekInfo := extractTime(info.Sksj)
		addressInfo := extractAddress(info.Jxdd)
		for _, w := range weekInfo {
			*table2IdCount++

			resTime = append(resTime, TimeInfo{
				ID:           uint32(*table2IdCount),
				CourseInfoId: uint32(*table1IdCount),
				WeekAndTime:  generator.WeekLesson2Bin(w.Weeks, w.Lessons),
				DayOfWeek:    uint8(w.Week),
				Area:         uint8(addressInfo.Area),
				Building:     addressInfo.Building,
				Classroom:    addressInfo.Room,
			})
		}

	}

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
