.PHONY: all
all:
	go run ariga.io/entimport/cmd/entimport -dsn "mysql://root:OrV6bilbR8WzCgk@tcp(127.0.0.1:3306)/ohdat" -tables "questionnaire_iqiyi,account"
	go generate ./ent