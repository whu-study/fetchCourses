package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type WeekInfo struct {
	Week    int   `json:"week"`
	Weeks   []int `json:"weeks"`
	Lessons []int `json:"lessons"`
}

func extractTime(s string) []WeekInfo {
	weekMapping := map[string]int{
		"星期一": 1,
		"星期二": 2,
		"星期三": 3,
		"星期四": 4,
		"星期五": 5,
		"星期六": 6,
		"星期日": 7,
	}

	// 修改后的正则表达式，包含对“周”字的匹配
	pattern := `([星期一二三四五六日]+)第(\d+)-(\d+)节\{([\d,-]+周(,[ \d,-]+周)*)\}`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(s, -1)

	var result []WeekInfo
	for _, match := range matches {
		week := weekMapping[match[1]]
		startLesson, _ := strconv.Atoi(match[2])
		endLesson, _ := strconv.Atoi(match[3])
		weekRanges := strings.Split(match[4], ",")

		var weeks []int
		for _, weekRange := range weekRanges {
			weekRange = strings.TrimSpace(weekRange)
			weekRange = strings.TrimSuffix(weekRange, "周")
			parts := strings.Split(weekRange, "-")
			if len(parts) == 1 {
				weekNum, _ := strconv.Atoi(parts[0])
				weeks = append(weeks, weekNum)
			} else if len(parts) == 2 {
				startWeek, _ := strconv.Atoi(parts[0])
				endWeek, _ := strconv.Atoi(parts[1])
				for w := startWeek; w <= endWeek; w++ {
					weeks = append(weeks, w)
				}
			}
		}

		lessons := []int{startLesson, endLesson}
		info := WeekInfo{
			Week:    week,
			Weeks:   weeks,
			Lessons: lessons,
		}
		result = append(result, info)
	}

	return result
}

type AddressInfo struct {
	Area     int    `json:"area"`
	Building string `json:"building"`
	Room     string `json:"room"`
}

var commonAddrRe = regexp.MustCompile(`(\d+)区(([\d` + "\u4e00-\u9fa5" + `]+)-|([A-Z]))(\d+)`)
var specialAddrRe = regexp.MustCompile(`国软(\d+)-(\d+)`)
var specialAddrRe2 = regexp.MustCompile(`新珈楼B(\d+)`)
var specialAddrRe3 = regexp.MustCompile(`(\d+)区(\d{2})(\d{3})`)

func extractAddress(s string) AddressInfo {

	res := AddressInfo{}
	if matches := commonAddrRe.FindStringSubmatch(s); len(matches) == 6 {
		res.Area, _ = strconv.Atoi(matches[1])
		if matches[3] != "" {
			res.Building = matches[3] + "-教学楼"
		} else {
			res.Building = matches[4] + "-教学楼"
		}
		res.Room = matches[5]
	} else if matches = specialAddrRe.FindStringSubmatch(s); len(matches) == 3 {
		res.Area = 3
		res.Building = "国软" + matches[1] + "-教学楼"
		res.Room = matches[2]
	} else if matches = specialAddrRe2.FindStringSubmatch(s); len(matches) == 2 {
		res.Area = 5
		res.Building = "新珈楼B"
		res.Room = matches[1]
	} else if matches = specialAddrRe3.FindStringSubmatch(s); len(matches) == 4 {
		res.Area, _ = strconv.Atoi(matches[1])
		res.Building = matches[2] + "-教学楼"
		res.Room = matches[3]
	} else if strings.Contains(s, "操场") || strings.Contains(s, "体育") || strings.Contains(s, "场") {
		res.Area = 6
		res.Building = s
		res.Room = "无"
	} else {
		panic(fmt.Sprintf("字符串格式不匹配: %s\n", s))
	}

	return res
}

type TeacherInfo struct {
	TeacherId    string `json:"teacherId"`
	TeacherName  string `json:"teacherName"`
	TeacherTitle string `json:"teacherTitle"`
}

// 可能有外语教师！
var teacherRe = regexp.MustCompile(`(\d+)/([\S\s]+?)/([` + "\u4e00-\u9fa5" + `]+)`)

func extractTeacher(s string) []TeacherInfo {
	// 00008396/杜莉/教授;00006329/李雪松/教授

	matches := teacherRe.FindAllStringSubmatch(s, -1)
	var result []TeacherInfo
	if len(matches) == 0 {
		panic(s)
	}
	for _, match := range matches {
		result = append(result, TeacherInfo{
			TeacherId:    match[1],
			TeacherName:  match[2],
			TeacherTitle: match[3],
		})
	}
	return result
}
