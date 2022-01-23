class Solution {
    public String reformatDate(String date) {
        HashMap<String, String> month = new HashMap<>();
        month.put("Jan", "01");
        month.put("Feb", "02");
        month.put("Mar", "03");
        month.put("Apr", "04");
        month.put("May", "05");
        month.put("Jun", "06");
        month.put("Jul", "07");
        month.put("Aug", "08");
        month.put("Sep", "09");
        month.put("Oct", "10");
        month.put("Nov", "11");
        month.put("Dec", "12");
        String[] dates = date.split(" ");
        ArrayList<String> ans = new ArrayList<>();
        ans.add(dates[2]);
        ans.add(month.get(dates[1]));
        ans.add(dates[0].length()==3 ? "0"+dates[0].substring(0, 1) : dates[0].substring(0, 2));
        return String.join("-", ans);
    }
}