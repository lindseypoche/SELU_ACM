package app

import "github.com/lindseypoche/SELU_ACM/api/controllers/ping"

func mapUrls() {

	router.GET("/ping", ping.Ping)
}
