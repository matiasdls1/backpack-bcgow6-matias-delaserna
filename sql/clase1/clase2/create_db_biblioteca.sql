CREATE DATABASE biblioteca;

CREATE TABLE autor
(
idAutor INT PRIMARY KEY NOT NULL,
Nombre VARCHAR(30) NOT NULL,
Nacionalidad VARCHAR(30) NOT NULL
);

CREATE TABLE estudiante
(
idLector INT PRIMARY KEY NOT NULL,
Nombre VARCHAR(30) NOT NULL,
Apellido VARCHAR(30) NOT NULL,
Direccion VARCHAR(30) NOT NULL,
Carrera VARCHAR(30) NOT NULL,
Edad INT NOT NULL
);

CREATE TABLE libro
(
idLibro INT PRIMARY KEY NOT NULL,
Titulo VARCHAR(30) NOT NULL,
Editorial VARCHAR(30) NOT NULL,
Area VARCHAR(30) NOT NULL
);

CREATE TABLE libroautor
(
idAutor INT NOT NULL,
idLibro INT NOT NULL,
FOREIGN KEY (idAutor) REFERENCES autor(idAutor),
FOREIGN KEY (idLibro) REFERENCES libro(idLibro)
);

CREATE TABLE prestamo
(
idLibro INT NOT NULL,
idLector INT NOT NULL,
FOREIGN KEY (idLector) REFERENCES estudiante(idLector),
FOREIGN KEY (idLibro) REFERENCES libro(idLibro),
FechaPrestamo VARCHAR(30) NOT NULL,
FechaDevolucion VARCHAR(30) NOT NULL,
Devuelto BOOL NOT NULL
)