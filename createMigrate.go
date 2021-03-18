package main

import (
	"github.com/NovikovRoman/gmdb/templates"
	"path/filepath"
)

func createMigrate() (err error) {
	const sqlDir = "/migrations"

	pathSql := filepath.Join(*migratePath, "..", sqlDir)

	if err = createDir(*migratePath); err != nil {
		return err
	}
	if err = createDir(pathSql); err != nil {
		return err
	}

	data := struct {
		Backtick string
	}{
		Backtick: backtick,
	}

	if err = saveTemplate(*migratePath+"/README.md", templates.MigrateReadme, data); err != nil {
		return
	}

	if err = saveTemplate(*migratePath+"/main.go", templates.MigrateMain, data); err != nil {
		return
	}

	err = saveTemplate(
		filepath.Join(pathSql, "202005041600_proxy.up.sql"), templates.MigrateUpSql, data)
	if err != nil {
		return
	}

	err = saveTemplate(
		filepath.Join(pathSql, "202005041600_proxy.down.sql"), templates.MigrateDownSql, data)
	return
}
