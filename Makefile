include .env

MYSQL_HOST ?= 127.0.0.1 # 外部から値を上書きできるようにする
MYSQL_PORT ?= 3306

MYSQL_DUMP_CMD = mysqldump -h $(MYSQL_HOST) -P $(MYSQL_PORT) -u $(MYSQL_USER) -p$(MYSQL_PASSWORD)

.PHONY: dump-schema update-schema # 既存のファイルがあっても上書きできるようにする

dump-schema: # structure.sqlに最新スキーマを書き出す(DROP文なし)
	$(MYSQL_DUMP_CMD) --no-data --skip-add-drop-table --no-tablespaces $(MYSQL_DATABASE) > db/structure.sql

# TODO: structure.sqlからテーブルを作れるようにする
