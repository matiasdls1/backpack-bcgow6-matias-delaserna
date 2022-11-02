-- 1 Mostrar el título y el nombre del género de todas las series.
select s.title, g.name from series s inner join genres g on s.genre_id = g.id;
-- 2 Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.
select * from actor_episode;
select * from episodes;
select * from actors;
select * from series;
select * from movies;
-- 3 Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.
select ser.title, count(sea.number) seasons from series ser inner join seasons sea on sea.serie_id = ser.id group by ser.title;
-- 4 Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, siempre que sea mayor o igual a 3.
select gen.name, count(mov.title) from genres gen inner join movies mov on gen.id = mov.genre_id group by gen.name having count(mov.title) >= 3;
-- 5 Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas de la guerra de las galaxias y que estos no se repitan.
select ac.first_name, ac.last_name from actors ac where ac.id in (select distinct actor_id from actor_movie where movie_id in (select mov.id as "movie_id" from movies mov where mov.title like 'La Guerra de las galaxias%'));