-- 1 Seleccionar el nombre, el puesto y la localidad de los departamentos donde trabajan los vendedores.
select emp.nombre, emp.apellido, emp.puesto, dep.localidad from empleado emp join departamento dep on emp.depto_nro = dep.depto_nro;


-- 2 Visualizar los departamentos con más de cinco empleados.
select dep.nombre_depto from empleado emp join departamento dep
on dep.depto_nro = emp.depto_nro
group by emp.depto_nro
having count(emp.depto_nro) > 1;

 
-- 3 Mostrar el nombre, salario y nombre del departamento de los empleados que tengan el mismo puesto que ‘Mito Barchuk’.
select emp.nombre, emp.apellido, emp.salario, dep.nombre_depto, emp.puesto from empleado emp 
join departamento dep on emp.depto_nro = dep.depto_nro 
where emp.puesto =
(select empleado.puesto from empleado where empleado.nombre like 'Mito' and empleado.apellido like 'Barchuk');


-- 4 Mostrar los datos de los empleados que trabajan en el departamento de contabilidad, ordenados por nombre.
select emp.* from empleado emp 
join departamento dep on emp.depto_nro = dep.depto_nro
where dep.nombre_depto like 'Contabilidad'
order by emp.nombre;


-- 5 Mostrar el nombre del empleado que tiene el salario más bajo.
select * from empleado emp
limit 1;


-- 6 Mostrar los datos del empleado que tiene el salario más alto en el departamento de ‘Ventas’.
select * from empleado emp
where emp.depto_nro =
(select dep.depto_nro from departamento dep
where dep.nombre_depto like 'Ventas')
and emp.salario = 
(select max(emp.salario) from empleado emp);
