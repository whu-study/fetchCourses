package main

type AllJson struct {
	TmpList []struct {
		Blyxrs             string `json:"blyxrs"`
		Blzyl              string `json:"blzyl"`
		Cxbj               string `json:"cxbj"`
		Date               string `json:"date"`
		DateDigit          string `json:"dateDigit"`
		DateDigitSeparator string `json:"dateDigitSeparator"`
		Day                string `json:"day"`
		Fxbj               string `json:"fxbj"`
		Jgpxzd             string `json:"jgpxzd"`
		JxbId              string `json:"jxb_id"`
		Jxbmc              string `json:"jxbmc"`
		Jxbzls             string `json:"jxbzls"`
		Kch                string `json:"kch"`
		KchId              string `json:"kch_id"`
		Kclxmc             string `json:"kclxmc"`
		Kcmc               string `json:"kcmc"`
		Kcrow              string `json:"kcrow"`
		Kklxdm             string `json:"kklxdm"`
		Kzmc               string `json:"kzmc"`
		Listnav            string `json:"listnav"`
		LocaleKey          string `json:"localeKey"`
		Month              string `json:"month"`
		PageTotal          int    `json:"pageTotal"`
		Pageable           bool   `json:"pageable"`
		QueryModel         struct {
			CurrentPage   int           `json:"currentPage"`
			CurrentResult int           `json:"currentResult"`
			EntityOrField bool          `json:"entityOrField"`
			Limit         int           `json:"limit"`
			Offset        int           `json:"offset"`
			PageNo        int           `json:"pageNo"`
			PageSize      int           `json:"pageSize"`
			ShowCount     int           `json:"showCount"`
			Sorts         []interface{} `json:"sorts"`
			TotalCount    int           `json:"totalCount"`
			TotalPage     int           `json:"totalPage"`
			TotalResult   int           `json:"totalResult"`
		} `json:"queryModel"`
		Rangeable   bool   `json:"rangeable"`
		Rwzxs       string `json:"rwzxs"`
		TotalResult string `json:"totalResult"`
		UserModel   struct {
			Monitor    bool   `json:"monitor"`
			RoleCount  int    `json:"roleCount"`
			RoleKeys   string `json:"roleKeys"`
			RoleValues string `json:"roleValues"`
			Status     int    `json:"status"`
			Usable     bool   `json:"usable"`
		} `json:"userModel"`
		Xf    string `json:"xf"`
		Xxkbj string `json:"xxkbj"`
		Year  string `json:"year"`
		Yxzrs string `json:"yxzrs"`
	} `json:"tmpList"`
	Sfxsjc string `json:"sfxsjc"`
}

type OneJson struct {
	Bxbj               string `json:"bxbj"`
	Date               string `json:"date"`
	DateDigit          string `json:"dateDigit"`
	DateDigitSeparator string `json:"dateDigitSeparator"`
	Day                string `json:"day"`
	DoJxbId            string `json:"do_jxb_id"`
	Dsfrl              string `json:"dsfrl"`
	Fxbj               string `json:"fxbj"`
	Jgpxzd             string `json:"jgpxzd"`
	Jsxx               string `json:"jsxx"`
	JxbId              string `json:"jxb_id"`
	Jxbrl              string `json:"jxbrl"`
	Jxdd               string `json:"jxdd"`
	Jxms               string `json:"jxms"`
	Kcgsmc             string `json:"kcgsmc"`
	Kclbmc             string `json:"kclbmc"`
	Kcxzmc             string `json:"kcxzmc"`
	Kkxymc             string `json:"kkxymc"`
	Listnav            string `json:"listnav"`
	LocaleKey          string `json:"localeKey"`
	Month              string `json:"month"`
	PageTotal          int    `json:"pageTotal"`
	Pageable           bool   `json:"pageable"`
	QueryModel         struct {
		CurrentPage   int           `json:"currentPage"`
		CurrentResult int           `json:"currentResult"`
		EntityOrField bool          `json:"entityOrField"`
		Limit         int           `json:"limit"`
		Offset        int           `json:"offset"`
		PageNo        int           `json:"pageNo"`
		PageSize      int           `json:"pageSize"`
		ShowCount     int           `json:"showCount"`
		Sorts         []interface{} `json:"sorts"`
		TotalCount    int           `json:"totalCount"`
		TotalPage     int           `json:"totalPage"`
		TotalResult   int           `json:"totalResult"`
	} `json:"queryModel"`
	Rangeable   bool   `json:"rangeable"`
	Sjsfsj      string `json:"sjsfsj"`
	Sksj        string `json:"sksj"`
	TotalResult string `json:"totalResult"`
	UserModel   struct {
		Monitor    bool   `json:"monitor"`
		RoleCount  int    `json:"roleCount"`
		RoleKeys   string `json:"roleKeys"`
		RoleValues string `json:"roleValues"`
		Status     int    `json:"status"`
		Usable     bool   `json:"usable"`
	} `json:"userModel"`
	XqhId string `json:"xqh_id"`
	Xqumc string `json:"xqumc"`
	Year  string `json:"year"`
	Yqmc  string `json:"yqmc"`
}
