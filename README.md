# SOLID Principle

### Five software design principles:
- **S** - Single Responsibility Principle
- **O** - Open / Closed Principle
- **L** - Liskov Substitution Principle
- **I** - Interface Segregation Principle
- **D** - Dependency Inversion Principle

### Why SOLID
To **avoid** typical issues in software development:
- rigidity (difficult to make change)
- fragility (change brings service and development interruption)
- immobility (poor re-usability)
- viscosity (highly design coupling, difficult to apply changes)

## SOLID in Details
### S - Single Responsibility Principle
*A class should have one, and only one, reason to change. - Robert C Martin*

**Example Explanation**




### O - Open / Closed Principle
*A module should be open for extensions, but closed for modification. - Robert C Martin*


### L - Liskov Substitution Principle
*Derrived methods should expect no more and provide no less. - Robert C Martin*

### I - Interface Segregation Principle
*Many client specific interfaces are better than one general purpose interface*


### D - Dependency Inversion Principle
*Depend upon abstractions. Do not depend upon concretions. - Robert C Martin*

### Source
- https://dave.cheney.net/2016/08/20/solid-go-design
- 


### NANTI DIRAPIHIHN
SRP Keuntungan :
- it makes your software easier to implement
- prevents unexpected side-effects of future changes

*Contoh implementasi
Ada 4 struct geometri cube, tube, cone, sphere
Memiliki class sendiri untuk menghitung
Memiliki class sendiri untuk menampilkan
