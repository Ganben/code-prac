select
    Department.name AS 'Department',
    Employee.name AS 'Employee',
    Salary
from
    Employee
        join
    Department on Employee.DepartmentId = Department.Id
where
    (Employee.DepartmentId, Salary) in
    (   select
            DepartmentId, MAX(Salary)
        from
            Employee
        group by DepartmentId
        )
;