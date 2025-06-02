package main

import database "fetchCourses/db"

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
