class SubrectangleQueries {
    private int[][] rectangle;
    private ArrayList<int[]> updates;

    public SubrectangleQueries(int[][] rectangle) {
        this.rectangle = rectangle;
        this.updates = new ArrayList<int[]>();
    }
    
    public void updateSubrectangle(int row1, int col1, int row2, int col2, int newValue) {
        updates.add(new int[]{row1, col1, row2, col2, newValue});
    }
    
    public int getValue(int row, int col) {
        for (int i = updates.size()-1; i >= 0; i--) {
            int[] update = updates.get(i);
            if (row >= update[0] && col >= update[1] && row <= update[2] && col <= update[3]) {
                return update[4];
            }
        }
        return rectangle[row][col];
    }
}

/**
 * Your SubrectangleQueries object will be instantiated and called as such:
 * SubrectangleQueries obj = new SubrectangleQueries(rectangle);
 * obj.updateSubrectangle(row1,col1,row2,col2,newValue);
 * int param_2 = obj.getValue(row,col);
 */