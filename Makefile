include .env

MYSQL_HOST ?= 127.0.0.1 # 外部から値を上書きできるようにする
MYSQL_PORT ?= 3306

MYSQL_DUMP_CMD = mysqldump -h $(MYSQL_HOST) -P $(MYSQL_PORT) -u $(MYSQL_USER) -p$(MYSQL_PASSWORD)
MYSQL_CONN = -h $(MYSQL_HOST) -P $(MYSQL_PORT) -u $(MYSQL_USER) -p$(MYSQL_PASSWORD) $(MYSQL_DATABASE)

.PHONY: dump-schema apply-schema

dump-schema: # 現在のDBスキーマをdump（DROPなし）
	$(MYSQL_DUMP_CMD) --no-data --skip-add-drop-table --no-tablespaces $(MYSQL_DATABASE) | \
	sed 's/ AUTO_INCREMENT=[0-9]*//' > db/structure.sql

apply-schema: # structure.sqlとDBの差分を検出して適用（破壊的でない）
	cat db/structure.sql | mysqldef $(MYSQL_CONN)

