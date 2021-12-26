PSQL := docker-compose exec postgis psql -U postgis
PSQL_COMMAND := $(PSQL) -c

psql:
	$(PSQL)
cleanup:
	docker volume rm JapanCityGeoJson_db-data -f && docker-compose build --no-cache && docker-compose up
up:
	docker-compose down && docker volume rm JapanCityGeoJson_db-data -f && docker-compose up --build
all:
	@make japan
	@make japan_count
	@make japan_count
	@make japan_ogaki
	@make prefectures
	@make prefectures_count
	@make prefectures_tokyo
japan_table:
	$(PSQL_COMMAND) "\d japan"
japan_count:
	$(PSQL_COMMAND) "select count(*) as japan_count from japan;"
japan_ogaki:
	$(PSQL_COMMAND) "select pref, regional, city1, city2, GeometryType(geom) from japan where city2 = '大垣市';"
prefectures:
	$(PSQL_COMMAND) "\d prefectures"
prefectures_count:
	$(PSQL_COMMAND) "select code, name, count(*) as count from prefectures group by code, name order by code;"
prefectures_tokyo:
	$(PSQL_COMMAND) "select code, name, GeometryType(geom) from prefectures where name = '東京都';"
