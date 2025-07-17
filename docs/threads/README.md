# Работа с потоками и синхронизация.

~~~java
import java.util.concurrent.atomic.AtomicInteger;

public class ThreadExample {
    private static AtomicInteger counter = new AtomicInteger(0); // Replace with AtomicInteger

    public static void main(String[] args) {
        Runnable task = () -> {
            for (int i = 0; i < 1000; i++) {
                counter.incrementAndGet(); // Atomic increment
            }
        };

        Thread thread1 = new Thread(task);
        Thread thread2 = new Thread(task);

        thread1.start();
        thread2.start();

        try {
            thread1.join();
            thread2.join();
        } catch (InterruptedException e) {
            e.printStackTrace();
        }

        System.out.println("Counter: " + counter.get()); // Use get() to retrieve value
    }
}
~~~