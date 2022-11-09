package sql

const (
	SAVE_MOVIE = "INSERT INTO movies (title, rating, awards, length, genre_id) VALUES (?, ?, ?, ?, ?);"

	GET_MOVIE = "SELECT id, title, rating, awards, length, genre_id FROM movies WHERE id=?;"

	GET_ALL_MOVIE = "SELECT id, title, rating, awards, length, genre_id FROM movies"

	GET_BY_TITLE = "SELECT id, title, rating, awards, length, genre_id FROM movies WHERE title=?;"

	UPDATE_MOVIE = "UPDATE movies SET title=?, rating=?, awards=?, length=?, genre_id=? WHERE id=?;"

	DELETE_MOVIE = "DELETE FROM movie where id=?"

	EXIST_MOVIE = "SELECT m.id FROM movies m WHERE m.id=?"

	GET_MOVIE_TIMEOUT = "SELECT SLEEP(20) FROM DUAL WHERE id=?;"
)
