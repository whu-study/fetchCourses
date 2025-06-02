package main

import database "fetchCourses/db"

var cookie = "_dx_uzZo5y=1641785697592RSOSLdVXre7YXItZA6NtZp84qUZqdcgX; _dx_app_111=; _dx_captcha_vid=sl6m8oebm5tdepow; SF_cookie_1=16391220; JSESSIONID=CE7F3A7D70E3AEC52FB3D29D6A93780A; iPlanetDirectoryPro=q0sa7YGfbRqmt7sFfEvWpu"

func main() {
	database.InitMysql()
	UpdateDB()
	//getMajor()
	//getAllMajorDetails()

	//getGeneral()
	//getAllGeneralDetails()

	//getPhysics()
	//getAllPhysicsDetails()
}
