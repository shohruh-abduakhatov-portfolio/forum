package config

func init() {
	conf = map[string]string{
		"EmailUsername":   "AlifCapitalHoldingsTest@gmail.com",
		"EmailPass":       "AlifCapitalHoldings2020",
		"EmailHost":       "smtp.gmail.com",
		"EmailPort":       "587",
		"DBName":          "forum.db",
		"DBUser":          "admin",
		"DBPass":          "admin",
		"GoogleClientId":  "658570880968-h0k49f8ah9g3ukciqrlihoc7r623aiqj.apps.googleusercontent.com",
		"GoogleSecretKey": "1wYo_dT6Yl5fHfyLF_aprhJN",
		"GoogleCallback":  "/googleoauth2callback",
		"GithubClientID":  "27ff55bdddc4e636df57",
		"GithubSecretKey": "95b2723757d89594132a72e14e8de8ff9fa8c797",
		"Host":            "localhost",
		"Port":            "8000",
		"Protocol":        "http",
	}
}
