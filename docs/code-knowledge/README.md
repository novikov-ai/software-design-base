# "Знание" кода

Три совсем простых примера. Главное: отнеситесь к ним не как к задачкам по полиморфизму или сборке, а как к вариантам скрытого влияния одних частей кода на другие, и старайтесь такого по возможности избегать.

## 1.

~~~java
class Animal {
    public void makeSound() {
        System.out.println("Some generic animal sound");
    }
}

class Cat extends Animal {
    // Переопределение метода makeSound
    public void makeSound() {
        System.out.println("Meow");
    }
}

public class Main {
    public static void main(String[] args) {
        Animal myCat = new Cat();
        myCat.makeSound();  // "Meow"
    }
}
~~~

После:

~~~java
class Animal {
    // Изменен метод в суперклассе
    public void makeGenericSound() {
        System.out.println("Some generic animal sound");
    }
}
~~~

Ошибка компиляции, так как не сможем определить метод `makeSound`. Метод `makeSound` теперь не переопределяет метод суперкласса, а является методом подкласса. Следовало бы использовать аннотацию @Override.

## 2.

~~~java
class Animal {
    public void makeSound() {
        System.out.println("Some generic animal sound");
    }
}

class Cat extends Animal {
    @Override
    public void makeSound(int numberOfSounds) {
        for (int i = 0; i < numberOfSounds; i++) {
            System.out.println("Meow");
        }
    }
    
    @Override
    public void makeSound() {
        System.out.println("Meow");
    }
}

public class Main {
    public static void main(String[] args) {
        Animal cat = new Cat();
        cat.makeSound();
        cat.makeSound(3);
    }
}
~~~

Ошибка компиляции. Попытка переопределить метод `makeSound(int numberOfSounds)`, который отсутствует в суперклассе. Незримая логическая ошибка: ошибка с аннотацией переопределяемого метода. Необходимо убрать переопределение у нового метода. 

## 3.

~~~java
/*
<dependency>
    <groupId>com.fasterxml.jackson.core</groupId>
    <artifactId>jackson-databind</artifactId>
    <version>2.9.10</version>
</dependency>
<dependency>
    <groupId>com.fasterxml.jackson.core</groupId>
    <artifactId>jackson-databind</artifactId>
    <version>2.12.5</version>
</dependency>
*/

import com.fasterxml.jackson.databind.ObjectMapper;

import java.io.IOException;
import java.util.HashMap;
import java.util.Map;

public class Main {
    public static void main(String[] args) {
        // Создаем объект ObjectMapper для парсинга JSON
        ObjectMapper objectMapper = new ObjectMapper();

        String jsonString = "{\"name\":\"John\", \"age\":30}";

        try {
            // Парсим JSON-строку в HashMap
            Map<String, Object> result = objectMapper.readValue(jsonString, HashMap.class);

            System.out.println("Name: " + result.get("name"));
        } catch (IOException e) {
            // Обработка ошибки парсинга
            e.printStackTrace();
        }

        try {
            String prettyJson = objectMapper.writerWithDefaultPrettyPrinter().writeValueAsString(result);
            System.out.println("Pretty JSON: " + prettyJson);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
~~~

Существует проблема со сборкой зависимости: дублирование зависимостей разных версий. В рантайме может загрузиться не та версия, которую ожидает разработчик, а у нее может поменяться API и поведение.

Также есть проблема с видимостью переменной `result` - она локально объявлена в первом блоке try, но недоступна в другом блоке try.