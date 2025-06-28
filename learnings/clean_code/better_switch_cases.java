// Anti Pattern

class EmployeePayCalculator {
  
  public Double calculatePay(Employee e) {
    return switch(e.emloyeeType) {
      case HOURLY -> calculateHourlyPay(e);
      case MONTLY -> calculateMonthlyPay(e);
      case DAILY -> calculateDailyPay(e);
      default: throw new IllegalArgumentException("unknown employee type");
  }
}

class Employee {
  private String employeeName;
  private EmployeeType employeeType;
}

class EmployeeType {
  HOURLY,
  MONTHLY,
  DAILY
}


class Driver {

  public static void main(String[] args) {

    Employee e = new Employee("foo", EmployeeTypeEnum.HOURLY);

    Double pay = calculatePay(e);

    Double taxes = calculateTaxes(e); // calculateTaxes may have further switch cases.
  }
}


// Better Pattern which complies with SPR, OCP

public abstract class Employee {
  private String employeeName;
  private String employeeId;
}

// Hourly Employee is a Employee
public class HourlyEmployee extends Employee {
  
  public HourlyEmployee(String employeeName, String employeeId) {
    super(employeeName, employeeId);
  }

  public Double calculatePay() {
    // perform computions
  }

  public Double calculateTaxes() {
    // perform computations
  }
}

public class DailyEmployee extends Employee {
  
  public DailyEmployee(String employeeName, String employeeId) {
    super(employeeName, employeeId);
  }

  public Double calculatePay() {
    // perform computions
  }

  public Double calculateTaxes() {
    // perform computations
  }
}

class Driver {

  public static void main(String[] args) {

    Employee e = new HourlyEmployee("foo", EmployeeTypeEnum.HOURLY);

    e.calculatePay();
    e.calculateTaxes();
  }
}
