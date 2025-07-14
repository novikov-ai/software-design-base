# Проблема с датой

## Исходный код

~~~java
import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.Date;

public class DateExample {
    public static void main(String[] args) {
        String dateString = "2024-05-13 14:30:00";
        SimpleDateFormat format = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss");
        try {
            Date date = format.parse(dateString);
            System.out.println("Date: " + date);
        } catch (ParseException e) {
            e.printStackTrace();
        }
    }
}
~~~

## Недостатки исходного Java-кода

1. Потоковая небезопасность
SimpleDateFormat не является потокобезопасным. Использование в многопоточных приложениях может вызвать гонки данных и неконсистентные результаты.

2. Устаревший API
Классы Date и SimpleDateFormat из пакета java.util считаются legacy (устаревшими) с Java 8. Они не поддерживают современные временные модели.

3. Неявная обработка часовых поясов
Код использует системный часовой пояс по умолчанию, что приводит:
- К неожиданным сдвигам времени при развертывании в разных средах
- Игнорированию проблем DST (летнее/зимнее время)
- Ошибкам в момент перевода часов

4. Слабая валидация данных

5. Автоматический пересчет невалидных дат (например, 2023-02-30 → 2023-03-02)

6. Нет обработки null/пустых строк

7. Принимает даты с неполной информацией (только год, месяц без дня)

8. Неинформативный вывод
System.out.println(date) использует формат по умолчанию (EEE MMM dd HH:mm:ss zzz yyyy), который не соответствует исходным данным.

9. Некорректная обработка ошибок
e.printStackTrace() не подходит для продакшн-кода:

10. Не предоставляет понятных сообщений пользователю

## Исправленный код

~~~java
import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;
import java.time.format.DateTimeParseException;

public class DateExample {
    public static void main(String[] args) {
        String dateString = "2024-05-13 14:30:00";
        DateTimeFormatter formatter = DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm:ss");
        
        try {
            LocalDateTime date = LocalDateTime.parse(dateString, formatter);
            
            // Вывод в исходном формате
            System.out.println("Original format: " + formatter.format(date));
            
            // Вывод в ISO-формате (альтернатива)
            System.out.println("ISO format: " + date);
        } catch (DateTimeParseException e) {
            System.err.println("Error parsing date: " + e.getMessage());
        }
    }
}
~~~