include .env

MYSQL_HOST ?= 127.0.0.1 # 外部から値を上書きできるようにする
MYSQL_PORT ?= 3306

export MYSQL_PWD=$(MYSQL_PASSWORD) # warningを出さないように環境変数 MYSQL_PWD を使う
MYSQL_DUMP_CMD = mysqldump -h $(MYSQL_HOST) -P $(MYSQL_PORT) -u $(MYSQL_USER)

.PHONY: dump-schema update-schema # 既存のファイルがあっても上書きできるようにする

dump-schema: # structure.sqlに最新スキーマを書き出す(DROP文なし)
	$(MYSQL_DUMP_CMD) --no-data --skip-add-drop-table --no-tablespaces $(MYSQL_DATABASE) > db/structure.sql
