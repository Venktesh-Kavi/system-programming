package list;

import java.util.Collection;
import java.util.Iterator;
import java.util.List;
import java.util.ListIterator;

public class CustomList<T> implements List<T> {
    private Object[] arr;
    private static final int DEFAULT_CAPACITY = 10;
    private static final int DEFAULT_LENGTH = 10;

    private static final int MAX_ARRAY_SIZE = Integer.MAX_VALUE - 8;
    private Object[] DEFAULT_ARR = {};
    private int size;

    public CustomList() {
        this.arr = DEFAULT_ARR;
    }

    public CustomList(Integer capacity) {
        if (capacity > 0) {
            this.arr = new Object[capacity];
        } else if (capacity == 0) {
            this.arr = DEFAULT_ARR;
        } else {
            throw new IllegalArgumentException("illegal capacity value");
        }
    }

    @Override
    public int size() {
        return size;
    }

    @Override
    public boolean isEmpty() {
        return false;
    }

    @Override
    public boolean contains(Object o) {
        return false;
    }

    @Override
    public Iterator<T> iterator() {
        return null;
    }

    @Override
    public Object[] toArray() {
        return new Object[0];
    }

    @Override
    public <T1> T1[] toArray(T1[] a) {
        return null;
    }

    @Override
    public boolean add(T t) {
        // if size reaches capacity need to grow and then insert
        // growth rate the current capacity << 1
        if (size() == arr.length) {
            arr = grow();
        }
        arr[size] = t;
        size++;
        return true;
    }

    private Object[] grow() {
        return grow(size + 1);
    }


    private Object[] grow(int minCap) {
        int newCap = minCap + minCap >> 1; // minCap + minCap / 2^1
        if (newCap - minCap <= 0) {
            // this is the first element getting added, so newCap = 0
            if (arr == DEFAULT_ARR) {
                Math.max(DEFAULT_CAPACITY, minCap);
            }
            if (minCap < 0) {
                throw new OutOfMemoryError();
            }
        }

        if (!(newCap - MAX_ARRAY_SIZE <= 0)) {
            newCap = hugeCap(minCap);
        }

        Object[] newArr = new Object[newCap];
        // copy to new array
        for (int i = 0; i < newCap; i++) {
            newArr[i] = arr[i];
        }
        return newArr;
    }

    private int hugeCap(int minCap) {
        if (minCap <= 0) {
            throw new OutOfMemoryError();
        } else {
            return minCap > MAX_ARRAY_SIZE ? Integer.MAX_VALUE : MAX_ARRAY_SIZE;
        }
    }

    @Override
    public boolean remove(Object o) {
        return false;
    }

    @Override
    public boolean containsAll(Collection<?> c) {
        return false;
    }

    @Override
    public boolean addAll(Collection<? extends T> c) {
        return false;
    }

    @Override
    public boolean addAll(int index, Collection<? extends T> c) {
        return false;
    }

    @Override
    public boolean removeAll(Collection<?> c) {
        return false;
    }

    @Override
    public boolean retainAll(Collection<?> c) {
        return false;
    }

    @Override
    public void clear() {

    }

    @Override
    public T get(int index) {
        return null;
    }

    @Override
    public T set(int index, T element) {
        return null;
    }

    @Override
    public void add(int index, T element) {

    }

    @Override
    public T remove(int index) {
        return null;
    }

    @Override
    public int indexOf(Object o) {
        return 0;
    }

    @Override
    public int lastIndexOf(Object o) {
        return 0;
    }

    @Override
    public ListIterator<T> listIterator() {
        return null;
    }

    @Override
    public ListIterator<T> listIterator(int index) {
        return null;
    }

    @Override
    public List<T> subList(int fromIndex, int toIndex) {
        return List.of();
    }

}