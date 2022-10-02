import java.util.ArrayList;
import java.util.List;

public class TypesOfVariables {
    public static void main(String[] args) {
        String stringVariable = "hello world";
        Integer integerVariable = 42;
        List<String> listOfStrings = new ArrayList<>();
        listOfStrings.add("one");
        listOfStrings.add("two");
        listOfStrings.add("three");
        System.out.println(stringVariable.getClass());
        System.out.println(stringVariable);
        System.out.println(integerVariable.getClass());
        System.out.println(integerVariable);
        System.out.println(listOfStrings.getClass());
        System.out.println(listOfStrings);
    }
}
