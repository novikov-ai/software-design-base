# Почему давно работающий код всё ещё содержит баги: Асинхронное программирование и многопоточность

## Race condition

### Проблема

В примере ниже гонка возникает из-за того, что каждая следующая итерация выполняется параллельно с предыдущей. Таким образом мы получаем борьбу за общий ресурс - в данном случае доступ к счетчику "counter", который одновременно обновляется сразу несколькими потоками. В примере операция инкремента не является атомарной. 

Чтобы исправить ситуацию, необходимо добавить, например, мьютекс или атомарный тип счетчика. 

~~~java
public class RaceConditionExample {

    private static int counter = 0;

    public static void main(String[] args) {
        int numberOfThreads = 10;
        Thread[] threads = new Thread[numberOfThreads];

        for (int i = 0; i < numberOfThreads; i++) {
            threads[i] = new Thread(() -> {
                for (int j = 0; j < 100000; j++) {
                    counter++;
                }
            });
            threads[i].start();
        }

        for (int i = 0; i < numberOfThreads; i++) {
            try {
                threads[i].join();
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }

        System.out.println("Final counter value: " + counter);
    }
}
~~~

### Решение

~~~java
import java.util.concurrent.atomic.AtomicInteger;

public class RaceConditionExample {
    private static final AtomicInteger counter = new AtomicInteger(0);

    public static void main(String[] args) throws InterruptedException {
        int numberOfThreads = 10;
        Thread[] threads = new Thread[numberOfThreads];

        for (int i = 0; i < numberOfThreads; i++) {
            threads[i] = new Thread(() -> {
                for (int j = 0; j < 100000; j++) {
                    counter.incrementAndGet(); // Атомарная операция
                }
            });
            threads[i].start();
        }

        for (Thread thread : threads) {
            thread.join();
        }

        System.out.println("Final counter value: " + counter.get());
    }
}
~~~

## Deadlock

### Проблема

Первй поток захватывает lock1, ждет какое-то время, за которое второй поток успевает захватить lock2, после чего мы в первом потоке хотим взять lock2, но он занят вторым потоком, который ожидает, когда мы отпустим lock1...

Решением будет служить унифицированный порядок блокировок для гарантии, что один поток дождется освобождения ресурса в другом и не будет пытаться его захватить.  

~~~java
public class DeadlockExample {

    private static final Object lock1 = new Object();
    private static final Object lock2 = new Object();

    public static void main(String[] args) {
        Thread thread1 = new Thread(() -> {
            synchronized (lock1) {
                System.out.println("Thread 1 acquired lock1");

                try { Thread.sleep(50); } 
                catch (InterruptedException e) { e.printStackTrace(); }

                synchronized (lock2) {
                    System.out.println("Thread 1 acquired lock2");
                }
            }
        });

        Thread thread2 = new Thread(() -> {
            synchronized (lock2) {
                System.out.println("Thread 2 acquired lock2");

                try { Thread.sleep(50); } 
                catch (InterruptedException e) { e.printStackTrace(); }

                synchronized (lock1) {
                    System.out.println("Thread 2 acquired lock1");
                }
            }
        });

        thread1.start();
        thread2.start();

        try {
            thread1.join();
            thread2.join();
        } catch (InterruptedException e) {
            e.printStackTrace();
        }

        System.out.println("Finished");
    }
}
~~~

### Решение

~~~java
public class DeadlockExample {

    private static final Object lock1 = new Object();
    private static final Object lock2 = new Object();

    public static void main(String[] args) {
        Thread thread1 = new Thread(() -> {
            synchronized (lock1) {
                System.out.println("Thread 1 acquired lock1");
                try { Thread.sleep(50); } 
                catch (InterruptedException e) { e.printStackTrace(); }
                
                synchronized (lock2) {
                    System.out.println("Thread 1 acquired lock2");
                }
            }
        });

        Thread thread2 = new Thread(() -> {
            // ВАЖНО: Изменён порядок блокировок - теперь совпадает с Thread1
            synchronized (lock1) {
                System.out.println("Thread 2 acquired lock1");
                try { Thread.sleep(50); } 
                catch (InterruptedException e) { e.printStackTrace(); }
                
                synchronized (lock2) {
                    System.out.println("Thread 2 acquired lock2");
                }
            }
        });

        thread1.start();
        thread2.start();

        try {
            thread1.join();
            thread2.join();
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
        System.out.println("Finished");
    }
}
~~~