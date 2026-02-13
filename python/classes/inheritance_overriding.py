# Exercise 3: Employee Hierarchy (Inheritance & Overriding)
# Scenario: A company has different types of employees with
# different salary structures.

# Task: Create a base class Employee with an __init__ for name
# and base_salary. Add a method calculate_pay().

# Requirement: Create a subclass Manager that gets a bonus. Create another
# subclass Developer that gets a stipend. Both should override calculate_pay()
# to include their extra compensation.

# Interview Focus: This checks your ability to use super() to call parent
# constructors and your understanding of method overriding.


class Employee:
    def __init__(self, name: str, base_salary: float):
        self.name = name
        self.base_salary = base_salary

    def calculate_pay(self) -> float:
        return self.base_salary


class Manager(Employee):
    def __init__(self, name: str, base_salary: float, bonus: float):
        # Correct way to call the parent constructor
        super().__init__(name, base_salary)
        self.bonus = bonus

    def calculate_pay(self) -> float:
        # You can use super().calculate_pay() here to get the base salary logic
        return super().calculate_pay() + self.bonus


# --- Testing the Hierarchy ---
employees = [
    Manager("Alice", 5000, 1000)
]

for emp in employees:
    print(f"{emp.name} is paid: ${emp.calculate_pay()}")
