import java.util.List;
import java.util.concurrent.Callable;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.Future;

// Read count_down_latch.md for more details
public class CountDownLatchPoc {
    public static void main(String[] args) {
        runCountDownLatchSample();
        runExeuctorServiceDiffSample();
    }

    public static void runExeuctorServiceDiffSample() {
        System.out.println("running executor service");
        ExecutorService es = Executors.newFixedThreadPool(2);
        // task of type 1
        Callable<Void> task1 = () -> {
            System.out.println("callable task 1 :: working");
            Thread.sleep(1000); // I/O task type 1
            return null;
        };

        // task of type 2
        Callable<Void> task2 = () -> {
            System.out.println("callable task 2 :: working");
            Thread.sleep(5000); // I/O task type 2
            return null;
        };
        try {
            List<Future<Void>> futures = es.invokeAll(List.of(task1, task2));
            for (Future<Void> future : futures) {
                try {
                    System.out.println("resolving future");
                    future.get();
                } catch (InterruptedException | ExecutionException e) {
                    e.printStackTrace();
                }
            }
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
        System.out.println("resolving complete");
        es.shutdown();
    }

    public static void runCountDownLatchSample() {
        CountDownLatch countDownLatch = new CountDownLatch(2); // assume that there is one parallel
                                                               // thread which is doing a operation.

        Thread t1 = new Thread(() -> {
            System.out.println("thread 1 :: working");
            try {
                Thread.sleep(1000); // I/O task type 1
            } catch (InterruptedException ex) {
                ex.printStackTrace();
            } finally {
                countDownLatch.countDown();
            }
        });
        Thread t2 = new Thread(() -> {
            System.out.println("thread 2 :: working");
            try {
                Thread.sleep(2000); // I/O task type 2
            } catch (InterruptedException ex) {
                ex.printStackTrace();
            } finally {
                countDownLatch.countDown();
            }
        });

        t1.start();
        t2.start();

        try {
            System.out.println("main thread::count down latch, awaiting for t1 & t2 to complete");
            countDownLatch.await(); // main thread awaits for the t1 & t2
            System.out.println("main thread::count down latch, t1 & t2 completed");
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
    }
}
