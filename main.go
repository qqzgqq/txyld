package main

import (
	"fmt"
	"math"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/kataras/iris"
)

// TxlydSQL mysql表txldy
type TxlydSQL struct {
	IDsql        int    `db:"id"`
	Starttime    string `db:"starttime"`
	Endtime      string `db:"endtime"`
	Question     string `db:"question"`
	Place        string `db:"place"`
	Questiontype string `db:"questiontype"`
	Count        string `db:"count"`
	Solvetime    string `db:"solvetime"`
	Note         string `db:"note"`
}

// MobaN mysql表moban
type MobaN struct {
	IDmoban          int    `db:"id"`
	QuestionList     string `db:"question_list"`
	PlaceList        string `db:"place_list"`
	QuestiontypeList string `db:"questiontype_list"`
	CountList        string `db:"count_list"`
}

func main() {
	db, _ := sqlx.Connect("mysql", "root:1q2w3e@tcp(127.0.0.1:3306)/guang?charset=utf8")
	app := iris.New()
	app.RegisterView(iris.HTML("./templates", ".html"))
	// sess := sessions.New(sessions.Config{
	// 	Cookie:                      "mysessionid",
	// 	Expires:                     time.Hour * 2,
	// 	DisableSubdomainPersistence: false,
	// })

	app.Get("/question_list", func(ctx iris.Context) {
		var page float64
		db.Get(&page, "SELECT count(*) FROM txlyd ")
		xsys := 3.0
		zxsys := math.Ceil(page / xsys)
		zxsZ := int(zxsys)
		var aaa []int
		for i := 0; i < zxsZ; i++ {
			aaa = append(aaa, i+1)

		}

		ctx.ViewData("sqlnum", aaa)
		//ctx.FormValue() must don't be null
		qunayeen := ctx.URLParams()["qunaye"]

		if qunayeen == "" {
			qunayeen = "1"
		}
		//string to int
		qunayeZ, _ := strconv.Atoi(qunayeen)
		qunayeS := qunayeZ*3 - 3

		sqlList := []TxlydSQL{}
		questionlist := []MobaN{}
		placelist := []MobaN{}
		questiontypelist := []MobaN{}
		countlist := []MobaN{}
		db.Select(&sqlList, "SELECT id,starttime,endtime,question,place,questiontype,count,solvetime,note FROM txlyd limit ?,?", qunayeS, 3)
		db.Select(&questionlist, "SELECT question_list FROM moban")
		db.Select(&placelist, "SELECT place_list FROM moban")
		db.Select(&questiontypelist, "SELECT questiontype_list FROM moban")
		db.Select(&countlist, "SELECT count_list FROM moban")

		ctx.ViewData("sqll", sqlList)
		ctx.ViewData("question_list", questionlist)
		ctx.ViewData("place_list", placelist)
		ctx.ViewData("questiontype_list", questiontypelist)
		ctx.ViewData("count_list", countlist)
		ctx.View("question_list.html")
	})
	app.Get("/delete", func(ctx iris.Context) {

		deleteidL := ctx.FormValue("deleteid_l")
		db.MustExec("delete from txlyd where id=?", deleteidL)
		fmt.Printf("id为：%s 删除成功", deleteidL)
		delid := ctx.FormValue("delid")
		db.MustExec("delete from moban where id=?", delid)
		fmt.Println("dell_moban_id:", delid)
		ctx.Redirect("/question_list", iris.StatusTemporaryRedirect)
	})
	app.Post("/insert", func(ctx iris.Context) {

		starttime := ctx.FormValue("starttime")
		endtime := ctx.FormValue("endtime")
		question := ctx.FormValue("question")
		place := ctx.FormValue("place")
		questiontype := ctx.FormValue("questiontype")
		count := ctx.FormValue("count")
		note := ctx.FormValue("note")
		if endtime == "" {
			solvetimeS := "-"
			db.MustExec("insert into txlyd(starttime, endtime, question, place, questiontype, count, solvetime,note) values(?,?,?,?,?,?,?,?)", starttime, endtime, question, place, questiontype, count, solvetimeS, note)
			fmt.Println(starttime, endtime, question, place, questiontype, count, solvetimeS, note)
			ctx.Redirect("/question_list/?qunaye=", iris.StatusTemporaryRedirect)
		} else {
			starttimeT, _ := time.Parse("2006-01-02 15:04", starttime)
			endtimeT, _ := time.Parse("2006-01-02 15:04", endtime)
			solvetime := (endtimeT.Unix() - starttimeT.Unix()) / 60
			solvetimeS := strconv.FormatInt(solvetime, 10) + "分钟"
			fmt.Println(solvetimeS)
			db.MustExec("insert into txlyd(starttime, endtime, question, place, questiontype, count, solvetime,note) values(?,?,?,?,?,?,?,?)", starttime, endtime, question, place, questiontype, count, solvetimeS, note)
			fmt.Println(starttime, endtime, question, place, questiontype, count, solvetimeS, note)
			ctx.Redirect("/question_list/?qunaye=", iris.StatusTemporaryRedirect)
		}

	})
	app.Post("/insert_moban", func(ctx iris.Context) {
		qtlbmb := ctx.FormValue("qtlbmb")
		qtddmb := ctx.FormValue("qtddmb")
		wtlx := ctx.FormValue("wtlx")
		fscs := ctx.FormValue("fscs")
		fmt.Println(qtlbmb, qtddmb, wtlx, fscs)
		db.MustExec("insert into moban(question_list,place_list,questiontype_list,count_list) values(?,?,?,?)", qtlbmb, qtddmb, wtlx, fscs)
		ctx.Redirect("/question_moban", iris.StatusTemporaryRedirect)
	})
	app.Get("/question_moban", func(ctx iris.Context) {
		questionxianshi := []MobaN{}
		db.Select(&questionxianshi, "select * from moban")
		fmt.Println("moban:", questionxianshi)
		ctx.ViewData("question_xianshi", questionxianshi)

		ctx.View("question_moban.html")

	})

	app.Get("/", func(ctx iris.Context) {
		ctx.Redirect("question_list", iris.StatusTemporaryRedirect)
	})

	app.StaticWeb("/static", "./static")
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
