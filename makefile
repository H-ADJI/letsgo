db:
	@docker run -p 3306:3306 --name test-mysql -d -e MYSQL_ROOT_PASSWORD=1234  mysql 

