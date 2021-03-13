class FrontMiddleBackQueue {
    private LinkedList<Integer> p;
    
    public FrontMiddleBackQueue() {
        p = new LinkedList<>();
    }
    
    public void pushFront(int val) {
        p.addFirst(val);
    }
    
    public void pushMiddle(int val) {
        p.add(p.size()/2, val);
    }
    
    public void pushBack(int val) {
        p.addLast(val);
    }
    
    public int popFront() {
        return p.isEmpty() ? -1 : p.removeFirst();
    }
    
    public int popMiddle() {
        return p.isEmpty() ? -1 : p.remove((p.size()-1)/2);
    }
    
    public int popBack() {
        return p.isEmpty() ? -1 : p.removeLast();
    }
}

/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
 * FrontMiddleBackQueue obj = new FrontMiddleBackQueue();
 * obj.pushFront(val);
 * obj.pushMiddle(val);
 * obj.pushBack(val);
 * int param_4 = obj.popFront();
 * int param_5 = obj.popMiddle();
 * int param_6 = obj.popBack();
 */